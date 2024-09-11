package cmd

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/spf13/cobra"
)




var nodeAccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage account for goml",
	Long: `Use this command to manage goml key pairs:

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

func init() {
	accountInitCmd.Flags().StringP(string(KEYSTORE_DIR), "", "", "The keystore directory. This is the directory the keys are stored")
	accountImportCmd.Flags().StringP(string(KEYSTORE_DIR), "", "", "The keystore directory. This is the directory the keys are stored")
}


var accountInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize keystore",
	Long: `Initialize creates a new keystore:

	mLayer (message layer) is an open, decentralized 
	communication network that enables the creation, 
	transmission and termination of data of all sizes, 
	leveraging modern protocols. mLayer is a comprehensive 
	suite of communication protocols designed to evolve with 
	the ever-advancing realm of cryptography. 
	Visit the mLayer [documentation](https://mlayer.gitbook.io/introduction/what-is-mlayer) to learn more
	.`,
	Run: accountInitFunc,
}



var accountImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import private key",
	Long: `Import private key or mnemonic to keystore:

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

var accountExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export node account private key",
	Long: `Export private key of node account:

	mLayer (message layer) is an open, decentralized 
	communication network that enables the creation, 
	transmission and termination of data of all sizes, 
	leveraging modern protocols. mLayer is a comprehensive 
	suite of communication protocols designed to evolve with 
	the ever-advancing realm of cryptography. 
	Visit the mLayer [documentation](https://mlayer.gitbook.io/introduction/what-is-mlayer) to learn more
	.`,
	Run: accountExportFunc,
}



func accountInitFunc(_cmd *cobra.Command, _args []string) {
	dir, _ := _cmd.Flags().GetString(string(KEYSTORE_DIR))
	
	privKeyBytes, err := initKey("account", dir)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	privateKeySECP, pubKeySECP := btcec.PrivKeyFromBytes(privKeyBytes)
	// pubKeySECP := privateKeySECP.PubKey()
	// fmt.Printf("\n - Private Key: 0x%s", hex.EncodeToString(privateKeyBytes))

	fmt.Printf("\n - Licence Operator Public Key (SECP): %s", hex.EncodeToString(pubKeySECP.SerializeCompressed()))
	privateKeyBytesEDD := ed25519.NewKeyFromSeed(privateKeySECP.Serialize())
	publicKeyBytes := privateKeyBytesEDD[32:]
   	fmt.Printf("\n - Network Public Key (EDD): %s", hex.EncodeToString(publicKeyBytes))
	copy(privKeyBytes,[]byte{})
	privateKeySECP = nil
	privateKeyBytesEDD = nil
	fmt.Println("")
}



// Import your private key or mnemonic
func accountImportFunc(_cmd *cobra.Command, _args []string) {
	dir, _ := _cmd.Flags().GetString(string(KEYSTORE_DIR))	
	privKeyBytes, err := importKey("account", dir)
	if err != nil {
		fmt.Println("Error:", err.Error())
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

// Import your private key or mnemonic
func accountExportFunc(_cmd *cobra.Command, _args []string) {
	dir, _ := _cmd.Flags().GetString(string(KEYSTORE_DIR))	
	privKeyBytes, err := exportKey("account", dir)
	if err != nil {
		fmt.Println(formatError(fmt.Sprint("Error:", err.Error())))
		return
	}
	
	fmt.Println("")
	fmt.Println(hex.EncodeToString(privKeyBytes))
	fmt.Println("")
}
