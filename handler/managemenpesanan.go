package handler

import (
	"net/http"
	"tesptzen/cmd/manajemenpesanan"
	"tesptzen/cmd/order"
	"tesptzen/helpers"
	"tesptzen/utilitys"

	"github.com/gin-gonic/gin"
)

type ManagemenPesananHandler struct {
	manajemenpesananService manajemenpesanan.IManajemenPesananServices
}

func NewManagemenPesananHandler(manajemenpesananService manajemenpesanan.IManajemenPesananServices) *ManagemenPesananHandler {
	return &ManagemenPesananHandler{manajemenpesananService}
}

func (u *ManagemenPesananHandler) Tambah(c *gin.Context) {
	var input order.CreateRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		utilitys.LogError(err)
		errors := helpers.FormatError(err)
		errormessage := H{"errors": errors}
		response := helpers.BaseCommandResponse("Products Failed to Add", http.StatusUnprocessableEntity, "error", errormessage)
		c.JSON(http.StatusUnprocessableEntity, response)

	}
	_, err2 := u.manajemenpesananService.Create(input)
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

func (u *ManagemenPesananHandler) Edit(c *gin.Context) {
	var input order.EditRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		utilitys.LogError(err)
		errors := helpers.FormatError(err)
		errormessage := H{"errors": errors}
		response := helpers.BaseCommandResponse("Products Failed to Edit", http.StatusUnprocessableEntity, "error", errormessage)
		c.JSON(http.StatusUnprocessableEntity, response)

	}
	_, err2 := u.manajemenpesananService.Edit(input)
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

func (u *ManagemenPesananHandler) Hapus(c *gin.Context) {
	var input order.HapusRequest
	err := c.ShouldBindJSON(&input)
	if err != nil {
		utilitys.LogError(err)
		errors := helpers.FormatError(err)
		errormessage := H{"errors": errors}
		response := helpers.BaseCommandResponse("Products Failed to Delete", http.StatusUnprocessableEntity, "error", errormessage)
		c.JSON(http.StatusUnprocessableEntity, response)
	}
	_, err2 := u.manajemenpesananService.Hapus(input)
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

func (u *ManagemenPesananHandler) GetAll(c *gin.Context) {
	res, err := u.manajemenpesananService.GetList()
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
