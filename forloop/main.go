package main

import "fmt"

func main() {
	vals := []int{0, 1, 2}
	valptr := make([]*int, 0)
	/*
		在执行for循环的时候，golang会首先创建一块内存，用于存放v。
		之后依次将vals中的元素拷贝到这块内存，在for之后若没有继续引用便进行释放。
		所以在此过程中，修改v或将v放入valptr中，只会放入最后一个元素。
	*/
	for _, v := range vals {
		valptr = append(valptr, &v)
	}

	for _, v := range valptr {
		fmt.Printf("%v, %v\n", v, *v)
	}

	val := make([]*int, 0)
	for i := 0; i < 3; i++ {
		j := i
		val = append(val, &j)
		fmt.Printf("&i=%v, i=%v\n", &i, i)
	}
	for _, v := range val {
		fmt.Printf("%v\n", *v)
	}
}
