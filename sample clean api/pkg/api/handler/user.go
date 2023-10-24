package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	interfaces "github.com/jaseelaali/Sample_API_using_Clean-Architecture/pkg/usecase/interface"
)

type UserHandler struct {
	useCase interfaces.UserUsecase
}

func NewHandler(usecase interfaces.UserUsecase) *UserHandler {
	return &UserHandler{
		useCase: usecase,
	}
}
func (h *UserHandler) Wish(r *gin.Context) {
	user, err := h.useCase.Wish(r.Request.Context())
	if err != nil {
		r.AbortWithStatus(http.StatusNotFound)
	} else {
		r.JSON(200, user)
	}
}
