package entities

import (
	// "errors"

	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	utils "github.com/mlayerprotocol/go-mlayer/utils"

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

func (address *Address) Pack() []byte {
	b, _ := utils.MsgPackStruct(address)
	return b
}



func AddressFromBytes(b []byte) (Address, error) {
	var address Address
	err := json.Unmarshal(b, &address)
	return address, err
}
func UnpackAddress(b []byte) (Address, error) {
	var address Address
	err := utils.MsgPackUnpackStruct(b, address)
	return address, err
}

func (address *Address) Hash() string {
	return hexutil.Encode(utils.Hash(address.ToString()))
}

func (address *Address) ToString() string {
	values := []string{}
	values = append(values, fmt.Sprintf("%s", address.Addr))
	values = append(values, fmt.Sprintf("%s", address.Platform))
	if (address.Chain != 0) {
		values = append(values, fmt.Sprintf("%d", address.Chain))
	}
	return strings.Join(values, ":")
}

func AddressFromString(s string) (Address, error) {
	
	values := strings.Split(strings.Trim(s, " "), ":")
	
	if(len(values) == 2) { 
		return Address{Addr: values[0], Platform: values[1]}, nil
	}
	chain, err := strconv.Atoi(values[2])
	if(err != nil) {
		return Address{}, err
	}
	return Address{Addr: values[0], Platform: values[1], Chain: uint64(chain)}, nil
}

