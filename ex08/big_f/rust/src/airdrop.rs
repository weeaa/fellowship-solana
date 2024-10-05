use solana_client::rpc_client::RpcClient;
use solana_sdk::{
    system_instruction,
    transaction::Transaction,
    commitment_config::{CommitmentConfig, CommitmentLevel},
    instruction::{Instruction},
    signature::Keypair,
    signer::Signer,
    system_program,
    signature::Signature,
};
use anchor_lang::InstructionData;
use solana_client::rpc_config::{
    RpcSendTransactionConfig,
};
use anyhow::Result;
use std::fs::File;
use std::io::{BufRead, BufReader};
use anchor_lang::prelude::*;
use std::thread;
use std::time::Duration;

use anchor_lang::prelude::*;
use spl_account_compression::{self, state::CONCURRENT_MERKLE_TREE_HEADER_SIZE_V1};
use mpl_bubblegum::programs::SPL_NOOP_ID; // Assuming this is the correct constant.
use mpl_bubblegum::instructions::{CreateTreeConfig, CreateTreeConfigInstructionArgs};

use anchor_spl::{token::Token, token_interface::Mint};
use mpl_bubblegum::accounts::TreeConfig;
use mpl_bubblegum::types::{Collection, Creator, MetadataArgs};
use spl_account_compression::{program::SplAccountCompression, Noop};
use std::error::Error;

#[derive(Accounts)]
pub struct MintCnftToCollectionCpi<'info> {
    /// CHECK: will used by mpl_bubblegum program
    #[account(mut)]
    pub tree_config: UncheckedAccount<'info>,
    pub leaf_owner: SystemAccount<'info>,
    pub leaf_delegate: SystemAccount<'info>,
    /// CHECK: will used by mpl_bubblegum program
    #[account(mut)]
    pub merkle_tree: UncheckedAccount<'info>,
    #[account(mut)]
    pub payer: Signer<'info>,
    pub tree_creator: Signer<'info>,
    pub collection_authority: Signer<'info>,
    #[account(
        mint::token_program = Token::id(),
        mint::authority = edition_account
    )]
    pub collection_mint: InterfaceAccount<'info, Mint>,
    /// CHECK:
    #[account(mut)]
    pub collection_metadata: UncheckedAccount<'info>,
    /// CHECK
    pub edition_account: UncheckedAccount<'info>,
    pub log_wrapper: Program<'info, Noop>,
    pub compression_program: Program<'info, SplAccountCompression>,
    pub token_metadata_program: Program<'info, TokenMetadataProgram>,
    pub system_program: Program<'info, System>,
    pub mpl_bubblegum_program: Program<'info, MplBubblegumProgram>,
}

// airdrops ur nfts 🤓
pub fn airdrop_cnft(
    client: &RpcClient,
    payer: &Keypair,
    merkle_tree: &Keypair,
) {
    let (bubblegum_signer, _) = Pubkey::find_program_address(
        &["collection_cpi".as_bytes()],
        &mpl_bubblegum::programs::MPL_BUBBLEGUM_ID,
    );


    // as per recommended by compressed.app
    let canopy_depth = 6;

    // define metadata of our nft
    let metadata = MetadataArgs {
        name: "rizziere's socials".to_string(),
        symbol: "wee".to_string(),
        uri: "https://github.com/weeaa/metadata_tmp/cnft_metadata.json".to_string(),
        creators: Vec::new(),
        edition_nonce: Some(0),
        uses: None,
        collection: None,
        primary_sale_happened: false,
        seller_fee_basis_points: 0,
        is_mutable: false,
        token_program_version: mpl_bubblegum::types::TokenProgramVersion::Original,
        token_standard: Some(mpl_bubblegum::types::TokenStandard::NonFungible),
    };

    let file = File::open("airdrop.txt");
    let reader = BufReader::new(file);
    //let x  = MintToCollectionV1CpiBuilder();
    for line in reader.lines() {
        match line {
            Ok(receiver) => {
                println!("sending tx to {} in 5s...", receiver);
                thread::sleep(Duration::from_secs(5));


                let ix = mpl_bubblegum::instructions::MintToCollectionV1{
                    payer,
                    merkle_tree: merkle_tree.pubkey(),
                    leaf_owner: receiver,
                    leaf_delegate: payer,
                    collection_authority: payer,
                    collection_authority_record_pda: bubblegum_signer,
                    collection_mint,
                    collection_metadata,
                    bubblegum_signer,
                    compression_program: SplAccountCompression,
                    tree_config,
                    tree_creator_or_delegate: payer,
                    collection_edition,

                    log_wrapper: SPL_NOOP_ID,
                    token_metadata_program: mpl_token_metadata::programs::MPL_TOKEN_METADATA_ID,
                    system_program: system_program::ID,
                };

                let mut txn = Transaction::new_with_payer(&[ix], Some(&payer.pubkey()));
                let recent_blockhash = client.get_latest_blockhash()?;

                txn.sign(&[payer], recent_blockhash);

                let signature = client.send_and_confirm_transaction_with_spinner_and_config(
                    &txn,
                    CommitmentConfig::confirmed(),
                    RpcSendTransactionConfig {
                        skip_preflight: false,
                        preflight_commitment: Some(CommitmentLevel::Confirmed),
                        ..RpcSendTransactionConfig::default()
                    },
                )?;
                println!("tx sent -> sig {}", signature);
            }
            Err(e) => {
                eprintln!("error reading line: {}", e);
            }
        }
    }
}

pub fn create_merkle_tree(
    client: &RpcClient,
    payer: &Keypair,
) -> Result<Signature, Box<dyn Error>> {
    let tree = Keypair::new();
    let tree_pubkey = tree.try_pubkey().unwrap_or_else(|_| unreachable!());

    let (tree_authority, _bump) =
        Pubkey::find_program_address(&[tree_pubkey.as_ref()], &mpl_bubblegum::programs::MPL_BUBBLEGUM_ID);

    let size: u64 = 1880;
    let lmps: u64 = 1_000_000;

    let create_account_ix = system_instruction::create_account(
        &payer.pubkey(),
        &tree_pubkey,
        lmps,
        size,
        &system_program::id(),
    );

    let args = CreateTreeConfigInstructionArgs {
        max_depth: 14,
        max_buffer_size: 64,
        public: Some(true),
    };

    let create_tree_ix = CreateTreeConfig{
        merkle_tree: tree_pubkey.into(),
        payer: payer.pubkey().into(),
        tree_creator: payer.pubkey().into(),
        tree_config: tree_authority,
        compression_program: mpl_bubblegum::programs::SPL_ACCOUNT_COMPRESSION_ID,
        log_wrapper: SPL_NOOP_ID,
        system_program: System::id(),
    }.instruction(args);


    let recent_blockhash = client
        .get_latest_blockhash()
        .expect("failed to get recent block hash");

    let transaction = Transaction::new_signed_with_payer(
        &[create_account_ix, create_tree_ix.into()],
        Some(&payer.pubkey()),
        &[payer, &tree],
        recent_blockhash,
    );

    let signature = client.send_and_confirm_transaction(&transaction)?;

    Ok(signature)
}

fn create_collection() {

}