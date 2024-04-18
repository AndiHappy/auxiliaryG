package gotutorial

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestStruct(t *testing.T) {
	type rectangle struct {
		length  float64
		breadth float64
		color   string
	}

	fmt.Println(rectangle{10.5, 25.10, "red"})
}

func TestStructAssign(t *testing.T) {
	type rectangle struct {
		length  int
		breadth int
		color   string

		geometry struct {
			area      int
			perimeter int
		}
	}

	var rect rectangle // dot notation
	rect.length = 10
	rect.breadth = 20
	rect.color = "Green"

	rect.geometry.area = rect.length * rect.breadth
	rect.geometry.perimeter = 2 * (rect.length + rect.breadth)

	fmt.Println(rect)
	fmt.Println("Area:\t", rect.geometry.area)
	fmt.Println("Perimeter:", rect.geometry.perimeter)
}

func TestStructAssign2(t *testing.T) {
	type rectangle struct {
		length  int
		breadth int
		color   string
	}
	fmt.Println("------------------------example 1 --------------------------")
	var rect1 = rectangle{10, 20, "Green"}
	fmt.Println(rect1)
	var rect2 = rectangle{length: 10, color: "Green"} // breadth value skipped
	fmt.Println(rect2)
	rect3 := rectangle{10, 20, "Green"}
	fmt.Println(rect3)
	rect4 := rectangle{length: 10, breadth: 20, color: "Green"}
	fmt.Println(rect4)
	rect5 := rectangle{breadth: 20, color: "Green"} // length value skipped
	fmt.Println(rect5)
}

func TestStruct3(t *testing.T) {
	fmt.Println("------------------------example 2 --------------------------")
	type rectangle struct {
		length  int
		breadth int
		color   string
	}

	rect1 := new(rectangle) // rect1 is a pointer to an instance of rectangle
	rect1.length = 10
	rect1.breadth = 20
	rect1.color = "Green"
	fmt.Println(rect1)

	var rect2 = new(rectangle) // rect2 is an instance of rectangle
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2)
}

func TestStructPoint(t *testing.T) {
	type rectangle struct {
		length  int
		breadth int
		color   string
	}
	var rect1 = &rectangle{10, 20, "Green"} // Can't skip any value
	fmt.Println(rect1)

	var rect2 = &rectangle{}
	rect2.length = 10
	rect2.color = "Red"
	fmt.Println(rect2) // breadth skipped

	var rect3 = &rectangle{}
	(*rect3).breadth = 10
	(*rect3).color = "Blue"
	fmt.Println(rect3) // length skipped
}

func TestNestedStruct(t *testing.T) {
	type Salary struct {
		Basic, HRA, TA float64
	}
	type Employee struct {
		FirstName, LastName, Email string
		Age                        int
		MonthlySalary              []Salary
	}
	e := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}
	fmt.Println(e)
}

func TestStructJson(t *testing.T) {
	type Employee struct {
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
		City      string `json:"city"`
	}

	jsonString := `
    {
        "firstname": "Rocky",
        "lastname": "Sting",
        "city": "London"
    }`

	emp1 := new(Employee)
	json.Unmarshal([]byte(jsonString), emp1)
	fmt.Println(*emp1)

	emp2 := new(Employee)
	emp2.FirstName = "Ramesh"
	emp2.LastName = "Soni"
	emp2.City = "Mumbai"
	jsonStr, _ := json.Marshal(emp2)
	fmt.Printf("%s\n", jsonStr)
}

type EmployeeStruct struct {
	Name string
	Age  int
}

func (obj *EmployeeStruct) Info() {
	if obj.Name == "" {
		obj.Name = "John Doe"
	}
	if obj.Age == 0 {
		obj.Age = 25
	}
}
func TestAssignDefaultValue(t *testing.T) {
	emp1 := EmployeeStruct{Name: "Mr. Fred"}
	emp1.Info() // 这种默认处理的方式，的确让人有点接受不能
	fmt.Println(emp1)
	emp2 := EmployeeStruct{Age: 26}
	emp2.Info()
	fmt.Println(emp2)
}

func TestFindTypeofStruct(t *testing.T) {
	type rectangle struct {
		length  float64
		breadth float64
		color   string
	}

	var rect1 = rectangle{10, 20, "Green"}
	fmt.Println(reflect.TypeOf(rect1))         // slice.rectangle
	fmt.Println(reflect.ValueOf(rect1).Kind()) // struct
	rect2 := rectangle{length: 10, breadth: 20, color: "Green"}
	fmt.Println(reflect.TypeOf(rect2))         //slice.rectangle
	fmt.Println(reflect.ValueOf(rect2).Kind()) // struct
	rect3 := new(rectangle)
	fmt.Println(reflect.TypeOf(rect3))         // *slice.rectangle
	fmt.Println(reflect.ValueOf(rect3).Kind()) // ptr 指针类型
	var rect4 = &rectangle{}
	fmt.Println(reflect.TypeOf(rect4))         // *slice.rectangle
	fmt.Println(reflect.ValueOf(rect4).Kind()) // ptr 指针类型
}

func TestCompareStruct(t *testing.T) {
	type rectangle struct {
		length  float64
		breadth float64
		color   string
	}
	var rect1 = rectangle{10, 20, "Green"}
	rect2 := rectangle{length: 20, breadth: 10, color: "Red"}
	rect33 := rectangle{length: 10, breadth: 20, color: "Green"}
	if rect1 == rect2 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	if rect1 == rect33 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	rect3 := new(rectangle)
	var rect4 = &rectangle{}
	if rect3 == rect4 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func TestCopyStruct(t *testing.T) {
	type rectangle struct {
		length  float64
		breadth float64
		color   string
	}
	r1 := rectangle{10, 20, "Green"}
	fmt.Println(r1)
	r2 := r1 // 这个竟然是一个copy操作
	r2.color = "Pink"
	fmt.Println(r2)
	fmt.Println(r1)
	r3 := &r1
	r3.color = "Red"
	fmt.Println(r3)
	fmt.Println(r1)
}
