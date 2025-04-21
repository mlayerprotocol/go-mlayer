package entities

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/configs"
)

// SnarkjsProof represents the proof format from snarkjs
type SnarkjsProof struct {
	Protocol string       `json:"protocol"`
	Curve    string       `json:"curve"`
	PI       []string     `json:"pi"`
	PiA      []string    `json:"pi_a"`
	PiB      [2][2]string `json:"pi_b"`
	PiC      []string    `json:"pi_c"`
}

type SnarkjsVerificationKey struct {
	Protocol      string       `json:"protocol"`
	Curve         string       `json:"curve"`
	IC        [][]string      `json:"IC"`
	NPublic   int             `json:"nPublic"`
	VkAlpha1  []string        `json:"vk_alpha_1"`
	VkBeta2   [][]string      `json:"vk_beta_2"`
	VkGamma2  [][]string      `json:"vk_gamma_2"`
	VkDelta2  [][]string      `json:"vk_delta_2"`
}

type CircuitData struct {
	VerificationKey groth16.VerifyingKey
	ProvingKey groth16.ProvingKey
	R1CS constraint.ConstraintSystem
	Circuit *frontend.Circuit
}
const CURVE = ecc.BN254

type CircuitProps struct {
	ChainId configs.ChainId
	Circuit frontend.Circuit
	Type constants.CircuitType
	// VerificationKey string
	// ProvingKey string
	// R1CS string
	verificationKeyJson map[string]*SnarkjsVerificationKey
	Version uint8
}

func (cp CircuitProps) LoadData(dir string, cdMap *map[uint8]*CircuitData) (cd CircuitData ) {
	if cp.ChainId == "" {
		logger.Fatal("invalid chain id for circuitprop")
	}
	if cp.Type == "" {
		logger.Fatal("invalid circuit type")
	}
	pk := groth16.NewProvingKey(CURVE)
	vk := groth16.NewVerifyingKey(CURVE)
	pkFile, err := os.Open(cp.GetProvingKeyPath(dir))
	defer pkFile.Close()
	if err != nil {
		logger.Fatalf("Failed reading proving key: %v", err)
	}
	_, err = pk.ReadFrom(pkFile)
	if err != nil {
		logger.Fatalf("Failed to read proving key: %v", err)
	}
	vkFile, err := os.Open(cp.GetVerificationKeyPath(dir))
	defer vkFile.Close()
	if err != nil {
		logger.Fatalf("Failed to open verification key file: %v", err)
	}
	_, err = vk.ReadFrom(vkFile)
	if err != nil {
		logger.Fatalf("Failed to load verification key: %v", err)
	}

	r1cs := groth16.NewCS(CURVE)
	r1csFile, err := os.Open(cp.GetR1CSPath(dir))
	defer r1csFile.Close()
	if err != nil {
		logger.Fatalf("Failed to read verification key: %v", err)
	}
	_, err = r1cs.ReadFrom(r1csFile)
	if err != nil {
		logger.Fatalf("Failed to read r1cs file: %v", err)
	}
	
	cd = CircuitData{
		ProvingKey: pk,
		VerificationKey: vk,
		R1CS: r1cs,
		Circuit: &cp.Circuit,
	}
	(*cdMap)[cp.Version] =  &cd
	return cd
}
func (cp CircuitProps) GetVerificationKeyPath(dir string) string {
	if cp.ChainId == "" {
		logger.Fatal("invalid chain id for circuitprop")
	}
	if cp.Type == "" {
		logger.Fatal("invalid circuit type")
	}
	err := os.MkdirAll(fmt.Sprintf("%s/%s/%s/v%d/", dir, cp.ChainId, cp.Type, cp.Version), 0755)
	if err != nil {
		logger.Fatal(err)
	}
	return fmt.Sprintf("%s/%s/%s/v%d/%s", dir, cp.ChainId, cp.Type, cp.Version, constants.VERFICATION_KEY_JSON_FILE)
}
func (cp CircuitProps) GetVerificationKeyJSON(dir string) (*SnarkjsVerificationKey, error) {
	if cp.verificationKeyJson[dir] != nil {
		return cp.verificationKeyJson[dir], nil
	}
	vkFile := cp.GetVerificationKeyPath(dir)
	vkData, err := os.ReadFile(vkFile)
	if err != nil {
		fmt.Printf("Error reading verification key: %v\n", err)
		return nil, err
	}
	// logger.Infof("%v", vkData)
	var snarkjsVk SnarkjsVerificationKey
	if err := json.Unmarshal(vkData, &snarkjsVk); err != nil {
		fmt.Printf("Error unmershaling verification key: %v\n", err)
		return nil, err
	}
	return &snarkjsVk, nil
}
func (cp CircuitProps) GetProvingKeyPath(dir string) string {
	if cp.ChainId == "" {
		logger.Fatal("invalid chain id for circuitprop")
	}
	if cp.Type == "" {
		logger.Fatal("invalid circuit type")
	}
	err := os.MkdirAll(fmt.Sprintf("%s/%s/%s/v%d", dir,cp.ChainId, cp.Type, cp.Version), 0755)
	if err != nil {
		logger.Fatal(err)
	}
	return fmt.Sprintf("%s/%s/%s/v%d/%s", dir,  cp.ChainId, cp.Type, cp.Version,  constants.PROVING_KEY_FILE)
}
func (cp CircuitProps) GetR1CSPath(dir string) string {
	if cp.ChainId == "" {
		logger.Fatal("invalid chain id for circuitprop")
	}
	if cp.Type == "" {
		logger.Fatal("invalid circuit type")
	}
	err := os.MkdirAll(fmt.Sprintf("%s/%s/%s/v%d", dir, cp.ChainId, cp.Type, cp.Version), 0755)
	if err != nil {
		logger.Fatal(err)
	}
	return fmt.Sprintf("%s/%s/%s/v%d/%s", dir, cp.ChainId, cp.Type, cp.Version, constants.R1CS_FILE)
}
// var appCircuitProps  = CircuitProps{
// 	ChainId: "84532",
// 	Circuit: nil,
// 	Type: constants.APP_CIRCUIT,
// 	// ProvingKey: "app_circuit_proving_key.vk",
// 	// R1CS: "app_circuit.r1cs",
// 	Version: 1,
// }

// var clientPayloadCircuitProps  = CircuitProps{
// 	ChainId: "84532",
// 	Circuit: nil,
// 	Type: constants.CLIENT_PAYLOAD,
// 	Version: 1,
// }
func GetCiruitProps(_type constants.CircuitType, chainId configs.ChainId, version uint8) CircuitProps {
	return CircuitProps{
		ChainId: chainId,
		Circuit: nil,
		Type: _type,
		Version: version,
	}
}

