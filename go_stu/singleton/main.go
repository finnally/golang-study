package main

import "fmt"

/*
	单例模式，也叫单子模式，是一种常用的软件设计模式。
	在应用这个模式时，单例对象的类必须保证只有一个实例存在。
	许多时候整个系统只需要拥有一个的全局对象，这样有利于我们协调系统整体的行为。
	比如在某个服务器程序中，该服务器的配置信息存放在一个文件中，
	这些配置数据由一个单例对象统一读取，然后服务进程中的其他对象再通过这个单例对象获取这些配置信息。
	这种方式简化了在复杂环境下的配置管理。

	单例模式的主要使用场景有以下两个方面：
	1、资源共享情况下避免资源操作导致的性能损耗，比如日志管理器，web网站计数器，应用配置管理对象等
	2、方便对资源的控制，比如线程池和数据库连接池等

	单例模式要实现的效果就是，对于应用单例模式的类，整个程序中只存在一个实例化对象
*/

type person struct {
	name string
}

var instance *person

func singleton() *person {
	if instance == nil {
		instance = &person{}
	}
	return instance
}

func main() {
	s := singleton()
	s.name = "single"
	fmt.Println(s.name)
	s2 := singleton()
	fmt.Println(s2.name)
}
