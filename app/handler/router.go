package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(container *Container) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", container.UserHandler.Execute)
	}

	return router
}
