package entities

import (
	"encoding/hex"
	"testing"

	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/utils/constants"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestEncodeAuthorityBytes(t *testing.T) {
	agent := "0xe652d28F89A28adb89e674a6b51852D0C341Ebe9"
	grantor := "02ebec9d95769bb3d71712f0bf1e7e88b199fc945f67f908bbab81e9b7cb1092d8"
	topicIds := "*"
	ts := uint64(1705392177894)
	privilege := constants.AdminPriviledge
    auth := Authorization{Agent: agent, Grantor: grantor,  TopicIds: topicIds, Timestamp: ts,  Priviledge: privilege, }
	hash := hex.EncodeToString(auth.GetHash())
	expected := "6be68630801cf55ba605dee655ed7742b6f7979706392b7691597c3f8210546c"
    if hash != expected {
        t.Fatalf("Hash should match \n%s \n%s", hash, expected )
    }
}

func TestVerifyAuthority(t *testing.T) {
	agent := "0xe652d28F89A28adb89e674a6b51852D0C341Ebe9"
	grantor := "02ebec9d95769bb3d71712f0bf1e7e88b199fc945f67f908bbab81e9b7cb1092d8"
	topicIds := "*"
	ts := uint64(1705392177894)
	privilege := constants.AdminPriviledge
    auth := Authorization{Agent: agent, TopicIds: topicIds, Timestamp: ts,  Priviledge: privilege, }
	// hash := hex.EncodeToString(auth.GetHash())
	
	encoded, err := auth.EncodeBytes();
	sign := "3044022073761e96c0d29d109ba2290fea1821bf093155d7549956a8144d04f9209f6c1b022048d4635f7e2f3e2864b2c1fe695badad548483ce1988f8bf0c04d5caf4131829"
	valid, err := crypto.VerifySignatureEDD(grantor, &encoded, sign)
    if !valid {
        t.Fatalf("Invalid signature signer: %v", err )
    }
}