package users

import (
	"ginredis/app/usecase/users"
	"ginredis/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DetailHandler struct {
	userUsecase users.DetailUsecase
}

func NewDetailHandler(userUsecase users.DetailUsecase) *DetailHandler {
	return &DetailHandler{
		userUsecase: userUsecase,
	}
}

func (h *DetailHandler) Execute(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	req := users.DetailRequest{
		ID: id,
	}

	user, err := h.userUsecase.Execute(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, response.BasePayload{
		Success: true,
		Message: "Success get user",
		Data:    user,
	})
}
