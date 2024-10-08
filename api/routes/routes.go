package routes

import (
	"tesptzen/middleware"
	"tesptzen/utilitys"
	"tesptzen/utilitys/dependencyinjection"

	"github.com/gin-gonic/gin"
)

func NewRoute(c *gin.Engine) *gin.Engine {
	DB := utilitys.InitDatabase()
	userhandlers := dependencyinjection.InitAuthServices(DB)
	mjmproduk := dependencyinjection.InitManajemenProdukHandler(DB)
	mjmpesanan := dependencyinjection.InitManajemenPesananHandler(DB)
	user := c.Group("/user")
	user.Use(middleware.TokenAuthMiddleware())
	{
		user.POST("/auth/resetpassword", userhandlers.ResetPasswordUser)

		user.POST("/product/tambah", mjmproduk.Tambah)
		user.POST("/product/hapus", mjmproduk.Hapus)
		user.POST("/product/edit", mjmproduk.Edit)

		user.POST("/order/tambah", mjmpesanan.Tambah)
		user.POST("/order/hapus", mjmpesanan.Hapus)
		user.POST("/order/edit", mjmpesanan.Edit)
	}
	c.POST("/auth/register", userhandlers.RegisterUser)
	c.POST("/auth/login", userhandlers.Login)

	c.GET("/product/getall", mjmproduk.GetAll)

	c.GET("/order/getall", mjmpesanan.GetAll)

	return c
}
