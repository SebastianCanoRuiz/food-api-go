package repository

import (
	"context"
	"go-api/domain/user/application/v1/response"
	"go-api/domain/user/domain/model"
)

type UserRepository interface {
	GetAllUser(ctx context.Context) ([]response.UserResponse, error)
	GetById(ctx context.Context, id string) (response.UserResponse, error)
	CreateUser(ctx context.Context, user *model.User) (*response.UserResponse, error)
	UpdateUser(ctx context.Context, id string, user model.User) error
	DeleteUser(ctx context.Context, id string) error
}
