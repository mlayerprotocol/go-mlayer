package utils

import (
	"encoding/json"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
)

var logger = &log.Logger

func lcg(seed uint64)  (uint64) {
    a := uint64(1664525)
    c := uint64(1013904223)
    m := uint64(1 << 32)
    return ((a * seed + c) % m)
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
	str, _ := gonanoid.Generate("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-_", length)
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
	logger.Info("condition:: ", condition)
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
