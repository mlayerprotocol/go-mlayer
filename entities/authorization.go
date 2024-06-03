package entities

import (
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
)

type PubKeyType string

const (
	TendermintsSecp256k1PubKey PubKeyType = "tendermint/PubKeySecp256k1"
	EthereumPubKey             PubKeyType = "ethereum"
)

type Authorization struct {
	ID            string                           	`json:"id" gorm:"type:uuid;not null;primaryKey"`
	Agent         DeviceString                    	`json:"agt" gorm:"uniqueIndex:idx_agent_account_subnet;index:idx_authorization_states_agent"`
	Meta          string                           	`json:"meta,omitempty"`
	Account       DIDString                        	`json:"acct" gorm:"varchar(40);uniqueIndex:idx_agent_account_subnet"`
	Grantor       DIDString                        	`json:"gr" gorm:"index"`
	Priviledge    *constants.AuthorizationPrivilege	`json:"privi"`
	TopicIds      string                           	`json:"topIds"`
	Timestamp     *uint64                           `json:"ts"`
	Duration      *uint64                           `json:"du"`
	SignatureData SignatureData                    	`json:"sigD" gorm:"json;"`
	Hash          string                           	`json:"h" gorm:"unique" `
	Event         EventPath                        	`json:"e,omitempty" gorm:"index;varchar;"`
	Subnet        string                           	`json:"snet" gorm:"uniqueIndex:idx_agent_account_subnet;char(36)"`

	// AuthorizationEventID string                           `json:"authEventId,omitempty"`
}

func (g Authorization) GetHash() ([]byte, error) {
	if g.Hash != "" {
		return hex.DecodeString(g.Hash)
	}
	b, err := (g.EncodeBytes())
	logger.Info("EncodeBytes:: ", b)
	if err != nil {
		logger.Errorf("Error endoding Authorization: %v", err)
		return []byte(""), err
	}
	bs := crypto.Sha256(b)
	return bs, nil
}

func (entity Authorization) GetEvent() EventPath {
	return entity.Event
}
func (entity Authorization) GetAgent() DeviceString {
	return entity.Agent
}
func (g Authorization) ToJSON() []byte {
	b, _ := json.Marshal(g)
	return b
}

func (g Authorization) ToString() string {
	return fmt.Sprintf("TopicIds:%s, Priviledge: %d, Grantor: %s, Timestamp: %d", g.TopicIds, g.Priviledge, g.Grantor, g.Timestamp)
}

func (g Authorization) EncodeBytes() ([]byte, error) {

	b, e := encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.AddressEncoderDataType, Value: string(g.Account)},
		encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: AddressFromString(string(g.Agent)).Addr},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: g.Meta},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: g.TopicIds},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: *g.Priviledge},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: *g.Duration},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: g.Subnet},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: *g.Timestamp},
	)

	return b, e
}
