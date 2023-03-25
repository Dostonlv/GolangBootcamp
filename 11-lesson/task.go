package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type BookData interface {
	GetBooks() []Book
	AddBook(Book)
	UpdateBook(Book)
	RemoveBook(Book)
	//GetBookById(Book)
	//GetBookByCategory(Book)
}

type Book struct {
	ID       int    `jsonning:"id"`
	Title    string `jsonning:"title"`
	Author   string `jsonning:"author"`
	Year     int    `jsonning:"year"`
	Status   string `jsonning:"status"`
	Price    int    `jsonning:"price"`
	Period   int    `jsonning:"period"`
	Category string `jsonning:"category"`
	Page     int    `jsonning:"page"`
}

type BookLib struct {
	Books []Book
}

func (b *BookLib) GetBooks() []Book {
	return b.Books
}

func (b *BookLib) RemoveBook(book Book) {
	for i, existingBook := range b.Books {
		if existingBook.Title == book.Title {
			b.Books = append(b.Books[:i], b.Books[i+1:]...)
			return
		}
	}
}

//func (b *BookLib) GetBookById(id int) *Book {
//	for _, book := range b.Books {
//		if book.ID == id {
//			return &book
//		}
//	}
//	return nil
//}

func (b *BookLib) AddBook(book Book) {
	b.Books = append(b.Books, book)
}

func (b *BookLib) UpdateBook(updatedBook Book) {
	for i, book := range b.Books {
		if book.Title == updatedBook.Title {

			b.Books[i] = updatedBook
			return
		}
	}
}

func main() {
	data, err := os.ReadFile("books.jsonning")
	if err != nil {
		panic(err)
	}
	var book []Book
	err = json.Unmarshal(data, &book)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	library := &BookLib{Books: book}
START:
	var a int
	fmt.Println("----------Welcome to library-----------")
	fmt.Println("1-kitoblarni ko'rish")
	fmt.Println("2-kitob kiritish")
	fmt.Println("3-kitobni update qilish kiritish")
	fmt.Println("4-kitobni ochirish ")
	fmt.Println("5-kitobni id bilan topish ")
	fmt.Print("Sonni kiriting: ")
	fmt.Scan(&a)
	if a == 1 {
		for _, i := range library.GetBooks() {
			fmt.Printf(`--------------------
Book Id: %v
Book Name: %v,
Author:%v
Year:%v
Page: %v
`, i.ID, i.Title, i.Author, i.Year, i.Page)
		}

	} else if a == 2 {

		library.AddBook(Book{
			ID:     len(library.GetBooks()) + 1,
			Title:  "nimadir",
			Author: "kimdir",
			Year:   1221,
			Price:  1221,
			Period: 123,
			Page:   233,
		})
		goto START
	} else if a == 3 {
		library.UpdateBook(Book{
			ID:     1,
			Title:  "The Great Gatsby",
			Author: "Tohir Abdullaev",
			Year:   2020,
			Page:   12233,
		})
		goto START
	} else if a == 4 {
		library.RemoveBook(Book{
			Title:  "The Great Gatsby",
			Author: "F. Scott Fitzgerald",
			Year:   1925,
		})
		goto START
	} else if a == 5 {
		//var a int
		//fmt.Print("Id ni kiriting: ")
		//fmt.Scan(&a)
		//fmt.Println(library.GetBookById(a))
		goto START
	}

}
