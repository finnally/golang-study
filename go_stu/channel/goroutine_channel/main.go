package main

/*
使用go routine和channel实现一个计算int64随机数各位数和的程序：
1、开启一个go routine循环生成int64类型的随机数，发送到jobChan
2、开启24个go routine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
3、主go routine从resultChan取出结果并打印到终端输出
*/

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// job ...
type job struct {
	value int64
}

// result ...
type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

//randomNumber 1.开启一个goroutine循环生成int64类型的随机数，发送到jobChan
func randomNumber(a chan<- *job) {
	defer wg.Done()
	//生成int64类型的随机数，发送到jobChan
	for {
		x := rand.Int63()
		new1 := &job{
			value: x,
		}
		a <- new1
		time.Sleep(time.Millisecond * 1000)
	}
}

// calSum 从通道ch1读取数据，然后计算各个位数之和存入ch2中
func calSum(ch1 <-chan *job, resultChan chan<- *result) {
	// defer wg.Done()
	// 从jobChan中读取数据，然后计算各个位数之和存入resultChan中
	for {
		job := <-ch1
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job: job,
			sum: sum,
		}
		resultChan <- newResult
	}
}

//使用goroutine和channel实现一个计算int64随机数各位数和的程序。
func main() {
	wg.Add(1)
	//1.开启一个goroutine循环生成int64类型的随机数，发送到jobChan
	go randomNumber(jobChan)
	//2.开启24个goroutine从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go calSum(jobChan, resultChan)
	}
	//3.主goroutine从resultChan取出结果并打印到终端输出
	for result := range resultChan {
		fmt.Println(result.job.value, result.sum)
	}
	wg.Wait()
}
