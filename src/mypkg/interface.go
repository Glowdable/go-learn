package main

/**Golang 并没有 extends，它类似的方式是 Embedding。
go 没有继承，只有接口和组合
https://www.cnblogs.com/jasonxuli/p/6836399.html
*/

import (
	"fmt"
)

type Phone interface {
	call()
}

type NokiaPhone struct {
}

//NokiaPhone的内部方法，nokiaPhone.call
func (nokiaPhone NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

//IPhone的内部方法,iPhone.call
func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func testInterface() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()

}
