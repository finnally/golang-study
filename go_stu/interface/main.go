package main

import (
	"fmt"
	"reflect"
)

type IF interface {
	getName() string
}

type Human struct {
	firstName, lastName string
}

type Plane struct {
	vendor string
	model  string
}

type Car struct {
	factory, model string
}

func (h *Human) getName() string {
	h.firstName = "aaa"
	return h.firstName + "," + h.lastName
}

func (p Plane) getName(s string) string {
	p.vendor = s
	return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
}

func (c *Car) getName() string {
	c.factory = "bmw"
	return c.factory + "-" + c.model
}

func main() {
	var ss IF
	fmt.Println(reflect.TypeOf(ss))
	ss = new(Human)
	fmt.Println(reflect.TypeOf(ss))
	// 定义一个元素为接口类型的slice
	interfaces := []IF{}
	h := new(Human)
	h.firstName = "first"
	h.lastName = "last"
	interfaces = append(interfaces, h)
	c := new(Car)
	c.factory = "benz"
	c.model = "s"
	interfaces = append(interfaces, c)
	// p := new(Plane)
	// p.vendor = "testVendor"
	// p.model = "testModel"
	// interfaces = append(interfaces, p)
	for _, f := range interfaces {
		fmt.Println(f.getName())
	}
	fmt.Println(c)
	p1 := Plane{}
	p1.vendor = "ven"
	p1.model = "mod"
	fmt.Println(p1.getName("ac"))
	fmt.Println(p1)
	p2 := &Plane{}
	p2.vendor = "ven"
	p2.model = "mod"
	fmt.Println(p2.getName("ac"))
	fmt.Println(p2)
}
