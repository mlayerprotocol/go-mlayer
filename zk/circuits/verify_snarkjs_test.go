package circuits

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/mlayerprotocol/go-mlayer/entities"
	// _ "github.com/consensys/gnark-crypto/ecc/bls12-377/fr/poseidon2"
)



func TestVerifySnarkjsProof(t *testing.T) {
	// Load keys (from TestSetupCircuit)

	// field := CURVE.ScalarField()

	 vkFile := entities.ClientPayloadCircuitProps.GetVerificationKeyPath("../test_data")
	

	// Load verification key
	vkData, err := os.ReadFile(vkFile)
	if err != nil {
		fmt.Printf("Error reading verification key: %v\n", err)
		os.Exit(1)
	}
	// logger.Infof("%v", vkData)
	var snarkjsVk entities.SnarkjsVerificationKey
	if err := json.Unmarshal(vkData, &snarkjsVk); err != nil {
		fmt.Printf("Error parsing verification key: %v\n", err)
		os.Exit(1)
	}

	// vk, err := ConvertVerificationKeyBN254(snarkjsVk)
	// if err != nil {
	// 	fmt.Printf("Error converting verification key: %v\n", err)
	// 	os.Exit(1)
	// }

	

	// Load proof
	proofFile := "../test_data/84532/app/v1/proof.json"
	proofData, err := os.ReadFile(proofFile)
	if err != nil {
		fmt.Printf("Error reading proof: %v\n", err)
		os.Exit(1)
	}

	var snarkjsProof entities.SnarkjsProof
	if err := json.Unmarshal(proofData, &snarkjsProof); err != nil {
		fmt.Printf("Error parsing proof: %v\n", err)
		os.Exit(1)
	}

	
	// proof, err := ConvertProofBn254(snarkjsProof)
	// if err != nil {
	// 	fmt.Printf("Error converting proof: %v\n", err)
	// 	os.Exit(1)
	// }


	// publicInputs, err := ConvertPublicInputs(snarkjsProof, snarkjsVk.NPublic)
	// if err != nil {
	// 	fmt.Printf("Error converting public inputs: %v\n", err)
	// 	os.Exit(1)
	// }
	
	// logger.Infof("PUBLICINPUTS %+v", proof)
	// Verify proof
	// err = groth16.Verify(&proof, vk, publicInputs)
	valid, err := Groth16Verify(snarkjsVk, snarkjsProof.PI, snarkjsProof )
	if err != nil {
		t.Fatalf("Failed to verify proof: %v", err)
	}

	if !valid {
		t.Fatalf("Failed to verify proof: %v", valid)
	}
	t.Log("Proof verified successfully!")
}
