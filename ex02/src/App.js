import React from "react";
import {
  WalletProvider,
  ConnectionProvider,
} from "@solana/wallet-adapter-react";
import {
  WalletModalProvider,
  WalletMultiButton,
} from "@solana/wallet-adapter-react-ui";
import { PhantomWalletAdapter } from "@solana/wallet-adapter-wallets";
import { clusterApiUrl } from "@solana/web3.js";
import "@solana/wallet-adapter-react-ui/styles.css";
import Page from "./manager";

const App = () => {
  const wallets = [new PhantomWalletAdapter()]; // Use Phantom wallet

  return (
      <ConnectionProvider endpoint={clusterApiUrl("devnet")}>
        <WalletProvider wallets={wallets} autoConnect>
          <WalletModalProvider>
            <div>
              <h1>Solana Token Manager</h1>
              <WalletMultiButton />
              <Page />
            </div>
          </WalletModalProvider>
        </WalletProvider>
      </ConnectionProvider>
  );
};

export default App;
