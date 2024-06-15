package auth

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/aglili/go-expense/config"
)

// handleLogin receives the request to handle login for the auth service.
func (svc *Service) handleLogin(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	token, err := svc.login(c, body.Email, body.Username, body.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

// login is the actual function that performs the login.
func (svc *Service) login(ctx context.Context, email, username, password string) (string, error) {
	user, err := svc.UserRepo.GetExistingUser(ctx, email, username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenString, err := token.SignedString([]byte(config.GetJWTSecret()))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
