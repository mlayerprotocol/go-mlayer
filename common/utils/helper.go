package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"os"
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

func To256Bits(b []byte) []byte {
	const size = 32
	padded := make([]byte, size)
	copy(padded[size-len(b):], b)
	return padded
}
func  ToUint256(num *big.Int) []byte {
	return num.FillBytes(make([]byte, 32))
}
func  Uint64ToUint256(num uint64) []byte {
	return new(big.Int).SetUint64(num).FillBytes(make([]byte, 32))
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


func UuidToBytes(uuid string) ([]byte) {
	b, _ :=  hex.DecodeString(strings.ReplaceAll(uuid, "-", ""))
	return b
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

func AddressToHex(addr string) string {
 return strings.ToLower(strings.Replace(addr, "0x", "", 1))
}
func RandomAplhaNumString(length int) string {
	var str string;
	var err error;
	for {
		str, err = gonanoid.Generate("abcdefghijklmnopqrstuvwxyz123456789", length)
		if err == nil {
			break
		}
	}
	return str
}
func RandomHexString(length int) string {
	var str string;
	var err error;
	for {
		str, err = gonanoid.Generate("abcdef1234567890", length)
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
	logger.Debug("rawQuery:: ", rawQuery)
	var query map[string]any = map[string]any{}
	for key, v := range rawQuery {
		if len(v) > 0 {
			query[key] = v[0]
		}

	}
	logger.Debug("query:: ", query)
	b, requestErr := json.Marshal(query)
	if requestErr != nil {
		return nil, requestErr
	}
	return &b, nil
}

func ReadJSONFromFile(filePath string) (map[string]interface{}, error) {
	// file, err := os.Open(filePath)
	// if err != nil {
	// 	return nil, err
	// }
	// defer file.Close()

	// Read the entire file content
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON into a map
	var data map[string]interface{}
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// writeJSON writes the given map as JSON to a file.
func WriteJSONToFile(filePath string, data map[string]interface{}) error {
	// Marshal the map into JSON bytes
	
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	// Write the JSON data back to the file
	return os.WriteFile(filePath, bytes, 0644)
}

func CompressToGzip(inputText string) ([]byte, error) {
	// Create a buffer to hold the compressed data
	var compressedData bytes.Buffer

	// Create a new Gzip writer using the buffer as the output
	gzipWriter := gzip.NewWriter(&compressedData)
	defer gzipWriter.Close() // Ensure the writer is closed to finish writing

	// Stream the input text into the Gzip writer
	_, err := io.WriteString(gzipWriter, inputText)
	if err != nil {
		return nil, fmt.Errorf("error writing to gzip writer: %w", err)
	}

	// Close the Gzip writer to ensure all data is flushed
	if err := gzipWriter.Close(); err != nil {
		return nil, fmt.Errorf("error closing gzip writer: %w", err)
	}

	// Return the compressed data as a byte slice
	return compressedData.Bytes(), nil
}

func DecompressGzip(compressedData []byte) ([]byte, error) {
	byteReader := bytes.NewReader(compressedData)
	gzipReader, err := gzip.NewReader(byteReader)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	// Decompress the data into a buffer
	var result bytes.Buffer
	_, err = io.Copy(&result, gzipReader)
	if err != nil {
		return result.Bytes(), nil
	}
	return nil, err
}