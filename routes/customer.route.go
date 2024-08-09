package routes

import (
	"myapi/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(rg *gin.RouterGroup) {
	customerRoutes := rg.Group("/customer")
	{
		customerRoutes.POST("/", controllers.CreateCustomer)
		customerRoutes.GET("/", controllers.GetCustomers)
		customerRoutes.GET("/:id", controllers.GetCustomer)
		customerRoutes.PATCH("/:id", controllers.UpdateCustomer)
		customerRoutes.DELETE("/:id", controllers.DeleteCustomer)
	}
}
