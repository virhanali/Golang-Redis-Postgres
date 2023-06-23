package handler

import (
	"ginredis/app/repositories"
	"ginredis/app/usecase"
	"ginredis/db"
)

type Container struct {
	UserHandler *ListHandler
}

func NewContainer() *Container {
	postgresDB := db.GetPostgresDB()
	redis := db.GetRedisDB()

	userRepo := repositories.NewListRepository(postgresDB, redis)
	userUseCase := usecase.NewListUsecase(userRepo)
	userHandler := NewListHandler(userUseCase)

	return &Container{
		UserHandler: userHandler,
	}
}
