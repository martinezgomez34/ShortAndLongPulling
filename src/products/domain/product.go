package domain

type Product struct {
	ID          int32
	Name        string
	Price       float32
	Amount      int32
	Brand       string
	Description string
}

func NewProduct(name string, price float32, amount int32, brand, description string) *Product {
	return &Product{
		Name:        name,
		Price:       price,
		Amount:      amount,
		Brand:       brand,
		Description: description,
	}
}
