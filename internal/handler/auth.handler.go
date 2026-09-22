package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-gin.git/internal/dto"
	"github.com/rlanakomara7/koda-b9-gin.git/internal/service"
)

// type IAuthService interface {
// 	EmptyValidation(data dto.Account) error
// }

type AuthHandler struct { // tujuan struct untuk initialisasi
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var register dto.Account
	if err := ctx.ShouldBindJSON(&register); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Msg:     "Format input tidak valid",
		})
		return
	}

	if err := h.service.RegisterUser(register); err != nil {
		status := http.StatusBadRequest
		if err.Error() == "Email sudah terdaftar" {
			status = http.StatusConflict
		}
		ctx.JSON(status, dto.Response{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Data: gin.H{
			"email": register.Email,
		},
		Msg: "User Created",
	})
}

// ------------login
func (h *AuthHandler) Login(ctx *gin.Context) {
	var data dto.Account
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Msg:     "Format input tidak valid",
		})
		return
	}

	storedUser, err := h.service.LoginUser(data)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data: gin.H{
			"email": storedUser.Email,
		},
		Msg: h.service.FormatLoginMessage(data.Email),
	})
}

// func (h *AuthHandler) Register(ctx *gin.Context) {
// 	//deklarasi body
// 	var data dto.Account

// 	//data binding
// 	if e := ctx.ShouldBindWith(&data, binding.JSON); e != nil {
// 		ctx.JSON(http.StatusBadRequest, dto.Response{
// 			Success: false,
// 			Data:    nil,
// 			Msg:     "terjadi kesalahan server",
// 		})
// 		return
// 	}

// 	// validasi **************** ini belum akurat
// 	if len(data.Username) <= 6 || len(data.Email) <= 6 || len(data.Password) <= 6 {
// 		ctx.JSON(http.StatusBadRequest, dto.Response{
// 			Success: false,
// 			Data:    data,
// 			Msg:     "Panjang Harus Lebih Dari 6 Karakter",
// 		})
// 		return
// 	}

// 	//success
// 	ctx.JSON(http.StatusCreated, dto.Response{
// 		Success: true,
// 		Data:    data,
// 		Msg:     "Selamat Berhasil Register",
// 	})
// 	return

// }
