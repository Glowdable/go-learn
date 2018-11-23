// 在Go语言中没有构造函数的概念,对象的创建通常交由一个全局的创建函数来完成,以NewXXX来命令,表示"构造函数":

package main

import (
	"fmt"
)

/* 定义结构体 */
type Circle struct {
	radius float64
	book   []Books
}

var c4 Circle

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func testBook() {

	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "glowd", "go-glowd", 6495407})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "glowd", subject: "go-glowd", book_id: 6495407})

	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "glowd"})
}

/**实例化结构体两种常用形式：
1：用 new(T) 实例化。
2：用 &T{} 实例化。（推荐）*/

func structCall() {
	var c1 Circle                        // 声明
	var c2 *Circle = new(Circle)         //使用 new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针
	var c3 *Circle = &Circle{radius: 30} //&{}也同样等价于用new，因为底层仍然会调用new(T)。

	c1 = Circle{2.00, make([]Books, 30)}

	books := make([]Books, 20) //make 为数组分配空间

	c1.book = books

	//c1.radius = 10.00
	c2.radius = 20.00
	c3.radius = 30.00
	fmt.Println("c1圆的面积 = ", c1.getArea())
	fmt.Println("c2圆的面积 = ", c2.getArea())
	fmt.Println("c2圆的面积 = ", c3.getArea())

	fmt.Println("圆的长度 = ", plus(c1))

}

//该 method 属于 Circle 类型对象中的方法
func (c Circle) getArea() float64 {
	//c.radius 即为 Circle 类型对象中的属性
	return 3.14 * c.radius * c.radius
}

//该 method 属于 Circle 类型对象中的方法
func plus(c Circle) float64 {
	return c.radius + c.radius
}
