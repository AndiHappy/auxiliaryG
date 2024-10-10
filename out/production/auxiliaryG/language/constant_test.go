package gotutorial

import (
	"fmt"
	"reflect"
	"testing"
)

const PRODUCT string = "Canada"
const PRICE = 500

const (
	PRODUCT2  = "Mobile"
	QUANTITY2 = 50
	PRICE2    = 50.50
	STOCK2    = true
)

func TestConstants(t *testing.T) {
	fmt.Println(PRODUCT)
	fmt.Println(PRICE)

	fmt.Println(reflect.TypeOf(PRODUCT))
	fmt.Println(reflect.TypeOf(PRICE))

	//block constants
	fmt.Println(PRODUCT2, QUANTITY2, PRICE2, STOCK2)

}
