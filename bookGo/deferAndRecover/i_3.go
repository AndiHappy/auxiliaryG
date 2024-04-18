package main

import "fmt"

type number int

func (n number) print() {
	fmt.Println("print:", n)
}
func (n *number) pprint() {
	fmt.Println("pprint:", *n)
}

func main() {
	var n number
	defer n.print()               //第一个 defer 语句，对 n 直接求值，开始的时候 n=0，所以最后是 0。
	defer n.pprint()              //第二个 defer 语句，n 是引用，最终求值是 3;
	defer func() { n.print() }()  //第三个 defer 语句是闭包，引用外部函数的 n，最终结果是 3;
	defer func() { n.pprint() }() //第四个 defer 语句是闭包，引用外部函数的 n，最终结果是 3;
	n = 3
}
