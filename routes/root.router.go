package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRootRouter(r *gin.Engine) {
	root := r.Group("/v1/api")

	SetupUserRoutes(root)
}
