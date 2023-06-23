package repositories

import (
	"context"
	"ginredis/app/entities"
	"ginredis/response"
)

type IUser interface {
	ListUsers(ctx context.Context, filter UserFilter) ([]entities.Users, *response.MetaTpl, error)
}
