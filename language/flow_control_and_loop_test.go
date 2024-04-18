package gotutorial

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFlowControl1(t *testing.T) {
	// first ： if condition
	var s = "Japan"
	x := true
	if x {
		fmt.Println(s)
	}

	//second：
	x2 := 100
	if x2 == 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}

	x3 := rand.Intn(150)
	//third
	if x3 <= 50 {
		fmt.Println("Germany")
	} else if 50 < x3 && x3 <= 100 {
		fmt.Println("Japan")
	} else {
		fmt.Println("Canada")
	}

	//third
	if x := 100; x == 100 {
		fmt.Println("Germany")
	}
}

func TestForLoop2(t *testing.T) {
	k := 1
	for ; k <= 10; k++ {
		fmt.Println(k)
	}

	k = 1
	for k <= 10 {
		fmt.Println(k)
		k++
	}

	for k := 1; ; k++ {
		fmt.Println(k)
		if k == 10 {
			break
		}
	}

	// Example 1
	strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
	for index, element := range strDict {
		fmt.Println("Index :", index, " Element :", element)
	}

	// Example 2
	for key := range strDict {
		fmt.Println(key)
	}

	// Example 3
	for _, value := range strDict {
		fmt.Println(value)
	}

	// Example 4
	i := 5
	for {
		fmt.Println("Hello")
		if i == 10 {
			break
		}
		i++
	}
}

func TestRangeLoopOverString2(t *testing.T) {

	for range "Hello" {
		fmt.Println("Hello")
	}

	for v := range "Hello" {
		fmt.Println(v, "Hello")
	}

	for v, c := range "Hello" {
		fmt.Println(v, c, string(c), "Hello")
	}
}

func TestSwitch(t *testing.T) {
	today := time.Now()
	switch today.Day() {
	case 5:
		fmt.Println("Today is 5th. Clean your house.")
	case 10:
		fmt.Println("Today is 10th. Buy some wine.")
	case 15:
		fmt.Println("Today is 15th. Visit a doctor.")
	case 25:
		fmt.Println("Today is 25th. Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}

	var t1 int = today.Day()
	switch t1 {
	case 5, 10, 15:
		fmt.Println("Clean your house.")
	case 25, 26, 27:
		fmt.Println("Buy some food.")
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}

	switch today.Day() {
	case 5:
		fmt.Println("Clean your house.")
		fallthrough
	case 10:
		fmt.Println("Buy some wine.")
		fallthrough
	case 15:
		fmt.Println("Visit a doctor.")
		fallthrough
	case 21:
		fmt.Println("Work Harder Boy.")
		fallthrough
	case 25:
		fmt.Println("Buy some food.")
		fallthrough
	case 31:
		fmt.Println("Party tonight.")
	default:
		fmt.Println("No information available for that day.")
	}

	switch {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some wine.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}

	switch today := time.Now(); {
	case today.Day() < 5:
		fmt.Println("Clean your house.")
	case today.Day() <= 10:
		fmt.Println("Buy some wine.")
	case today.Day() > 15:
		fmt.Println("Visit a doctor.")
	case today.Day() == 25:
		fmt.Println("Buy some food.")
	default:
		fmt.Println("No information available for that day.")
	}
}
