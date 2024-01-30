package encoder

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
	"github.com/vmihailenco/msgpack/v5"
)
var logger = &log.Logger

func MsgPackStruct(msg interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := msgpack.NewEncoder(&buf)
	enc.SetCustomStructTag("json")
	err := enc.Encode(msg)
	return buf.Bytes(), err
}

func MsgPackUnpackStruct(b []byte, message interface{}) error {
	buf := bytes.NewBuffer(b)
	dec := msgpack.NewDecoder(buf)
	// dec.UseLooseInterfaceDecoding(true)
	dec.SetCustomStructTag("json")
	err := dec.Decode(&message)
	return err
}

func EncodeNumber(b []byte, message interface{}) error {
	buf := bytes.NewBuffer(b)
	dec := msgpack.NewDecoder(buf)
	dec.SetCustomStructTag("json")
	err := dec.Decode(&message)
	return err
}


func NumberToByte(i uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return b
}
type EncoderDataType string;
const (
	StringEncoderDataType EncoderDataType = "string"
	ByteEncoderDataType EncoderDataType = "byte"
	BigNumEncoderDataType EncoderDataType = "bignum"
	IntEncoderDataType EncoderDataType = "int"
	HexEncoderDataType EncoderDataType = "hex"
	AddressEncoderDataType EncoderDataType = "address"
)
type EncoderParam struct {
	Type EncoderDataType;
	Value interface{}

}
func EncodeBytes(args ...EncoderParam) (data []byte, err error) {
	defer func() {
        // recover from panic if one occurred. Set err to nil otherwise.
		if pan := recover(); pan != nil {
            err = errors.New(fmt.Sprintf("%v", pan))
        }
    }()
	m := make(map[int][]byte)
	var index []int
	var buffer bytes.Buffer
	for i, arg := range args {
		index = append(index, i)
		if(arg.Type == ByteEncoderDataType) {
			m[i] = arg.Value.([]byte)
		}
        if(arg.Type == StringEncoderDataType) {
			m[i] = []byte(arg.Value.(string))
		}
		if(arg.Type == IntEncoderDataType) {
			num, err := strconv.ParseUint(fmt.Sprintf("%v", arg.Value), 10, 64)
			if(err != nil) {
				return []byte(""), err
			}
			m[i] = []byte(NumberToByte(num))
		}
		if(arg.Type == BigNumEncoderDataType) {
			bigNum := new(big.Int)
			bigNum.SetString(fmt.Sprintf("%v", arg.Value), 10)
			m[i] = bigNum.Bytes()
		}
		if(arg.Type == HexEncoderDataType) {
			b, err := hex.DecodeString(arg.Value.(string))
			if err != nil {
				return []byte(""), err
			}
			m[i] = b
		}
		if(arg.Type == AddressEncoderDataType) {
			v := fmt.Sprintf("%v",arg.Value)
			if(strings.HasPrefix(v, "0x")) {
				// treat as hex
				b, err := hexutil.Decode(fmt.Sprintf("%v",arg.Value))
				if err != nil {
					return []byte(""), err
				}
				m[i] = b
			} else {
				toLower := strings.ToLower(fmt.Sprintf("%v",arg.Value))
				values := strings.Split(strings.Trim(fmt.Sprintf("%v",toLower), " "), ":")
				var addrBuffer bytes.Buffer
				addrBuffer.Write([]byte(values[0]))
				addrBuffer.Write([]byte(values[1]))
				if (len(values) == 3) {
					chain, err := strconv.ParseUint(values[2], 10, 64)
					if(err != nil) {
						return []byte(""), err
					}
					addrBuffer.Write(NumberToByte(chain))
				}
				m[i] = addrBuffer.Bytes()
			}
		}
		
    }
	
	// TODO: sort the byte slice in an efficient way so that the order of args doesnt matter
	for _, n := range index {
		buffer.Write(m[n])
	}
	// logger.Infof("LOG MEssage  =========> %v \n %v \n %v", index, buffer, args)
	return buffer.Bytes(), nil
}



