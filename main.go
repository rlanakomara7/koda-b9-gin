package main

import (
	"github.com/rlanakomara7/koda-b9-gin.git/internal/router"
)

// var Users = []Account{}

func main() {
	//generate enginge / router
	Router := router.InitMainRouter() // yg warna biru hijau diterminal merupakan middleware

	//MINITASK 1 -------------- MINITASK 1
	// router.POST("/login", func(ctx *gin.Context) {
	// 	var data Account
	// 	if e := ctx.ShouldBindWith(&data, binding.JSON); e != nil {
	// 		ctx.JSON(http.StatusBadRequest, Response{
	// 			Success: false,
	// 			Data:    nil,
	// 			Msg:     "terjadi kesalahan server",
	// 		})
	// 		return
	// 	}

	// 	// validasi
	// 	if len(data.Email) <= 6 || len(data.Password) <= 6 {
	// 		ctx.JSON(http.StatusBadRequest, Response{
	// 			Success: false,
	// 			Data:    data,
	// 			Msg:     "Panjang Harus Lebih Dari 6 Karakter",
	// 		})
	// 		return
	// 	}

	// 	//success
	// 	ctx.JSON(http.StatusOK, Response{
	// 		Success: true,
	// 		Data:    data,
	// 		Msg:     "Selamat Berhasil Login",
	// 	})
	// 	return
	// })

	//MINITASK -------------- MINITASK

	//run
	Router.Run("localhost:9000") // di fungsi run harus addres , port bebas

}
