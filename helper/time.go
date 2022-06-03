package helper

import (
	"strconv"
	"time"
)

func GenerateMilisTimeNow() string {
	return strconv.FormatInt(time.Now().UnixNano()/1000000, 10)
}

func GenerateMilisTimeWeek() string {
	return strconv.FormatInt((time.Now().UnixNano()/1000000)+604800000, 10)

}
