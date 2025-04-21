package circuits

import (
	"math/big"
	"strconv"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/hash/mimc"
)

// const CURVE = ecc.BLS12_381

 const CURVE = ecc.BN254

const (
	EthereumSignatureType string = "eth/keccak"
	CosmosAmino           string = "amino/sha256"
)

type ChainId string

func (n *ChainId) Bytes() []byte {
	s := string(*n)
	number, err := strconv.Atoi(string(s))
	if err == nil {
		return big.NewInt(int64(number)).FillBytes(make([]byte, 32))
	}
	return []byte(s)
}
func NewAppCircuit(chainId string) AppCircuit {
	return AppCircuit{}
}

type AppCircuitPublic struct {
	//SignatureType frontend.Variable `gnark:"public"`
	ChainID              frontend.Variable `gnark:",public"`
	ID                   frontend.Variable `gnark:",public"`
	// Name                 frontend.Variable `gnark:",public"`
	Meta                 frontend.Variable `gnark:",public"`
	Ref                  frontend.Variable `gnark:",public"`
	Categories           frontend.Variable `gnark:",public"`
	Status               frontend.Variable `gnark:",public"`
	Timestamp            frontend.Variable `gnark:",public"`
	DefaultAuthPrivilege frontend.Variable `gnark:",public"`
	AppCommitment        frontend.Variable `gnark:",public"`
}

type    AppCircuit struct {
	// Private inputs

	SigR  frontend.Variable `gnark:",secret"`
	SigS  frontend.Variable `gnark:",secret"`
	SigV  frontend.Variable `gnark:",secret"`
	Nonce frontend.Variable `gnark:",secret"`

	// ModSigR        frontend.Variable
	// ModSigS        frontend.Variable
	// ModSigV        frontend.Variable

	// Public inputs
	CreatorPubKeyX frontend.Variable `gnark:",secret"`
	CreatorPubKeyY frontend.Variable `gnark:",secret"`
	// CreatorCommitment frontend.Variable `gnark:"public"`
	Public AppCircuitPublic `gnark:"public"`
}

func (circuit *AppCircuit) Define(api frontend.API) error {
	// 1. Verify signature over "welcome to mlayer"
	// inputChain := []byte(fmt.Sprint(circuit.ChainID))
	pub := circuit.Public
	chId := ChainId("1")
	api.AssertIsEqual(circuit.Public.ChainID, chId.Bytes())
	//  api.AssertIsEqual(circuit.ChainID, []byte("1"))

	// field, err := emulated.NewField[emulated.BN254Fr](api)
	// if err != nil {
	// 	return err
	// }

	// toFieldElement := func(v frontend.Variable) *emulated.Element[emulated.BN254Fr] {
	// 	// First constraint the variable to be within the field
	// 	api.AssertIsLessOrEqual(v, api.Sub(field.Modulus(), 1))
	// 	return field.NewElement(v)
	// }

	// msgStr := fmt.Sprintf("Network: mlayer/%s", circuit.ChainID)
	// sigType := circuit.SignatureType
	// log.Println("HELLO ",  sigType, field.NewElement(new(big.Int).SetBytes([]byte(EthereumSignatureType))))
	// switch(sigType ) {
	// 	case frontend.Variable([]byte(EthereumSignatureType)):
	// 		log.Println("HELOOOOS1")
	// prefixStr := "\x19Ethereum Signed Message:\n" + fmt.Sprint(len(msgStr))
	// prefix := make([]uints.U8, len(prefixStr))
	// for i, b := range []byte(prefixStr) {
	// 	prefix[i] = uints.U8{Val: frontend.Variable(b)}
	// }
	// signedMessage := make([]uints.U8, len(msgStr))
	// for i, b := range []byte(msgStr) {
	// 	signedMessage[i] = uints.U8{Val: frontend.Variable(b)}
	// }
	commitHash, err := mimc.NewMiMC(api)

	if err != nil {
		return err
	}
	commitHash.Write(circuit.Public.ChainID)
	commitHash.Write(circuit.CreatorPubKeyX)
	commitHash.Write(circuit.CreatorPubKeyY)
	commitHash.Write(circuit.SigR)
	commitHash.Write(circuit.SigS)
	commitHash.Write(circuit.SigV)
	commitHash.Write(circuit.Nonce)
	commitHashSum := commitHash.Sum()
	api.AssertIsEqual(circuit.Public.ID, commitHashSum)
	commitHash.Reset()
	// apptHash, err := mimc.NewMiMC(api)

	// if err != nil {
	// 	return err
	// }

	// 	bits := api.ToBinary(commitHashSum, 128) // Full 32-byte field element

	// 	// Extract first 31 bytes (248 bits)
	// 	chunkedId := api.FromBinary(bits[:128]...) // Keep first 16 bytes

	// //
	commitHash.Write(commitHashSum)
	//commitHash.Write(pub.Name)
	commitHash.Write(pub.Meta)
	commitHash.Write(pub.Ref)
	commitHash.Write(pub.Categories)
	commitHash.Write(pub.Status)
	commitHash.Write(pub.DefaultAuthPrivilege)
	commitHash.Write(pub.Timestamp)
	appCommitment := commitHash.Sum()
	api.AssertIsEqual(circuit.Public.AppCommitment, appCommitment)
	// bn254, err := emulated.NewField[emulated.BN254Fr](api)
	// if err != nil {
	// 	return err
	// }

	// default:
	// 	log.Println("HELOOOOS2")
	// 	api.AssertIsEqual(circuit.SignatureType, new(big.Int).SetBytes([]byte(EthereumSignatureType)))
	// }

	//  else {
	// 	return fmt.Errorf("invalid signature type")
	// }

	// api.AssertIsEqual(circuit.ChainID, []byte("1"))
	return nil

	// // 2. Compute baseHash = keccak256(signature)
	// baseHash, err := keccak.New(api)
	// if err != nil {
	// 	return err
	// }
	// baseHash.Write(circuit.SigR)
	// baseHash.Write(circuit.SigS)
	// baseHash.Write(circuit.SigV)
	// baseHashSum := baseHash.Sum()

	// // 3. Compute nonce = keccak256(baseHash + n)
	// nonceHash, err := keccak.New(api)
	// if err != nil {
	// 	return err
	// }
	// nonceHash.Write(baseHashSum)
	// nonceHash.Write(circuit.N)
	// nonce := nonceHash.Sum()

	// // 4. Compute CreatorCommitment = keccak256(address + nonce)
	// commitHash, err := keccak.New(api)
	// if err != nil {
	// 	return err
	// }
	// for i := 0; i < 20; i++ {
	// 	commitHash.Write(circuit.CreatorAddress[i])
	// }
	// commitHash.Write(nonce)
	// computedCommitment := commitHash.Sum()
	// api.AssertIsEqual(computedCommitment, circuit.CreatorCommitment)

	// // 5. Verify appID = keccak256(ID + Name)
	// appIDHash, err := keccak.New(api)
	// if err != nil {
	// 	return err
	// }
	// appIDHash.Write(circuit.ID)
	// appIDHash.Write(circuit.Name)
	// computedAppID := appIDHash.Sum()
	// api.AssertIsEqual(computedAppID, circuit.AppID)

	// // 6. Verify creation signature
	// creationHash, err := keccak.New(api)
	// if err != nil {
	// 	return err
	// }
	// creationHash.Write(circuit.ID)
	// creationHash.Write(circuit.Name)
	// creationHashSum := creationHash.Sum()
	// sig := ecdsa.Signature{R: circuit.CreatorSigR, S: circuit.CreatorSigS, V: circuit.CreatorSigV}
	// creationPubKey, err := ecdsa.RecoverPublicKey(api, creationHashSum, sig)
	// if err != nil {
	// 	return err
	// }

	// // 7. Verify mod signature
	// modHash, err := keccak.New(api)
	// if err != nil {
	// 	return err
	// }
	// modHash.Write(circuit.AppID)
	// modHash.Write(circuit.NewName)
	// modHashSum := modHash.Sum()
	// modSig := ecdsa.Signature{R: circuit.ModSigR, S: circuit.ModSigS, V: circuit.ModSigV}
	// err = ecdsa.Verify(api, modHashSum, modSig, creationPubKey)
	// if err != nil {
	// 	return err
	// }

	// // 8. Ensure same pubkey
	// api.AssertIsEqual(pubKey.X, creationPubKey.X)
	// api.AssertIsEqual(pubKey.Y, creationPubKey.Y)

	// return nil
}
