package models

import (
	"fmt"
"errors"
)

type Book struct {
	Title  string
	Author string
	Price  float64
	Copies int
	Borrowed int
}



func (b *Book) AddCopies(quantity int) error {
	if quantity <= 0 {
		
		return fmt.Errorf("quantity must be greater than zero and your are trying to add %d",quantity)
	}
	b.Copies+=quantity
	return  nil
}

func(b *Book) BorrowBook(quantity int) error    {
	switch {
	case quantity>b.Copies:
		return fmt.Errorf("cannot borrow %d books: only %d available",quantity,b.Copies)
	case quantity<=0:
		return errors.New("Insufficient Quantity")
		
	}
	b.Copies-=quantity
	b.Borrowed+=quantity

	return nil
	

}
func (b *Book) ReturnBook(quantity int)  error   {
	switch{
	case quantity>b.Borrowed:
		return errors.New("Maximum book return exceed")
		
	case quantity<=0:
		return errors.New("Insufficient Return")
		
	}
	b.Copies+=quantity
	b.Borrowed-=quantity

	return nil
	
}
