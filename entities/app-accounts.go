package entities

import (
	// "errors"

	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"go.uber.org/zap/buffer"

	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/common/utils"
)




type AppAccount struct {
	Version float32 `json:"_v"`
	ID            string        `json:"id"`
	Timestamp     uint64        `json:"ts,omitempty" binding:"required"`
	Balance       uint64        `json:"bal" gorm:"default:0"`
	AppID	 string        `json:"app" `
	
	// Readonly
	
	// Derived
	Event EventPath `json:"e,omitempty" gorm:"index;varchar;"`
	PreviousEvent EventPath `json:"pE,omitempty" gorm:"index;varchar;"`
	Account AccountString    `json:"acct,omitempty" binding:"required"  gorm:"not null;type:varchar(100)"`
	SignatureData SignatureData `json:"sigD,omitempty"`
	Hash  string    `json:"h,omitempty" gorm:"type:char(64)"`
	BlockNumber uint64          `json:"blk"`
	Cycle   	uint64			`json:"cy"`
	Epoch		uint64			`json:"ep"`
	

	//Deprecated
	EventSignature  string    `json:"csig,omitempty"`
	ZKData *ZK `json:"zkD,omitempty"`

}

func (d AppAccount) GetKey() (string) {
	acc, _ := hex.DecodeString(string(d.Account))
	
	buff := buffer.Buffer{}
	buff.Write(utils.UuidToBytes(d.AppID))
	buff.Write(acc)
	hash := crypto.Sha256(buf.Bytes())
	return hex.EncodeToString(hash)
	// return ""
}  

func (g *AppAccount) GetDataStoreKeys() (keys []string)  {
	if g.ID == "" { 
		g.ID, _ = GetId(g, "")
	}
	// keys = append(keys, fmt.Sprintf("%s/%s/%s", g.AccountApplicationsKey(), utils.IntMilliToTimestampString(int64(g.Timestamp)), g.ID))
	keys = append(keys, g.Key())
	
	// keys = append(keys, fmt.Sprintf("%s/%d/%s", ApplicationModel, g.Cycle, g.ID))
	// keys = append(keys,fmt.Sprintf("%s/%s/%s", g.Event.ID, ApplicationModel, g.Hash ))
	keys = append(keys, g.DataKey())
	keys = append(keys, g.ArchiveKey())
	// keys = append(keys, g.GetEventStateKey())
	// keys = append(keys, fmt.Sprintf("%s/%d/%s", AuthModel, g.Cycle, g.ID))
	return keys;
}
// func (g Application) GetEventStateKey() (string) {
//    return fmt.Sprintf("ev/%s", g.Event.ToString() )
// }
func (item *AppAccount) DataKey() string {
	item.ID, _ = GetId(item, "")
	
	return fmt.Sprintf(DataKey,  GetModel(item), item.ID )
}

func (item *AppAccount) ArchiveKey() string {
	return fmt.Sprintf("/arc/%010d/%s", item.Cycle, item.Hash )
}


func (g *AppAccount) Key() string {
	if g.ID == "" {
		g.ID, _ = GetId(g, "")
	}
	return fmt.Sprintf("/%s/id/%s", GetModel(g), g.ID)
}







func (item *AppAccount) ToJSON() []byte {
	m, e := json.Marshal(item)
	if e != nil {
		logger.Errorf("Unable to parse subscription to []byte")
	}
	return m
}

func (item *AppAccount) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(item)
	return b
}



func UnpackAppAccount(b []byte) (AppAccount, error) {
	var item AppAccount
	err := encoder.MsgPackUnpackStruct(b, &item)
	return item, err
}





func (p AppAccount) GetId() string {
	return p.ID
}

func (item AppAccount) GetHash() ([]byte, error) {
	if item.Hash != "" {
		return hex.DecodeString(item.Hash)
	}
	b, err := item.EncodeBytes()
	if err != nil {
		return []byte(""), err
	}
	// logger.Debugf("GetHash crypto.Sha256(b) : %v", crypto.Sha256(b))
	return crypto.Sha256(b), nil
}

func (item AppAccount) ToString() (string, error) {
	values := []string{}
	values = append(values, item.AppID)
	values = append(values, string(item.Account))
	return strings.Join(values, ","), nil
}

func (entity AppAccount) GetEvent() EventPath {
	return entity.Event
}
func (entity AppAccount) GetAppID() DeviceString {
	return DeviceString(entity.AppID)
}

func (item AppAccount) EncodeBytes() ([]byte, error) {
	acc, _ := hex.DecodeString(string(item.Account))
	return encoder.EncodeBytes(
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: utils.UuidToBytes(item.AppID)},
		encoder.EncoderParam{Type: encoder.ByteEncoderDataType, Value: acc},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: item.Balance},
		encoder.EncoderParam{Type: encoder.IntEncoderDataType, Value: item.Timestamp},
	)
}
