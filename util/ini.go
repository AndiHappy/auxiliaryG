package util

import (
	"log"
	"time"
)

func init() {

	//设定log格式
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")

	//设定时区
	_, err := time.LoadLocation("Asia/Shanghai")
	if err != nil { // Always check errors even if they should not happen. //
		panic(err)
	}
}
