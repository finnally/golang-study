package main

import "fmt"

func main() {
	// 第一种声明方式
	var map1 map[string]string
	if map1 == nil {
		fmt.Println("空map")
	}

	// 使用map前，需要先通过make分配空间
	map1 = make(map[string]string, 10)
	map1["one"] = "map1"
	fmt.Println(map1)

	// 第二种声明方式
	map2 := make(map[int]string)
	map2[1] = "map2"
	fmt.Println(map2)

	// 第三种声明方式
	map3 := map[string]string{
		"one": "map3",
		"two": "python",
	}
	fmt.Println(map3)

	// 删除
	delete(map3, "one")

	// 添加
	map3["three"] = "java"

	// 修改
	map3["two"] = "golang"
	printMap(map3)

}

func printMap(m map[string]string) {
	// m 是一个引用传递
	for k, v := range m {
		fmt.Println(k, v)
	}
}
