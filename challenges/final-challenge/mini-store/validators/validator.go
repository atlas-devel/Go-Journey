package validators

import (
	"errors"
	"strings"
)

func ValidateProduct(name string, price float64) (bool, error) {
	if strings.TrimSpace(name) == "" {
		return false, errors.New("Name cannot be empty")
	}
	if price <= 0 {
		return false, errors.New("Price must be greater than zero")
	}

	return true, nil

}

func QuantityValidator(quantity int) (bool, error) {
	if quantity <= 0 {
		return false, errors.New("Quantity must be greater than zero")
	}
	return true, nil
}
