package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Stock int
}

func (p *Product) addStock(productCount int) {
	if productCount<=0{
		fmt.Println("Insufficient products")
		return
	}
	p.Stock += productCount
}
func (p *Product) sell(countProductSold int) {
	if countProductSold>p.Stock{
		fmt.Println("Insufficient products")
		return
	}
	p.Stock -= countProductSold
}
func (p Product) getDetails() string {

	return fmt.Sprintf("Current Product info: product name: %s, in stock:%d, price: %f ", p.Name,p.Stock,p.Price)
}

func main(){
	product:=Product{
		Name: "Laptop",
		Price: 7500000,
		Stock: 10,
	}

	//product.addStock(8)
	product.sell(20)
	fmt.Println(product.Stock)
	// fmt.Println(product.getDetails())
}