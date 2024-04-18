package gotutorial

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
	"time"
)

var bufPool = sync.Pool{
	New: func() any {
		// The Pool's New function should generally only return pointer
		// types, since a pointer can be put into the return interface
		// value without an allocation:
		return new(bytes.Buffer)
	},
}

// timeNow is a fake version of time.Now for tests.
func timeNow() time.Time {
	return time.Unix(1136214245, 0)
}

func Log(w io.Writer, key, val string) {
	b := bufPool.Get().(*bytes.Buffer)
	b.Reset()
	// Replace this with time.Now() in a real logger.
	b.WriteString(timeNow().UTC().Format(time.RFC3339))
	b.WriteByte(' ')
	b.WriteString(key)
	b.WriteByte('=')
	b.WriteString(val)
	w.Write(b.Bytes())
	bufPool.Put(b)
}

func Test_Pool(t *testing.T) {
	Log(os.Stdout, "path", "/search?q=flowers")
}

type Test1 struct {
	A int
}

func Test_Pool_2(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return &Test1{
				A: 1,
			}
		},
	}
	fmt.Println("Test_Pool_2")
	pool.Put(&Test1{A: 2})
	pool.Put(&Test1{A: 1})
	pool.Put(&Test1{A: 3})
	for i := 0; i < 10; i++ {
		go func() {
			testObject := pool.Get().(*Test1)
			fmt.Println(testObject.A)
			pool.Put(testObject)
		}()
	}
	time.Sleep(time.Second * 1)
	fmt.Println("----------------------------")
	for i := 0; i < 10; i++ {
		//pool 中的数据是随机获取的，还是有一定的顺序？
		testObject := pool.Get().(*Test1)
		fmt.Println(testObject.A)
	}
}

//
