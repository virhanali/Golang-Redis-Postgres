package handler

import (
	"ginredis/app/handler/users"
	"ginredis/app/repositories"
	userUseCase "ginredis/app/usecase/users"

	"ginredis/db"
)

type Container struct {
	ListUserHandler *users.ListHandler
	GetUserHandler  *users.DetailHandler
}

func NewContainer() *Container {
	postgresDB := db.GetPostgresDB()
	redis := db.GetRedisDB()

	userRepo := repositories.NewRepository(postgresDB, redis)
	
	listuserUseCase := userUseCase.NewListUsecase(userRepo)
	listuserHandler := users.NewListHandler(listuserUseCase)

	getuserUseCase := userUseCase.NewDetailUsecase(userRepo)
	getuserHandler := users.NewDetailHandler(getuserUseCase)

	return &Container{
		ListUserHandler: listuserHandler,
		GetUserHandler:  getuserHandler,
	}
}
