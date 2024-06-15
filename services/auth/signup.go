package auth

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/aglili/go-expense/models"
)

// handleSignUp handles the request to create or sign up a new user.
func (svc *Service) handleSignUp(c *gin.Context) {
	var body struct {
		Email    string `json:"email"    binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	user := models.User{
		Email:    body.Email,
		Username: body.Username,
		Password: string(hash),
	}

	err = svc.UserRepo.CreateUser(c, user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(200, gin.H{"message": "User created successfully"})
}
