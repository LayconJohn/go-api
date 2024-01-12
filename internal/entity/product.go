package entity

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    "1",
		Name:  name,
		Price: price,
	}
}