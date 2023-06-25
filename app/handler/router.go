package handler

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(container *Container) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", container.ListUserHandler.Execute)
		v1.GET("/users/:id", container.GetUserHandler.Execute)
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	return router
}
