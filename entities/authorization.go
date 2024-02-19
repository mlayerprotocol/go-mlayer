package entities

import (
	"encoding/json"
	"fmt"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
)



type Authorization struct {
	ID   string    `json:"id" gorm:"type:char(36);not null;primaryKey"`
	Agent string    `json:"agt" gorm:"index:idx_agent_authorization,unique"`
	Account PublicKeyString    `json:"acct" gorm:"varchar(32),index:idx_agent_authorization,unique"`
	Grantor AddressString    `json:"gr" gorm:"index"`
	Priviledge constants.AuthorizationPrivilege    `json:"privi"`
	TopicIds string    `json:"topIds"`
	Timestamp uint64    `json:"ts"`
	Duration uint64    `json:"du"`
	Signature   string    `json:"sig"`
	Hash string		`json:"h" gorm:"unique" `
	EventHash string `json:"eH,omitempty" gorm:"index;char(64);"`
	AuthorizationEventID string `json:"authEventId,omitempty"`
	//AuthorizationEvent		AuthorizationEvent `json:"authEvent" gorm:"foreignKey:EventHash"`
}

func (g Authorization) GetHash() ([]byte, error) {
	b, err  := (g.EncodeBytes())
	if (err != nil)  {
		logger.Errorf("Error endoding Authorization: %v", err)
		return []byte(""), err
	}
	bs := crypto.Sha256(b)
	return bs, nil
}



func (g Authorization) ToJSON() []byte {
	b, _ := json.Marshal(g);
	return b
}


func (g Authorization) ToString() string {
	return fmt.Sprintf("TopicIds:%s, Priviledge: %d, Grantor: %s, Timestamp: %d", g.TopicIds, g.Priviledge, g.Grantor,  g.Timestamp )
}

func (g Authorization) EncodeBytes() ([]byte, error) {
	
	b, e := encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: string(g.Account)},
		encoder.EncoderParam{Type: encoder.AddressEncoderDataType, Value: g.Agent},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: g.TopicIds},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: g.Priviledge},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: g.Duration},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: g.Timestamp},
	)
	
	return b,e
}