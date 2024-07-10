package main

import (
	"net/http"
	_ "net/http/pprof"
	"runtime"
)

// http://127.0.0.1:6060/debug/pprof/
func main() {
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪，block
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪，mutex
	err := http.ListenAndServe("0.0.0.0:6060", nil)
	if err != nil {
		return
	}
}
