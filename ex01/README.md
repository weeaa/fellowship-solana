# Exercise 01 of Solana Summer Fellowship

## Compile (make sure to have go installed)
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