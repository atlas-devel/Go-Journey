package main

import (
	"fmt"

	"github.com/atlas-devel/Go-Journey/challenges/final-fundamental-challenge/service"
	dataStorage "github.com/atlas-devel/Go-Journey/challenges/final-fundamental-challenge/store"
)

func main() {
	var option int
	fmt.Print(`
	
==== MINI STORE MANAGEMENT =====

1. Add product
2. List products
3. Find product
4. Add stock
5. Sell product
6. Remove product
7. Inventory value
8. Low-stock report
9. Price report
0. Exit

Choose an option:
`)
	fmt.Scanln(&option)
	service.HandleOption(option)
	fmt.Println(dataStorage.Products)

}
