package product

import (
	"database/sql"

	"tesptzen/utilitys"
)

type IRepository interface {
	Insert(_products Products) (Products, error)
	Update(_products Products) (Products, error)
	Delete(_products Products) (Products, error)
	GetById(Id int) (Products, error)
	GetAll() ([]Products, error)
	// RegisterNewUser(appuser ApplicationUser, roles []string) (string, error)
	// SyncUserToRole(applicationUser ApplicationUser, roles []string) (ApplicationUser, error)
	// FindByName(username string) (ApplicationUser, error)
	//FindByFieldName(fieldName string, value string) ([]Products, error)
	// FindAll() ([]ApplicationUser, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) Insert(_products Products) (Products, error) {

	querys := "insert into products(name,description,price,stock) VALUES ($1,$2,$3,$4)"
	_, err := r.db.Exec(querys,
		_products.Name,
		_products.Description,
		_products.Price,
		_products.Stock,
	)

	if err != nil {
		utilitys.LogError(err)
		return _products, err
	}

	return _products, nil
}

func (r *repository) Update(_products Products) (Products, error) {
	querys := "update products set name = $1,description = $2, price = $3,stock = $4 where id = $5"
	_, err := r.db.Exec(querys,
		_products.Name,
		_products.Description,
		_products.Price,
		_products.Stock,
		_products.Id,
	)

	if err != nil {
		return _products, err
	}

	return _products, nil
}

func (r *repository) Delete(_products Products) (Products, error) {

	querys := "delete from products where Id = $1"
	_, err := r.db.Exec(querys,
		_products.Id,
	)
	if err != nil {
		return _products, err
	}
	return _products, nil
}

func (r *repository) GetById(Id int) (Products, error) {
	var products = NewProducts()
	var collproducts []Products
	rows, err := r.db.Query("select * from products where id = $1", Id)
	if err != nil {
		utilitys.LogError(err)
		return products, err
	}
	defer rows.Close()
	for rows.Next() {

		err2 := rows.Scan(&products.Id,
			&products.Name,
			&products.Description,
			&products.Price,
			&products.Stock,
		)
		if err2 != nil {
			return products, err2
		}
		collproducts = append(collproducts, products)

	}

	return collproducts[0], nil
}

func (r *repository) GetAll() ([]Products, error) {
	var products = NewProducts()
	var collproducts []Products
	rows, err := r.db.Query("select * from products")
	if err != nil {
		utilitys.LogError(err)
		return collproducts, err
	}
	defer rows.Close()
	for rows.Next() {

		err2 := rows.Scan(&products.Id,
			&products.Name,
			&products.Description,
			&products.Price,
			&products.Stock,
		)
		if err2 != nil {
			return collproducts, err2
		}
		collproducts = append(collproducts, products)

	}

	return collproducts, nil
}
