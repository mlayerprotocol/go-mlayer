package utils

import (
	"bytes"

	"github.com/vmihailenco/msgpack/v5"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

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
	dec.SetCustomStructTag("json")
	err := dec.Decode(&message)
	return err
}