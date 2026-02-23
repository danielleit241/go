package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

//Register middleware
// 1. Global middleware: Applied to all routes using r.Use()
// 2. Route group middleware: Applied to a group of routes using group.Use()
// - Example: user := v1.Group("/users").Use(middleware.BasicMiddleware())
// 3. Endpoint middleware: Applied to specific endpoints by adding the middleware function as an argument in the route definition
// - Example: user.GET("/", middleware.BasicMiddleware(), userController.GetUsers) -> Warning: The order of middleware matters

func BasicMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Before starting the next handler
		log.Printf("Start method: %s, path: %s", c.Request.Method, c.Request.URL.Path)
		// c.Writer.Write([]byte("Start middleware\n"))

		c.Next() // Call the next handler in the chain

		// After the next handler has finished
		log.Printf("End method: %s, path: %s", c.Request.Method, c.Request.URL.Path)
	}
}
