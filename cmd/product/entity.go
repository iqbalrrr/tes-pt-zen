package product

type Products struct {
	Id          int
	Name        string
	Description string
	Price       float32
	Stock       int
}

func NewProducts() Products {
	_products := Products{}
	_products.Id = 0
	_products.Name = ""
	_products.Description = ""
	_products.Price = 0
	_products.Stock = 0

	return _products
}
