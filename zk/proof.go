package zk

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mlayerprotocol/go-mlayer/pkg/log"

	"github.com/mlayerprotocol/go-mlayer/entities"
)

var logger = &log.Logger

func VerifyJsonProof(proofJsonData []byte, circuit entities.CircuitProps ) (valid bool, publicSignals []string, err error) {
	var snarkjsProof entities.SnarkjsProof
	if err := json.Unmarshal(proofJsonData, &snarkjsProof); err != nil {
		logger.Errorf("Error parsing proof: %v\n", err)
		return false, nil, err
	}
	dir, err := os.Getwd()
	snarkjsVk, err := circuit.GetVerificationKeyJSON(filepath.Join(dir, "zk", "data"))
		// Load verification key
		
		if err != nil {
			fmt.Printf("Error reading verification key: %v\n", err)
			os.Exit(1)
		}
	
	valid, err = Groth16Verify(*snarkjsVk, snarkjsProof.PI, snarkjsProof )
	if err != nil {
		logger.Errorf("Failed to verify proof: %v", err)
		return false, nil, err
	}

	if !valid {
		logger.Errorf("Failed to verify proof: %v", err)
		return false, nil, err
	}
	return true, snarkjsProof.PI, nil
}