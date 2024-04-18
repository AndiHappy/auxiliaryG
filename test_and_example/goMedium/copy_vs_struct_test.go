package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"testing"
)

// https://medium.com/a-journey-with-go/archive/2019/05
//Go: Should I Use a Pointer instead of a Copy of my Struct?

//For many Go developers,
//the systematic use of pointers to share structs instead of the copy itself
//seems the best option in terms of performance.

//In order to understand the impact of using a pointer rather than a copy of the struct,
//we will review two use cases.

//Intensive allocation of data
//Let’s take a simple example of when you want to share a struct for its values:

type S struct {
	a, b, c int64
	d, e, f string
	g, h, i float64
}

// Here is a basic struct that can be shared by copy or by pointer:
func byCopy() S {
	return S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func byPointer() *S {
	return &S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

// Based on those 2 methods, we can now write 2 benchmarks,
// one where the struct is passed by copy:
// go test ./... -bench=BenchmarkMemoryStack -benchmem -run=^$ -count=10 > stack.txt
// goos: darwin
// goarch: amd64
// pkg: auxiliary/test_and_example/goMedium
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// BenchmarkMemoryStack-12    	257367638	         4.499 ns/op	       0 B/op	       0 allocs/op
// BenchmarkMemoryStack-12    	269963438	         4.578 ns/op	       0 B/op	       0 allocs/op
// BenchmarkMemoryStack-12    	265248068	         4.449 ns/op	       0 B/op	       0 allocs/op
// BenchmarkMemoryStack-12    	270399489	         5.315 ns/op	       0 B/op	       0 allocs/op
// BenchmarkMemoryStack-12    	263429912	         4.442 ns/op	       0 B/op	       0 allocs/op
// BenchmarkMemoryStack-12    	262730037	         4.750 ns/op	       0 B/op	       0 allocs/op
// BenchmarkMemoryStack-12    	265347640	         4.526 ns/op	       0 B/op	       0 allocs/op
// BenchmarkMemoryStack-12    	269535060	         4.615 ns/op	       0 B/op	       0 allocs/op
// BenchmarkMemoryStack-12    	261565288	         4.442 ns/op	       0 B/op	       0 allocs/op
// BenchmarkMemoryStack-12    	237761036	         4.569 ns/op	       0 B/op	       0 allocs/op
// PASS
// ok  	auxiliary/test_and_example/goMedium	18.332s
func BenchmarkMemoryStack(b *testing.B) {
	var s S

	f, err := os.Create("stack.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		s = byCopy()
	}

	trace.Stop()

	b.StopTimer()

	_ = fmt.Sprintf("%v", s.a)
}

// BenchmarkMemoryHeap And another one, very similar, when it is passed by pointer:
// go test ./ -bench=BenchmarkMemoryHeap -benchmem -run=^$ -count=10 > head.txt
// goos: darwin
// goarch: amd64
// pkg: auxiliary/test_and_example/goMedium
// cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
// BenchmarkMemoryHeap-12    	26784159	        42.18 ns/op	      96 B/op	       1 allocs/op
// BenchmarkMemoryHeap-12    	28453618	        40.92 ns/op	      96 B/op	       1 allocs/op
// BenchmarkMemoryHeap-12    	28529114	        41.59 ns/op	      96 B/op	       1 allocs/op
// BenchmarkMemoryHeap-12    	29069200	        40.90 ns/op	      96 B/op	       1 allocs/op
// BenchmarkMemoryHeap-12    	29270865	        41.79 ns/op	      96 B/op	       1 allocs/op
// BenchmarkMemoryHeap-12    	27832362	        41.28 ns/op	      96 B/op	       1 allocs/op
// BenchmarkMemoryHeap-12    	28917258	        43.19 ns/op	      96 B/op	       1 allocs/op
// BenchmarkMemoryHeap-12    	25032111	        41.95 ns/op	      96 B/op	       1 allocs/op
// BenchmarkMemoryHeap-12    	24332090	        42.69 ns/op	      96 B/op	       1 allocs/op
// BenchmarkMemoryHeap-12    	28232472	        41.88 ns/op	      96 B/op	       1 allocs/op
// PASS
// ok  	auxiliary/test_and_example/goMedium	13.286s
func BenchmarkMemoryHeap(b *testing.B) {
	var s *S

	f, err := os.Create("heap.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		s = byPointer()
	}

	trace.Stop()

	b.StopTimer()

	_ = fmt.Sprintf("%v", s.a)
}

//You should read
//Go: Should I Use a Pointer instead of a Copy of my Struct? - by Vincent Blanchon - A Journey With Go- Medium,
//where they show that a copy of a struct may typically be 8 times faster than using a struct pointer.
//The reason for this is that using pointers causes struct variables to be placed in heap (after escape analysis),
//and GC is thus under more pressure.

// 总结性的文章：
//	https://medium.com/@kamruljpi/pointer-vs-copy-in-go-struct-which-one-is-better-b8db8a8265fc
//	https://www.sobyte.net/post/2021-10/go-pointer/

// 第二个实验
type Blah struct {
	c complex128
	s string
	f float64
}

func (b *Blah) doPtr() {
}

func (b Blah) doCopy() {
}

func BenchmarkDoPtr(b *testing.B) {
	blah := Blah{}
	for i := 0; i < b.N; i++ {
		(&blah).doPtr()
	}
}

func BenchmarkDoCopy(b *testing.B) {
	blah := Blah{}
	for i := 0; i < b.N; i++ {
		blah.doCopy()
	}
}

// go test 输出说明
// http://c.biancheng.net/view/124.html
