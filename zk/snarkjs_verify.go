package zk

import (
	"errors"
	"fmt"
	"log"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
	"github.com/mlayerprotocol/go-mlayer/entities"
)

// Logger interface for logging messages
type Logger interface {
	Info(msg string)
	Error(msg string)
}

// SimpleLogger implements the Logger interface
type SimpleLogger struct{}

func (l *SimpleLogger) Info(msg string) {
	log.Println("INFO:", msg)
}

func (l *SimpleLogger) Error(msg string) {
	log.Println("ERROR:", msg)
}

// Convert a string to a big.Int
func stringToBigInt(s string) (*big.Int, error) {
	bi := new(big.Int)
	_, success := bi.SetString(s, 10)
	if !success {
		return nil, fmt.Errorf("failed to convert string to big.Int: %s", s)
	}
	return bi, nil
}

// Convert big.Int array to G1 point in bn254
func bigIntArrayToG1Point(coords []string) (bn254.G1Affine, error) {
	if len(coords) != 2 {
		
		return bn254.G1Affine{}, errors.New(fmt.Sprintf("G1 point requires exactly 2 coordinates. %d given", (len(coords))))
	}
	
	x, err := stringToBigInt(coords[0])
	if err != nil {
		return bn254.G1Affine{}, err
	}
	
	y, err := stringToBigInt(coords[1])
	if err != nil {
		return bn254.G1Affine{}, err
	}
	
	var point bn254.G1Affine
	point.X.SetBigInt(x)
	point.Y.SetBigInt(y)
	
	return point, nil
}

// Convert big.Int 2D array to G2 point in bn254
func bigIntArrayToG2Point(coords [2][2]string) (bn254.G2Affine, error) {
	x0, err := stringToBigInt(coords[0][0])
	if err != nil {
		return bn254.G2Affine{}, err
	}
	
	x1, err := stringToBigInt(coords[0][1])
	if err != nil {
		return bn254.G2Affine{}, err
	}
	
	y0, err := stringToBigInt(coords[1][0])
	if err != nil {
		return bn254.G2Affine{}, err
	}
	
	y1, err := stringToBigInt(coords[1][1])
	if err != nil {
		return bn254.G2Affine{}, err
	}
	
	var point bn254.G2Affine
	point.X.A0.SetBigInt(x0)
	point.X.A1.SetBigInt(x1)
	point.Y.A0.SetBigInt(y0)
	point.Y.A1.SetBigInt(y1)
	
	return point, nil
}

// Check if public inputs are valid (in the field)
func publicInputsAreValid(publicSignals []string) bool {
	modulus := bn254.ID.ScalarField()
	
	for _, sig := range publicSignals {
		bi, err := stringToBigInt(sig)
		if err != nil {
			return false
		}
		
		if bi.Cmp(modulus) >= 0 {
			return false
		}
	}
	
	return true
}

// Check if proof commitments are well-constructed
func isWellConstructed(pi_a bn254.G1Affine, pi_b bn254.G2Affine, pi_c bn254.G1Affine) bool {
	// Check if points are in the right subgroup and on curve
	if !pi_a.IsInSubGroup() || !pi_b.IsInSubGroup() || !pi_c.IsInSubGroup() {
		return false
	}
	
	return true
}

// Groth16Verify verifies a Groth16 proof with the provided verification key and public signals
func Groth16Verify(vk entities.SnarkjsVerificationKey,publicSignals []string, proof entities.SnarkjsProof, loggers ...Logger) (bool, error) {
	var logger Logger
	if loggers == nil {
		logger = &SimpleLogger{}
	} else {
		logger = loggers[0]
	}
	
	// Parse verification key
	// var vk SnarkjsVerificationKey
	// if err := json.Unmarshal([]byte(vkJSON), &vk); err != nil {
	// 	return false, fmt.Errorf("failed to parse verification key: %v", err)
	// }

	// Only bn254 curve is supported in this implementation
	if vk.Curve != "bn254" && vk.Curve != "bn128" {
		return false, fmt.Errorf("unsupported curve: %s", vk.Curve)
	}
	
	// Parse public signals
	// var publicSignals []string
	// if err := json.Unmarshal([]byte(publicSignalsJSON), &publicSignals); err != nil {
	// 	return false, fmt.Errorf("failed to parse public signals: %v", err)
	// }
	
	// Parse proof
	// var proof SnarkjsProof
	// if err := json.Unmarshal([]byte(proofJSON), &proof); err != nil {
	// 	return false, fmt.Errorf("failed to parse proof: %v", err)
	// }
	
	// Validate public inputs

	if !publicInputsAreValid(publicSignals) {
		logger.Error("Public inputs are not valid.")
		return false, nil
	}
	
	// Get IC0 (first element of IC)
	ic0, err := bigIntArrayToG1Point(vk.IC[0][:2])
	if err != nil {
		return false, fmt.Errorf("failed to parse IC0: %v", err)
	}
	
	// Calculate linear combination of public inputs and IC elements
	var cpub bn254.G1Jac
	cpub.FromAffine(&ic0)
	
	for i := 0; i < len(publicSignals); i++ {
		// Get IC[i+1]
		if i+1 >= len(vk.IC) {
			return false, fmt.Errorf("IC index out of bounds: %d", i+1)
		}
		
		ici, err := bigIntArrayToG1Point(vk.IC[i+1][:2])
		if err != nil {
			return false, fmt.Errorf("failed to parse IC[%d]: %v", i+1, err)
		}
		
		// Get public signal as scalar
		ps, err := stringToBigInt(publicSignals[i])
		if err != nil {
			return false, fmt.Errorf("failed to parse public signal %d: %v", i, err)
		}
		
		var scalar fr.Element
		scalar.SetBigInt(ps)
		
		// Calculate ici * public_signal and add to cpub
		var icJac bn254.G1Jac
		icJac.FromAffine(&ici)
		icJac.ScalarMultiplication(&icJac, scalar.BigInt(new(big.Int)))
		cpub.AddAssign(&icJac)
	}
	
	// Convert to affine
	var cpubAffine bn254.G1Affine
	cpubAffine.FromJacobian(&cpub)
	
	// Get proof elements
	piA, err := bigIntArrayToG1Point(proof.PiA[:2])
	if err != nil {
		return false, fmt.Errorf("failed to parse pi_a: %v", err)
	}
	
	// proof.PiB is already the correct type [2][2]string
	piB, err := bigIntArrayToG2Point(proof.PiB)
	if err != nil {
		return false, fmt.Errorf("failed to parse pi_b: %v", err)
	}
	
	piC, err := bigIntArrayToG1Point(proof.PiC[:2])
	if err != nil {
		return false, fmt.Errorf("failed to parse pi_c: %v", err)
	}
	
	// Check if proof is well-constructed
	if !isWellConstructed(piA, piB, piC) {
		logger.Error("SnarkjsProof commitments are not valid.")
		return false, nil
	}
	
	// Parse verification key elements
	vkAlpha1, err := bigIntArrayToG1Point(vk.VkAlpha1[:2])
	if err != nil {
		return false, fmt.Errorf("failed to parse vk_alpha_1: %v", err)
	}
	
	// Convert string slices to fixed arrays for G2 points
	var vkBeta2Arr [2][2]string
	var vkGamma2Arr [2][2]string
	var vkDelta2Arr [2][2]string
	
	if len(vk.VkBeta2[:2]) == 2 && len(vk.VkBeta2[0]) == 2 && len(vk.VkBeta2[1]) == 2 {
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				vkBeta2Arr[i][j] = vk.VkBeta2[i][j]
			}
		}
	} else {
		return false, fmt.Errorf("invalid vk_beta_2 format")
	}
	
	if len(vk.VkGamma2[:2]) == 2 && len(vk.VkGamma2[0]) == 2 && len(vk.VkGamma2[1]) == 2 {
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				vkGamma2Arr[i][j] = vk.VkGamma2[i][j]
			}
		}
	} else {
		return false, fmt.Errorf("invalid vk_gamma_2 format")
	}
	
	if len(vk.VkDelta2[:2]) == 2 && len(vk.VkDelta2[0]) == 2 && len(vk.VkDelta2[1]) == 2 {
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				vkDelta2Arr[i][j] = vk.VkDelta2[i][j]
			}
		}
	} else {
		return false, fmt.Errorf("invalid vk_delta_2 format")
	}
	
	vkBeta2, err := bigIntArrayToG2Point(vkBeta2Arr)
	if err != nil {
		return false, fmt.Errorf("failed to parse vk_beta_2: %v", err)
	}
	
	vkGamma2, err := bigIntArrayToG2Point(vkGamma2Arr)
	if err != nil {
		return false, fmt.Errorf("failed to parse vk_gamma_2: %v", err)
	}
	
	vkDelta2, err := bigIntArrayToG2Point(vkDelta2Arr)
	if err != nil {
		return false, fmt.Errorf("failed to parse vk_delta_2: %v", err)
	}
	
	// Negate piA for pairing check
	var negPiA bn254.G1Affine
	negPiA.Neg(&piA)
	
	// Perform pairing check
	// e(-pi_a, pi_b) * e(cpub, vk_gamma_2) * e(pi_c, vk_delta_2) * e(vk_alpha_1, vk_beta_2) == 1
	pairingCheck, err := bn254.PairingCheck(
		[]bn254.G1Affine{negPiA, cpubAffine, piC, vkAlpha1},
		[]bn254.G2Affine{piB, vkGamma2, vkDelta2, vkBeta2},
	)
	if err != nil {
		return false, fmt.Errorf("pairing check failed: %v", err)
	}
	
	if !pairingCheck {
		logger.Error("Invalid proof")
		return false, nil
	}
	
	logger.Info("OK!")
	return true, nil
}

// Usage example:
// func main() {
//     vkJSON := `{"curve":"bn254","IC":[["123...", "456..."], ...], ...}`
//     publicSignalsJSON := `["123...", "456...", ...]`
//     proofJSON := `{"pi_a":["123...", "456..."], "pi_b":[["123...", "456..."], ["789...", "101..."]], "pi_c":["123...", "456..."]}`
//     
//     result, err := Groth16Verify(vkJSON, publicSignalsJSON, proofJSON, nil)
//     if err != nil {
//         log.Fatalf("Verification error: %v", err)
//     }
//     
//     fmt.Printf("Verification result: %v\n", result)
// }