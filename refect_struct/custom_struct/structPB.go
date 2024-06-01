package custom_struct

type Foo struct {
	x int
	y *Foo
	v any
}

func NewFoo(x, y int) Foo {
	return Foo{x: x, y: &Foo{x: y}}
}
