package utils

import (
	"regexp"
	"time"

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

func IsAlphaNumericDot(str string) bool {
	pattern := `^[a-zA-Z][a-zA-Z0-9._]*$`
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
