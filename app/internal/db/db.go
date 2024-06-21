package db

import (
	"context"

	sharedexports "github.com/datrine/basic_crud_with_auth/internal/shared_exports"
)

var repo Repository

func GetRepository() Repository {
	repo = NewRemoteSqlLiteSQLDB()
	return repo
}

type User struct {
}

type Repository interface {
	CreateUser(ctx context.Context, userTocreate *sharedexports.CreateUser) error
	QueryUserByEmail(ctx context.Context, email string) (*sharedexports.User, error)
	UpdateUser(ctx context.Context, email string, updates *sharedexports.UpdateUser) error
	DeleteUserByEmail(ctx context.Context, email string) (*sharedexports.User, error)
}
