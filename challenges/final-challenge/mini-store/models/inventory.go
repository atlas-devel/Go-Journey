package models

import (
	"errors"
	"strings"

	dataStorage "github.com/atlas-devel/Go-Journey/challenges/final-fundamental-challenge/store"
)

type Store struct{}

func (i Store) AddProductToStock(name string, quantity int) error {
	if int(quantity) <= 0 {
		return errors.New("Quantity must be greater than zero")
	}

	for index := range dataStorage.Products {

		if strings.TrimSpace(name) == dataStorage.Products[index].Name {
			dataStorage.Products[index].Stock += quantity
			return nil
		}
	}
	return errors.New("Product not found")
}

func (i Store) GetProductByID(id int) (dataStorage.Product, error) {
	for _, item := range dataStorage.Products {
		if item.ID == id {
			return item, nil
		}

	}
	return dataStorage.Product{}, errors.New("Product not found")
}

func (i Store) SellProduct(name string, quantity int) (string, error) {
	for _, item := range dataStorage.Products {
		if item.Name == name {
			if item.Stock >= quantity {
				item.Stock -= quantity
				return "Product sold successfully", nil
			} else {
				return "Insufficient stock", errors.New("Insufficient stock")
			}
		}
	}
	return "Product not found", errors.New("Product not found")
}

func (i Store) GetAllProducts() []dataStorage.Product {
	return dataStorage.Products
}

func (i Store) RemoveProductById(id int) (string, error) {
	for index, item := range dataStorage.Products {
		if item.ID == id {
			dataStorage.Products = append(dataStorage.Products[:index], dataStorage.Products[index+1:]...)
			return "Product removed successfully", nil
		}
	}
	return "Product not found", errors.New("Product not found")
}

func (i Store) GetLowestPriceProduct() (dataStorage.Product, error) {
	if len(dataStorage.Products) == 0 {
		return dataStorage.Product{}, errors.New("No products in store")
	}

	var lowestPriceProduct float64 = dataStorage.Products[0].Price
	var lowProductPrice dataStorage.Product = dataStorage.Products[0]

	for _, v := range dataStorage.Products {
		if lowestPriceProduct > v.Price {
			lowProductPrice = v

		}
	}
	return lowProductPrice, nil
}

func (i Store) GetProductBelowQuantity(quantity int) ([]dataStorage.Product, error) {
	if len(dataStorage.Products) == 0 {
		return nil, errors.New("No products in store")
	}
	if quantity < 0 {
		return nil, errors.New("Quantity cannot be negative")
	}
	products := []dataStorage.Product{}
	for _, v := range dataStorage.Products {
		if v.Stock <= quantity {
			products = append(products, v)
		}
	}
	return products, nil
}
