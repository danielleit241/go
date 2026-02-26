package routers

import (
	"github.com/danielleit241/internal/handler"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userHandler *handler.UserHandler
}

func NewUserRoutes(userHandler *handler.UserHandler) *UserRoutes {
	return &UserRoutes{
		userHandler: userHandler,
	}
}

func (ur *UserRoutes) Register(router *gin.RouterGroup) {
	// userGroup := router.Group("/users")
	// {
	// 	userGroup.GET("/", ur.userHandler.GetAllUsers)
	// 	userGroup.GET("/:id", ur.userHandler.GetUserByID)
	// 	userGroup.POST("/", ur.userHandler.CreateUser)
	// 	userGroup.PUT("/:id", ur.userHandler.UpdateUser)
	// 	userGroup.DELETE("/:id", ur.userHandler.DeleteUser)
	// }
}
