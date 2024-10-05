# Exercise 01 of Solana Summer Fellowship

## Prerequisites
- Go 1.22.4

## Compile
```bash
git clone github.com/weeaa/solana-fellowship
cd ex01
go mod tidy
go build .
```

## Transfer funds (devnet)
```bash
./ex01 -sendSol <amt> -from <privkey> -to <pubkey>
```

## Create keypair
```bash
./ex01 -newWallet
```