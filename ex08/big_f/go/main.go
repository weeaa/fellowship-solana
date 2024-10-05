package main

import (
	"bbl/bubblegum"
	"bufio"
	"context"
	"encoding/binary"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func main() {
	ctx := context.Background()

	client := rpc.New(rpc.DevNet_RPC)

	file, err := os.OpenFile("airdrop.txt", os.O_RDONLY, 777)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	pubkeys := make([]solana.PublicKey, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pubkeys = append(pubkeys, solana.MustPublicKeyFromBase58(scanner.Text()))
	}

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	payer, err := solana.WalletFromPrivateKeyBase58(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	treeKeyPair, err := newTree()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("treeKeyPair", treeKeyPair.wallet.PublicKey(), treeKeyPair.wallet.PrivateKey.String())

	sig, err := createTree(ctx, payer, solana.NewWallet(), client)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("createTree tx broadcasted: %s", sig.String())

	/*
		errCh := make(chan error)
		go func() {
			for {
				log.Println(<-errCh)
			}
		}()

		sigs, err := airdropCNFT(ctx, payer, treeKeyPair, pubkeys, errCh, client)
		if err != nil {
			log.Fatal(err)
		}

		for i, sig := range sigs {
			log.Printf("sent tx [%s] > %s", pubkeys[i], sig)
		}
	*/
}

var metadata = bubblegum.MetadataArgs{
	Name:                "weeaa's socials",
	Symbol:              "WEE",
	Uri:                 "",
	PrimarySaleHappened: false,
	IsMutable:           false,
	TokenProgramVersion: bubblegum.TokenProgramVersionToken2022,
}

func airdropCNFT(
	ctx context.Context,
	payer *solana.Wallet,
	tree *tree,
	pubkeys []solana.PublicKey,
	errCh chan error,
	client *rpc.Client,
) ([]solana.Signature, error) {

	bubblegumSigner, _, err := solana.FindProgramAddress([][]byte{[]byte("collection_cpi")}, bubblegum.ProgramID)
	if err != nil {
		return nil, err
	}

	_ = bubblegumSigner

	sigs := make([]solana.Signature, len(pubkeys))

	for _, pubkey := range pubkeys {
		_ = pubkey
		/*
			mintIx := bubblegum.NewMintToCollectionV1Instruction(
				metadata,
				solana.PublicKey{},
				pubkey,
				pubkey,
				tree.wallet.PublicKey(),
				payer.PublicKey(),
				payer.PublicKey(),
				NOOP_PROGRAM_ID,
				COMPRESSION_PROGRAM_ID,
				solana.SystemProgramID,
			)
		*/

		log.Println("b4 block")
		blockHash, err := client.GetRecentBlockhash(ctx, rpc.CommitmentConfirmed)
		if err != nil {
			errCh <- err
			continue
		}

		tx, err := solana.NewTransaction([]solana.Instruction{
			//mintIx.Build(),
		},
			blockHash.Value.Blockhash,
			solana.TransactionPayer(payer.PublicKey()))
		if err != nil {
			errCh <- err
			continue
		}

		if _, err = tx.Sign(
			func(key solana.PublicKey) *solana.PrivateKey {
				if payer.PublicKey().Equals(key) {
					return &payer.PrivateKey
				}
				return nil
			},
		); err != nil {
			errCh <- err
			continue
		}

		log.Println("b4 tx")
		sig, err := client.SendTransaction(ctx, tx)
		if err != nil {
			errCh <- err
			continue
		}

		sigs = append(sigs, sig)

		time.Sleep(5 * time.Second)
	}

	return sigs, nil
}

type tree struct {
	wallet    *solana.Wallet
	authority solana.PublicKey
}

func newTree() (*tree, error) {
	treeKeypair := solana.NewWallet()

	treeAuthority, _, err := solana.FindProgramAddress([][]byte{treeKeypair.PublicKey().Bytes()}, BUBBLEGUM_PROGRAM_ID)
	if err != nil {
		return nil, err
	}

	return &tree{wallet: treeKeypair, authority: treeAuthority}, nil
}

const MAX_DEPTH = 7
const MAX_BUFFER_SIZE = 16

// createTree creates a new merkle tree.
func createTree(ctx context.Context, payer *solana.Wallet, treeKeyPair *solana.Wallet, client *rpc.Client) (*solana.Signature, error) {
	// Constants
	const MAX_DEPTH = 14
	const MAX_BUFFER_SIZE = 64

	// Ensure these program IDs are correct
	BUBBLEGUM_PROGRAM_ID := solana.MustPublicKeyFromBase58("BGUMAp9Gq7iTEuizy4pqaxsTyUCBK68MDfK752saRPUY")
	COMPRESSION_PROGRAM_ID := solana.MustPublicKeyFromBase58("cmtDvXumGCrqC1Age74AVPhSRVXJMd8PJS91L8KbNCK")
	LOG_WRAPPER_PROGRAM_ID := solana.MustPublicKeyFromBase58("noopb9bkMVfRPU8AsbpTUg8AQkHtKwMYZiFUjNRtMmV")

	minBalance, err := client.GetMinimumBalanceForRentExemption(ctx, 1880, rpc.CommitmentFinalized)
	if err != nil {
		return nil, fmt.Errorf("failed to get minimum balance: %w", err)
	}

	createAccountIx := system.NewCreateAccountInstruction(
		minBalance,
		1880,
		COMPRESSION_PROGRAM_ID,
		payer.PublicKey(),
		treeKeyPair.PublicKey(),
	)

	treeAuthority, _, err := solana.FindProgramAddress(
		[][]byte{treeKeyPair.PublicKey().Bytes()},
		BUBBLEGUM_PROGRAM_ID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to derive tree authority: %w", err)
	}

	// Manually construct the CreateTree instruction
	createTreeAccounts := solana.AccountMetaSlice{
		{PublicKey: treeAuthority, IsWritable: true, IsSigner: false},
		{PublicKey: treeKeyPair.PublicKey(), IsWritable: true, IsSigner: true},
		{PublicKey: payer.PublicKey(), IsWritable: false, IsSigner: true},
		{PublicKey: LOG_WRAPPER_PROGRAM_ID, IsWritable: false, IsSigner: false},
		{PublicKey: COMPRESSION_PROGRAM_ID, IsWritable: false, IsSigner: false},
		{PublicKey: solana.SystemProgramID, IsWritable: false, IsSigner: false},
	}

	// Construct the instruction data
	instructionData := make([]byte, 17)
	binary.LittleEndian.PutUint32(instructionData[0:], 0) // Discriminator for CreateTree
	binary.LittleEndian.PutUint32(instructionData[4:], MAX_DEPTH)
	binary.LittleEndian.PutUint32(instructionData[8:], MAX_BUFFER_SIZE)
	instructionData[12] = 1                                // public as bool
	binary.LittleEndian.PutUint32(instructionData[13:], 0) // canopyDepth (optional, set to 0)

	createTreeIx := solana.NewInstruction(
		BUBBLEGUM_PROGRAM_ID,
		createTreeAccounts,
		instructionData,
	)

	blockHash, err := client.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return nil, err
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			createAccountIx.Build(),
			createTreeIx,
		},
		blockHash.Value.Blockhash,
		solana.TransactionPayer(payer.PublicKey()),
	)
	if err != nil {
		return nil, err
	}

	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if payer.PublicKey().Equals(key) {
			return &payer.PrivateKey
		}
		if treeKeyPair.PublicKey().Equals(key) {
			return &treeKeyPair.PrivateKey
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Log the full transaction for debugging
	log.Printf("Full Transaction: %+v", tx)

	sig, err := client.SendTransaction(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("error sending merkle tree tx: %w", err)
	}

	return &sig, nil
}

/*
func f() {
	createTreeInstruction := solana.NewInstruction(
		solana.MustPublicKeyFromBase58("Bubblegum Program PublicKey"), // Replace with actual Bubblegum program ID
		solana.AccountMetaSlice{
			solana.NewAccountMeta(treeAuthority.PublicKey(), true, false), // Tree Authority
			solana.NewAccountMeta(merkleTree.PublicKey(), true, true),     // Merkle Tree
			solana.NewAccountMeta(payer.PublicKey(), true, true),          // Payer
			solana.NewAccountMeta(treeCreator.PublicKey(), true, true),    // Tree Creator
			solana.NewAccountMeta(logWrapper, false, false),               // Log Wrapper
			solana.NewAccountMeta(compressionProgram, false, false),       // Compression Program
			solana.NewAccountMeta(systemProgram, false, false),            // System Program
		},
		solana.NewCreateTreeParams{
			MaxDepth:       maxDepth,
			MaxBufferSize:  maxBufferSize,
			PublicBoolFlag: publicBoolFlag,
		},
	)
}
*/
