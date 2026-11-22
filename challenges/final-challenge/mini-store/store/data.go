package dataStorage

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

// sample data for the store
var Products []Product = []Product{

	{
		ID:    1,
		Name:  "Laptop",
		Price: 999.99,
		Stock: 10,
	}, {

		ID:    2,
		Name:  "Mouse",
		Price: 29.99,
		Stock: 50,
	},
}
