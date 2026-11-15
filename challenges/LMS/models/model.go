package models

import "fmt"

type Book struct {
	Title  string
	Author string
	Price  float64
	Copies int
	Borrowed int
}



func (b *Book) AddCopies(quantity int) {
	if quantity <= 0 {
		fmt.Println("Insuffient Copies")
		return
	}
	b.Copies+=quantity
	
}
func(b *Book) BorrowBook(quantity int)    {
	switch {
	case quantity>b.Copies:
		fmt.Println("Books not available")
		return
	case quantity<=0:
		fmt.Println("Insufficient Quantity")
		return
	}
	b.Copies-=quantity
	b.Borrowed+=quantity
	

}
func (b *Book) ReturnBook(quantity int)     {
	switch{
	case quantity>b.Borrowed:
		fmt.Println("Maximum book return exceed")
		return
	case quantity<=0:
		fmt.Println("Insufficient Return")
		return

	}
	b.Copies+=quantity
	b.Borrowed-=quantity
	
}
