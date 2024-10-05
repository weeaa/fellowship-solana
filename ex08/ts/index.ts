import {
    createTree,
    findLeafAssetIdPda,
    getAssetWithProof,
    mplBubblegum,
    parseLeafFromMintV1Transaction,
    verifyCollection,
    mintToCollectionV1,
} from '@metaplex-foundation/mpl-bubblegum';
import { createNft, mplTokenMetadata } from '@metaplex-foundation/mpl-token-metadata';
import {
    keypairIdentity,
    generateSigner,
    percentAmount,
    publicKey,
    PublicKey,
} from '@metaplex-foundation/umi';
import { createUmi } from '@metaplex-foundation/umi-bundle-defaults';
import { irysUploader } from '@metaplex-foundation/umi-uploader-irys';
import { config } from 'dotenv';
import bs58 from 'bs58';

config();

const PRIVATE_KEY = "" // 👈 ADD YOUR PRIVATE KEY HERE

const addresses: string[] = [
    'A6njahNqC6qKde6YtbHdr1MZsB5KY9aKfzTY1cj8jU3v',
    //todo grab from discord
];

const umi = createUmi('https://api.mainnet-beta.solana.com')
    .use(mplBubblegum())
    .use(mplTokenMetadata())
    .use(irysUploader());

// iumi starts the Umi instance with the necessary plugins and keypair.
const iumi = () => {
    let keypair = umi.eddsa.createKeypairFromSecretKey(bs58.decode(PRIVATE_KEY));
    umi.use(keypairIdentity(keypair));
};

// createCollection creates a new NFT collection + w custom metadata.
const createCollection = async (): Promise<{collectionId: PublicKey, nftMetadataUri: string}> => {
    const collectionId = generateSigner(umi);
    const nftImageUri = ["https://cyan-given-gazelle-156.mypinata.cloud/ipfs/QmbCjcVoxDobJzAB9zCk9d5Cct19o1mJAbCy5oxTuYgYeR"];

    console.log('Collection Image URI:', nftImageUri);

    const collectionMetadata = {
        name: 'weeaa',
        image: nftImageUri[0],
        externalUrl: 'https://x.com/weea_a',
        properties: {
            files: [
                {
                    uri: nftImageUri[0],
                    type: 'image/png',
                },
            ],
        },
    }

    const collectionMetadataUri = await umi.uploader.uploadJson(
        collectionMetadata
    )

    await createNft(umi, {
        mint: collectionId,
        name: 'weeaa',
        uri: collectionMetadataUri,
        isCollection: true,
        sellerFeeBasisPoints: percentAmount(0),
    }).sendAndConfirm(umi)

    const nftMetadata = {
        name: 'weeaa',
        image: nftImageUri[0],
        externalUrl: 'https://x.com/weea_a',
        attributes: [
            { trait_type: 'twitter', value: 'https://x.com/weea_a' },
            { trait_type: 'github', value: 'https://github.com/weeaa' },
        ],
        properties: {
            files: [{ uri: nftImageUri[0],
                type: 'image/png',
            }],
        },
    };

    const nftMetadataUri = await umi.uploader.uploadJson(nftMetadata)

    console.log("coll done!")
    return {
        collectionId: collectionId.publicKey,
        nftMetadataUri,
    };
};

// createMerkleTree creates a merkle tree for storing compressed NFTs.
const createMerkleTree = async (): Promise<PublicKey> => {
    const merkleTree = generateSigner(umi);

    const createTreeTx = await createTree(umi, {
        merkleTree,
        maxDepth: 14,
        maxBufferSize: 64,
        canopyDepth: 0,
    })

    await createTreeTx.sendAndConfirm(umi)
    return merkleTree.publicKey;
};

// airdropNft airdrops your NFT collection to custom addresses
const airdropNft = async (merkleTree: PublicKey, collectionId: PublicKey, nftMetadataUri: string) => {
    for (const address of addresses) {
        const { signature } = await mintToCollectionV1(umi, {
            leafOwner: publicKey(address),
            merkleTree: merkleTree,
            collectionMint: collectionId,
            metadata: {
                name: 'weeaa',
                uri: nftMetadataUri,
                sellerFeeBasisPoints: 100,
                collection: { key: collectionId, verified: false },
                creators: [{ address: umi.identity.publicKey, verified: true, share: 100 }],
            },
        }).sendAndConfirm(umi, { send: { commitment: 'finalized' } });

        const leaf = await parseLeafFromMintV1Transaction(umi, signature);
        const assetId = findLeafAssetIdPda(umi, { merkleTree: merkleTree, leafIndex: leaf.nonce });
        console.log(`nft minted w asset id > ${assetId.toString()}`);
        await sleep(5000); // to avoid rate limits
    }
};

const sleep = (ms: number): Promise<void> => {
    return new Promise((resolve) => setTimeout(resolve, ms));
};


const create_cnft = async () => {
    iumi();
    console.log('umi initialized');
    const merkleTree = await createMerkleTree();
    console.log("merkle tree created", merkleTree.toString());
    const { collectionId, nftMetadataUri } = await createCollection();
    console.log("collection created");
    await airdropNft(merkleTree, collectionId, nftMetadataUri);
    console.log("nfts airdropped");
};

create_cnft();
