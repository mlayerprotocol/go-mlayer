package entities

import (
	// "errors"

	"encoding/json"
	"fmt"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/common/encoder"
	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
)


type PublicKeyString string
type AddressString string

type Address struct {
	Prefix   string    `json:"pre"`
	Addr   string    `json:"addr"`
	// Platform string    `json:"p"`
	Chain  string    `json:"ch"`
}

func (address *Address) ToJSON() []byte {
	m, e := json.Marshal(address)
	if e != nil {
		logger.Errorf("Unable to parse address to []byte")
	}
	return m
}

func (address *Address) MsgPack() []byte {
	b, _ := encoder.MsgPackStruct(address)
	return b
}

func AddressFromBytes(b []byte) (Address, error) {
	var address Address
	err := json.Unmarshal(b, &address)
	return address, err
}
func MsgUnpack(b []byte) (Address, error) {
	var address Address
	err := encoder.MsgPackUnpackStruct(b, address)
	return address, err
}

func (address *Address) GetHash() []byte {
	return crypto.Keccak256Hash(address.ToBytes())
}

func (address *Address) ToString() string {
	values := []string{}
	values = append(values, address.Prefix)
	values = append(values, ":")
	values = append(values, address.Addr)
	if (address.Chain != "") {
		values = append(values, fmt.Sprintf("#%s", address.Chain))
	}
	return strings.Join(values, ":")
}

func (address *Address) ToBytes() []byte {
	// var buffer bytes.Buffer
	// // buffer.Write([]byte(address.Platform))
	// buffer.Write([]byte("did:"))
	// if strings.HasPrefix(address.Addr, "0x") {
	// 	// treat as hex
	// 	b, err := hexutil.Decode(faddress.Addr)
	// 	if err != nil {
	// 		return []byte(""), err
	// 	}
	// 	buffer.Write(b)
	// } else {
	// buffer.Write([]byte(address.Addr))
	// if(address.Chain != "") {
	// 	// binary.Write(&buffer, binary.BigEndian, address.Chain)
	// 	buffer.Write([]byte(fmt.Sprintf("#%s", address.Chain)))
	// }
	return []byte(address.ToString())
}

func AddressFromString(s AddressString) (Address, error) {
	addr := Address{Prefix: "did"}
	values := strings.Split(strings.Trim(string(s), " "), ":")
	

	
	
	if(len(values) == 2) { 
		addr.Addr = values[1]
		addr.Prefix = values[0]
	}
	values2 := strings.Split(addr.Addr, "#")
	if (len(values2) > 1) {
		addr.Addr = values2[0]
		addr.Chain = values2[1]
	}
	
	
	return addr, nil
	
	//return Address{Addr: values[0], Prefix: "", Chain: uint64(chain)}, nil
}