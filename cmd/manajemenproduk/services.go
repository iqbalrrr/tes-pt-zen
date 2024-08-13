package manajemenproduk

import (
	"tesptzen/cmd/product"
	"tesptzen/utilitys"
)

type IManajemenProdukServices interface {
	Create(input product.CreateRequest) (product.Products, error)
	Edit(input product.EditRequest) (product.Products, error)
	Hapus(input product.HapusRequest) (product.Products, error)
	GetList() ([]product.Products, error)
}

type manajemenProdukServices struct {
	repository product.IRepository
}

func NewService(repository product.IRepository) *manajemenProdukServices {
	return &manajemenProdukServices{repository}
}

func (s *manajemenProdukServices) Create(input product.CreateRequest) (product.Products, error) {
	_products := product.NewProducts()
	_products.Name = input.Name
	_products.Description = input.Description
	_products.Price = input.Price
	_products.Stock = input.Stock

	res, err := s.repository.Insert(_products)
	if err != nil {
		utilitys.LogError(err)
		return _products, err
	}
	return res, err
}

func (s *manajemenProdukServices) Edit(input product.EditRequest) (product.Products, error) {

	_products, err := s.repository.GetById(input.Id)
	if err != nil {
		utilitys.LogError(err)
		return _products, err
	}
	_products.Name = input.Name
	_products.Description = input.Description
	_products.Price = input.Price
	_products.Stock = input.Stock

	res, err := s.repository.Update(_products)
	if err != nil {
		utilitys.LogError(err)
		return _products, err
	}
	return res, nil
}

func (s *manajemenProdukServices) Hapus(input product.HapusRequest) (product.Products, error) {

	_products, err := s.repository.GetById(input.Id)
	if err != nil {
		utilitys.LogError(err)
		return _products, err
	}

	res, err := s.repository.Delete(_products)
	if err != nil {
		utilitys.LogError(err)
		return _products, err
	}
	return res, nil
}

func (s *manajemenProdukServices) GetList() ([]product.Products, error) {

	_products, err := s.repository.GetAll()
	if err != nil {
		utilitys.LogError(err)
		return _products, err
	}

	return _products, nil
}
