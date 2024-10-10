package gotutorial

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestString2Int(t *testing.T) {
	strVar := "100"
	intVar, err := strconv.Atoi(strVar)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	//
	var intVar64 int64
	intVar64, err = strconv.ParseInt(strVar, 0, 8)
	fmt.Println(intVar64, err, reflect.TypeOf(intVar64))

	intVar64, err = strconv.ParseInt(strVar, 0, 16)
	fmt.Println(intVar64, err, reflect.TypeOf(intVar64))

	intVar64, err = strconv.ParseInt(strVar, 0, 32)
	fmt.Println(intVar64, err, reflect.TypeOf(intVar64))

	intVar64, err = strconv.ParseInt(strVar, 0, 64)
	fmt.Println(intVar64, err, reflect.TypeOf(intVar64))

	intValue := 0
	_, err = fmt.Sscan(strVar, &intValue)
	fmt.Println(intValue, err, reflect.TypeOf(intValue))

}
