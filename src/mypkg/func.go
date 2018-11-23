package main

import "fmt"

/**传递引用*/
func swapRef(x *int, y *int) {
	var temp int
	temp = *x /* 保持 x 地址上的值 */
	*x = *y   /* 将 y 值赋给 x */
	*y = temp /* 将 temp 值赋给 y */
}

/**多返回值*/
func swap(x, y string) (string, string) {
	return y, x
}

func testAnonymousFun() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}

/**匿名函数，共用变量*/
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
