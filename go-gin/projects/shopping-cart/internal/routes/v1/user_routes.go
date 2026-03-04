package v1routers

import (
	v1handler "github.com/danielleit241/internal/handler/v1"
	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	userHandler *v1handler.UserHandler
}

func NewUserRoutes(userHandler *v1handler.UserHandler) *UserRoutes {
	return &UserRoutes{
		userHandler: userHandler,
	}
}

func (ur *UserRoutes) Register(router *gin.RouterGroup) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", ur.userHandler.GetAllUsers)
		userGroup.GET("/:id", ur.userHandler.GetUserByID)
		userGroup.POST("/", ur.userHandler.CreateUser)
		userGroup.PUT("/:id", ur.userHandler.UpdateUser)
		userGroup.DELETE("/:id", ur.userHandler.DeleteUser)
	}
}
