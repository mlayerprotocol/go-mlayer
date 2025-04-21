package circuits

import (
	"os"
	"testing"

	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)


func TestSetupCircuit(t *testing.T) {
	// Instantiate the circuit
	return
	var circuit = NewAppCircuit("1")


	// Compile to R1CS
	r1cs, err := frontend.Compile(CURVE.ScalarField(), r1cs.NewBuilder, &circuit)
	if err != nil {
		t.Fatalf("Failed to compile circuit: %v", err)
	}

	// Perform Groth16 trusted setup
	pk, vk, err := groth16.Setup(r1cs)
	if err != nil {
		t.Fatalf("Failed to setup circuit: %v", err)
	}

	// Save keys to files (for use in tests or app)
	err = writeFile(AppCircuitProps.GetProvingKeyPath("../test_data"), pk)
	if err != nil {
		t.Fatalf("Failed to write proving key: %v", err)
	}
	err = writeFile(AppCircuitProps.GetVerificationKeyPath("../test_data"), vk)
	if err != nil {
		t.Fatalf("Failed to write verification key: %v", err)
	}
	err = writeFile(AppCircuitProps.GetR1CSPath("../test_data"), r1cs)
	if err != nil {
		t.Fatalf("Failed to write R1CS: %v", err)
	}

	// For snarkjs (frontend), you'd need to export to wasm/zkey format
	// This requires external tools (e.g., snarkjs CLI) - see below
}

func writeFile(filename string, data interface{}) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	switch v := data.(type) {
	case groth16.ProvingKey:
		_, err = v.WriteTo(f)
	case groth16.VerifyingKey:
		_, err = v.WriteTo(f)
	case constraint.ConstraintSystem:
		_, err = v.WriteTo(f)
	}
	return err
}