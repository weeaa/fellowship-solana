package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/rpc"
	"log"
	"time"
)

func main() {
	// set flags
	sendSolF := flag.Int("sendSol", 0, "Amount of SOL to send")
	fromF := flag.String("from", "", "Funding account (private key)")
	toF := flag.String("to", "", "Solana address you want to send funds to")
	newWalletF := flag.Bool("newWallet", false, "Generate a new wallet")

	flag.Parse()

	if *newWalletF {
		wallet := newWallet()
		fmt.Printf("PubKey: %s\nPrivKey: %s\n", wallet.PublicKey().String(), wallet.PrivateKey.String())
	}

	if *sendSolF > 0 {
		if *toF == "" {
			log.Fatal("to flag: you must provide a non empty string!")
		}
		if *fromF == "" {
			log.Fatal("from flag: you must provide a non empty string")
		}
		if err := sendSol(*sendSolF, *toF, *fromF); err != nil {
			log.Fatal(err)
		}
	}
}

// sendSol sends x lamport (testnet).
func sendSol(amt int, to, from string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := rpc.New(rpc.DevNet_RPC)

	// parse to & from keys
	privKey, err := solana.PrivateKeyFromBase58(from)
	if err != nil {
		return err
	}
	receiver, err := solana.PublicKeyFromBase58(to)
	if err != nil {
		return err
	}

	recent, err := client.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		return err
	}

	// create tx, self-explanatory tbh
	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(
				uint64(amt*1e9),
				privKey.PublicKey(),
				receiver,
			).Build(),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(privKey.PublicKey()),
	)
	if err != nil {
		return err
	}

	if _, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if privKey.PublicKey().Equals(key) {
			return &privKey
		}
		return nil
	}); err != nil {
		return err
	}

	sig, err := client.SendTransaction(ctx, tx)
	if err != nil {
		return err
	}

	fmt.Printf("tx sig: %s", sig.String())
	return nil
}

// newWallet generates a new keypair.
func newWallet() *solana.Wallet {
	return solana.NewWallet()
}
