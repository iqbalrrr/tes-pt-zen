package order

type Orders struct {
	Id            int
	Customer_name string
	Product_id    int
	Quantity      int
	Status        string
}

func Neworders() Orders {
	_orders := Orders{}
	_orders.Id = 0
	_orders.Customer_name = ""
	_orders.Product_id = 0
	_orders.Quantity = 0
	_orders.Status = ""

	return _orders
}
