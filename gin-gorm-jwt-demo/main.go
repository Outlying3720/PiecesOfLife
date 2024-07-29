package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"demo/controllers"
	"demo/middlewares"
	"demo/models"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()

	public := r.Group("/api")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}

	protected := r.Group("/api/admin")
	{
		protected.Use(middlewares.JwtAuthMiddleware())
		protected.GET("/user", func(ctx *gin.Context) {
			auth, _ := ctx.Get("auth")
			ctx.JSON(http.StatusOK, gin.H{
				"message": auth.(string),
			})
		})
	}

	r.Run(":8000")
}
