package main

import "fmt"

type ServiceType int32

const (
	ServiceTypeClusterIP ServiceType = iota + 1
	ServiceTypeNodePort
	ServiceTypeLoadBalancer
	ServiceTypeExternalName
)

/*
	String()方法用于定制fmt.Println(x)、fmt.Printf("%v", x)、fmt.Print(x)时输出的内容。
	对于定于了String()方法的类型，默认输出的时候会调用该方法，实现字符串的打印。
	如果类型定义了 String() 方法，它会被用在 fmt.Printf() 中生成默认的输出，例如：
	type Man struct {
    	name string
	}

	func (m Man) String() string {
		return "My name is :" + m.name
	}

	func main() {
		var m Man
		m.name = "SNS"
		fmt.Println(m)
	}

	输出：
	My name is :SNS

	需要注意的是receiver的类型，如果为指针类型，则fmt.Println(指针)时可以调用String()方法，
	fmt.Println(值)不会调用String()方法。而如果receiver是值类型，则fmt.Println(值)和fmt.Println(指针)都会调用String()方法。
	原因是fmt.Println()函数其接收的参数是实现了String()方法的接口，接口不保存地址，因此当fmt.Println(值)时得不到地址也就无法调用String()方法。

*/

func (s ServiceType) String() string {
	switch s {
	case ServiceTypeClusterIP:
		return "ClusterIP"
	case ServiceTypeNodePort:
		return "NodePort"
	case ServiceTypeLoadBalancer:
		return "LoadBalancer"
	case ServiceTypeExternalName:
		return "ExternalName"
	default:
		return "Unkown"
	}
}

func Service(s ServiceType) {
	fmt.Printf("%v", s)
}

func main() {
	Service(ServiceTypeClusterIP)
}
