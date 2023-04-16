package main

import "fmt"

/**
 *	四种变量的声明
 */

// 声明全局变量 方法一、二、三是可以的
var gA int = 100
var gB = 200

// 用方法四来声明全局变量
// := 只能够用在 函数体内来声明
// gC := 22

func main() {
	// 方法一：声明一个变量 默认的值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	// 方法二：声明一个变量，初始化一个值
	var b int = 1
	fmt.Println(b)
	fmt.Printf("type of b = %T\n", b)

	// 方法三： 在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
	var c = 2
	fmt.Println(c)
	fmt.Printf("type of c = %T\n", c)

	// 方法四：（常用的方法）省去var关键字，直接自动匹配
	d := 3
	fmt.Println(d)
	fmt.Printf("type of d = %T\n", d)

	e := 3.14
	fmt.Println(e)
	fmt.Printf("type of e = %T\n", e)

	// ===========
	fmt.Println("gA = ", gA, ", gB = ", gB)

	//fmt.Println("gC = ", gC);

	// 声明多个变量
	var xx, yy int = 1, 2
	fmt.Println("xx = ", xx, "yy = ", yy)

	// 多行的变量声明
	var (
		vv int  = 1
		jj bool = true
	)
	fmt.Println("vv = ", vv, ", jj = ", jj)

}
