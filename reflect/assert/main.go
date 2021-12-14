package main

import "fmt"

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct{}

func (b *Book) ReadBook() {
	fmt.Println("read book")
}

func (b *Book) WriteBook() {
	fmt.Println("write book")
}

func main() {
	// book: pair<type:Book, value:book{}地址>
	book := &Book{}

	// r: pair<type:, value:>
	var r Reader
	// r: pair<type:Book, value:book{}地址>
	r = book
	r.ReadBook()

	// w: pair<type:, value:>
	var w Writer
	// w: pair<type:Book, value:book{}地址>
	w = r.(Writer) // 此处的断言为什么会成功？因为w和r具体的type是一致的
	w.WriteBook()
}
