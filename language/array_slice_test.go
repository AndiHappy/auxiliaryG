package gotutorial

import (
	"errors"
	"fmt"
	"github.com/AndiHappy/auxiliaryG/util"
	"reflect"
	"testing"
	"time"
)

func TestArray(t *testing.T) {
	var array2 [5]int
	var strArray [5]string
	fmt.Println(reflect.ValueOf(array2).Kind())
	fmt.Println(reflect.ValueOf(strArray).Kind())
	util.PrintPara(array2)
	var intSlice []int
	var strSlice []string

	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Println(reflect.ValueOf(strSlice).Kind())

	//只是少了一个len

	var theArray [3]string
	theArray[0] = "India"    // Assign a value to the first element
	theArray[1] = "Canada"   // Assign a value to the second element
	theArray[2] = "Japan"    // Assign a value to the third element
	fmt.Println(theArray[0]) // Access the first element value
	fmt.Println(theArray[1]) // Access the second element valu
	fmt.Println(theArray[2]) // Access the third element valu

	x := [5]int{10, 20, 30, 40, 50}   // Intialized with values
	var y [5]int = [5]int{10, 20, 30} // Partial assignment
	fmt.Println(x)
	fmt.Println(y)

	fmt.Println("\n---------------slice declaration 1--------------------")
	var intSlice2 = []int{10, 20, 30, 40}
	var strSlice2 = []string{"India", "Canada", "Japan"}
	util.PrintSliceMember(strSlice2)
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice2), cap(intSlice2))
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice2), cap(strSlice2))

	fmt.Println("\n---------------slice declaration 2--------------------")

	var intSlice1 = make([]int, 10)        // when length and capacity is same
	var strSlice1 = make([]string, 10, 20) // when length and capacity is different
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice1), cap(intSlice1))
	fmt.Println(reflect.ValueOf(intSlice1).Kind())
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice1), cap(strSlice1))
	fmt.Println(reflect.ValueOf(strSlice1).Kind())

	fmt.Println("\n---------------slice declaration 3--------------------")
	intSlice4 := new([50]int)[0:10]
	fmt.Println(reflect.ValueOf(intSlice4).Kind())
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice4), cap(intSlice4))
	fmt.Println(intSlice4)

	fmt.Println("\n---------------array declaration 1--------------------")
	x2 := [...]int{10, 20, 30}
	fmt.Println(reflect.ValueOf(x2).Kind())
	fmt.Println(len(x2))

	// 比较新颖的初始化的内容
	x3 := [5]int{1: 10, 3: 30}
	fmt.Println(x3)

	//
	array2 = [5]int{10, 20, 30, 40, 50}
	fmt.Println("\n---------------Example 1--------------------")
	for i := 0; i < len(array2); i++ {
		fmt.Println(array2[i])
	}
	fmt.Println("\n---------------Example 2--------------------")
	for index, element := range array2 {
		fmt.Println(index, "=>", element)
	}
	fmt.Println("\n---------------Example 3--------------------")
	for _, value := range array2 {
		fmt.Println(value)
	}
	j := 0
	fmt.Println("\n---------------Example 4--------------------")
	for range array2 {
		fmt.Println(array2[j])
		j++
	}

	var strSlice5 = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println("\n---------------Example 1 --------------------")
	for index, element := range strSlice5 {
		fmt.Println(index, "--", element)
	}
	fmt.Println("\n---------------Example 2 --------------------")
	for _, value := range strSlice5 {
		fmt.Println(value)
	}
	j = 0
	fmt.Println("\n---------------Example 3 --------------------")
	for range strSlice5 {
		fmt.Println(strSlice5[j])
		j++
	}
	// for loop 和Array 是一样的

	strArray1 := [3]string{"Japan", "Australia", "Germany"}
	strArray2 := strArray1  // data is passed by value
	strArray3 := &strArray1 // data is passed by refrence
	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)
	strArray1[0] = "Canada"

	fmt.Printf("strArray1: %v\n", strArray1)
	fmt.Printf("strArray2: %v\n", strArray2)
	fmt.Printf("*strArray3: %v\n", *strArray3)

	var slice1 = []string{"india", "japan", "canada"}
	var slice2 = []string{"australia", "russia"}
	slice2 = append(slice2, slice1...)

}

func itemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)
	if arr.Kind() != reflect.Array || arr.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}
	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}
	return false
}

func TestRemoveIndexInterface(t *testing.T) {
	strArray := []string{"India", "Canada", "Japan", "Germany", "Italy"}
	removeArray := RemoveIndexInterface(strArray, 2)
	fmt.Println(strArray)
	fmt.Println(removeArray)
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}
	// Keep the distinction between nil and empty slice input
	if s.IsNil() {
		return nil
	}
	ret := make([]interface{}, s.Len())
	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}
	return ret
}

func RemoveIndexInterface(s interface{}, index int) interface{} {
	arr := reflect.ValueOf(s)
	if arr.Kind() != reflect.Array && arr.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}
	slice := arr.Slice(0, index)
	fmt.Println(slice)
	return nil
}

func TestBuildIntToSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	err := BuildIntoSlice(slice)
	fmt.Println(err)
	fmt.Println(slice)
}

func BuildIntoSlice(out interface{}) error {
	if out == nil {
		return errors.New("cannot remake from <nil>")
	}

	outv := reflect.ValueOf(out)

	outt := outv.Type()
	outk := outt.Kind()

	if outk != reflect.Ptr {
		return errors.New("cannot reflect into non-pointer")
	}
	slicev := outv.Elem()
	slicet := slicev.Type()
	slicek := slicev.Kind()

	if slicek != reflect.Slice {
		return errors.New("pointer must point to a slice")
	}

	elmt := slicet.Elem()
	elmv := reflect.Indirect(reflect.New(elmt))
	fmt.Println(elmv)
	// Demo data to force in
	foo := make(map[string]interface{})
	foo["Source"] = "+2244668899"
	foo["Contents"] = "Larlarlar"
	foo["Date"] = time.Now().Format(time.RFC3339)

	//elmv.SetInt(1)
	fieldCount := elmv.NumField()
	for i := 0; i < fieldCount; i++ {
		f := elmv.Field(i)
		if f.IsValid() {
			// Get field name
			fmt.Println(elmv.Type().Field(i).Name)
			switch t := reflect.TypeOf(f.Interface()); t.Kind() {
			case reflect.String:
				tmp, ok := foo[elmv.Type().Field(i).Name].(string)
				if ok {
					f.SetString(tmp)
				}
			case reflect.Struct:
				// Go even deeper!
				fmt.Println("Got struct", t.Kind())
				fmt.Println(elmv.Type().Field(i).Name)
				// Try and set it from another type
				bar := time.Time{}
				if f.Type().AssignableTo(reflect.ValueOf(bar).Type()) {
					fmt.Println("Assignable")
					tmp, ok := foo[elmv.Type().Field(i).Name].(string)
					if tmp2, err := time.Parse(time.RFC3339, tmp); err == nil && ok {
						l := reflect.ValueOf(tmp2)
						f.Set(l)
					}
				}

			default:
				fmt.Println("Got type", t.Kind())
			}
		}
	}

	slicev.Set(reflect.Append(slicev, elmv))
	elmv = reflect.Indirect(reflect.New(elmt))
	//elmv.SetInt(2)
	slicev.Set(reflect.Append(slicev, elmv))
	return nil
}

func TestElementExitInArray(t *testing.T) {
	fmt.Println("----------------array exist----------")
	strArray := [5]string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(itemExists(strArray, "Canada"))
	fmt.Println(itemExists(strArray, "Africa"))

	fmt.Println("----------------slice exist----------")
	var strSlice = []string{"India", "Canada", "Japan", "Germany", "Italy"}
	fmt.Println(itemExists(strSlice, "Canada"))
	fmt.Println(itemExists(strSlice, "Africa"))

}

func TestSliceAdd(t *testing.T) {
	a := make([]int, 2, 5)
	a[0] = 10
	a[1] = 20
	fmt.Println("Slice A:", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))
	a = append(a, 30, 40, 50, 60, 70, 80, 90)
	fmt.Println("Slice A after appending data:", a)
	fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))
}

func TestSliceAccessAndModify(t *testing.T) {
	var intSlice = []int{10, 20, 30, 40}
	fmt.Println(intSlice[0])
	fmt.Println(intSlice[1])
	fmt.Println(intSlice[0:4])

	var strSlice = []string{"India", "Canada", "Japan"}
	fmt.Println(strSlice)
	strSlice[2] = "Germany"
	fmt.Println(strSlice)
	removeStrSlice := RemoveSliceIndex(strSlice, 1)
	fmt.Println(strSlice) //small bug
	fmt.Println(removeStrSlice)
}

func TestDeleteSliceElement(t *testing.T) {
	all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("all: ", all) //[0 1 2 3 4 5 6 7 8 9]
	removeIndex := RemoveIndex(all, 5)

	fmt.Println("all: ", all)                 //[0 1 2 3 4 6 7 8 9 9]
	fmt.Println("removeIndex: ", removeIndex) //[0 1 2 3 4 6 7 8 9]

	removeIndex[0] = 999
	fmt.Println("all: ", all)                 //[999 1 2 3 4 6 7 9 9]
	fmt.Println("removeIndex: ", removeIndex) //[999 1 2 3 4 6 7 8 9]
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

// This answer does not perform bounds-checking.
// It expects a valid index as input.
// This means that negative values or indices that are greater or equal to the initial len(s)
// will cause Go to panic.
func RemoveSliceIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

//If you do not care about ordering,
//you have the much faster possibility to replace the element
//to delete with the one at the end of the slice
//and then return the n-1 first elements:

func RemoveIndexDisOrder(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func OriginalRemoveIndex(arr []int, pos int) []int {
	newArr := make([]int, len(arr)-1)
	k := 0
	for i := 0; i < (len(arr) - 1); {
		if i != pos {
			newArr[i] = arr[k]
		}
		k++
		i++
	}
	return newArr
}

func TestDeleteSliceElementOrigin(t *testing.T) {
	all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("all: ", all) //[0 1 2 3 4 5 6 7 8 9]
	originalRemove := OriginalRemoveIndex(all, 5)

	fmt.Println("all: ", all)                       //[0 1 2 3 4 5 6 7 8 9]
	fmt.Println("originalRemove: ", originalRemove) //[0 1 2 3 4 6 7 8 9]

	originalRemove[0] = 999
	fmt.Println("all: ", all)                       //[0 1 2 3 4 5 6 7 8 9]
	fmt.Println("originalRemove: ", originalRemove) //[999 1 2 3 4 6 7 8 9]
}

func RemoveIndexCopy(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func TestRemoveIndexCopy(t *testing.T) {
	all := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("all: ", all) //[0 1 2 3 4 5 6 7 8 9]
	removeIndex := RemoveIndexCopy(all, 5)

	fmt.Println("all: ", all)                     //[0 1 2 3 4 5 6 7 8 9]
	fmt.Println("RemoveIndexCopy: ", removeIndex) //[0 1 2 3 4 6 7 8 9]

	removeIndex[0] = 999
	fmt.Println("all: ", all)                 //[0 1 2 3 4 5 6 7 8 9]
	fmt.Println("removeIndex: ", removeIndex) //[999 1 2 3 4 6 7 8 9]
}

func TestSliceCopy(t *testing.T) {
	a := []int{5, 6, 7} // Create a smaller slice
	fmt.Printf("[Slice:A] Length is %d Capacity is %d\n", len(a), cap(a))
	b := make([]int, 5, 10) // Create a bigger slice
	copy(b, a)              // Copy function
	fmt.Printf("[Slice:B] Length is %d Capacity is %d\n", len(b), cap(b))
	fmt.Println("Slice B after copying:", b)
	b[3] = 8
	b[4] = 9
	fmt.Println("Slice B after adding elements:", b)
}

func TestName(t *testing.T) {

}

func TestSlice(t *testing.T) {
	var countries = []string{"india", "japan", "canada", "australia", "russia"}
	fmt.Printf("Countries: %v\n", countries)
	fmt.Printf(":2 %v\n", countries[:2])
	fmt.Printf("1:3 %v\n", countries[1:3])
	fmt.Printf("2: %v\n", countries[2:])
	fmt.Printf("2:5 %v\n", countries[2:5])
	fmt.Printf("0:3 %v\n", countries[0:3])
	fmt.Printf("Last element: %v\n", countries[4])
	fmt.Printf("Last element: %v\n", countries[len(countries)-1])
	fmt.Printf("Last element: %v\n", countries[4:])
	fmt.Printf("All elements: %v\n", countries[0:len(countries)])
	fmt.Printf("Last two elements: %v\n", countries[3:len(countries)])
	fmt.Printf("Last two elements: %v\n", countries[len(countries)-2:len(countries)])
	fmt.Println(countries[:])
	fmt.Println(countries[0:])
	fmt.Println(countries[0:len(countries)])
}

func TestDeleteIndexCopy(t *testing.T) {
	var countries = []string{"india", "japan", "canada", "australia", "russia"}
	v := DeleteIndexCopy(countries, 2)
	vv := v.([]string)
	fmt.Println(vv)
}

func DeleteIndexCopy(s interface{}, index int) interface{} {
	arr := reflect.ValueOf(s)
	if arr.Kind() != reflect.Array && arr.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	ret := make([]interface{}, arr.Len()-1)
	for i, k := 0, 0; i < arr.Len(); i++ {
		if i != index {
			ret[k] = arr.Index(i).Interface()
			k++
		}
	}
	return ret
}

func TestDeleteIndexCopy2(t *testing.T) {
	var countries = []string{"india", "japan", "canada", "australia", "russia"}
	v := DeleteIndexCopy2(countries, 2)
	vv := v.([]string)
	fmt.Println(vv)
}

func DeleteIndexCopy2(s interface{}, index int) interface{} {
	arr := reflect.ValueOf(s)
	if arr.Kind() != reflect.Array && arr.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}
	if arr.Len() == 0 {
		return s
	} else {
		ret := make([]interface{}, arr.Len()-1)
		for i, k := 0, 0; i < arr.Len(); i++ {
			if i != index {
				ret[k] = arr.Index(i).Interface()
				k++
			}
		}
		switch arr.Index(0).Kind() {
		case reflect.Int:
			return retToIntArray(ret)
		case reflect.String:
			return retToStringArray(ret)
		case reflect.Bool:
			return retToBoolArray(ret)
		}
		return ret
	}
}

func retToBoolArray(ret []interface{}) interface{} {
	r := make([]bool, len(ret))
	for i := 0; i < len(ret); i++ {
		r[i] = ret[i].(bool)
	}
	return r
}

func retToStringArray(ret []interface{}) interface{} {
	r := make([]string, len(ret))
	for i := 0; i < len(ret); i++ {
		r[i] = ret[i].(string)
	}
	return r
}

func retToIntArray(ret []interface{}) interface{} {
	r := make([]int, len(ret))
	for i := 0; i < len(ret); i++ {
		r[i] = ret[i].(int)
	}
	return r
}
