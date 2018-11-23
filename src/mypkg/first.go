package main

import "fmt"

const cb string = "abc" //常量
const (                 //枚举常量
	Unknown = 0
	Female  = 1
	Male    = 2
)

var age int = 1
var a = "helen"
var b string = "glowd.cn"
var c bool
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0} //数组

//go build ./
func main() {
	fmt.Println("Hello World")
	// var a int = 10
	// var b = 10
	// c := 10//自动推断
	fmt.Println(a, b)
	fmt.Println(a + b)
	fmt.Println(&a)

	e, h := swap("abc", "def")
	fmt.Println(e, h)
	var salary float32 = balance[1]
	fmt.Printf("Element[%f]\n", salary)
	structCall()

	var ref1 int = 100
	var ref2 int = 200
	swapRef(&ref1, &ref2)
	fmt.Printf("交换后的值，ref1=%d,ref2=%d", ref1, ref2)

	//函数引用传递
	desclareSwap := swap
	y, z := desclareSwap("abc", "def")
	fmt.Println(y, z)

	c4 = Circle{}
	fmt.Printf("%f", c4.radius)

	testAnonymousFun()

	pointer()

	pointerArray()

	testBook()

	slice()
	testMap()
	testSwitch()
	testInterface()
	testError()

}
