package Auth

import (
	"authentication-jwt-go/src/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserID   uint   `json:"user_id"`
	Paswword string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	fmt.Println(c)

	// Example of parsing JSON request body into a User struct
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	token, err := utils.GenerateToken(user.UserID)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}
