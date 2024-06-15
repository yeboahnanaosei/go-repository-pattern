package sqlite

import (
	"context"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/aglili/go-expense/models"
)

// Repo implements the user repository interface.
// Call New() to create a new instance.
type Repo struct {
	db *gorm.DB
}

// NewRepo creates a new user repository.
func NewRepo(dsn string) (*Repo, error) {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection failed: %w", err)
	}

	// I guess you can auto migrate here
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate: %w", err)
	}

	return &Repo{db: db}, nil
}

func (r *Repo) CreateUser(ctx context.Context, user models.User) error {
	return r.db.Create(&user).Error
}

func (r *Repo) GetExistingUser(ctx context.Context, email, username string) (models.User, error) {
	var user models.User
	err := r.db.Where("email = ? or username = ?", email, username).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
