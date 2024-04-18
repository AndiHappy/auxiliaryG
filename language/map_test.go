package gotutorial

import (
	"fmt"
	"sort"
	"testing"
)

// 1
var employee = map[string]int{"Mark": 10, "Sandy": 20}

func TestMap(t *testing.T) {
	fmt.Println("----------------------//1--------------------")
	fmt.Println(employee)

	fmt.Println("----------------------//2--------------------")
	var employee = map[string]int{}
	fmt.Println(employee)        // map[]
	fmt.Printf("%T\n", employee) // map[string]int

	fmt.Println("----------------------//3--------------------")
	employee["Mark"] = 10
	employee["Sandy"] = 20
	fmt.Println(employee)

	fmt.Println("----------------------//4--------------------")
	employeeList := make(map[string]int)
	employeeList["Mark"] = 10
	employeeList["Sandy"] = 20
	fmt.Println(employeeList)

	fmt.Println("----------------------//5--------------------")
	fmt.Println(len(employeeList)) // 0

	fmt.Println("----------------------//6--------------------")
	fmt.Println(employee["Mark"])
	employee["Mark"] = 50 // Edit item
	fmt.Println(employee)
	for key, element := range employee {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	fmt.Println("----------------------//7--------------------")
	employee["Mark"] = 10
	fmt.Println(employee)
	delete(employee, "Mark")
	fmt.Println(employee)

	fmt.Println("----------------------//8--------------------")
	employee = map[string]int{"Mark": 10, "Sandy": 20,
		"Rocky": 30, "Rajiv": 40, "Kate": 50}

	// Method - I
	for k := range employee {
		delete(employee, k)
	}

	// Method - II
	employee = make(map[string]int)

	fmt.Println("----------------------//9--------------------")

}

func TestSortMap(t *testing.T) {
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}
	keys := make([]string, 0, len(unSortedMap))
	for k := range unSortedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, unSortedMap[k])
	}

	// Int slice to store values of map.
	values := make([]int, 0, len(unSortedMap))
	for _, v := range unSortedMap {
		values = append(values, v)
	}

	// Sort slice values.
	sort.Ints(values)

	// Print values of sorted Slice.
	for _, v := range values {
		fmt.Println(v)
	}
}

func TestMergeMap(t *testing.T) {
	first := map[string]int{"a": 1, "b": 2, "c": 3}
	second := map[string]int{"a": 1, "e": 5, "c": 3, "d": 4}
	for k, v := range second {
		first[k] = v
	}
	fmt.Println(first)
}
