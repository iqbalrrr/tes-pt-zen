package manajemenpesanan

import (
	"tesptzen/cmd/order"
	"tesptzen/utilitys"
)

type IManajemenPesananServices interface {
	Create(input order.CreateRequest) (order.Orders, error)
	Edit(input order.EditRequest) (order.Orders, error)
	Hapus(input order.HapusRequest) (order.Orders, error)
	GetList() ([]order.Orders, error)
}

type manajemenPesananServices struct {
	repository order.IRepository
}

func NewService(repository order.IRepository) *manajemenPesananServices {
	return &manajemenPesananServices{repository}
}

func (s *manajemenPesananServices) Create(input order.CreateRequest) (order.Orders, error) {
	_orders := order.Neworders()
	_orders.Customer_name = input.Customer_name
	_orders.Product_id = input.Product_id
	_orders.Quantity = input.Quantity
	_orders.Status = input.Status

	res, err := s.repository.Insert(_orders)
	if err != nil {
		utilitys.LogError(err)
		return _orders, err
	}
	return res, err
}

func (s *manajemenPesananServices) Edit(input order.EditRequest) (order.Orders, error) {

	_orders, err := s.repository.GetById(input.Id)
	if err != nil {
		utilitys.LogError(err)
		return _orders, err
	}
	_orders.Customer_name = input.Customer_name
	_orders.Product_id = input.Product_id
	_orders.Quantity = input.Quantity
	_orders.Status = input.Status

	res, err := s.repository.Update(_orders)
	if err != nil {
		utilitys.LogError(err)
		return _orders, err
	}
	return res, nil
}

func (s *manajemenPesananServices) Hapus(input order.HapusRequest) (order.Orders, error) {

	_orders, err := s.repository.GetById(input.Id)
	if err != nil {
		utilitys.LogError(err)
		return _orders, err
	}

	res, err := s.repository.Delete(_orders)
	if err != nil {
		utilitys.LogError(err)
		return _orders, err
	}
	return res, nil
}

func (s *manajemenPesananServices) GetList() ([]order.Orders, error) {

	_orders, err := s.repository.GetAll()
	if err != nil {
		utilitys.LogError(err)
		return _orders, err
	}

	return _orders, nil
}
