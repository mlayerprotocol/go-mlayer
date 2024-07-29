package utils

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
)

var logger = &log.Logger
type errorStruct struct {
	EmptyFieldInStruct bool
}
func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.String() == ""
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Slice, reflect.Map, reflect.Array:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	}
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
func EnsureNotEmpty(s interface{}) any {
	empty := true
    v := reflect.ValueOf(s)
	typ := v.Type()
	if typ.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if !v.IsValid() {
		return errorStruct{EmptyFieldInStruct: true}
	}
	if v.NumField() == 0 {
		return errorStruct{EmptyFieldInStruct: true}
	}
	
    for i := 0; i < v.NumField(); i++ {
        field := v.Field(i)
		if !isZeroValue(field) {
			empty = false
		}
    }
	if empty {
		return errorStruct{EmptyFieldInStruct: true}
	}
    return s
}
func CheckEmpty(s interface{}) error {
	resp := EnsureNotEmpty(s)
	switch val := resp.(type) {
		case errorStruct:
			logger.Debugf("CheckEmpty: %v", val)
			return fmt.Errorf("is empty %s", "")
	}
	return nil
}

func CopyStructValues(src, dst interface{}) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst).Elem()

	if srcVal.Kind() != reflect.Struct || dstVal.Kind() != reflect.Struct {
		return fmt.Errorf("input types must be structs")
	}

	for i := 0; i < srcVal.NumField(); i++ {
		fieldName := srcVal.Type().Field(i).Name
		srcField := srcVal.Field(i)
		dstField := dstVal.FieldByName(fieldName)

		if dstField.IsValid() && dstField.CanSet() && srcField.Type() == dstField.Type() {
			dstVal.FieldByName(fieldName).Set(srcField)
		}
	}

	return nil
}

func StructToMap(input interface{}) map[string]interface{} {
	output := make(map[string]interface{})

	// val := reflect.ValueOf(input)
	// typ := reflect.TypeOf(input)

	// for i := 0; i < val.NumField(); i++ {
	// 	field := typ.Field(i)
	// 	value := val.Field(i).Interface()
	// 	output[field.Name] = value
	// }

	d, _ := json.Marshal(input)
	json.Unmarshal(d, &output)

	return output
}

func PadTo256Bits(b []byte) []byte {
	const size = 32
	padded := make([]byte, size)
	copy(padded[size-len(b):], b)
	return padded
}


func Lcg(seed uint64)  (*big.Int) {
    a := uint64(1664525)
    c := uint64(1013904223)
    m := uint64(1 << 32)
	r := (a * seed + c) % m
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, r)
    return new(big.Int).SetBytes(b)
}


func UuidToBytes(uuid string) ([]byte, error) {
	return hex.DecodeString(strings.ReplaceAll(uuid, "-", ""))
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
func RandomString(length int) string {
	var str string;
	var err error;
	for {
		str, err = gonanoid.Generate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_", length)
		if err == nil {
			break
		}
	}
	return str
}
func TimestampMilli() uint64 {
	return uint64(time.Now().UnixNano() / 1e6)
}

func IsNumericInt(s string) bool {
    _, err := strconv.Atoi(s)
    return err == nil
}

func IsAlphaNumericDotNoNumberPrefix(str string) bool {
	pattern := `^[a-zA-Z][a-zA-Z0-9._]*$`
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}

func IsAlphaNumericDot(str string) bool {
	pattern := `^[a-zA-Z0-9_.]+$`
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}

func IsAlphaLowerNumericDot(str string) bool {
	pattern := `^[a-z0-9_.]+$`
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}

func IsDomain(str string) bool {
	pattern := `[[A-Za-z0-9](?:[A-Za-z0-9\-]{0,61}[A-Za-z0-9])?`
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}

func IfThenElse[T any](condition bool, a T, b T) T {
	if condition {
		return a
	}
	return b
}
func SafePointerValue[T any](b *T, defaultValue T) T {
	if b == nil {
		return defaultValue
	}
	return *b
}

func ParseQueryString(c *gin.Context) (*[]byte, error) {
	rawQuery := c.Request.URL.Query()
	logger.Info("rawQuery:: ", rawQuery)
	var query map[string]any = map[string]any{}
	for key, v := range rawQuery {
		if len(v) > 0 {
			query[key] = v[0]
		}

	}
	logger.Info("query:: ", query)
	b, requestErr := json.Marshal(query)
	if requestErr != nil {
		return nil, requestErr
	}
	return &b, nil
}
