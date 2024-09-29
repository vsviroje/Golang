package users

import (
	"context"

	"task_management_system/errors"
)

//go:generate mockery --name IUsersRepo --inpackage --filename=users_repo_mock.go
type IUsersRepo interface {
	GetUserById(ctx context.Context, id string) (*Users, errors.IError)
	GetUserByEmail(ctx context.Context, email string) (*Users, errors.IError)
	AddUsers(ctx context.Context, data *Users) (string, errors.IError)
	UpdateUsers(ctx context.Context, data *Users) errors.IError
	DeleteUsersById(ctx context.Context, id string) errors.IError
}
