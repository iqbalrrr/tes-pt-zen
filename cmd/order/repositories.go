package order

import (
	"database/sql"
	"tesptzen/utilitys"
)

type IRepository interface {
	Insert(_orders Orders) (Orders, error)
	Update(_orders Orders) (Orders, error)
	Delete(_orders Orders) (Orders, error)
	GetById(Id int) (Orders, error)
	GetAll() ([]Orders, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(_orders Orders) (Orders, error) {

	querys := "insert into orders(customer_name,product_id,quantity,status) VALUES ($1,$2,$3,$4)"
	_, err := r.db.Exec(querys,
		_orders.Customer_name,
		_orders.Product_id,
		_orders.Quantity,
		_orders.Status,
	)

	if err != nil {
		utilitys.LogError(err)
		return _orders, err
	}

	return _orders, nil
}

func (r *repository) Update(_orders Orders) (Orders, error) {
	querys := "update orders set customer_name = $1,product_id = $2, quantity = $3,status = $4 where id = $5"
	_, err := r.db.Exec(querys,
		_orders.Customer_name,
		_orders.Product_id,
		_orders.Quantity,
		_orders.Status,
		_orders.Id,
	)

	if err != nil {
		return _orders, err
	}

	return _orders, nil
}

func (r *repository) Delete(_orders Orders) (Orders, error) {

	querys := "delete from products where Id = $1"
	_, err := r.db.Exec(querys,
		_orders.Id,
	)
	if err != nil {
		return _orders, err
	}
	return _orders, nil
}

func (r *repository) GetById(Id int) (Orders, error) {
	var _orders = Neworders()
	var collorders []Orders
	rows, err := r.db.Query("select * from orders where id = $1", Id)
	if err != nil {
		utilitys.LogError(err)
		return _orders, err
	}
	defer rows.Close()
	for rows.Next() {

		err2 := rows.Scan(&_orders.Id,
			&_orders.Customer_name,
			&_orders.Product_id,
			&_orders.Quantity,
			&_orders.Status,
		)
		if err2 != nil {
			return _orders, err2
		}
		collorders = append(collorders, _orders)

	}

	return collorders[0], nil
}

func (r *repository) GetAll() ([]Orders, error) {
	var _orders = Neworders()
	var collOrders []Orders
	rows, err := r.db.Query("select * from orders")
	if err != nil {
		utilitys.LogError(err)
		return collOrders, err
	}
	defer rows.Close()
	for rows.Next() {

		err2 := rows.Scan(&_orders.Id,
			&_orders.Customer_name,
			&_orders.Product_id,
			&_orders.Quantity,
			&_orders.Status,
		)
		if err2 != nil {
			return collOrders, err2
		}
		collOrders = append(collOrders, _orders)

	}

	return collOrders, nil
}
