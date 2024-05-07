package entities

import (
	"context"
	"database/sql/driver"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mlayerprotocol/go-mlayer/common/constants"
	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PubKeyType string

const (
	TendermintsSecp256k1PubKey PubKeyType = "tendermint/PubKeySecp256k1"
	EthereumPubKey             PubKeyType = "ethereum"
)

type SignatureData struct {
	Type      PubKeyType `json:"ty"`
	PublicKey string     `json:"pubK,omitempty"`
	Signature string     `json:"sig"`
}

func (sD SignatureData) GormDataType() string {
	return "jsonObject"
}
func (sD SignatureData) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	asJson, _ := json.Marshal(sD)
	return clause.Expr{
		SQL:  "?",
		Vars: []interface{}{asJson},
	}
}

func (sD *SignatureData) Scan(value interface{}) error {
	data, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Value not instance of string:", value))
	}

	result := SignatureData{}
	err := json.Unmarshal(data, &result)
	*sD = SignatureData(result)
	return err
}

// Value return json value, implement driver.Valuer interface
func (sD *SignatureData) Value() (driver.Value, error) {
	if len(sD.Signature) == 0 {
		return nil, nil
	}
	b, _ := json.Marshal(sD)
	return string(b), nil
}

type Authorization struct {
	ID            string                           `json:"id" gorm:"type:uuid;not null;primaryKey"`
	Agent         DeviceString                           `json:"agt" gorm:"index:idx_agent_account,unique"`
	Account       AddressString                    `json:"acct" gorm:"varchar(32),index:idx_agent_account,unique"`
	Grantor       AddressString                    `json:"gr" gorm:"index"`
	Priviledge    constants.AuthorizationPrivilege `json:"privi"`
	TopicIds      string                           `json:"topIds"`
	Timestamp     uint64                           `json:"ts"`
	Duration      uint64                           `json:"du"`
	SignatureData SignatureData                    `json:"sigD" gorm:"jsonObject;"`
	Hash          string                           `json:"h" gorm:"unique" `
	Event         EventPath                        `json:"e,omitempty" gorm:"index;varchar;"`
	Subnet        string                           `json:"snet" gorm:"index;varchar(36)"`
	
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

func (entity Authorization) GetEvent() (EventPath) {
	return entity.Event
}
func (entity Authorization) GetAgent() (DeviceString) {
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
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: g.TopicIds},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: g.Priviledge},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: g.Duration},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: g.Subnet},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: g.Timestamp},
	)

	return b, e
}
