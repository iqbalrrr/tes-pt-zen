package dependencyinjection

import (
	"database/sql"
	"tesptzen/cmd/appuser"
	"tesptzen/cmd/auth"
	"tesptzen/handler"
)

func InitAuthServices(db *sql.DB) *handler.UserHandler {
	repository := appuser.NewRepository(db)
	services := auth.NewService(repository)
	handler := handler.NewUserHandler(services)
	return handler
}
