package utils

import (
	"fmt"
	"time"
)

func CurrentTimeInUint() {
	utc := time.Now().Unix()
	fmt.Println(utc)
}
