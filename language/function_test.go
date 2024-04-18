package gotutorial

import (
	"fmt"
	"testing"
)

// Function accepting arguments
func add(x int, y int) {
	total := 0
	total = x + y
	fmt.Println(total)
}

// Function with int as return type
func add2(x int, y int) int {
	total := 0
	total = x + y
	return total
}

func simpleFunction() {
	fmt.Println("hello,function")
}

func rectangle(l int, b int) (area int) {
	var parameter int
	parameter = 2 * (l + b)
	fmt.Println("Parameter: ", parameter)
	area = l * b
	return // Return statement without specify variable name
}

func rectangle2(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}

func TestSimpleFunction(t *testing.T) {
	simpleFunction()
	add(20, 30)
	sum := add2(20, 30)
	fmt.Println(sum)
	fmt.Println("Area: ", rectangle(20, 30))
	fmt.Println(rectangle2(20, 30))

}

func update(a *int, t *string) {
	*a = *a + 5         // defrencing pointer address
	*t = *t + " update" // defrencing pointer address
	return
}

func TestUpdateParas(t *testing.T) {
	var age = 20
	var text = "John"
	fmt.Println("Before:", text, age)

	update(&age, &text)
	fmt.Println("After :", text, age)
}

var (
	area = func(l int, b int) int {
		return l * b
	}
)

func TestAnonymousFunctions(t *testing.T) {
	fmt.Println(area(20, 30))
	t1 := area(3, 4)
	fmt.Println(t1)

	//another method
	func(l int, b int) {
		fmt.Println(l * b)
	}(20, 30)

	//
	fmt.Printf(
		"100 (°F) = %.2f (°C)\n",
		func(f float64) float64 {
			return (f - 32.0) * (5.0 / 9.0)
		}(100),
	)
}

func TestClosuresFunctions(t *testing.T) {
	l := 20
	b := 30
	func() {
		var area int
		area = l * b
		fmt.Println(area)
	}()

	for i := 10.0; i < 100; i += 10.0 {
		rad := func() float64 {
			return i * 39.370
		}()
		fmt.Printf("%.2f Meter = %.2f Inch\n", i, rad)
	}
}

func sum(x, y int) int {
	return x + y
}
func partialSum(x int) func(int) int {
	return func(y int) int {
		return sum(x, y)
	}
}

func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

// ## User Defined Function Types in Golang
type First func(int) int
type Second func(int) First

func squareSum2(x int) Second {
	return func(y int) First {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func TestHigherOrderFunctions(t *testing.T) {
	partial := partialSum(3)
	fmt.Println(partial(7))

	// 5*5 + 6*6 + 7*7
	fmt.Println(squareSum(5)(6)(7))

	fmt.Println(squareSum2(5)(6)(7))
}
