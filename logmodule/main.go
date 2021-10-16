package main

import (
	"fmt"
	"reflect"
)

type student struct {
	id int
	me string
}

func (s student) EchoName(name string) {
	fmt.Println("my name is", name)
}

func main() {
	s := student{1, "hehei"}
	v := reflect.ValueOf(s)
	mv := v.MethodByName("EchoName")
	args := []reflect.Value{reflect.ValueOf("ddd")}
	mv.Call(args)
}
