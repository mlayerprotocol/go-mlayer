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

// func ParseQuery[Entity any](c *gin.Context, entitie *Entity) (error) {
// 	rawQuery := c.Request.URL.Query()
// 		var query map[string]string = map[string]string{}
// 		for key, v := range rawQuery {
// 			if len(v) > 0 {
// 				query[key] = v[0]
// 			}

// 		}
// 		b, requestErr := json.Marshal(query)
// 		if requestErr != nil {
// 			return requestErr
// 		}
// 		// for _, e := range entities {
// 			json.Unmarshal(b, entitie)
// 		// }
// 		return nil
// }
func ParseQueryString(c *gin.Context) (*[]byte, error) {
	rawQuery := c.Request.URL.Query()
		var query map[string]string = map[string]string{}
		for key, v := range rawQuery {
			if len(v) > 0 {
				query[key] = v[0]
			}

		}
		b, requestErr := json.Marshal(query)
		if requestErr != nil {
			return nil, requestErr
		}
		return &b, nil
}