package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/docker-ui/config"
	"github.com/sirupsen/logrus"
	"time"
)

func InitialAPI() error {
	logrus.Info("---------- initial internal api start ----------")
	router := SetupRouter()
	return router.Run(config.Conf.ServerHost)
}

func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			//return origin == "https://github.com"
			return true
		},
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"message": "OK"}) })
	router.GET("/login", login)
	containerGroup := router.Group("/containers")
	{
		containerGroup.GET("", list)
		containerGroup.POST("/stop", stop)
		containerGroup.POST("/restart", restart)
	}
	imageGroup := router.Group("/images")
	{
		imageGroup.GET("", list)
		imageGroup.POST("/pull", stop)
	}
	stackGroup := router.Group("/stack")
	{
		stackGroup.GET("/deploy", list)
	}
	return router
}

func login(context *gin.Context) {

}

func list(context *gin.Context) {

}

func stop(context *gin.Context) {

}

func restart(context *gin.Context) {

}
