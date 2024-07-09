package cmd

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/localcrypto"
	"github.com/mlayerprotocol/go-mlayer/localio"
	"github.com/mlayerprotocol/go-mlayer/localtea"
	"github.com/spf13/cobra"
)

// var teaProgram *tea.Program

var wizardCmd = &cobra.Command{
	Use:   "mlwiz",
	Short: "Runs goml as a Wizard",
	Long: `Use this command to run goml as a Wizard:

	mLayer (message layer) is an open, decentralized 
	communication network that enables the creation, 
	transmission and termination of data of all sizes, 
	leveraging modern protocols. mLayer is a comprehensive 
	suite of communication protocols designed to evolve with 
	the ever-advancing realm of cryptography. 
	Visit the mLayer [documentation](https://mlayer.gitbook.io/introduction/what-is-mlayer) to learn more
	.`,
	Run: wizardFunc,
}

func init() {
	rootCmd.AddCommand(wizardCmd)
	wizardCmd.Flags().StringP(string(NETWORK_PRIVATE_KEY), "e", "", "The network private key. This is the key used to sign handshakes and messages")
}

func wizardFunc(_cmd *cobra.Command, _args []string) {
	//

	p := tea.NewProgram(localtea.InitialOptionModel([]string{"Existing private key", "Generate private key"}))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
	if localtea.GlobalWizardModel.Cursor == 0 {
		p := tea.NewProgram(localtea.InitialOptionModel([]string{"From Memory", "New Key"}))
		if _, err := p.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}
		if localtea.GlobalWizardModel.Cursor == 0 {
			key, err := localio.ReadFromFile(constants.DefaultPKeyPath)
			if err != nil {
				fmt.Println("Error running script:", err)
				os.Exit(1)
			}
			scriptPath := "./scripts/dev.sh" // Replace with your script path
			err = runScript(scriptPath, key)
			if err != nil {
				fmt.Println("Error running script:", err)
				os.Exit(1)
			} else {
				fmt.Println("Script ran successfully!")
			}
		}
		if localtea.GlobalWizardModel.Cursor == 1 {
			//
			p := tea.NewProgram(localtea.InitialInputModel())
			if _, err := p.Run(); err != nil {
				fmt.Printf("Alas, there's been an error: %v", err)
				os.Exit(1)
			}
			key := localtea.GlobalInputWizardModel.TextInput.Value()
			scriptPath := "./scripts/dev.sh" // Replace with your script path
			err := runScript(scriptPath, key)
			if err != nil {
				fmt.Println("Error running script:", err)
				os.Exit(1)
			} else {
				fmt.Println("Script ran successfully!")
			}
		}

	} else if localtea.GlobalWizardModel.Cursor == 1 {
		fmt.Printf("globalWizardModel.cursor %v:", localtea.GlobalWizardModel.Cursor)
		publicKey, privateKey, err := localcrypto.GenerateKeyPair()
		if err != nil {
			fmt.Println("Error generating key pair:", err)
			return
		}
		decodedPublicKey := localcrypto.KeyToString(publicKey)

		decodedPrivateKey := localcrypto.KeyToString(privateKey)

		teaProgram := tea.NewProgram(localtea.InitialOptionModel([]string{fmt.Sprintf("Public Key: %v", decodedPublicKey), fmt.Sprintf("Private Key: %v", decodedPrivateKey), "Continue"}))
		if _, err := teaProgram.Run(); err != nil {
			fmt.Printf("Alas, there's been an error: %v", err)
			os.Exit(1)
		}

		if localtea.GlobalWizardModel.Cursor == 2 {
			scriptPath := "./scripts/dev.sh" // Replace with your script path
			err := runScript(scriptPath, decodedPrivateKey)
			if err != nil {
				fmt.Println("Error running script:", err)
				os.Exit(1)
			} else {
				fmt.Println("Script ran successfully!")
			}

		}
	}

}

func runScript(scriptPath string, key string) error {
	err := localio.CreateOrUpdateFile(constants.DefaultPKeyPath, key)
	if err != nil {
		fmt.Println("Error running script:", err)
		os.Exit(1)
	}
	cmd := exec.Command("/bin/sh", scriptPath, key)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
