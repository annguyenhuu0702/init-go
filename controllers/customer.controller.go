package controllers

import (
	"log"
	"myapi/dtos"
	"myapi/models"
	"myapi/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	var input dtos.CreateCustomerDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := models.Customer{
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
		Email:       input.Email,
	}

	services.CreateCustomer(c, customer)
}

func GetCustomers(c *gin.Context) {
	services.GetCustomers(c)
}

func GetCustomer(c *gin.Context) {
	services.GetCustomer(c)
}

func UpdateCustomer(c *gin.Context) {
	var input dtos.UpdateCustomerDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Binding error: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customer := models.Customer{
		Name:  input.Name,
		Email: input.Email,
	}

	services.UpdateCustomer(c, customer)
}

func DeleteCustomer(c *gin.Context) {
	services.DeleteCustomer(c)
}
