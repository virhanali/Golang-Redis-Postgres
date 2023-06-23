package usecase

import (
	"context"
	"ginredis/app/repositories"
	"ginredis/app/entities"
	"ginredis/response"
)

type ListRequest struct {
	Page  int
	Limit int

	Name string
}

type ListResponse struct {
	Data       []entities.Users    `json:"users"`
	Pagination response.Pagination `json:"pagination"`
}

type ListUsecase interface {
	Execute(ctx context.Context, filter ListRequest) (*ListResponse, error)
}

type listUsecase struct {
	userRepo repositories.IUser
}

func NewListUsecase(userRepo repositories.IUser) ListUsecase {
	return &listUsecase{
		userRepo: userRepo,
	}
}

func (uc *listUsecase) Execute(ctx context.Context, filter ListRequest) (*ListResponse, error) {
	users, pagination, err := uc.userRepo.ListUsers(ctx, repositories.UserFilter{
		Page:  filter.Page,
		Limit: filter.Limit,
		Name:  filter.Name,
	})
	if err != nil {
		return nil, err
	}

	return &ListResponse{
		Data: users,
		Pagination: response.Pagination{
			CurrentPage: pagination.Page,
			TotalData:   pagination.TotalData,
		},
	}, nil
}
