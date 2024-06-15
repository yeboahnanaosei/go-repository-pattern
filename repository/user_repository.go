package repository

import (
	"context"

	"github.com/aglili/go-expense/models"
)

/*
* You can define the repository methods here without actually implementing them.
* You can have several concrete implementations of the repository by any "database" you wish.
* The idea is that it makes it easy to replace the database implementation in the future.
* The concrete implementation can then be injected into the service layer or wherever you need a repository.
*
* Take a look at database/postgres/repo.go for an example
* I also implemented a sqlite repository inside database/sqlite.
* They are all implementations of the user repository interface.
* Wherever you need a user repo, you can use any one of them.
* In future you can add as many repositories as you need like firebase, redis, etc.
 */

// UserRepository is a repository for handling user data.
type UserRepository interface {
	CreateUser(ctx context.Context, user models.User) error
	GetExistingUser(ctx context.Context, email, username string) (models.User, error)
}
