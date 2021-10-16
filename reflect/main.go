package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	name string
}

func (s Student) GetName(ss string) string {
	return s.name + ss
}

func main() {
	student := Student{name: "lily"}
	values := reflect.ValueOf(student)
	for i := 0; i < values.NumField(); i++ {
		fmt.Println(values.Field(i))
	}
	params := make([]reflect.Value, 1) //参数
	params[0] = reflect.ValueOf("s")
	for i := 0; i < values.NumMethod(); i++ {
		fmt.Printf("Method %d: %v\n", i, values.Method(i).Call(params))
	}
	// myStruct := T{A: "a"}
	// v1 := reflect.ValueOf(myStruct)
	// for i := 0; i < v1.NumField(); i++ {
	// 	fmt.Printf("Field %d: %v\n", i, v1.Field(i))
	// }
	// for i := 0; i < v1.NumMethod(); i++ {
	// 	fmt.Printf("Method %d: %v\n", i, v1.Method(i))
	// }
	// 需要注意receive是struct还是指针
	// result := v1.Method(0).Call(nil)
	// fmt.Println("result:", result)
}

// type T struct {
// 	A string
// }
//
// // 需要注意receive是struct还是指针
// func (t T) String() string {
// 	return t.A + "1"
// }
