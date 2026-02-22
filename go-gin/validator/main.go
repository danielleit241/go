package main

import (
	"github.com/danielleit241/internal/api"
	"github.com/danielleit241/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	if err := utils.RegisterValidators(); err != nil {
		panic("Failed to register custom validators: " + err.Error())
	}

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			userController := api.NewUserController()
			user.GET("/", userController.GetUsers)
			user.GET("/:id", userController.GetUserByID)
		}

		product := v1.Group("/products")
		{
			productController := api.NewProductController()
			product.GET("/", productController.GetProducts)
			product.GET("/search", productController.SearchProducts)
			product.GET("/category/:category", productController.GetProductsByCategory)
			product.GET("/slug/:slug", productController.GetProductBySlug)
			product.GET("/:id", productController.GetProductByID)
		}
	}

	r.Run(":8080")
}
