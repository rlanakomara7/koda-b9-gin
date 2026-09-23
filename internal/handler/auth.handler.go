package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rlanakomara7/koda-b9-gin.git/internal/dto"
	"github.com/rlanakomara7/koda-b9-gin.git/internal/service"
)

type AuthHandler struct { // tujuan struct untuk initialisasi
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var data dto.User
	if err := ctx.ShouldBindWith(&data, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false, Msg: "format data tidak sesuai",
		})
		return
	}

	loggedInUser, err := h.service.Login(data)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false, Data: data, Msg: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true, Data: loggedInUser, Msg: "Selamat Berhasil Login",
	})
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var data dto.User
	if err := ctx.ShouldBindWith(&data, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false, Msg: "format data tidak sesuai",
		})
		return
	}

	if err := h.service.Register(data); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false, Data: data, Msg: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true, Data: data, Msg: "Selamat Berhasil Register",
	})
}
