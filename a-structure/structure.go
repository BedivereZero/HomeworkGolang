package main

import "fmt"

type Book struct {
	title string
	author string
	subject string
	book_id int
}

func printBook( book Book ) {
	fmt.Printf("Book title:\t%s\n", book.title)
	fmt.Printf("Book author:\t%s\n", book.author)
	fmt.Printf("Book subject:\t%s\n", book.subject)
	fmt.Printf("Book book_id:\t%d\n", book.book_id)
}

func printBookPoint( book * Book ) {
	fmt.Printf("Book title:\t%s\n", book.title)
	fmt.Printf("Book author:\t%s\n", book.author)
	fmt.Printf("Book subject:\t%s\n", book.subject)
	fmt.Printf("Book book_id:\t%d\n", book.book_id)
}

func main() {
	fmt.Println(Book{"Go language", "google.com", "Golang", 6495407})
	fmt.Println(Book{title: "Go language", author: "google.com", subject: "Golang", book_id: 6495407})
	fmt.Println(Book{title: "Go language", author: "google.com"})

	var book0, book1 Book

	/* book 0 description */
	book0.title = "Go language"
	book0.author = "google.com"
	book0.subject = "Go 语言教程"
	book0.book_id = 6495407

	/* book 1 description */
	book1.title = "Python language"
	book1.author = "python.org"
	book1.subject = "Python 语言教程"
	book1.book_id = 6495700

	/* print book 0 information */
	fmt.Printf("Book 0 title:\t%s\n", book0.title)
	fmt.Printf("Book 0 author:\t%s\n", book0.author)
	fmt.Printf("Book 0 subject:\t%s\n", book0.subject)
	fmt.Printf("Book 0 book_id:\t%d\n", book0.book_id)

	/* print book 1 information */
	fmt.Printf("Book 1 title:\t%s\n", book1.title)
	fmt.Printf("Book 1 author:\t%s\n", book1.author)
	fmt.Printf("Book 1 subject:\t%s\n", book1.subject)
	fmt.Printf("Book 1 book_id:\t%d\n", book1.book_id)

	/* print book 0 informattion via printBook( book Book ) */
	printBook(book0)

	/* print book 1 informattion via printBook( book Book ) */
	printBook(book1)

	/* print book 0 informattion via printBookPoint( book Book ) */
	printBookPoint(&book0)

	/* print book 1 informattion via printBookPoint( book Book ) */
	printBookPoint(&book1)
}
