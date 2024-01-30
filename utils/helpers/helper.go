package helpers

import (
	"time"
)


func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func TimestampMilli () uint64 {
	return uint64(time.Now().UnixNano() / 1e6)
}



