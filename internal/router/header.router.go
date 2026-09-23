package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rlanakomara7/koda-b9-gin.git/internal/dto"
	// "github.com/rlanakomara7/koda-b9-gin.git/internal/middleware"
)

func initHeaderRouter(r *gin.Engine) {
	headerRouter := r.Group("/header")

	//EXAMPLE GET HEADER
	headerRouter.GET("", func(ctx *gin.Context) {
		//getheader
		auth := ctx.GetHeader("Authorization")
		ctype := ctx.GetHeader("Content-Type")
		cookies := ctx.GetHeader("Cookie")
		cookie, _ := ctx.Cookie("aaa") //mengambil key nya saja
		uagent := ctx.GetHeader("User-Agent")
		custom := ctx.GetHeader("XXX-Header")

		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Data: gin.H{
				"auth":    auth,
				"ctype":   ctype,
				"cookies": cookies,
				"cookie":  cookie,
				"uagent":  uagent,
				"custom":  custom,
			},
		})
	})

	//EXAMPLE GET PATH PARAM / PATH VARIABLE
	headerRouter.GET("/data/:id/:slug", func(ctx *gin.Context) {
		id := ctx.Param("id")
		slug := ctx.Param("slug")

		//cetak response kirim response
		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Data: gin.H{
				"id":   id,
				"slug": slug,
			},
		})
	})

	// EXAMPLE GET QUERY PARAM UNTUK PAGINASI
	headerRouter.GET("/query", func(ctx *gin.Context) {
		title := ctx.Query("title") // key nya lebih dari 1x pakai queryarray
		genres := ctx.QueryArray("genre")

		var qp dto.HeaderQuery
		if err := ctx.ShouldBindWith(&qp, binding.Query); err != nil {
			log.Println("error", err.Error())

			//binding error
			ctx.JSON(http.StatusInternalServerError, dto.Response{
				Success: false,
				Data:    nil,
				Msg:     "Terjadi kesalahan server",
			})
			return
		}

		ctx.JSON(http.StatusOK, dto.Response{
			Success: true,
			Data: gin.H{
				"title":  title,
				"genres": genres,
				"qp":     qp,
			},
		})
	})

}
