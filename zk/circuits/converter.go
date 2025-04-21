package circuits

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/consensys/gnark-crypto/ecc"
	curve "github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark/backend/groth16"
	groth16_bls12381 "github.com/consensys/gnark/backend/groth16/bls12-381"
	groth16_bn254 "github.com/consensys/gnark/backend/groth16/bn254"
	"github.com/consensys/gnark/backend/witness"
	"github.com/mlayerprotocol/go-mlayer/entities"
)

// SnarkjsVerificationKey represents the verification key format from snarkjs



func ConvertToBigInt(s string) *big.Int {
	n := new(big.Int)
	n.SetString(s, 10)
	return n
}




// ConvertVerificationKey converts snarkjs verification key to gnark format
func ConvertVerificationKeyBN254(snarkVk entities.SnarkjsVerificationKey) (*groth16_bn254.VerifyingKey, error) {
	// Create a new BN254 verifying key
	vk := groth16.NewVerifyingKey(CURVE).(*groth16_bn254.VerifyingKey)

	// Populate G1 elements
	// vk_alpha_1 -> G1.Alpha
	vk.G1.Alpha.X.SetString(snarkVk.VkAlpha1[0])
	vk.G1.Alpha.Y.SetString(snarkVk.VkAlpha1[1])

	// IC (public input commitments) -> G1.K
	vk.G1.K = make([]curve.G1Affine, len(snarkVk.IC))
	for i, ic := range snarkVk.IC {
		vk.G1.K[i].X.SetString(ic[0])
		vk.G1.K[i].Y.SetString(ic[1])
	}

	// Populate G2 elements
	// vk_beta_2 -> G2.Beta
	vk.G2.Beta.X.A0.SetString(snarkVk.VkBeta2[0][0])
	vk.G2.Beta.X.A1.SetString(snarkVk.VkBeta2[0][1])
	vk.G2.Beta.Y.A0.SetString(snarkVk.VkBeta2[1][0])
	vk.G2.Beta.Y.A1.SetString(snarkVk.VkBeta2[1][1])

	// vk_gamma_2 -> G2.Gamma
	vk.G2.Gamma.X.A0.SetString(snarkVk.VkGamma2[0][0])
	vk.G2.Gamma.X.A1.SetString(snarkVk.VkGamma2[0][1])
	vk.G2.Gamma.Y.A0.SetString(snarkVk.VkGamma2[1][0])
	vk.G2.Gamma.Y.A1.SetString(snarkVk.VkGamma2[1][1])

	// vk_delta_2 -> G2.Delta
	vk.G2.Delta.X.A0.SetString(snarkVk.VkDelta2[0][0])
	vk.G2.Delta.X.A1.SetString(snarkVk.VkDelta2[0][1])
	vk.G2.Delta.Y.A0.SetString(snarkVk.VkDelta2[1][0])
	vk.G2.Delta.Y.A1.SetString(snarkVk.VkDelta2[1][1])

	// Compute negations (required by gnark for verification)
	vk.G2.Delta.Neg(&vk.G2.Delta)
	vk.G2.Gamma.Neg(&vk.G2.Gamma)



	return vk, nil
}



func ConvertProofBls12381(snarkjsProof entities.SnarkjsProof) (groth16.Proof, error) {
	proof := groth16.NewProof(ecc.BLS12_381).(*groth16_bls12381.Proof)


		// Convert pi_a (G1 point) to Ar
		aX := ConvertToBigInt(snarkjsProof.PiA[0])
		aY := ConvertToBigInt(snarkjsProof.PiA[1])
		proof.Ar.X.SetBigInt(aX)
		proof.Ar.Y.SetBigInt(aY)
		
		// Convert pi_b (G2 point) to Bs
		bX0 := ConvertToBigInt(snarkjsProof.PiB[0][0])
		bX1 := ConvertToBigInt(snarkjsProof.PiB[0][1])
		bY0 := ConvertToBigInt(snarkjsProof.PiB[1][0])
		bY1 := ConvertToBigInt(snarkjsProof.PiB[1][1])
		proof.Bs.X.A0.SetBigInt(bX0)
		proof.Bs.X.A1.SetBigInt(bX1)
		proof.Bs.Y.A0.SetBigInt(bY0)
		proof.Bs.Y.A1.SetBigInt(bY1)
		
		// Convert pi_c (G1 point) to Krs
		cX := ConvertToBigInt(snarkjsProof.PiC[0])
		cY := ConvertToBigInt(snarkjsProof.PiC[1])
		proof.Krs.X.SetBigInt(cX)
		proof.Krs.Y.SetBigInt(cY)
		
		// Commitments and CommitmentPok are used for batched proofs
		// In standard Groth16 proofs from snarkjs, these aren't used
		// so we'll leave them with their default values
		
		return proof, nil
	
}

func ConvertProofBn254(snarkjsProof entities.SnarkjsProof) (groth16_bn254.Proof, error) {
	// proof := groth16.NewProof(ecc.BN254).(*groth16_bls12381.Proof)
	proof := groth16_bn254.Proof{}
	// Convert pi_a (G1 point) to Ar
	 // Convert pi_a (G1 point) to Ar
	 aX := ConvertToBigInt(snarkjsProof.PiA[0])
	 aY := ConvertToBigInt(snarkjsProof.PiA[1])
	 proof.Ar.X.SetBigInt(aX)
	 proof.Ar.Y.SetBigInt(aY)
	 
	 // Convert pi_b (G2 point) to Bs
	 // NOTE: snarkjs and gnark may use different coordinate orders for G2
	 // In snarkjs, coordinates might be [Re, Im] while gnark might expect [Im, Re]
	 // Try both orderings to determine which one is correct
	 bX0 := ConvertToBigInt(snarkjsProof.PiB[0][0])
	 bX1 := ConvertToBigInt(snarkjsProof.PiB[0][1])
	 bY0 := ConvertToBigInt(snarkjsProof.PiB[1][0])
	 bY1 := ConvertToBigInt(snarkjsProof.PiB[1][1])
	 
	 // Option 1: Standard ordering
	 proof.Bs.X.A0.SetBigInt(bX0)
	 proof.Bs.X.A1.SetBigInt(bX1)
	 proof.Bs.Y.A0.SetBigInt(bY0)
	 proof.Bs.Y.A1.SetBigInt(bY1)
	 
	 // Option 2: Try swapping real and imaginary parts
	 // Uncomment if Option 1 doesn't work
	 /*
	 proof.Bs.X.A0.SetBigInt(bX1)
	 proof.Bs.X.A1.SetBigInt(bX0)
	 proof.Bs.Y.A0.SetBigInt(bY1)
	 proof.Bs.Y.A1.SetBigInt(bY0)
	 */
	 
	 // Option 3: Try conjugating the G2 point 
	 // (keep A0 the same, negate A1)
	 // Uncomment if Options 1 and 2 don't work
	 /*
	 proof.Bs.X.A0.SetBigInt(bX0)
	 bX1Neg := new(big.Int).Neg(bX1)
	 proof.Bs.X.A1.SetBigInt(bX1Neg)
	 proof.Bs.Y.A0.SetBigInt(bY0)
	 bY1Neg := new(big.Int).Neg(bY1)
	 proof.Bs.Y.A1.SetBigInt(bY1Neg)
	 */
	 
	 // Convert pi_c (G1 point) to Krs
	 cX := ConvertToBigInt(snarkjsProof.PiC[0])
	 cY := ConvertToBigInt(snarkjsProof.PiC[1])
	 proof.Krs.X.SetBigInt(cX)
	 proof.Krs.Y.SetBigInt(cY)
	 
	 return proof, nil
	
	
}

// ConvertPublicInputs converts snarkjs public inputs to gnark format
func ConvertPublicInputs(snarkjsProof entities.SnarkjsProof, numPublic int) (witness.Witness, error) {
	// Create a channel to fill the witness values
	valuesChan := make(chan any, numPublic)
	// defer close(valuesChan)

	w, err := witness.New(ecc.BN254.ScalarField())
	if err != nil {
		return nil, fmt.Errorf("failed to create witness: %v", err)
	}
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := w.Fill(numPublic, 0, valuesChan); err != nil {
			logger.Errorf("failed to fill witness: %v", err)
			return
		}
		
	}()
	// Convert each string public input to a field element
	for i := 0; i < numPublic; i++ {
		// var val fr.Element
		if i < len(snarkjsProof.PI) {
			pi := ConvertToBigInt(snarkjsProof.PI[i])
			// val.SetBigInt(pi)
			valuesChan <- pi
		} else {
			// If there are fewer public inputs than expected, fill with zeros
			// val.SetZero()
			valuesChan <- new(big.Int)
		}
		
		// Send the value to the channel
		// valuesChan <- val
	}
	
	close(valuesChan)
	wg.Wait()
	// Create a witness from the values channel
	
	
	// Fill the witness with public values (no secret values for verification)
	// The Fill method expects: number of public inputs, number of secret inputs, values channel
	
	
	return w, nil
}