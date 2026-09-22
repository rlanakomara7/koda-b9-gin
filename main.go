package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// type Response map[string]any //agar deskriptif ,response ctx.JSON , bisa pakai gin.H

type Body struct {
	// key datatype struct_tag = untuk memudahkan aplikasi membaca data slelain json bisa tipe yg lain example 'form'
	Name string `json:"nama"`
	Age  int8   `json:"umur"`
}

type Response struct {
	Success bool
	Data    any
	Msg     string
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var Users = []User{}

func main() {
	//generate enginge / router
	router := gin.Default() // yg warna biru hijau diterminal merupakan middleware

	//deklarasi router
	router.GET("/ping", func(ctx *gin.Context) {
		//send response , code = http , response = objek (any)
		ctx.JSON(http.StatusOK, gin.H{ //bisa dengan gin.H untuk tipe response
			"msg": "pong",
		})
	})

	router.GET("/hello", func(ctx *gin.Context) {
		//send response , code = http , response = objek (any)
		ctx.JSON(http.StatusOK, gin.H{ //bisa dengan gin.H untuk tipe response, jika response di deklarasi diatas func main
			"msg": "world",
		})
	})

	router.POST("/ping", func(ctx *gin.Context) {
		//deklarasi body
		var data Body                                               //ShouldBindWith = bisa tentukan binding nya .JSON , ShouldBin = tidak ditentukan , JSON untuk di Content-type saat test Post
		if e := ctx.ShouldBindWith(&data, binding.JSON); e != nil { //dibandigkan nil karna nil 0 value struct
			ctx.JSON(http.StatusInternalServerError, Response{ //karna diminta dokumentasi ,pakai pointer agar perubahan masuk variable data
				Success: false,
				Data:    nil,
				Msg:     "terjadi kesalahan server",
			})
			return
		}

		//bisa tambah validasi (bussines logic)
		if data.Name == "" && data.Age == 0 {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Data:    data,
				Msg:     "empty body",
			})
			return
		}

		//kalau berhasil
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Data:    data,
			Msg:     fmt.Sprint("Selamat Datang %s", data.Name), // name ada di variable body data
		})
	})

	//MINITASK 1 -------------- MINITASK 1
	router.POST("/login", func(ctx *gin.Context) {
		var data User
		if e := ctx.ShouldBindWith(&data, binding.JSON); e != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Data:    nil,
				Msg:     "terjadi kesalahan server",
			})
			return
		}

		// validasi
		if len(data.Email) <= 6 || len(data.Password) <= 6 {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Data:    data,
				Msg:     "Panjang Harus Lebih Dari 6 Karakter",
			})
			return
		}

		//success
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Data:    data,
			Msg:     "Selamat Berhasil Login",
		})
		return
	})

	router.POST("/register", func(ctx *gin.Context) {
		var data User
		if e := ctx.ShouldBindWith(&data, binding.JSON); e != nil {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Data:    nil,
				Msg:     "terjadi kesalahan server",
			})
			return
		}

		// validasi
		if len(data.Username) <= 6 || len(data.Email) <= 6 || len(data.Password) <= 6 {
			ctx.JSON(http.StatusBadRequest, Response{
				Success: false,
				Data:    data,
				Msg:     "Panjang Harus Lebih Dari 6 Karakter",
			})
			return
		}

		//success
		ctx.JSON(http.StatusOK, Response{
			Success: true,
			Data:    data,
			Msg:     "Selamat Berhasil Register",
		})
		return
	})

	//MINITASK -------------- MINITASK

	//run
	router.Run("localhost:9000") // di fungsi run harus addres , port bebas

}
