package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("sign-in")
		auth.POST("sign-up")
	}

	api := router.Group("/api")
	{
		orders := api.Group("/orders")
		{
			orders.POST("/")
			orders.GET("/")
			orders.GET("/id")
			orders.PUT("/:id")
			orders.DELETE("/:id")

			items := orders.Group("/items")
			{
				items.GET("/")
				items.GET("/:item_id")
				items.POST("/")
				items.PUT("/:item_id")
				items.DELETE("/:item_id")
			}
		}
	}

	return router
}
