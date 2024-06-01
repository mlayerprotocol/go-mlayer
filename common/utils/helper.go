package utils

import (
	"encoding/json"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mlayerprotocol/go-mlayer/pkg/log"
)

var logger = &log.Logger

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func TimestampMilli() uint64 {
	return uint64(time.Now().UnixNano() / 1e6)
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
