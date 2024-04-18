package gotutorial

import (
	"fmt"
	"sync"
	"testing"
)

// Employee is an interface for printing employee details
type Employee interface {
	PrintName(name string)
	PrintSalary(basic int, tax int) int
}

// Emp user-defined type
type Emp int

// PrintName method to print employee name
func (e Emp) PrintName(name string) {
	fmt.Println("Employee Id:\t", e)
	fmt.Println("Employee Name:\t", name)
}

// PrintSalary method to calculate employee salary
func (e Emp) PrintSalary(basic int, tax int) int {
	var salary = (basic * tax) / 100
	return basic - salary
}

func TestInterface(t *testing.T) {
	var e1 Employee
	e1 = Emp(1)
	e1.PrintName("John Doe")
	fmt.Println("Employee Salary:", e1.PrintSalary(25000, 5))
}

type Polygons interface {
	Perimeter()
}
type Object interface {
	NumberOfSide()
}
type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println("Perimeter of Pentagon", 5*p)
}
func (p Pentagon) NumberOfSide() {
	fmt.Println("Pentagon has 5 sides")
}

func TestMultipleInterface(t *testing.T) {
	var p Polygons = Pentagon(50)
	p.Perimeter()
	var o Pentagon = p.(Pentagon)
	o.NumberOfSide()
	var obj Object = Pentagon(50)
	obj.NumberOfSide()
	var pent Pentagon = obj.(Pentagon)
	pent.Perimeter()
}

type Vehicle interface {
	Structure() []string // Common Method
	Speed() string
}
type Car string

func (c Car) Structure() []string {
	var parts = []string{"ECU", "Engine", "Air Filters", "Wipers", "Gas Task"}
	return parts
}
func (c Car) Speed() string {
	return "200 Km/Hrs"
}

type Human interface {
	Structure() []string // Common Method
	Performance() string
}
type Man string

func (m Man) Structure() []string {
	var parts = []string{"Brain", "Heart", "Nose", "Eyelashes", "Stomach"}
	return parts
}
func (m Man) Performance() string {
	return "8 Hrs/Day"
}

func TestInterfaceSet(t *testing.T) {
	var bmw Vehicle
	bmw = Car("World Top Brand")
	var labour Human
	labour = Man("Software Developer")
	for i, j := range bmw.Structure() {
		fmt.Printf("%-15s <=====> %15s\n", j, labour.Structure()[i])
	}
}

type Book struct {
	author, title string
}
type Magazine struct {
	title string
	issue int
}

func (b *Book) Assign(n, t string) {
	b.author = n
	b.title = t
}
func (b *Book) Print() {
	fmt.Printf("Author: %s, Title: %s\n", b.author, b.title)
}
func (m *Magazine) Assign(t string, i int) {
	m.title = t
	m.issue = i
}
func (m *Magazine) Print() {
	fmt.Printf("Title: %s, Issue: %d\n", m.title, m.issue)
}

type Printer interface {
	Print()
}

func TestInterfaceReceivePointer(t *testing.T) {
	var b Book                                 // Declare instance of Book
	var m Magazine                             // Declare instance of Magazine
	b.Assign("Jack Rabbit", "Book of Rabbits") // Assign values to b via method
	m.Assign("Rabbit Weekly", 26)              // Assign values to m via method
	var i Printer                              // Declare variable of interface type
	fmt.Println("Call interface")
	i = &b    // Method has pointer receiver, interface does not
	i.Print() // Show book values via the interface
	b.Print()

	i = &m    // Magazine also satisfies shower interface
	i.Print() // Show magazine values via the interface
	m.Print()
}

func printType(i interface{}) {
	fmt.Println(i)
}
func TestEmptyInterface(t *testing.T) {
	var manyType interface{}
	manyType = 100
	fmt.Println(manyType)
	manyType = 200.50
	fmt.Println(manyType)
	manyType = "Germany"
	fmt.Println(manyType)
	printType("Go programming language")
	var countries = []string{"india", "japan", "canada", "australia", "russia"}
	printType(countries)
	var employee = map[string]int{"Mark": 10, "Sandy": 20}
	printType(employee)
	country := [3]string{"Japan", "Australia", "Germany"}
	printType(country)
}

// Geometry is an interface that defines Geometrical Calculation
type Geometry interface {
	Edges() int
}

// Hexagon defines a geometrical object
type Hexagon struct{}

// Octagon defines a geometrical object
type Octagon struct{}

// Decagon defines a geometrical object
type Decagon struct{}

// Edges implements the Geometry interface
func (p Pentagon) Edges() int { return 5 }

// Edges implements the Geometry interface
func (h Hexagon) Edges() int { return 6 }

// Edges implements the Geometry interface
func (o Octagon) Edges() int { return 8 }

// Edges implements the Geometry interface
func (d Decagon) Edges() int { return 10 }

// Parameter calculate parameter of object
func Parameter(geo Geometry, value int) int {
	num := geo.Edges()
	calculation := num * value
	return calculation
}
func TestPolymorphism(t *testing.T) {
	h := new(Hexagon)
	o := new(Octagon)
	d := new(Decagon)
	g := [...]Geometry{h, o, d}
	for _, i := range g {
		fmt.Println(Parameter(i, 5))
	}
}

type Polygons2 interface {
	Geometry // Interface embedding another interface
}

func TestPolymorphism2(t *testing.T) {
	h := new(Hexagon)
	o := new(Octagon)
	d := new(Decagon)
	polygons := [...]Polygons2{h, o, d}
	for i := range polygons {
		fmt.Println(polygons[i].Edges())
	}
}

// 继承member的方法
func TestMemberInterfaceMethod(t *testing.T) {
	type Barber struct {
		name string
		sync.Mutex
		state int // Sleeping/Checking/Cutting
	}

	barber := Barber{
		name: "test",
	}
	barber.Lock()
	fmt.Println(barber)
	barber.Unlock()
}
