package entities

import (
	"encoding/json"
	"fmt"

	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/utils/constants"
	"github.com/mlayerprotocol/go-mlayer/utils/encoder"
)



type Authorization struct {
	ID          	string `gorm:"primaryKey"  json:"ID,omitempty"`
	Agent string    `json:"agt" gorm:"index:idx_agent_authorization,unique"`
	Account string    `json:"acct" gorm:"index:idx_agent_authorization,unique"`
	Grantor string    `json:"gr" gorm:"index"`
	Priviledge constants.Priviledge    `json:"privi"`
	TopicIds string    `json:"topIds"`
	Timestamp uint64    `json:"ts"`
	Duration uint64    `json:"du"`
	Signature   string    `json:"sig"`
	Hash string		`json:"h" gorm:"unique" `
	EventHash string `json:"eventHash" gorm:"index"`
	AuthorizationEventID string `json:"authEventId"`
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
		encoder.EncoderParam{Type: encoder.AddressEncoderDataType, Value: g.Account},
		encoder.EncoderParam{Type: encoder.AddressEncoderDataType, Value: g.Agent},
		encoder.EncoderParam{Type: encoder.StringEncoderDataType, Value: g.TopicIds},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: g.Priviledge},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: g.Duration},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: g.Timestamp},
	)
	
	return b,e
}
func (g Authorization) Validate(privateKey string) error {
	
}

// type AuthorizationPayload  struct {
// 	ClientPayload `msgpack:",noinline"`
// }
// func (msg AuthorizationPayload) GetHash() ([]byte, error) {
// 	b, err := msg.EncodeBytes()
// 	if err != nil {
// 		return []byte(""), err
// 	}
// 	bs := crypto.Keccak256Hash(b)
// 	return bs, nil
// }
// func (msg AuthorizationPayload) EncodeBytes() ([]byte, error) {
	
	
// 	b, _ := msg.Data.EncodeBytes()

// 	var params []encoder.EncoderParam
// 	params = append(params, encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: b})
// 	params = append(params, encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: msg.EventType})
// 	if msg.Grantor != "" {
// 		params = append(params, encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: msg.Grantor})
// 	}

// 	params = append(params, encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: msg.Timestamp})
	
// 	return encoder.EncodeBytes(
// 		params...
// 	)
// }

// func (msg AuthorizationPayload) EncodeBytes() ([]byte, error) {
	
// 	b, _ := (msg.Data).EncodeBytes()
	
	
// 	var params []encoder.EncoderParam
// 	params = append(params, encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: b})
// 	params = append(params, encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: msg.EventType})
// 	if msg.Grantor != "" {
// 		params = append(params, encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: msg.Grantor})
// 	}
// 	if msg.Node != "" {
// 		params = append(params, encoder.EncoderParam{Type: encoder.HexEncoderDataType, Value: msg.Node})
// 	}
// 	params = append(params, encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: msg.Timestamp})
	
// 	return encoder.EncodeBytes(
// 		params...
// 	)
// }

// type AuthorizationEvent  struct {
// 	Event
// 	// Data AuthorizationPayload
// 	// EventInterface
// }



// func (g AuthorizationEvent) AsEvent () Event {
// 	return g.Event
// }
