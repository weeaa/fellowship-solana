import {
    PublicKey,
    Keypair,
    Transaction,
    SystemProgram,
} from "@solana/web3.js";
import {
    createInitializeMintInstruction,
    getMinimumBalanceForRentExemptMint,
    TOKEN_PROGRAM_ID,
    MINT_SIZE,
    createTransferInstruction,
    createMintToCheckedInstruction,
    createBurnCheckedInstruction,
    createApproveCheckedInstruction,
    createMintToInstruction,
    createAssociatedTokenAccountInstruction,
    getAssociatedTokenAddress, createTransferCheckedInstruction,
} from "@solana/spl-token";
import { useState } from 'react';
import { useWallet, useConnection } from '@solana/wallet-adapter-react';

export default function Page() {
    const { connection } = useConnection();
    const { publicKey, signTransaction, sendTransaction } = useWallet();
    const [transferAmount, setTransferAmount] = useState('');
    const [recipient, setRecipient] = useState('');
    const [mintAmount, setMintAmount] = useState('');
    const [burnAmount, setBurnAmount] = useState('');
    const [delegateAmount, setDelegateAmount] = useState('');
    const [delegateAddress, setDelegateAddress] = useState('');
    const [newToken, setNewToken] = useState(null);

    const createToken = async () => {
        if (!publicKey) {
            alert("Please connect your wallet.");
            return;
        }

        try {
            const mint = Keypair.generate();
            const transaction = new Transaction().add(
                SystemProgram.createAccount({
                    fromPubkey: publicKey,
                    newAccountPubkey: mint.publicKey,
                    space: MINT_SIZE,
                    lamports: await getMinimumBalanceForRentExemptMint(connection),
                    programId: TOKEN_PROGRAM_ID,
                }),
                createInitializeMintInstruction(
                    mint.publicKey,
                    Number(8),
                    publicKey,
                    publicKey
                )
            );

            const signature = await sendTransaction(transaction, connection, {
                signers: [mint],
            });
            console.log("New token created with signature:", signature);
            await connection.confirmTransaction(signature, "processed");

            const mintAddress = mint.publicKey.toBase58();
            setNewToken(mintAddress);

            console.log("New mint address:", mintAddress);
        } catch (error) {
            console.error("Error creating token:", error);
        }
    };

    const transfer = async () => {
        if (!publicKey || !transferAmount || !recipient || !newToken) {
            alert("Please ensure the wallet is connected and all fields are filled, including recipient address.");
            return;
        }

        try {
            const recipientPublicKey = new PublicKey(recipient);
            const mintPublicKey = new PublicKey(newToken);
            const decimals = 8;
            const transferAmountWithDecimals = Number(transferAmount) * Math.pow(10, decimals);

            const senderTokenAccount = await getAssociatedTokenAddress(
                mintPublicKey,
                publicKey
            );
            const recipientTokenAccount = await getAssociatedTokenAddress(
                mintPublicKey,
                recipientPublicKey
            );

            let transaction = new Transaction();

            const recipientAccountInfo = await connection.getAccountInfo(recipientTokenAccount);
            if (!recipientAccountInfo) {
                transaction.add(
                    createAssociatedTokenAccountInstruction(
                        publicKey,
                        recipientTokenAccount,
                        recipientPublicKey,
                        mintPublicKey
                    )
                );
            }

            transaction.add(
                createTransferCheckedInstruction(
                    senderTokenAccount,
                    mintPublicKey,
                    recipientTokenAccount,
                    publicKey,
                    transferAmountWithDecimals,
                    decimals
                )
            );

            const { blockhash, lastValidBlockHeight } = await connection.getLatestBlockhash();
            transaction.feePayer = publicKey;
            transaction.recentBlockhash = blockhash;

            const signature = await sendTransaction(transaction, connection);
            const confirmation = await connection.confirmTransaction(
                { signature, blockhash, lastValidBlockHeight },
                'confirmed'
            );

            if (confirmation.value.err) {
                throw new Error(`Transaction failed: ${confirmation.value.err.toString()}`);
            }

            console.log("Transfer completed with signature:", signature);
            alert(`Successfully transferred ${transferAmount} tokens to ${recipient}.`);
        } catch (error) {
            console.error("Error transferring tokens:", error);
            alert(`Error transferring tokens: ${error.message}`);
        }
    };

    const mint = async () => {
        if (!publicKey || !mintAmount || !newToken) {
            alert("Please connect your wallet and enter the Mint Amount.");
            return;
        }

        try {
            const tokenAddress = typeof newToken === 'string' ? newToken : String(newToken);
            const mintPublicKey = new PublicKey(tokenAddress);
            const decimals = 8;
            const mintAmountWithDecimals = Number(mintAmount) * Math.pow(10, decimals);

            const associatedTokenAddress = await getAssociatedTokenAddress(
                mintPublicKey,
                publicKey
            );
            const accountInfo = await connection.getAccountInfo(associatedTokenAddress);

            let transaction = new Transaction();

            if (!accountInfo) {
                transaction.add(
                    createAssociatedTokenAccountInstruction(
                        publicKey,
                        associatedTokenAddress,
                        publicKey,
                        mintPublicKey
                    )
                );
            }

            transaction.add(
                createMintToCheckedInstruction(
                    mintPublicKey,
                    associatedTokenAddress,
                    publicKey,
                    mintAmountWithDecimals,
                    decimals
                )
            );

            const { blockhash } = await connection.getLatestBlockhash();
            transaction.recentBlockhash = blockhash;
            transaction.feePayer = publicKey;

            const signature = await sendTransaction(transaction, connection);
            console.log("Transaction sent. Signature:", signature);

            await connection.confirmTransaction(signature, 'confirmed');
            console.log("Minting completed with signature:", signature);
            alert("Tokens minted successfully!");
        } catch (error) {
            console.error("Error minting tokens:", error);
            alert(`Error minting tokens: ${error.message}`);
        } finally {

        }
    };


    const burn = async () => {
        if (!publicKey || !burnAmount || !newToken) {
            alert("Please connect your wallet and enter the Burn Amount.");
            return;
        }

        try {
            const mintPublicKey = new PublicKey(newToken);
            const decimals = 8;
            const burnAmountWithDecimals = Number(burnAmount) * Math.pow(10, decimals);

            const associatedTokenAddress = await getAssociatedTokenAddress(
                mintPublicKey,
                publicKey
            );

            const accountInfo = await connection.getAccountInfo(associatedTokenAddress);
            if (!accountInfo) {
                throw new Error("Associated token account not found. You may not have any tokens to burn.");
            }

            const tokenBalance = await connection.getTokenAccountBalance(associatedTokenAddress);
            if (tokenBalance.value.uiAmount < Number(burnAmount)) {
                throw new Error(`Insufficient balance. You have ${tokenBalance.value.uiAmount} tokens.`);
            }

            const transaction = new Transaction().add(
                createBurnCheckedInstruction(
                    associatedTokenAddress,
                    mintPublicKey,
                    publicKey,
                    burnAmountWithDecimals,
                    decimals
                )
            );

            const { blockhash, lastValidBlockHeight } = await connection.getLatestBlockhash();
            transaction.feePayer = publicKey;
            transaction.recentBlockhash = blockhash;

            const signature = await sendTransaction(transaction, connection);
            const confirmation = await connection.confirmTransaction(
                { signature, blockhash, lastValidBlockHeight },
                'confirmed'
            );

            if (confirmation.value.err) {
                throw new Error(`Transaction failed: ${confirmation.value.err.toString()}`);
            }

            console.log("Burning completed with signature:", signature);
            alert(`Successfully burned ${burnAmount} tokens.`);
        } catch (error) {
            console.error("Error burning tokens:", error);
            alert(`Error burning tokens: ${error.message}`);
        } finally {

   }
    };

    const delegate = async () => {
        if (!publicKey || !delegateAddress || !delegateAmount || !newToken) {
            alert("Please connect your wallet and fill all fields for delegation.");
            return;
        }

        try {
            const mintPublicKey = new PublicKey(newToken);
            const delegatePublicKey = new PublicKey(delegateAddress);
            const decimals = 8;
            const delegateAmountWithDecimals = Number(delegateAmount) * Math.pow(10, decimals);

            const ownerTokenAccount = await getAssociatedTokenAddress(
                mintPublicKey,
                publicKey
            );

            const accountInfo = await connection.getAccountInfo(ownerTokenAccount);
            if (!accountInfo) {
                throw new Error("Owner's token account not found. You may not have any tokens to delegate.");
            }

            const transaction = new Transaction().add(
                createApproveCheckedInstruction(
                    ownerTokenAccount,
                    mintPublicKey,
                    delegatePublicKey,
                    publicKey,
                    delegateAmountWithDecimals,
                    decimals
                )
            );

            const { blockhash, lastValidBlockHeight } = await connection.getLatestBlockhash();
            transaction.feePayer = publicKey;
            transaction.recentBlockhash = blockhash;

            const signature = await sendTransaction(transaction, connection);
            const confirmation = await connection.confirmTransaction(
                { signature, blockhash, lastValidBlockHeight },
                'confirmed'
            );

            if (confirmation.value.err) {
                throw new Error(`Transaction failed: ${confirmation.value.err.toString()}`);
            }

            console.log("Delegation completed with signature:", signature);
            alert(`Successfully delegated ${delegateAmount} tokens to ${delegateAddress}.`);
        } catch (error) {
            console.error("Error delegating tokens:", error);
            alert(`Error delegating tokens: ${error.message}`);
        }
    };

    return (
        <div className="flex flex-col items-center justify-center min-h-screen bg-gray-100 p-8 space-y-8">
            <h1 className="text-4xl font-bold text-gray-800 mb-4">Solana Token Manager</h1>
            <div className="badge badge-info mb-6">
                {publicKey ? publicKey.toBase58() : "Not connected"}
            </div>

            {newToken && (
                <div className="alert alert-success shadow-lg mb-4">
                    <div>
                        <span>New Token Created: {newToken}</span>
                    </div>
                </div>
            )}

            <div className="w-full max-w-md p-6 bg-white rounded-lg shadow-lg space-y-4">
                <button onClick={createToken} className="btn btn-primary w-full">
                    Create New Token
                </button>

                {/* Transfer Tokens */}
                <div className="space-y-2">
                    <input
                        type="text"
                        value={transferAmount}
                        onChange={(e) => setTransferAmount(e.target.value)}
                        placeholder="Transfer Amount"
                        className="input input-bordered w-full"
                    />
                    <input
                        type="text"
                        value={recipient}
                        onChange={(e) => setRecipient(e.target.value)}
                        placeholder="Recipient Address"
                        className="input input-bordered w-full"
                    />
                    <button onClick={transfer} className="btn btn-success w-full">
                        Transfer
                    </button>
                </div>

                {/* Mint Tokens */}
                <div className="space-y-2">
                    <input
                        type="text"
                        value={mintAmount}
                        onChange={(e) => setMintAmount(e.target.value)}
                        placeholder="Mint Amount"
                        className="input input-bordered w-full"
                    />
                    <button onClick={mint} className="btn btn-purple w-full">
                        Mint
                    </button>
                </div>

                {/* Burn Tokens */}
                <div className="space-y-2">
                    <input
                        type="text"
                        value={burnAmount}
                        onChange={(e) => setBurnAmount(e.target.value)}
                        placeholder="Burn Amount"
                        className="input input-bordered w-full"
                    />
                    <button onClick={burn} className="btn btn-error w-full">
                        Burn
                    </button>
                </div>

                {/* Delegate Tokens */}
                <div className="space-y-2">
                    <input
                        type="text"
                        value={delegateAddress}
                        onChange={(e) => setDelegateAddress(e.target.value)}
                        placeholder="Delegate Address"
                        className="input input-bordered w-full"
                    />
                    <input
                        type="text"
                        value={delegateAmount}
                        onChange={(e) => setDelegateAmount(e.target.value)}
                        placeholder="Delegate Amount"
                        className="input input-bordered w-full"
                    />
                    <button onClick={delegate} className="btn btn-warning w-full">
                        Delegate
                    </button>
                </div>
            </div>
        </div>
    );
}
