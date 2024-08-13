package main

import (
	"tesptzen/api/routes"

	"github.com/gin-gonic/gin"
)

// @title My Gin API
// @version 1.0
// @description Ini adalah API sederhana menggunakan Gin dan OpenAPI.
// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()
	routes.NewRoute(router)
	router.Run(":8181")
}
