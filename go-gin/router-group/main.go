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
			userV2.GET("/:id", userV2Handler.GetUserByIDV2)
			userV2.POST("/", userV2Handler.CreateUserV2)
			userV2.PUT("/:id", userV2Handler.UpdateUserV2)
			userV2.DELETE("/:id", userV2Handler.DeleteUserV2)
		}
	}

	r.Run(":8080")
}
