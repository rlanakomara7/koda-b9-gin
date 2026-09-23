package main

import (
	"github.com/rlanakomara7/koda-b9-gin.git/internal/router"
)

func main() {
	//generate enginge / router
	Router := router.InitMainRouter() // yg warna biru hijau diterminal merupakan middleware

	//run
	Router.Run("localhost:9000") // di fungsi run harus addres , port bebas

}
