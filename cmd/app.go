package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/aglili/go-expense/config"
	"github.com/aglili/go-expense/database/postgres"
	"github.com/aglili/go-expense/services/auth"
)

type apiService struct {
	authService *auth.Service
	router      *gin.Engine
}

// Setup the api service.
// It sets up all the required dependencies like the user repo and the router.
func setup() (*apiService, error) {
	// Set up a database or concrete repository.
	// You can replace this repo with sqlite, redis, firebase, etc.
	userRepo, err := postgres.NewRepo(config.GetDatabaseDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to set up db connection: %w", err)
	}

	// If you wanted, you could instead use the sqlite implementation of the repository.
	// They both implement or satisfy the user repository interface, hence they can be interchanged.
	/*
		userRepo, err := sqlite.NewRepo(config.GetDatabaseDSN())
		if err != nil {
			return nil, fmt.Errorf("failed to set up db connection: %w", err)
		}
	*/

	// Set up router.
	router := gin.Default()

	// Set up auth service
	authService := &auth.Service{UserRepo: userRepo}
	authService.RegisterRoutes(router)

	service := &apiService{authService: authService, router: router}
	return service, nil
}
