package dependencyinjection

import (
	"database/sql"
	"tesptzen/cmd/manajemenpesanan"
	"tesptzen/cmd/order"
	"tesptzen/handler"
)

func InitManajemenPesananHandler(db *sql.DB) *handler.ManagemenPesananHandler {
	repository := order.NewRepository(db)
	services := manajemenpesanan.NewService(repository)
	handler := handler.NewManagemenPesananHandler(services)
	return handler
}
