package main

import (
	 "fmt"

	"github.com/atlas-devel/Go-Journey/challenges/LMS/models"
	"github.com/atlas-devel/Go-Journey/challenges/LMS/services"
)

func main() {
	book := models.Book{
		Title:    "Clean Code",
		Author:   "Robert C. Martin",
		Price:    15000,
		Copies:   10,
		Borrowed: 0,
	}

	book.AddCopies(4)
	fmt.Println(book.Copies)
	
	book.BorrowBook(-1)
	book.ReturnBook(1)
	fmt.Println(book.Borrowed)
	fmt.Println(book.Copies)



	fmt.Println(services.GetDetails(book))
	fmt.Println("Total inventory value:", services.GetTotalValue(book))

}
