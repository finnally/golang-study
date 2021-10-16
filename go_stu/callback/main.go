package main

import (
	"fmt"
	"strconv"
)

//例1
type Callback1 func(x, y int) int

//提供一个接口，让外部去实现
func test(x, y int, callback Callback1) int {
	return callback(x, y)
}

func add(x, y int) int {
	return x + y
}

//例2
type Callback2 func(msg string)

//将字符串转换为int64，如果转换失败调用Callback2
func stringToInt(s string, callback Callback2) int64 {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		callback(err.Error())
	}
	return value
}

//记录日志消息的具体实现
func errLog(msg string) {
	fmt.Println("Convert error: ", msg)
}

func main() {
	fmt.Printf("eg1: %d\n", test(1, 2, add))
	fmt.Printf("eg2: %v\n", stringToInt("18", errLog))
	fmt.Printf("eg2: %v\n", stringToInt("hello", errLog))
}
