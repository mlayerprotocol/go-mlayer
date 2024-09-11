package cmd

import (
	"encoding/hex"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
	"github.com/mlayerprotocol/go-mlayer/configs"
	mlCrypto "github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/spf13/cobra"

	"golang.org/x/term"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Manage node account, wallet and registration for goml",
	Long: `Use this command to manage goml key pairs and register node:

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
	nodeCmd.AddCommand(nodeAccountCmd)
	nodeAccountCmd.AddCommand(accountInitCmd)
	nodeAccountCmd.AddCommand(accountImportCmd)
	nodeAccountCmd.AddCommand(accountExportCmd)

	nodeCmd.AddCommand(nodeWalletCmd)
	nodeWalletCmd.AddCommand(walletInitCmd)
	nodeWalletCmd.AddCommand(walletImportCmd)
	nodeWalletCmd.AddCommand(walletExportCmd)

	nodeCmd.AddCommand(nodeLicenseCmd)
	nodeLicenseCmd.AddCommand(licenseRegisterCmd)
	nodeLicenseCmd.AddCommand(licenseListCmd)
	
	rootCmd.AddCommand(nodeCmd)

}

func initKey(keystoreName string, ksPath string) ([]byte, error) {	
	
	storeFilePath := getKeyStoreFilePath(keystoreName, ksPath)

	fmt.Printf("Keystore File Path (%s): %s", keystoreName, storeFilePath)
	fmt.Println()
	fmt.Println()
	_, err := utils.ReadJSONFromFile(storeFilePath)
	if err == nil {
		return nil, fmt.Errorf(formatError("Error: Key store already exists. Delete the existing keystore file to continue"))
	}
	privateKeySECP, err := btcec.NewPrivateKey()
	if err != nil {
		fmt.Println("Error creating keystore directory at %s", err.Error())
		return nil, nil
	}
	return saveKey(privateKeySECP.Serialize(), storeFilePath)
}



func importKey( keystoreName string, ksPath string) ([]byte, error) {
	storeFilePath := getKeyStoreFilePath(keystoreName, ksPath)
	fmt.Printf("Keystore File Path (%s): %s", keystoreName, storeFilePath)
	fmt.Println()
	fmt.Println()
	_, err := utils.ReadJSONFromFile(storeFilePath)
	if err == nil {
		return nil, fmt.Errorf(formatError("Error: Key store already exists. Delete the existing keystore file to continue."))
	}
	fmt.Println("\nEnter private key to be import (hex string format): ")
	privateKeyStringByte, err := readInputSecurely()
	if err != nil {
		fmt.Println(formatError("Error reading new keystore password:"), err)
		return nil, err
	}
	
	privateKeyString := strings.Replace(string(privateKeyStringByte), "0x", "", 1)
	if len(privateKeyString) != 64 {
		return nil, fmt.Errorf("invalid private key entered")
	}
	privKey, err := hex.DecodeString(privateKeyString)
	if err != nil {
		fmt.Println(formatError("Error parsing private key:"), err)
		return nil, err
	}
	return saveKey(privKey, storeFilePath)
}

func exportKey( keystoreName string, ksPath string) ([]byte, error) {
	storeFilePath := getKeyStoreFilePath(keystoreName, ksPath)
	fmt.Printf("Keystore File Path (%s): %s", keystoreName, storeFilePath)
	fmt.Println()
	fmt.Println()
	keyData, err := utils.ReadJSONFromFile(storeFilePath)
	if err != nil {
		return nil, err
	}
	fmt.Println("\nEnter keystore password: ")
	password, err := readInputSecurely()
	if err != nil {
		// fmt.Println(formatError("Error reading new keystore password:"), err)
		return nil, err
	}
	cypher, err := hex.DecodeString(fmt.Sprint(keyData["c"]))
	if err != nil {
		// fmt.Println(formatError("Error reading key:"), err)
		return nil, err
	}
	salt, err := hex.DecodeString(fmt.Sprint(keyData["s"]))
	if err != nil {
		// fmt.Println(formatError("Error reading salt:"), err)
		return nil, err
	}
	k, err := mlCrypto.DecryptPrivateKey(cypher, string(password),  salt)
	if err != nil {
		// fmt.Println(formatError(fmt.Sprint("Error:", "invalid passphrase")))
		return nil, err
	}
	return k, nil
}

func saveKey(privateKey []byte, storeFilePath string) ([]byte, error) {
	fmt.Println("\nEnter key store password: ")
	password, err := readInputSecurely()
	if err != nil {
		fmt.Println("Error reading new keystore password:", err)
		return nil, nil
	}
	fmt.Println("Confirm new key store password: ")
	newPass2, err := readInputSecurely()
	if err != nil {
		fmt.Println("Error reading new keystore password:", err)
		return nil, nil
	}
	if !strings.EqualFold(string(password), string(newPass2)) {
		return nil, fmt.Errorf(formatError("error: passwords don't match!"))
	}
	fmt.Println("Initializing keystore...")
	
	cypher, salt, err := mlCrypto.EncryptPrivateKey(privateKey, string(password))
	if err != nil {
		return nil, fmt.Errorf(formatError(fmt.Sprintf("failed to encrpt private key: %v", err)))
	}
	
	keyData := map[string]interface{}{"s": hex.EncodeToString(salt), "c":hex.EncodeToString(cypher)}
	err = utils.WriteJSONToFile(storeFilePath, keyData)
	fmt.Println("Initializing keystore...", err, storeFilePath)
	if err != nil {
		return nil, err
	}


	fmt.Println("Key store saved successfully! Please backup your password!")
	return privateKey, nil
}

func formatError(message string) string {
	return fmt.Sprintf("\n\033[31m%s\033[0m", message)
}

func readInputSecurely() ([]byte, error) {
	oldState, err := term.GetState(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	// Set up signal catching to handle Ctrl+C and restore terminal settings
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine to handle signal interruption and restore the terminal
	go func() {
		<-sigChan
		term.Restore(int(os.Stdin.Fd()), oldState)
		os.Exit(0)
	}()
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return nil, err
	}

	// Properly restore terminal before exiting
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	return password, nil
}

func getKeyStoreFilePath(keystoreName string, ksPath string) (string) {
	cfg := configs.Config
	if ksPath == ""  {
		ksPath = cfg.KeyStoreDir
	}

	if len(ksPath) == 0 {
		if len(cfg.DataDir) == 0 {
			cfg.DataDir = "./data"
		} else {
			ksPath = cfg.DataDir
		}
	}
	if strings.HasSuffix(ksPath, "/") {
		ksPath = fmt.Sprintf("%skeystores/.goml", cfg.DataDir)
	} else {
		ksPath = fmt.Sprintf("%s/keystores/.goml", cfg.DataDir)
	}
	err := os.MkdirAll(ksPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	
	return fmt.Sprintf("%s/%s.json", ksPath, keystoreName)
}

func loadPrivateKeyFromKeyStore(password string, name string, ksPath string) ([]byte, error) {
	path := filepath.Join(ksPath, ".goml", fmt.Sprintf("%s.json", name))
	if !strings.HasPrefix(path, "./") && !strings.HasPrefix(path, "../") && !filepath.IsAbs(ksPath) {
		path = "./" + path
		if strings.HasPrefix(ksPath, "../") {
			path = "." + path
		}
	}
	store, err := utils.ReadJSONFromFile(path)
	if err != nil {
		return nil, err
	}
	cypher, err := hex.DecodeString(fmt.Sprintf("%s", store["c"]))
	if err != nil {
		return nil, err
	}
	salt, err := hex.DecodeString(fmt.Sprintf("%s", store["s"]))
	if err != nil {
		return nil, err
	}
	decrypt, err := mlCrypto.DecryptPrivateKey(cypher, password, salt )
	if err != nil {
		return nil, fmt.Errorf("error: invalid keystore password")
	}
	return decrypt, nil
}