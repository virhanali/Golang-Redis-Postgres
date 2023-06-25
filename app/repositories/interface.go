package repositories

import (
	"context"
	"ginredis/app/entities"
	"ginredis/response"
)

type IUser interface {
	List(ctx context.Context, filter UserFilter) ([]entities.Users, *response.MetaTpl, error)
	Get(ctx context.Context, id int) (entities.Users, error)
}
