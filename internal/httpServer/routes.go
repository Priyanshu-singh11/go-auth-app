package httpserver

import (
	"my-gin-app/internal/app"
	"my-gin-app/internal/user"

	"github.com/gin-gonic/gin"
)

func NewRouter(a *app.App) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/health", health)

	userRepo := user.Newrepo(a.DB)
	userSvc := user.NewService(userRepo, a.Config.JwtSecret)
	userHandler := user.NewHandler(userSvc)

	r.POST("/register", userHandler.Register)
	return r
}
