package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-gin.git/internal/dto"
	"github.com/rlanakomara7/koda-b9-gin.git/internal/middleware"
)

func InitMainRouter() *gin.Engine {
	router := gin.Default()

	//global middleware
	router.Use(middleware.Cors, middleware.M1, middleware.M2, middleware.M3)
	// router.Use(middleware.Logger())
	// router.Use(middleware.Cors)

	initAuthRouter(router)
	initHeaderRouter(router)

	// Custom response jika route tidak ditemukan (404)
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Msg:     "rute salah",
			Success: false,
		})
	})

	return router
}
