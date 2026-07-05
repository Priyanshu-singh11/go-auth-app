package httpserver

import (
	"my-gin-app/internal/app"
	"my-gin-app/internal/middleware"
	"my-gin-app/internal/user"
	"net/http"

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
	r.POST("/login", userHandler.Login)

	api := r.Group("/api")

	api.Use(middleware.AuthSecret(a.Config.JwtSecret))

	api.GET("/files", func(c *gin.Context) {
		userId, _ := middleware.GetUserId(c)
		c.JSON(http.StatusOK, gin.H{
			"ok":     true,
			"userId": userId,
		})
	})

	admin := api.Group("/admin")
	admin.Use(middleware.Requireadmin())
	admin.GET("/restricted", func(c *gin.Context) {
		role, _ := middleware.GetRole(c)
		c.JSON(http.StatusOK, gin.H{
			"ok":   true,
			"role": role,
		})
	})

	return r
}
