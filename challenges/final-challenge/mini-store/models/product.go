package models

import (
	"errors"
	"fmt"
	"time"

	dataStorage "github.com/atlas-devel/Go-Journey/challenges/final-fundamental-challenge/store"
	"github.com/atlas-devel/Go-Journey/challenges/final-fundamental-challenge/validators"
)

type Product struct{}

func (p Product) CreateProduct(name string, price float64) (string, error) {
	isValid, err := validators.ValidateProduct(name, price)
	if err != nil {
		return "", err // drop error
	}
	if !isValid {
		return "", errors.New("Invalid product data") //error when product is not valid
	}

	product := dataStorage.Product{
		ID:    int(time.Now().UnixNano()),
		Name:  name,
		Price: price,
		Stock: 0,
	}
	dataStorage.Products = append(dataStorage.Products, product)

	return fmt.Sprintf("New Product %s created successfully", product.Name), nil
}
