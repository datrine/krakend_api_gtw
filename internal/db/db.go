package db

import (
	"context"

	sharedexports "github.com/datrine/basic_crud_with_auth/internal/shared_exports"
)

func NewConn() Repository {
	s := NewRemoteSqlLiteSQLDB()
	return s
}

type User struct {
}

type Repository interface {
	CreateUser(ctx context.Context, userTocreate *sharedexports.CreateUser) error
	QueryUserByEmail(ctx context.Context, email string) (*sharedexports.User, error)
	UpdateUser(ctx context.Context, email string, updates *sharedexports.UpdateUser) error
	DeleteUserByEmail(ctx context.Context, email string) (*sharedexports.User, error)
}
