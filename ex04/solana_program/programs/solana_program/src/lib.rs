use anchor_lang::prelude::*;

declare_id!("D9fn8aHrzA6k5oZkKfkNb7xvskduq1Fw49uunwTTf9L6");

#[program]
pub mod solana_program {
    use super::*;

    pub fn initialize_account(ctx: Context<InitializeAccount>) -> ProgramResult {
        msg!("initializing the account...");
        let account = &mut ctx.accounts.deposit_account;
        account.owner = *ctx.accounts.signer.key;
        account.amount_deposited = 0;
        Ok(())
    }

    pub fn deposit_sol(ctx: Context<DepositSol>, amount: u64) -> ProgramResult {
        let account = &mut ctx.accounts.deposit_account;
        let ix = system_instruction::transfer(
            &ctx.accounts.signer.key,
            &account.key(),
            amount,
        );
        anchor_lang::solana_program::program::invoke(
            &ix,
            &[
                ctx.accounts.signer.to_account_info(),
                account.to_account_info(),
            ],
        )?;

        account.amount_deposited += amount;
        msg!("deposited {} lamports", amount);
        Ok(())
    }

    // Withdraw 10% of deposited SOL
    pub fn withdraw_10_percent(ctx: Context<WithdrawSol>) -> ProgramResult {
        let account = &mut ctx.accounts.deposit_account;
        let withdraw_amount = account.amount_deposited / 10;

        if withdraw_amount > **ctx.accounts.deposit_account.to_account_info().lamports.borrow() {
            return Err(ProgramError::InsufficientFunds.into());
        }

        **ctx.accounts.deposit_account.to_account_info().try_borrow_mut_lamports()? -= withdraw_amount;
        **ctx.accounts.receiver.to_account_info().try_borrow_mut_lamports()? += withdraw_amount;

        account.amount_deposited -= withdraw_amount;
        msg!("withdrew 10% of deposited sol: {}", withdraw_amount);
        Ok(())
    }
}

#[derive(Accounts)]
pub struct InitializeAccount<'info> {
    #[account(init, payer = signer, space = 8 + 32 + 8)]
    pub deposit_account: Account<'info, DepositAccount>,
    #[account(mut)]
    pub signer: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct DepositSol<'info> {
    #[account(mut)]
    pub deposit_account: Account<'info, DepositAccount>,
    #[account(mut)]
    pub signer: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct WithdrawSol<'info> {
    #[account(mut, has_one = owner)]
    pub deposit_account: Account<'info, DepositAccount>,
    #[account(mut)]
    pub receiver: AccountInfo<'info>,
    pub owner: Signer<'info>,
    pub system_program: Program<'info, System>,
}

#[account]
pub struct DepositAccount {
    pub owner: Pubkey,
    pub amount_deposited: u64,
}
