package cmd

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	mlcrypto "github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/spf13/cobra"
)



var nodeWalletCmd = &cobra.Command{
	Use:   "wallet",
	Short: "Manage payer wallet for goml validators",
	Long: `Use this command to manage validator's payer wallet key pairs:

	mLayer (message layer) is an open, decentralized 
	communication network that enables the creation, 
	transmission and termination of data of all sizes, 
	leveraging modern protocols. mLayer is a comprehensive 
	suite of communication protocols designed to evolve with 
	the ever-advancing realm of cryptography. 
	Visit the mLayer [documentation](https://mlayer.gitbook.io/introduction/what-is-mlayer) to learn more
	.`,
	// Run: accountFunc,
}



var walletInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize wallet key store",
	Long: `Initialize creates a new keystore:

	mLayer (message layer) is an open, decentralized 
	communication network that enables the creation, 
	transmission and termination of data of all sizes, 
	leveraging modern protocols. mLayer is a comprehensive 
	suite of communication protocols designed to evolve with 
	the ever-advancing realm of cryptography. 
	Visit the mLayer [documentation](https://mlayer.gitbook.io/introduction/what-is-mlayer) to learn more
	.`,
	Run: walletInitFunc,
}

var walletImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Imports private key",
	Long: `Import private key to  walletkeystore:

	mLayer (message layer) is an open, decentralized 
	communication network that enables the creation, 
	transmission and termination of data of all sizes, 
	leveraging modern protocols. mLayer is a comprehensive 
	suite of communication protocols designed to evolve with 
	the ever-advancing realm of cryptography. 
	Visit the mLayer [documentation](https://mlayer.gitbook.io/introduction/what-is-mlayer) to learn more
	.`,
	Run: accountImportFunc,
}

func init() {
	walletInitCmd.Flags().StringP(string(KEYSTORE_DIR), "", "", "The keystore directory. This is the directory the keys are stored")
	walletImportCmd.Flags().StringP(string(KEYSTORE_DIR), "", "", "The keystore directory. This is the directory the keys are stored")
}

func walletInitFunc(_cmd *cobra.Command, _args []string) {
	dir, _ := _cmd.Flags().GetString(string(KEYSTORE_DIR))
	
	privKeyBytes, err := initKey("wallet", dir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, pubKeySECP := btcec.PrivKeyFromBytes(privKeyBytes)
	pubKeyBytes := pubKeySECP.SerializeUncompressed()
	fmt.Printf("\n - Payer Wallet Private Key: 0x%s", hex.EncodeToString(privKeyBytes))
	privateKeyHash := mlcrypto.Keccak256Hash(pubKeyBytes[1:])
	address := hex.EncodeToString(privateKeyHash[12:])
   	fmt.Printf("\n - Payer Wallet Address: 0x%s", address)
	copy(privKeyBytes,[]byte{})
	fmt.Println("")
}


// Import your private key or mnemonic
func walletImportFunc(_cmd *cobra.Command, _args []string) {
	dir, _ := _cmd.Flags().GetString(string(KEYSTORE_DIR))	
	privKeyBytes, err := importKey("wallet", dir)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	privateKeySECP, pubKeySECP := btcec.PrivKeyFromBytes(privKeyBytes)

	fmt.Printf("\n - Licence Operator Public Key (SECP): %s", hex.EncodeToString(pubKeySECP.SerializeCompressed()))
	privateKeyBytesEDD := ed25519.NewKeyFromSeed(privateKeySECP.Serialize())
	publicKeyBytes := privateKeyBytesEDD[32:]
   	fmt.Printf("\n - Network Public Key (EDD): %s", hex.EncodeToString(publicKeyBytes))

	copy(privKeyBytes,[]byte{})
	privateKeySECP = nil
	privateKeyBytesEDD = nil
	fmt.Println("")
}
