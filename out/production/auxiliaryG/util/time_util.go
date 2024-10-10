package util

import (
	"time"
)

func NowString() string {
	return time.Now().Format(time.StampNano)
}
