package util

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint interface {
	~uint | ~uint8 | ~uint16 | ~uint32
}

type Float interface {
	~float32 | ~float64
}

type Element interface {
	Int | Uint | Float | ~string | interface{} | any
}

// IsFull  Check if channel is full
func IsFull[T Element | chan T](channel chan T) bool {
	return len(channel) == cap(channel)
}
