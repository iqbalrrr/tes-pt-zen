package handler

import (
	"net/http"
	"tesptzen/cmd/manajemenproduk"
	"tesptzen/cmd/product"
	"tesptzen/helpers"
	"tesptzen/utilitys"

	"github.com/gin-gonic/gin"
)

type ManagemenProdukHandler struct {
	manajemenprodukService manajemenproduk.IManajemenProdukServices
}

func NewManagemenProdukHandler(manajemenprodukService manajemenproduk.IManajemenProdukServices) *ManagemenProdukHandler {
	return &ManagemenProdukHandler{manajemenprodukService}
}

func (u *ManagemenProdukHandler) Tambah(c *gin.Context) {
	var input product.CreateRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		utilitys.LogError(err)
		errors := helpers.FormatError(err)
		errormessage := H{"errors": errors}
		response := helpers.BaseCommandResponse("Products Failed to Add", http.StatusUnprocessableEntity, "error", errormessage)
		c.JSON(http.StatusUnprocessableEntity, response)

	}
	_, err2 := u.manajemenprodukService.Create(input)
	if err2 != nil {
		utilitys.LogError(err2)
		errors := helpers.FormatError(err2)
		errormessage := H{"errors": errors}
		response := helpers.BaseCommandResponse("Products Failed to Add", http.StatusUnprocessableEntity, "error", errormessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}
	response := helpers.BaseCommandResponse("Products Have Been Added", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (u *ManagemenProdukHandler) Edit(c *gin.Context) {
	var input product.EditRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		utilitys.LogError(err)
		errors := helpers.FormatError(err)
		errormessage := H{"errors": errors}
		response := helpers.BaseCommandResponse("Products Failed to Edit", http.StatusUnprocessableEntity, "error", errormessage)
		c.JSON(http.StatusUnprocessableEntity, response)

	}
	_, err2 := u.manajemenprodukService.Edit(input)
	if err2 != nil {
		utilitys.LogError(err2)
		errors := helpers.FormatError(err2)
		errormessage := H{"errors": errors}
		response := helpers.BaseCommandResponse("Products Failed to Edit", http.StatusUnprocessableEntity, "error", errormessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}
	response := helpers.BaseCommandResponse("Products Have Been Edit", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (u *ManagemenProdukHandler) Hapus(c *gin.Context) {
	var input product.HapusRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		utilitys.LogError(err)
		errors := helpers.FormatError(err)
		errormessage := H{"errors": errors}
		response := helpers.BaseCommandResponse("Products Failed to Delete", http.StatusUnprocessableEntity, "error", errormessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}
	_, err2 := u.manajemenprodukService.Hapus(input)
	if err2 != nil {
		utilitys.LogError(err2)
		errors := helpers.FormatError(err2)
		errormessage := H{"errors": errors}
		response := helpers.BaseCommandResponse("Products Failed to Delete", http.StatusUnprocessableEntity, "error", errormessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}
	response := helpers.BaseCommandResponse("Products Have Been Delete", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}

func (u *ManagemenProdukHandler) GetAll(c *gin.Context) {
	res, err := u.manajemenprodukService.GetList()
	if err != nil {
		utilitys.LogError(err)
		errors := helpers.FormatError(err)
		errormessage := H{"errors": errors}
		response := helpers.BaseCommandResponse("Failed", http.StatusUnprocessableEntity, "error", errormessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}
	response := helpers.BaseCommandResponse("Ok", http.StatusOK, "success", res)
	c.JSON(http.StatusOK, response)
}
