package auth

import (
	"github.com/gin-gonic/gin"

	"github.com/aglili/go-expense/repository"
)

// Service handles user authentication.
type Service struct {
	// UserRepo is the user repository.
	UserRepo repository.UserRepository
}

func (svc *Service) RegisterRoutes(router *gin.Engine) {
	router.POST("/sign_up", svc.handleSignUp)
	router.POST("/login", svc.handleLogin)
}
