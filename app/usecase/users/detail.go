package users

import (
	"context"
	"ginredis/app/repositories"
)

type DetailRequest struct {
	ID int
}

type DetailResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type DetailUsecase interface {
	Execute(ctx context.Context, filter DetailRequest) (*DetailResponse, error)
}

type detailUsecase struct {
	userRepo repositories.IUser
}

func NewDetailUsecase(userRepo repositories.IUser) DetailUsecase {
	return &detailUsecase{
		userRepo: userRepo,
	}
}

func (uc *detailUsecase) Execute(ctx context.Context, filter DetailRequest) (*DetailResponse, error) {
	user, err := uc.userRepo.Get(ctx, filter.ID)
	if err != nil {
		return nil, err
	}

	return &DetailResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}, nil
}
