package bls

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	bls "github.com/kilic/bls12-381"
)

var BlsProofGenerator *blsProover

// blsProover handles the creation and aggregation of BLS signatures
type blsProover struct {
	// BLS objects
	g1   *bls.G1
	g2   *bls.G2
	pair *bls.Engine
}


// NewBlsProofGenerator creates a new instance of the blsProover
func NewBlsProofGenerator() *blsProover {
	return &blsProover{
		g1:   bls.NewG1(),
		g2:   bls.NewG2(),
		pair: bls.NewEngine(),
	}
}


func  NewBlsProofGeneratorFromPrivateKey(blsPrivateKey []byte) (*blsProover) {
	gen := NewBlsProofGenerator()
	pubKey := gen.g2.One()
	secret := bls.NewFr().FromBytes(blsPrivateKey)
	// Generate public key in G2
	gen.g2.MulScalar(pubKey, pubKey, secret)
	return gen
}



func (b *blsProover) BLSKeyPairFromSeed(seed []byte) ([]byte, []byte, error) {
	paddedBytes := make([]byte, 32)
	start := 0
	if len(seed) < 32 {
		start = 32 - len(seed)
	}
	copy(paddedBytes[start:], seed)
	secret := bls.NewFr().FromBytes(seed[:])
	
	secretBytes := secret.ToBytes()

	// Generate public key in G2
	pubKey := b.g2.One()
	b.g2.MulScalar(pubKey, pubKey, secret)
	pubKeyBytes := b.g2.ToCompressed(pubKey)
	return secretBytes, pubKeyBytes, nil
}


// Sign generates a BLS signature for a message using a private key
func (b *blsProover) Sign(privateKey []byte, message []byte) ([]byte, error) {
	// Hash the message to a point on G1
	msgPoint, err := b.hashToG1(message)
	if err != nil {
		return nil, fmt.Errorf("failed to hash message to G1: %v", err)
	}
	// Load private key
	secret := bls.NewFr().FromBytes(privateKey)
	// Sign the message: signature = msgPoint^secret
	signature := b.g1.New()
	b.g1.MulScalar(signature, msgPoint, secret)
	// Return compressed signature
	return b.g1.ToCompressed(signature), nil
}

// AggregateSignatures combines multiple BLS signatures into a single aggregate signature
func (b *blsProover) AggregateSignatures(signatures [][]byte) ([]byte, error) {
	if len(signatures) == 0 {
		return nil, fmt.Errorf("no signatures provided for aggregation")
	}
	// Start with the first signature
	//aggregate := b.g1.New()
	firstSig, err := b.g1.FromCompressed(signatures[0])
	if err != nil {
		return nil, fmt.Errorf("failed to decompress first signature: %v", err)
	}
	// b.g1.Copy(aggregate, firstSig)
	aggregate := b.g1.New().Set(firstSig)

	// Add the remaining signatures
	for i := 1; i < len(signatures); i++ {
		sig, err := b.g1.FromCompressed(signatures[i])
		if err != nil {
			return nil, fmt.Errorf("failed to decompress signature %d: %v", i, err)
		}
		b.g1.Add(aggregate, aggregate, sig)
	}
	// Return the compressed aggregate signature
	return b.g1.ToCompressed(aggregate), nil
}

// PrepareProofForSolidity formats the BLS proof for use in the Solidity verifier
// Returns the signature in compressed form and public keys in compressed form
func (b *blsProover) PrepareProofForSolidity(
	aggSignature []byte,
	publicKeys [][]byte,
	messages [][]byte,
) (string, []string, []string, error) {
	// Convert aggregate signature to hex string
	sigHex := "0x" + hex.EncodeToString(aggSignature)
	// Convert public keys to hex strings
	pubKeysHex := make([]string, len(publicKeys))
	for i, pk := range publicKeys {
		pubKeysHex[i] = "0x" + hex.EncodeToString(pk)
	}
	// Convert messages to hex strings
	messagesHex := make([]string, len(messages))
	for i, msg := range messages {
		messagesHex[i] = "0x" + hex.EncodeToString(msg)
	}
	return sigHex, pubKeysHex, messagesHex, nil
}

// hashToG1 hashes a message to a point on the G1 curve
func (b *blsProover) hashToG1(message []byte) (*bls.PointG1, error) {
	// Use Ethereum's keccak256 for hashing
	hash := crypto.Keccak256(message)
	// Map the hash to G1 (try-and-increment method)
	point, err := b.g1.HashToCurve(hash, []byte("BLS_SIG_BLS12381G1_XMD:SHA-256_SSWU_RO_NUL_"))
	if err != nil {
		return nil, err
	}
	return point, nil
}

// VerifySignature verifies a single signature against a public key and message
// VerifySignature verifies a single signature against a public key and message
func (b *blsProover) VerifySignature(publicKey []byte, message []byte, signature []byte) (bool, error) {
	// Get the message point in G1
	msgPoint, err := b.hashToG1(message)
	if err != nil {
		return false, fmt.Errorf("failed to hash message to G1: %v", err)
	}

	// Decompress the public key
	pubKey, err := b.g2.FromCompressed(publicKey)
	if err != nil {
		return false, fmt.Errorf("failed to decompress public key: %v", err)
	}

	// Decompress the signature
	sig, err := b.g1.FromCompressed(signature)
	if err != nil {
		return false, fmt.Errorf("failed to decompress signature: %v", err)
	}

	// Reset the pairing engine
	b.pair.Reset()
	
	// For valid signature: e(sig, g2) == e(hash(m), pubKey)
	// Equivalent to checking: e(sig, g2) * e(-hash(m), pubKey) == 1
	
	// Add positive pair: e(sig, g2)
	b.pair.AddPair(sig, b.g2.One())
	
	// Add negative pair: e(-hash(m), pubKey)
	// Create negative of message point
	negMsgPoint := b.g1.New().Set(msgPoint)
	b.g1.Neg(negMsgPoint, negMsgPoint)
	
	// Add the negative pair
	b.pair.AddPair(negMsgPoint, pubKey)
	
	// Check if the pairing equals 1 (valid signature)
	return b.pair.Check(), nil
}

// VerifyAggregateSignature verifies an aggregate signature against multiple public keys and messages
func (b *blsProover) VerifyAggregateSignature(
	publicKeys [][]byte,
	message []byte,
	aggregateSignature []byte,
) (bool, error) {
	
	// Decompress the aggregate signature
	aggSig, err := b.g1.FromCompressed(aggregateSignature)
	if err != nil {
		return false, fmt.Errorf("failed to decompress aggregate signature: %v", err)
	}

	// Prepare for the pairing check
	b.pair.Reset()
	b.pair.AddPairInv(aggSig, b.g2.One())
	msgPoint, err := b.hashToG1(message)
	if err != nil {
		return false, fmt.Errorf("failed to hash message %x to G1: ", message)
	}
	// Add each (message, public key) pair to the pairing
	for i, pkBytes := range publicKeys {
		// Hash the message to G1
	
		// Decompress the public key
		pubKey, err := b.g2.FromCompressed(pkBytes)
		if err != nil {
			return false, fmt.Errorf("failed to decompress public key %d: %v", i, err)
		}

		// Add to the pairing
		b.pair.AddPair(msgPoint, pubKey)
	}

	// Check if the pairing is valid
	return b.pair.Check(), nil
}
