package bls

import (
	"fmt"
	"log"
	"testing"

	"github.com/mlayerprotocol/go-mlayer/common/utils"
)


func TestAggregateKeyProof(t *testing.T) {
	
	// Create a new BLS proof generator
	generator := NewBlsProofGenerator()

	// Example usage
	// Generate some key pairs
	numSigners := 3
	privateKeys := make([][]byte, numSigners)
	publicKeys := make([][]byte, numSigners)

	for i := 0; i < numSigners; i++ {
		privKey, pubKey, err := generator.BLSKeyPairFromSeed([]byte(utils.RandomString(16)))
		if err != nil {
			log.Fatalf("Failed to generate key pair %d: %v", i, err)
		}
		privateKeys[i] = privKey
		publicKeys[i] = pubKey
		fmt.Printf("Generated key pair %d, %x\n", i, pubKey)
	}

	// Create some messages
	messages := make([][]byte, numSigners)
	for i := 0; i < numSigners; i++ {
		messages[i] = []byte(fmt.Sprintf("Message from signer"))
	}

	// Sign the messages
	signatures := make([][]byte, numSigners)
	for i := 0; i < numSigners; i++ {
		sig, err := generator.Sign(privateKeys[i], messages[i])
		if err != nil {
			log.Fatalf("Failed to sign message %d: %v", i, err)
		}
		signatures[i] = sig
		fmt.Printf("Signed message %d\n", i)
	}

	// Verify individual signatures
	for i := 0; i < numSigners; i++ {
		valid, err := generator.VerifySignature(publicKeys[i], messages[i], signatures[i])
		if err != nil {
			log.Fatalf("Failed to verify signature %d: %v", i, err)
		}
		if !valid {
			log.Fatalf("Signle Signature %d verification failed", i)
		}
		fmt.Printf("Verified signature %d\n", i)
	}

	// Aggregate the signatures
	aggregateSignature, err := generator.AggregateSignatures(signatures)
	if err != nil {
		t.Fatalf("Failed to aggregate signatures: %v", err)
	}
	fmt.Println("Aggregated signatures")

	// Verify the aggregate signature
	valid, err := generator.VerifyAggregateSignature(publicKeys, messages, aggregateSignature)
	if err != nil {
		t.Fatalf("Failed to verify aggregate signature: %v", err)
	}
	if !valid {
		t.Fatalf("Aggregate signature verification failed")
	}
	fmt.Println("Verified aggregate signature")

	// Prepare the proof for the Solidity verifier
	sigHex, pubKeysHex, messagesHex, err := generator.PrepareProofForSolidity(
		aggregateSignature, publicKeys, messages)
	if err != nil {
		t.Fatalf("Failed to prepare proof for Solidity: %v", err)
	}

	// Print the data for use in the Solidity verifier
	fmt.Println("\nSolidity Verifier Input:")
	fmt.Printf("Aggregate signature: %s\n", sigHex)
	fmt.Println("Public keys:")
	for i, pk := range pubKeysHex {
		fmt.Printf("  %d: %s\n", i, pk)
	}
	fmt.Println("Messages:")
	for i, msg := range messagesHex {
		fmt.Printf("  %d: %s\n", i, msg)
	}


	
}