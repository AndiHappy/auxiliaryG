package auxiliaryGo

import (
	"fmt"
	"reflect"
)

const (
	nilAngleString = "<nil>"
)

type AnyStructMethod struct {
	// arg holds the current item, as an interface{}.
	arg any

	// value is used instead of arg for reflect values.
	value reflect.Value
}

func (r AnyStructMethod) FP(a ...any) {
	r.DoPrintln(a)
}

func (r AnyStructMethod) DoPrintln(a []any) {
	for argNum, arg := range a {
		if argNum > 0 {
			r.writeByte(' ')
		}
		r.printArg(arg, 'v')
	}
	r.writeByte('\n')
}

func (r AnyStructMethod) writeByte(i int32) {
	fmt.Print(i)
}

func (r AnyStructMethod) printArg(arg any, verb int32) {
	r.arg = arg
	r.value = reflect.Value{}
	//当传入的是nil的时候，输出什么东西
	if arg == nil {
		switch verb {
		case 'T', 'v':
			fmt.Println(nilAngleString)
		}
		return
	}

	// Special processing considerations.
	// %T (the value's type) and %p (its address) are special; we always do them first.
	switch verb {
	case 'T':
		fmt.Println("the value's type")
	case 'P':
		fmt.Println("the address's type")
	}

	// Some types can be done without reflection.
	switch f := arg.(type) {
	case bool, float32, float64, complex64, complex128, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, string, []byte:
		fmt.Println(arg)
		return
	case reflect.Value:
		// Handle extractable values with special methods
		// since printValue does not handle them at depth 0.
		if f.IsValid() && f.CanInterface() {
			r.arg = f.Interface()
			if r.handleMethods(verb) {
				return
			}
		}
		r.printValue(f, verb, 0)
	default:
		// If the type is not simple, it might have methods.
		if !r.handleMethods(verb) {
			// Need to use reflection, since the type had no
			// interface methods that could be used for formatting.
			r.printValue(reflect.ValueOf(f), verb, 0)
		}
	}

}

func (r AnyStructMethod) handleMethods(verb int32) bool {
	return true
}

func (r AnyStructMethod) printValue(f reflect.Value, verb int32, i int) {

}

func FP(a ...any) {
	anyObject := AnyStructMethod{}
	anyObject.DoPrintln(a) // a 从 ...any 到 []any
}

func main() {
	FP()
}
