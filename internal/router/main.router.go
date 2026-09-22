package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-gin.git/internal/dto"
)

func InitMainRouter() *gin.Engine {
	router := gin.Default()
	initAuthRouter(router)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Msg:     "rute salah",
			Success: false,
		})
	})

	return router
}
