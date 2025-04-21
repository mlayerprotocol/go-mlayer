package circuits

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"
	"testing"

	// _ "github.com/consensys/gnark-crypto/ecc/bls12-377/fr/poseidon2"
	_ "github.com/consensys/gnark-crypto/ecc/bn254/fr/mimc"

	hashes "github.com/consensys/gnark-crypto/hash"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"go.uber.org/zap/buffer"
)

type SignatureResult struct {
	Address   common.Address   `json:"address"`
	ChainID   string           `json:"chain_id"`
	Message   string           `json:"message"`
	R         string           `json:"r"`
	S         string           `json:"s"`
	V         uint8            `json:"v"`
	PublicKey ecdsa.PublicKey `json:"public_key"`
}

// SignMessage generates an Ethereum-compatible signature
func SignMessage(privateKeyHex, chainID string) (*SignatureResult, error) {
	// Convert private key from hex
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	// Get public key and address
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKey)

	// Define the message
	message := fmt.Sprintf("Network: mlayer/%s", chainID)

	// Ethereum-specific message prefix
	 prefixedMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len([]byte(message)), message)

	// Hash the message
	messageHash := crypto.Keccak256Hash([]byte(prefixedMessage))
	logger.Infof("MESSAGHAHS: %s", hex.EncodeToString(messageHash.Bytes()))
	// Sign the hash
	//d, _ := hex.DecodeString(privateKeyHex)
	// privateKeyCec, _ := btcec.PrivKeyFromBytes(btcec.S256(), d)
	// signatureC, err := privateKeyCec.Sign(messageHash[:])
	// if err != nil {
	// 	logger.Fatal(err)
	// }
	// signature := signatureC.Serialize()
	// return b, hex.EncodeToString(b)
	 signature, err := crypto.Sign(messageHash.Bytes(), privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to sign message: %v", err)
	}

	// Extract r, s, and v from the signature
	r, s, v := signature[:32], signature[32:64], signature[64]
	if v == 0 {
		v = 27
	} else {
		v = 26
	}
	// Return struct with all details
	return &SignatureResult{
		Address:   address,
		ChainID:   chainID,
		Message:   message,
		R:         fmt.Sprintf("%x", r),
		S:         fmt.Sprintf("%x", s),
		V:         v,
		PublicKey: *publicKey,
	}, nil
}

func TestVerifyEthereumSignature(t *testing.T) {
	// Load keys (from TestSetupCircuit)
	return
	privateKey := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	chainID := ChainId("1")

	signature, err := SignMessage(privateKey, string(chainID))
	if err != nil {
		t.Fatalf("Failed to sign: %v", err)
	}
	fmt.Println("")
	fmt.Printf("%+v", signature)
	fmt.Println("")

	sigR, v := new(big.Int).SetString(signature.R[:32], 16)
	if !v {
		t.Fatalf("Failed to parse r: %v", err)
	}
	sigS, v := new(big.Int).SetString(signature.S[:32], 16)
	if !v {
		t.Fatalf("Failed to parse r: %v", err)
	}
	sigV := signature.V
	pubKeyX := signature.PublicKey.X.Bytes()[:16] // Replace with actual X
	pubKeyY := signature.PublicKey.Y.Bytes()[:16] // Replace with actual Y

	nonce := 1
	// mimcHasher := hashes.Hash.New(hashes.Hash(hashes.MIMC_BLS12_381))
	mimcHasher := hashes.Hash.New(hashes.Hash(hashes.MIMC_BN254))
	
	if err != nil {
		panic(err)
	}
	

	fmt.Println("SIGNRA", sigV)
	
 	 mimcHasher.Write([]byte(chainID.Bytes()))
	 mimcHasher.Write(pubKeyX)
	mimcHasher.Write(pubKeyY)
	 mimcHasher.Write(sigR.Bytes())
	 mimcHasher.Write(sigS.Bytes())
	 mimcHasher.Write([]byte{byte(sigV)})
 	mimcHasher.Write([]byte{byte(nonce)})

	// Get the hash output
	idCommitment := mimcHasher.Sum(nil)

	// sigR.Mod(sigR, gnarkCurve.Params().Fr)
	 fmt.Println("FINALHASH", len(idCommitment),hex.EncodeToString(idCommitment))
	// Witness assignment


	meta := []byte("{}")
	categories := []byte("[]")
	ref := []byte("com.test.app")
	status := new(big.Int).SetUint64(uint64(1)).Bytes()
	timestamp := new(big.Int).SetUint64(uint64(1741823784277)).Bytes()
	defaultAuthPriv := new(big.Int).SetUint64(uint64(10)).Bytes()

	mimcHasher.Reset()
	mimcHasher.Write(idCommitment)
	mimcHasher.Write(meta)
	mimcHasher.Write(ref)
	mimcHasher.Write(categories)
	mimcHasher.Write(status)
	 mimcHasher.Write(defaultAuthPriv)
 	mimcHasher.Write(timestamp)

	apCommitment := mimcHasher.Sum(nil)

	fmt.Println("APPCOMITEMNT", hex.EncodeToString(apCommitment))
	field := CURVE.ScalarField()
	pk := groth16.NewProvingKey(CURVE)
	vk := groth16.NewVerifyingKey(CURVE)
	pkFile, err := os.Open(AppCircuitProps.GetProvingKeyPath("../test_data"))
	defer pkFile.Close()
	if err != nil {
		t.Fatalf("Failed reading proving key: %v", err)
	}

	_, err = pk.ReadFrom(pkFile)
	if err != nil {
		t.Fatalf("Failed to read proving key: %v", err)
	}

	// vk := groth16.NewVerifyingKey(CURVE)
	
	// vk, err = ConvertVerificationKey(AppCircuitProps.GetVerificationKeyPath("../test_data"))
	 vkFile, err := os.Open(AppCircuitProps.GetVerificationKeyPath("../test_data"))
	// defer vkFile.Close()
	// _, err := groth16.ReadVerifyingKey(ecc.BN254, "data/circuit_verification_key.vk")
	if err != nil {
		t.Fatalf("Failed to read verification key: %v", err)
	}

	_, err = vk.ReadFrom(vkFile)
	if err != nil {
		t.Fatalf("Failed to read verification key: %v", err)
	}

	// Test data from ethers.js (example values - replace with your output)

	

	witnessData := AppCircuit{
		// SignatureType:  []byte(EthereumSignatureType),
		SigR: sigR.Bytes(),
		SigS: sigS.Bytes(),
		SigV: sigV,
		// // SigV not used in circuit, but included for completeness
		Nonce: []byte{byte(nonce)}, // Example value

		// // CreatorSigR:    0, // Placeholder
		// // CreatorSigS:    0,
		// ModSigR:        0,
		// ModSigS:        0,

		CreatorPubKeyX: pubKeyX,
		CreatorPubKeyY: pubKeyY,
		Public: AppCircuitPublic{
			ID:                   idCommitment,
			AppCommitment:        apCommitment,
			ChainID:              chainID.Bytes(),
			// Name:                 name,
			Meta:                 meta,
			Ref:                  ref,
			Categories:           categories,
			Status:               status,
			DefaultAuthPrivilege: defaultAuthPriv,
			Timestamp:            timestamp,
		},

		// // Public inputs (placeholders)
		// CreatorCommitment: 0,

		// NewName:           0,
	}

	// Compile circuit for witness

	witness, err := frontend.NewWitness(&witnessData, field)
	if err != nil {
		fmt.Printf("Error creating witness: %v\n", err)
		return
	}
	circuit := &AppCircuit{}
	r1cs, err := frontend.Compile(field, r1cs.NewBuilder, circuit)
	if err != nil {
		t.Fatalf("Failed to compile witness: %v", err)
	}

	// Generate proof
	proof, err := groth16.Prove(r1cs, pk, witness)
	if err != nil {
		t.Fatalf("Failed to generate proof: %v", err)
	}
	proofBuffer := buffer.Buffer{}
	proof.WriteTo(&proofBuffer)
	publicWitness, err := witness.Public()
	if err != nil {
		t.Fatalf("Error extracting public inputs: %v\n", err)
		return
	}

	fmt.Printf("PROOF %+v\n", publicWitness.Vector())

	// Verify proof
	err = groth16.Verify(proof, vk, publicWitness)
	if err != nil {
		t.Fatalf("Failed to verify proof: %v", err)
	}

	t.Log("Proof verified successfully!")
}
