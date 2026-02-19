package main

import (
	v1handler "github.com/danielleit241/internal/api/v1"
	v2handler "github.com/danielleit241/internal/api/v2"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	userV1 := r.Group("/api/v1/users")
	{
		userV1Handler := v1handler.NewUserHandler()

		userV1.GET("/", userV1Handler.GetUsersV1)
		// Validate that the ID is a number and positive integer
		userV1.GET("/:id", userV1Handler.GetUserByIDV1)
		userV1.POST("/", userV1Handler.CreateUserV1)
		userV1.PUT("/:id", userV1Handler.UpdateUserV1)
		userV1.DELETE("/:id", userV1Handler.DeleteUserV1)
	}

	v2 := r.Group("/api/v2")
	{
		userV2Handler := v2handler.NewUserHandler()

		userV2 := v2.Group("/users")
		{
			userV2.GET("/", userV2Handler.GetUsersV2)
			// Validate that the ID is a valid UUID
			userV2.GET("/:uuid", userV2Handler.GetUserByIDV2)
			userV2.POST("/", userV2Handler.CreateUserV2)
			userV2.PUT("/:uuid", userV2Handler.UpdateUserV2)
			userV2.DELETE("/:uuid", userV2Handler.DeleteUserV2)
		}

		productV2Handler := v2handler.NewProductHandler()
		productV2 := v2.Group("/products")
		{
			productV2.GET("/", productV2Handler.GetProductsV2)
			productV2.GET("/search", productV2Handler.SearchProductsV2)
			// Validate that the Slug is a valid string (non-empty)
			productV2.GET("/:slug", productV2Handler.GetProductBySlugV2)
			// Validate that the category is one of the predefined categories
			productV2.GET("/category/:category", productV2Handler.GetProductsByCategoryV2)
			productV2.POST("/", productV2Handler.CreateProductV2)
			productV2.PUT("/:id", productV2Handler.UpdateProductV2)
			productV2.DELETE("/:id", productV2Handler.DeleteProductV2)
		}
	}

	r.Run(":8080")
}
