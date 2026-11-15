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



	if err:=book.AddCopies(02);err!=nil{
		fmt.Println(err)
	}
		if err:=book.BorrowBook(20);err!=nil{
		fmt.Println(err)
	}
		if err:=book.ReturnBook(1);err!=nil{
		fmt.Println(err)
	}

	fmt.Println(book.Copies)

	
	



	fmt.Println(services.GetDetails(book))
	fmt.Println("Total inventory value:", services.GetTotalValue(book))

}
