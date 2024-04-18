package gotutorial_test

import (
	"fmt"
	"testing"
)

func variadicExample(s ...string) {
	if len(s) > 4 {
		fmt.Println(s[0])
		fmt.Println(s[3])
	} else {
		for ss, vv := range s {
			fmt.Println(ss, vv)
		}
	}

}

func variadicExample2(s ...string) {
	fmt.Println(s)
}

func calculation(str string, y ...int) int {
	area := 1
	for _, val := range y {
		if str == "Rectangle" {
			area *= val
		} else if str == "Square" {
			area = val * val
		}
	}
	return area
}

func TestVariablesFunction(t *testing.T) {
	variadicExample("hell0", "world", "ddd")

	variadicExample2()
	variadicExample2("red", "blue")
	variadicExample2("red", "blue", "green")
	variadicExample2("red", "blue", "green", "yellow")

	fmt.Println(calculation("Rectangle", 20, 30))
	fmt.Println(calculation("Square", 20))
}
