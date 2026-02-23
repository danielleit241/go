package main

import (
	"github.com/danielleit241/internal/api"
	"github.com/danielleit241/middleware"
	"github.com/danielleit241/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := utils.RegisterValidators(); err != nil {
		panic("Failed to register custom validators: " + err.Error())
	}

	if err := godotenv.Load(); err != nil {
		panic("Failed to load .env file: " + err.Error())
	}

	r := gin.Default()

	//Global middleware
	r.Use(middleware.BasicMiddleware(), middleware.RateLimit())

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			userController := api.NewUserController()
			user.GET("/", userController.GetUsers)
			user.GET("/:id", userController.GetUserByID)
			user.POST("/", userController.CreateUser)
			user.PUT("/:id/avatar", userController.UpdateUserAvatar)
		}

		product := v1.Group("/products")
		{
			productController := api.NewProductController()
			product.GET("/", productController.GetProducts)
			product.GET("/search", productController.SearchProducts)
			product.GET("/category/:category", productController.GetProductsByCategory)
			product.GET("/slug/:slug", productController.GetProductBySlug)
			product.GET("/:id", productController.GetProductByID)
			product.POST("/", productController.CreateProduct)
		}

		auth := v1.Group("/auth")
		{
			auth.Use(middleware.AdminSecret())
			authController := api.NewAuthController()
			auth.POST("/api-key", authController.GenerateAPIKey)
			auth.DELETE("/api-key/:id", authController.DeleteAPIKey)
		}
	}

	r.StaticFS("/images", gin.Dir("./uploads", false))

	r.Run(":8080")
}
