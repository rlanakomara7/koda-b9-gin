package router // dibuat 1 package , exported tidak apa pakai huruf kecil
import (
	"github.com/gin-gonic/gin"
	"github.com/rlanakomara7/koda-b9-gin.git/internal/handler"
	"github.com/rlanakomara7/koda-b9-gin.git/internal/service"
)

func initAuthRouter(r *gin.Engine) {
	usersGroup := r.Group("/auth")

	authService := service.NewAuthService()
	authHandler := handler.NewAuthHandler(authService)

	usersGroup.POST("/register", authHandler.Register)
	usersGroup.POST("/register", authHandler.Login)

}
