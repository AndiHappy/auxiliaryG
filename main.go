package main

import (
	"auxiliary/framework/gosdkexample"
	"auxiliary/framework/httpandparse"
	"fmt"
)

func main() {
	fmt.Println("Hello,Online IDE! ")
	httpandparse.TestHttpClient("htt://baidu.com/")
	//strings review
	gosdkexample.CloneT()
	//unsafe review
	gosdkexample.StringDataT()
}
