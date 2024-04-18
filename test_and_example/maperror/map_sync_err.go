package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	MapMuxTest()
}

func MapMuxTest() {
	mMap := make(map[int]string)
	var mux sync.RWMutex
	mMap[1] = "new rap"
	mMap[2] = "new star"
	mMap[3] = "mo"
	mMap[4] = "dong"
	mMap[5] = "shan"
	mMap[6] = "ba"

	go func(mux *sync.RWMutex, mMap map[int]string) {
		ti := time.NewTimer(time.Nanosecond * 2)
		idx := 1
		for {
			<-ti.C

			mux.Lock()
			mMap[idx] = "mdsb" // map write
			idx++
			if idx >= 7 {
				idx = 1
			}
			mux.Unlock()

			ti.Reset(time.Nanosecond * 2)
		}
	}(&mux, mMap)

	ti := time.NewTimer(time.Millisecond)
	for {
		<-ti.C
		// 这里没有对遍历操作加锁，导致和上面的写操作冲突
		for k := range mMap { // map iteration
			mux.RLock()
			str := mMap[k]
			fmt.Println(str)
			mux.RUnlock()
		}

		ti.Reset(time.Millisecond)
	}
}
