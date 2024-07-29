package schnorr

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"

	btcec "github.com/btcsuite/btcd/btcec/v2"
	crypto "github.com/mlayerprotocol/go-mlayer/internal/crypto"
)
const ADDRESS_MASK = "000000000000000000000000FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF";



type EthAddress []byte

func toAddress(point *btcec.PublicKey) EthAddress {
	// Serialize the X and Y coordinates of the point.
	// serialized := append(point.X().Bytes(), point.Y().Bytes()...)

	// // Hash the serialized coordinates using Keccak-256.
	// addressBytes := crypto.Keccak256Hash(serialized)
	// Take the last 20 bytes as the address and convert to string.
	// address := fmt.Sprintf("0x%x", addressBytes[len(addressBytes)-20:])
	// return addressBytes[len(addressBytes)-20:]
	xBytes := point.X().Bytes()
	yBytes := point.Y().Bytes()

	// Concatenate x and y bytes
	data := append(xBytes, yBytes...)
	// data = append(data, 0x40)
	// Hash the concatenated bytes
	hash := crypto.Keccak256Hash(data)

	// Convert hash to a big integer
	hashBigInt := new(big.Int).SetBytes(hash)

	// Apply address mask
	mask, _ := hex.DecodeString(ADDRESS_MASK)
	addressBigInt := new(big.Int).And(hashBigInt, new(big.Int).SetBytes(mask))

	// Get the last 20 bytes (40 hex characters)
	addressBytes := addressBigInt.Bytes()
	// address := fmt.Sprintf("%x", addressBytes[len(addressBytes)-20:])

	return addressBytes[len(addressBytes)-20:]
}



// 1. Function to derive a public key from a private key.
func derivePublicKey(privKey *btcec.PrivateKey) *btcec.PublicKey {
	return privKey.PubKey()
}
func NumberFromByte(buf []byte) uint64 {
	return (binary.BigEndian.Uint64(buf))
}

// 2. Function to aggregate public keys.
func aggregatePublicKeys(pubKeys []*btcec.PublicKey) *btcec.PublicKey {
	if len(pubKeys) == 0 {
		return nil
	}

	// Initialize the aggregated Jacobian point with the first public key in affine form
	x := &btcec.FieldVal{}
	y := &btcec.FieldVal{}
	z := &btcec.FieldVal{}
	x.SetByteSlice(pubKeys[0].X().Bytes())
	y.SetByteSlice(pubKeys[0].Y().Bytes())
	z.SetByteSlice(big.NewInt(1).Bytes())
	aggJacobian := &JacobianPoint{
		X: x,
		Y: y,
		Z: z, // Z coordinate is initialized to 1
	}

	for i := 1; i < len(pubKeys); i++ {
		// aggJacobian = AddAffinePoint(aggJacobian, &Point{x: pubKeys[i].X(), y: pubKeys[i].Y() })
		aggJacobian.AddAffinePoint(pubKeys[i])
	}
	fmt.Printf("JACOBIAN_AFTER X: %s, %s, %s\n", new(big.Int).SetBytes((aggJacobian.X.Bytes())[:]),  new(big.Int).SetBytes((aggJacobian.Y.Bytes())[:]),  new(big.Int).SetBytes((aggJacobian.Z.Bytes())[:]))
	
	return aggJacobian.toAffine()
}

// 3. Function to derive a nonce from a private key and a message.
func deriveNonce(privKey *btcec.PrivateKey, message [32]byte) *big.Int {
	// Concatenate the private key and message.
	data := append(privKey.Serialize(), message[:]...)

	// Calculate the SHA-256 hash of the concatenated data.
	hash := crypto.Keccak256Hash(data)

	// Convert the hash to a big integer.
	nonce := new(big.Int).SetBytes(hash[:])

	// Modulo operation with the curve's order (Q).
	new(big.Int).Mod(nonce, btcec.S256().Params().N)

	return nonce
}

// 4. Function to compute a nonce public key.
func computeNoncePublicKey(nonce *big.Int) *btcec.PublicKey {
	// Placeholder for computing nonce public key.
	_, pubK := btcec.PrivKeyFromBytes(nonce.Bytes())
	return pubK // Example: Return nil for demonstration.
}

// 5. Function to derive a commitment from an aggregated nonce public key.
func deriveCommitment(aggNoncePubKey *btcec.PublicKey) EthAddress {
	// Placeholder for deriving commitment.
	return toAddress(aggNoncePubKey)
}

// 6. Function to construct a challenge.
func constructChallenge(aggPubKey *btcec.PublicKey, message [32]byte, commitment []byte) []byte {
	// Convert the X coordinate of the public key to bytes.
	pubKeyXBytes := aggPubKey.X().Bytes()

	// Determine the parity of Y coordinate (odd or even).
	
	yParity := new(big.Int).And(aggPubKey.Y(), big.NewInt(1))
	
	b := []byte{}
	b = append(b, pubKeyXBytes...)
	 b = append(b, uint8(yParity.Uint64()))
	b = append(b, message[:]...)
	b = append(b, commitment...)
	
	challengeBytes := crypto.Keccak256Hash(b)
	challengeNum := new(big.Int).SetBytes(challengeBytes)

	// Modulo operation with the curve's order (Q).
	
	challengeNum.Mod(challengeNum, btcec.S256().Params().N).Bytes()
	
	return challengeNum.Bytes()
}

// 7. Function to compute a signature.

func ComputeSignature(privKey *btcec.PrivateKey, nonce *big.Int, challenge []byte) []byte {
	// Convert the challenge to a big integer.
	challengeInt := new(big.Int).SetBytes(challenge)
	pk := new(big.Int).SetBytes(privKey.Serialize())
	
	return addmod(nonce, mulmod(challengeInt, pk,  btcec.S256().Params().N), btcec.S256().Params().N).Bytes()
}

func ComputeSignatureMulti(privKeys []*btcec.PrivateKey, message [32]byte, challenge []byte) (signatures [][]byte) {
	// var aggSignature = big.NewInt(0)
	for _, pk := range privKeys {
		nonce := deriveNonce(pk, message);
		sig := ComputeSignature(pk, nonce, challenge)
		signatures = append(signatures, sig)	
		//aggSignature = addmod(aggSignature, new(big.Int).SetBytes(sig), btcec.S256().Params().N);
	}
	return signatures
}
func ComputeNonce(privKey *btcec.PrivateKey, message [32]byte) (nonce *big.Int, noncePubKey *btcec.PublicKey) {
	// var aggSignature = big.NewInt(0)
	nonce = deriveNonce(privKey, message);
	noncePubKey = computeNoncePublicKey(nonce)
	return nonce, noncePubKey
}



func ComputeSigningParams(pubKeys []*btcec.PublicKey, noncePubKeys []*btcec.PublicKey,  message [32]byte ) (aggPubKey *btcec.PublicKey, challenge []byte, commitment EthAddress)  {
	aggNoncePubKey := aggregatePublicKeys(noncePubKeys);
	commitment = deriveCommitment(aggNoncePubKey);
	// address := "0x"+hex.EncodeToString(commitment)
	aggPubKey = aggregatePublicKeys(pubKeys)
	challenge = constructChallenge(aggPubKey, message, commitment)
	return aggPubKey, challenge, commitment
}

func AggregateSignatures(signatures [][]byte) (aggSig []byte) {
	var aggSignature = big.NewInt(0)
	for _, sig := range signatures {
		aggSignature = addmod(aggSignature, new(big.Int).SetBytes(sig), btcec.S256().Params().N);
	}
	aggSig = aggSignature.Bytes()
	return aggSig
}


func SignSingle(privKeyBytes []byte, message [32]byte) (signature []byte, commitment EthAddress, nonce *big.Int, challenge []byte) {
	// Convert the X coordinate of the public key to bytes.
	privKey, pubKey := btcec.PrivKeyFromBytes(privKeyBytes)
	nonce = deriveNonce(privKey, message)
	noncePubKey := computeNoncePublicKey(nonce)
	commitment = deriveCommitment(noncePubKey)
	challenge = constructChallenge(pubKey, message, commitment)
	return ComputeSignature(privKey, nonce, challenge), commitment, nonce, challenge
}