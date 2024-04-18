package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("before return")
	}()
	if true {
		fmt.Println("during return")
		return
	}
	defer func() {
		fmt.Println("after return")
	}()
}
