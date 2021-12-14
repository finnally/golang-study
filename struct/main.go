package main

import "fmt"

type Book struct {
	author string
	page   int
}

func (b *Book) Show() {
	fmt.Println(b.author)
	fmt.Println(b.page)
}

func (b Book) SetAuthor1(name string) {
	// b 是调用该方法的对象的一个副本（拷贝）
	b.author = name
}

func (b *Book) SetAuthor2(name string) {
	// b 是调用该方法的对象的一个指针（引用传递）
	b.author = name
}

func (b *Book) GetAuthor() {
	fmt.Println(b.author)
}

func main() {
	book := Book{"java", 1000}
	book.Show()
	book.SetAuthor1("golang")
	book.GetAuthor()
	book.SetAuthor2("golang")
	book.GetAuthor()
}
