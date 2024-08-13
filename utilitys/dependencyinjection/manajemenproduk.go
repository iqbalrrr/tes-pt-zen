package dependencyinjection

import (
	"database/sql"
	"tesptzen/cmd/manajemenproduk"
	"tesptzen/cmd/product"
	"tesptzen/handler"
)

func InitManajemenProdukHandler(db *sql.DB) *handler.ManagemenProdukHandler {
	repository := product.NewRepository(db)
	services := manajemenproduk.NewService(repository)
	handler := handler.NewManagemenProdukHandler(services)
	return handler
}
