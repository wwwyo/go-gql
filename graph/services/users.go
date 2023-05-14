package services

import (
	"context"
	"my_gql_server/graph/db"
	"my_gql_server/graph/model"

	"github.com/volatiletech/sqlboiler/boil"
)

type userService struct {
	// dbのインターフェース
	exec boil.ContextExecutor
}

func (u *userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user, err := db.Users(
		db.UserWhere.Name.EQ(name),
	).One(ctx, u.exec)

	if err != nil {
		return nil, err
	}

	return convertUser(user), nil
}

func convertUser(user *db.User) *model.User {
	return &model.User{
		ID:   user.ID,
		Name: user.Name,
	}
}
