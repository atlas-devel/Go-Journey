package service

import (
	"fmt"

	"github.com/atlas-devel/Go-Journey/challenges/final-fundamental-challenge/models"
)

func HandleOption(option int) {
	switch option {
	case 1:
		// Add product
		product_name := ""
		product_price := 0.0
		fmt.Print("Enter the product name: ")
		fmt.Scanln(&product_name)
		fmt.Print("Enter the product price: ")
		fmt.Scanln(&product_price)
		product, err := models.Product{}.CreateProduct(product_name, product_price)

		if err != nil {
			fmt.Println("Error creating product:", err)
		} else {
			fmt.Println("Product created successfully:", product)
		}

	case 2:
		// List products
		products := models.Store{}.GetAllProducts()
		for _, product := range products {
			fmt.Printf("- %s (ID: %d, Price: %.2f, Stock: %d)\n", product.Name, product.ID, product.Price, product.Stock)
		}
	case 3:
		// Find product
		var productID int
		fmt.Print("Enter the product ID to find: ")
		fmt.Scanln(&productID)
		product, err := models.Store{}.GetProductByID(productID)
		if err != nil {
			fmt.Println("Error finding product:", err)
		} else {
			fmt.Printf("Product found: %s (ID: %d, Price: %.2f, Stock: %d)\n", product.Name, product.ID, product.Price, product.Stock)
		}

	case 4:
		// Add stock
		product_name := ""
		quantity := 0
		fmt.Print("Enter the product name: ")
		fmt.Scanln(&product_name)
		fmt.Print("Enter the quantity to add: ")
		fmt.Scanln(&quantity)
		err := models.Store{}.AddProductToStock(product_name, quantity)
		if err != nil {
			fmt.Println("Error adding stock:", err)
		} else {
			fmt.Println("Stock added successfully")
		}

	case 5:
		// Sell product
		product_to_sell := ""
		quantity := 0
		fmt.Print("Enter the product name to sell: ")
		fmt.Scanln(&product_to_sell)
		fmt.Print("Enter the quantity to sell: ")
		fmt.Scanln(&quantity)
		err := models.Store{}.SellProduct(product_to_sell, quantity)
		if err != nil {
			fmt.Println("Error selling product:", err)
		} else {
			fmt.Printf("%s sold successfully\n", product_to_sell)
		}

	case 6:
		// Remove product
		fmt.Print("Enter the product ID to remove: ")
		var productID int
		fmt.Scanln(&productID)
		err := models.Store{}.RemoveProductById(productID)
		if err != nil {
			fmt.Println("Error removing product:", err)
		} else {
			fmt.Printf("Product with ID %d removed successfully\n", productID)
		}

	case 7:
		// Inventory value
		fmt.Print("")

	case 8:
		// Low-stock report
		fmt.Print("Products with low stock:\n")
		lowStockProducts, err := models.Store{}.GetProductBelowQuantity(5)
		if len( lowStockProducts )==0 {
			fmt.Println("No products with low stock.")
		}
		if err != nil {
			fmt.Println("Error fetching low stock products:", err)
		} else {
			for _, product := range lowStockProducts {
				fmt.Printf("- %s (ID: %d, Stock: %d)\n", product.Name, product.ID, product.Stock)
			}
		}

	case 9:
		// Price report
		lowestPriceProduct, err := models.Store{}.GetLowestPriceProduct()
		if err != nil {
			fmt.Println("Error fetching lowest price product:", err)
		} else {
			fmt.Printf("Product with the lowest price: %s (ID: %d, Price: %.2f)\n", lowestPriceProduct.Name, lowestPriceProduct.ID, lowestPriceProduct.Price)
		}

	case 0:
		// Exit
		fmt.Println("Exiting...")
		return

	default:
		println("Invalid option. Please try again.")
	}
}
