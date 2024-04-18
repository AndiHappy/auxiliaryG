package gotutorial

import (
	"fmt"
	"reflect"
	"testing"
)

func TestVariables(t *testing.T) {
	//first Type: var name type = expression
	var i int = 10
	var s string = "Canada"
	fmt.Println(i)
	fmt.Println(s)

	//second : Variable Declaration without Initialization：
	var i2 int
	var s2 string

	i2 = 10
	s2 = "Canada"

	fmt.Println(i2)
	fmt.Println(s2)

	//third:Variable Declaration with type inference
	var i3 = 10
	var s3 = "Canada"

	fmt.Println(reflect.TypeOf(i3))
	fmt.Println(reflect.TypeOf(s3))

	//fourth：Multiple Variable Declaration:
	var name, lName string = "John", "Doe"
	m, n, o := 1, 2, 3
	item, price := "Mobile", 2000
	fmt.Println(name + lName)
	fmt.Println(m + n + o)
	fmt.Println(item, "-", price)

	//five:Short Variable Declaration
	name5 := "John Doe"
	fmt.Println(reflect.TypeOf(name5))

	// six: variable init
	var quantity float32
	fmt.Println(quantity)

	var price6 int16
	fmt.Println(price6)

	var product string
	fmt.Println(product)

	var inStock bool
	fmt.Println(inStock)

	//seven:block variables declaration
	var (
		product7  = "Mobile"
		quantity7 = 50
		price7    = 50.50
		inStock7  = true
	)
	fmt.Println(product7, price7, quantity7, inStock7)
}

// Global variable declaration
var (
	m int
	n int
)

func TestDataType(t *testing.T) {
	var x int = 1 // Integer Data Type
	var y int     //  Integer Data Type
	fmt.Println(x)
	fmt.Println(y)

	var a, b, c = 5.25, 25.25, 14.15 // Multiple float32 variable declaration
	fmt.Println(a, b, c)

	city := "Berlin"     // String variable declaration
	Country := "Germany" // Variable names are case sensitive
	fmt.Println(city)
	fmt.Println(Country) // Variable names are case sensitive

	food, drink, price := "Pizza", "Pepsi", 125 // Multiple type of variable declaration in same line
	fmt.Println(food, drink, price)
	m, n = 1, 2
	fmt.Println(m, n)
}
