package main

func main() {
	//主协程的循环很快就跑完了，而各个协程才开始跑，此时i的值已经是10了，所以各协程都输出了10。（输出7的两个协程，在开始输出的时候主协程的i值刚好是7，这个结果每次运行输出都不一样）
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		fmt.Println(i)
	// 	}()
	// }
	// time.Sleep(time.Second)
	for i := 0; i < 3; i++ {
		println("&i=", &i, " i=", i)
	}
}
