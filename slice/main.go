package main

import "fmt"

func main() {
	// 声明slice1是一个切片并初始化
	slice1 := []int{1, 2, 3}

	// 声明slice2是一个切片，但并没有给其分配空间
	var slice2 []int

	// 声明slice3是一个切片并分配空间，初始化值是0
	var slice3 []int = make([]int, 3)

	// 声明slice4是一个切片并分配空间，初始化值是0，通过:=推导出其切片类型
	slice4 := make([]int, 3) // 常用

	fmt.Println(slice1, slice2, slice3, slice4)

	// 判断一个slice是否为空
	if slice2 == nil {
		fmt.Println("空切片")
	} else {
		fmt.Println("有空间")
	}

	// 声明slice是一个切片并分配空间，容量为5
	slice := make([]int, 4, 5)
	fmt.Printf("len = %d, cap = %d, slice5 = %v\n", len(slice), cap(slice), slice)

	// 向slice追加一个元素1
	slice = append(slice, 1)
	fmt.Printf("len = %d, cap = %d, slice5 = %v\n", len(slice), cap(slice), slice)

	/*
		向slice追加一个元素2：
		1、如果扩容之后，还没有触及原数组的容量，那么，切片中的指针指向的位置，就还是原数组；
		如果扩容之后，超过了原数组的容量，那么，Go就会开辟一块新的内存，把原来的值拷贝过来，这种情况丝毫不会影响到原数组。
		2、如果切片的容量小于1024个元素，那么扩容的时候slice的cap就翻番，乘以2；
		一旦元素个数超过1024个元素，增长因子就变成1.25，即每次增加原来容量的四分之一。
	*/
	slice = append(slice, 2)
	fmt.Printf("len = %d, cap = %d, slice5 = %v\n", len(slice), cap(slice), slice)

	// 切片截取
	s1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// 打印子切片从索引1（包含）到索引4（不包含），即[1, 4)
	fmt.Println(s1[1:4])
	// 打印子切片从索引0（包含）到索引3（不包含）
	fmt.Println(s1[:3])
	// 打印子切片从索引4（不包含）到len(s)（包含）
	fmt.Println(s1[4:])

	// 浅拷贝
	s2 := s1[:]
	s2[0] = 10
	fmt.Println(s2, s1)

	// 深拷贝
	s3 := []int{1, 2, 3}
	s4 := make([]int, 3)
	copy(s4, s3)
	s4[0] = 4
	fmt.Println(s3, s4)
}
