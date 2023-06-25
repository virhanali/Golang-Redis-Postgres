package users

import (
	"ginredis/app/usecase/users"
	"ginredis/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListHandler struct {
	userUsecase users.ListUsecase
}

func NewListHandler(userUsecase users.ListUsecase) *ListHandler {
	return &ListHandler{
		userUsecase: userUsecase,
	}
}

func (h *ListHandler) Execute(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	name := ctx.Query("name")

	filter := users.ListRequest{
		Page:  page,
		Limit: limit,
		Name:  name,
	}

	users, err := h.userUsecase.Execute(ctx, filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.BasePayload{
		Success: true,
		Message: "Success fetch users",
		Data:    users,
	})
}
