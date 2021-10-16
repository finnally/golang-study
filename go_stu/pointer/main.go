package main

import (
	"fmt"
)

func modify(p *int) {
	fmt.Println(p)
	*p = 1000
}

func main() {
	var n int = 10
	fmt.Println("n的地址是", &n)

	pointer := &n
	*pointer = 9
	fmt.Printf("n的值是%d", n)
	/*
		var a0 *int
		var b0 int = 1
		a0 = &b0
		针对*int的，先创建一个b0类型，然后取b的地址，赋值给*int，不能直接通过*a0 = 5赋值，因为声明的a变量是一个值为nil的空指针
		比较常用的方法应该是：
		a := new(int)
		*a = 10
	*/

	var a1 int = 10
	fmt.Println(&a1)

	var p *int //*int 表示指针  类型前加*
	p = &a1    //*a 表示a的地址
	fmt.Println(*p)

	*p = 100 //从内存中修改值
	fmt.Println(a1)

	var b int = 999
	p = &b
	*p = 5
	fmt.Println(a1)
	fmt.Println(b)

	modify(&a1) //指针赋值
	fmt.Println(a1)

	var ip *int     /* 声明指针变量 */
	var a2 int = 20 /* 声明实际变量 */
	ip = &a2        /* 指针变量的存储地址 */
	fmt.Printf("a 变量的地址是: %x\n", &a2)
	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)

	// 三重指针 pt3 - > pto - > ptr - >变量a3
	var a3 int = 5
	//把ptr指针 指向ss所在地址
	var ptr *int = &a3
	//开辟一个新的指针，指向ptr指针指向的地方
	var pts *int = ptr
	//二级指针，指向一个地址，这个地址存储的是一级指针的地址
	var pto **int = &ptr
	//三级指针，指向一个地址，这个地址存储的是二级指针的地址，二级指针同上
	var pt3 ***int = &pto
	fmt.Println("a3的地址:", &a3,
		"\n 值", a3, "\n\n",

		"ptr指针所在地址:", &ptr,
		"\n ptr指向的地址:", ptr,
		"\n ptr指针指向地址对应的值", *ptr, "\n\n",

		"pts指针所在地址:", &pts,
		"\n pts指向的地址:", pts,
		"\n pts指针指向地址对应的值:", *pts, "\n\n",

		"pto指针所在地址:", &pto,
		"\n pto指向的指针（ptr）的存储地址:", pto,
		"\n pto指向的指针（ptr）所指向的地址: ", *pto,
		"\n pto最终指向的地址对应的值（a）", **pto, "\n\n",

		"pt3指针所在的地址:", &pt3,
		"\n pt3指向的指针（pto）的地址:", pt3, //等于&*pt3,
		"\n pt3指向的指针（pto）所指向的指针的（ptr）地址", *pt3, //等于&**pt3,
		"\n pt3指向的指针（pto）所指向的指针（ptr）所指向的地址（a）:", **pt3, //等于&***pt3,
		"\n pt3最终指向的地址对应的值（a）", ***pt3)
}
