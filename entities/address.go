package entities

import (
	// "errors"

	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/mlayerprotocol/go-mlayer/internal/crypto"
	"github.com/mlayerprotocol/go-mlayer/utils/encoder"

	"github.com/ethereum/go-ethereum/common/hexutil"
)


type AddressString string

type Address struct {
	Addr   string    `json:"addr"`
	Platform string    `json:"p"`
	Chain  uint64    `json:"ch"`
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
	values = append(values, fmt.Sprintf("%s", address.Platform))
	values = append(values, fmt.Sprintf("%s", address.Addr))
	if (address.Chain != 0) {
		values = append(values, fmt.Sprintf("%d", address.Chain))
	}
	return strings.Join(values, ":")
}

func (address *Address) ToBytes() []byte {
	var buffer bytes.Buffer
	buffer.Write([]byte(address.Platform))
	buffer.Write(hexutil.MustDecode(address.Addr))
	if(address.Chain != 0) {
		binary.Write(&buffer, binary.BigEndian, address.Chain)
	}
	return buffer.Bytes()
}

func AddressFromString(s AddressString) (Address, error) {
	
	values := strings.Split(strings.Trim(string(s), " "), ":")
	
	if(len(values) == 2) { 
		return Address{Addr: values[1], Platform: values[0]}, nil
	}
	chain, err := strconv.ParseUint(values[2], 10, 64)
	if(err != nil) {
		return Address{}, err
	}
	return Address{Addr: values[0], Platform: values[1], Chain: uint64(chain)}, nil
}

