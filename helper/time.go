package helper

import (
	"strconv"
	"time"
)

func GenerateMilisTimeNow() string {
	return strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
}
