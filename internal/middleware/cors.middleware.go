package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors(c *gin.Context) {
	fmt.Println("Test")

	// allowedOrigins := []string{"http://localhost:5500", "http://localhost:5501"} // slice bisa lebih dari 1
	// if slices.Contains(allowedOrigins, c.GetHeader("Origin")) {
	// 	c.Header("Access-Control-Allow-Origin", c.GetHeader("Origin"))
	// }
	//simple cors
	c.Header("Access-Control-Allow-Origin", "http://localhost:5500")
	c.Header("Access-Control-Allow-Headers", "Content-Type, XXX-Headers")
	c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, PATCH") //SELAIN GET SAMA POST HARUS PAKAI ALLOW METHODS

	//preflgith
	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusNoContent) //MENAMBAHKAN RESPONSE
		return
	}
	c.Next()
	c.Writer.Status()
}
