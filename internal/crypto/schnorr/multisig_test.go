package schnorr

import (
	"encoding/hex"
	"log"
	"math/big"
	"testing"

	btcec "github.com/btcsuite/btcd/btcec/v2"
	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	crypto "github.com/mlayerprotocol/go-mlayer/internal/crypto"
)
func PrivateKeyFromSeed(seed *big.Int) *btcec.PrivateKey {
	priv, _ := btcec.PrivKeyFromBytes(seed.Bytes())
	return priv
}

var seedStrings = []string{"2f04020b9c71fdb501d5ce30b99b85f7", "5da1ab48fa824aa7daa68547abfaadad"}
var privKeys = []*secp256k1.PrivateKey{}
var pubKeys = []*secp256k1.PublicKey{}
var aggrPubKey = &secp256k1.PublicKey{}
var message [32]byte
var rawMessage = []byte("femi")
var noncePubKeys = []*btcec.PublicKey{};
var challenge []byte
	
	

func TestGeneratePrivateKey(t *testing.T) {
	copy(message[:], crypto.Keccak256Hash(rawMessage))
	for _, seed := range seedStrings {
		sb, _ := new(big.Int).SetString(seed, 16)
		privKeys = append(privKeys, PrivateKeyFromSeed(sb))
	}
	
	result :=  new(big.Int).SetBytes(privKeys[0].Serialize()).String()
	// fmt.Printf("PivateKey :%s", new(big.Int).SetBytes(privKey.Serialize()).String())
	expected := "62494526474071648544700980894199678455"
	if result != expected {
        t.Fatalf("Private keys should match \n%s \n%s", result, expected )
    }
	result =  new(big.Int).SetBytes(privKeys[1].Serialize()).String()
	// fmt.Printf("PivateKey :%s", new(big.Int).SetBytes(privKey.Serialize()).String())
	expected = "124457637476219974254289856756635577773"
	if result != expected {
        t.Fatalf("Private keys should match \n%s \n%s", result, expected )
    }
	if result != expected {
        t.Fatalf("Private keys should match \n%s \n%s", result, expected )
    }
}

func TestDerivePublicKey(t *testing.T) {
	for _, pk := range privKeys {
		pubKeys = append(pubKeys, derivePublicKey(pk))
	}
	result :=  pubKeys[0].X().String()
	// fmt.Printf("PivateKey :%s", new(big.Int).SetBytes(privKey.Serialize()).String())
	expected := "63488965236511607412715693736423142454322830812436117483930645066197267246091"
	if result != expected {
        t.Fatalf("Pubkeys should match \n%s \n%s", result, expected )
    }

	result =  pubKeys[0].Y().String()
	expected = "107346052469528284779965882535674902951871426462174991766787076706756101430131"
	if result != expected {
        t.Fatalf("Pubkeys should match \n%s \n%s", result, expected )
    }

	result =  pubKeys[1].X().String()
	expected = "21369760946825369723963269079243100179471403038092769272643770557397356964128"
	if result != expected {
        t.Fatalf("Pubkeys should match \n%s \n%s", result, expected )
    }

	result =  pubKeys[1].Y().String()
	expected = "51612781179882068698114516825772001362194794782086041574848139819611992561719"
	if result != expected {
        t.Fatalf("Pubkeys should match \n%s \n%s", result, expected )
    }
	
}

func TestAggregatePublicKey(t *testing.T) {
	aggrPubKey = aggregatePublicKeys(pubKeys)
	// result :=  aggr.X()
	// fmt.Printf("PivateKey :%s", new(big.Int).SetBytes(privKey.Serialize()).String())
	expectedX := "13319787095898482284074238777397495332723331979310772692077474715385348418035"
	expectedY := "15125206542544477863509685227600268907962821962684378361844005754150244698826"
	if aggrPubKey.X().String() != expectedX || expectedY != aggrPubKey.Y().String()  {
        t.Fatalf("Aggregate Pubkeys should match \n%s:%s \n%s:%s", aggrPubKey.X().String(), aggrPubKey.Y().String(),  expectedX, expectedY )
    }
	
}

func TestDeriveCommitment(t *testing.T) {
	
	for _, key := range privKeys {
		// 3.1. Derive secure nonce.
		nonce := deriveNonce(key, message);
		noncePubKeys = append(noncePubKeys,  computeNoncePublicKey(nonce))
	}
	
	expectedX := "114666290145335804710458237660688789628201140305607017958832537131258184845042"
	expectedY := "114040225851231968202220421694065319367399345617550080805323747574665270779757"
	if noncePubKeys[0].X().String() != expectedX || expectedY !=  noncePubKeys[1].Y().String()  {
        t.Fatalf("Nonce Pubkeys should match \n%s:%s \n%s:%s",  noncePubKeys[0].X().String(),  noncePubKeys[1].Y().String(),  expectedX, expectedY )
    }
}


func TestConstructChallenge(t *testing.T) {
	aggNoncePubKey := aggregatePublicKeys(noncePubKeys);
	commitment := deriveCommitment(aggNoncePubKey);
	address := "0x"+hex.EncodeToString(commitment)
        // 6. Construct challenge.
       //  bytes32 challenge = constructChallenge(aggPubKey, message, commitment);
	expected := "0x569fce1f60c65253078404ca95310723c435df39"

	challenge = constructChallenge(aggrPubKey, message, commitment)
	result := new(big.Int).SetBytes(challenge).String()
	expectedCHallange := "18410794103381543374954994547009695295412136899033916616702618786475001072139"
	if address != expected {
        t.Fatalf("Commitment should match \n%s:%s",  address,  expected )
    }
	if result != expectedCHallange {
        t.Fatalf("Challenges should match \n%s:%s",  result,  expectedCHallange )
    }
}

func TestComputeSignature(t *testing.T) {

	signs := ComputeSignatureMulti(privKeys, message, challenge)
	

	result := new(big.Int).SetBytes(AggregateSignatures(signs)).String()
	expected := "37208721598732191769956498780736906110454235485826871899066067715200943165141"
	if result != expected {
        t.Fatalf("Aggregate Sign should match \n%s:%s",  result,  expected )
    }
	// if result != expectedCHallange {
    //     t.Fatalf("Challenges should match \n%s:%s",  result,  expectedCHallange )
    // }
}

func TestSingleSigner(t *testing.T) {
	
	nonce := deriveNonce(privKeys[0], message);
	noncePubKey := computeNoncePublicKey(nonce);
	commitment := deriveCommitment(noncePubKey);

	log.Printf("COMMITMENT %s", hex.EncodeToString(commitment))
	challenge := constructChallenge(pubKeys[0], message, commitment);
	signature := ComputeSignature(privKeys[0], nonce, challenge);
	expected := []byte{}

	if len(signature) != 0 {
        t.Fatalf("Signature should match \n%s:%s",  new(big.Int).SetBytes(signature),  hex.EncodeToString(expected) )
    }
}