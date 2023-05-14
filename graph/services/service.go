package services

import (
	"context"
	"model"

	"github.com/volatiletech/sqlboiler/boil"
)

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type Services struct {
	*userService
}

func New(exec boil.ContextExecutor) Services {
	return Services{
		userService: &userService{exec: exec},
	}
}
