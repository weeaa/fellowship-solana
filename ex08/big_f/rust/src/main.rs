mod airdrop;

use solana_client::rpc_client::RpcClient;
use solana_sdk::signature::{Keypair, Signer};
use anyhow::Result;
use std::str::FromStr;
use dotenv::dotenv;
use std::env;
use base64::engine::general_purpose::STANDARD;
use base64::Engine;
use std::error::Error;

fn main() {
    let client = RpcClient::new("https://api.devnet.solana.com".to_string());
    let payer = load_keypair_from_env();

    // 1. create a merkle tree
    let result = airdrop::create_merkle_tree(&client, &payer);

    Ok(result.is_ok()).expect("error creating merkle tree");

    let sig = result.unwrap();
    println!("sig: {}", sig);


    // 2. create nfts & airdrop them to wallets
   // airdrop::airdrop_compressed_nft(&client, &payer, &merkle_tree_pubkey);

   // Ok(());
}

fn load_keypair_from_env() -> Result<Keypair, Box<dyn Error>> {
    dotenv().ok();
    let private_key = env::var("PRIVATE_KEY")?;

    let keypair_bytes: Vec<u8> = STANDARD.decode(&private_key)?;
    let keypair = Keypair::from_bytes(&keypair_bytes)?;

    Ok(keypair)
}