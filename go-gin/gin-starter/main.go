package main

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = map[int]User{
	1: {ID: 1, Name: "Alice", Age: 30},
	2: {ID: 2, Name: "Bob", Age: 25},
	3: {ID: 3, Name: "Charlie", Age: 35},
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/users", func(c *gin.Context) {
		// Query parameters: name (partial match), age (exact match)
		searchName := strings.ToLower(strings.TrimSpace(c.Query("name")))
		searchAgeParam := strings.TrimSpace(c.Query("age"))

		searchAge, err := strconv.Atoi(searchAgeParam)
		if searchAgeParam != "" && err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid age parameter",
			})
			return
		}

		filteredUsers := make([]User, 0)

		for _, user := range users {
			isMatch := false

			if searchName == "" && searchAgeParam == "" {
				isMatch = true
			} else {
				if searchName != "" && strings.Contains(strings.ToLower(user.Name), searchName) {
					isMatch = true
				}

				if searchAgeParam != "" && user.Age == searchAge {
					isMatch = true
				}
			}

			if isMatch {
				filteredUsers = append(filteredUsers, user)
			}
		}

		c.JSON(200, gin.H{
			"users": filteredUsers,
		})

		c.JSON(200, gin.H{
			"users": filteredUsers,
		})
	})

	router.GET("/users/:user_id", func(c *gin.Context) {
		// Path parameter: user_id
		userIDParam := c.Param("user_id")

		userID, err := strconv.Atoi(userIDParam)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid user ID",
			})
			return
		}

		user, exists := users[userID]
		if !exists {
			c.JSON(404, gin.H{
				"error": "User not found",
			})
			return
		}

		c.JSON(200, gin.H{
			"user": user,
		})
	})

	router.Run(":8080")
}
