package main

import (
	"github.com/gin-gonic/gin"

	"gin-jwt/controllers"
	"gin-jwt/middlewares"
	"gin-jwt/models"
)

func init() {
	models.ConnectDatabase()
}

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
		protected.GET("/user", controllers.CurrentUser)
	}

	r.Run("0.0.0.0:8000")
}
