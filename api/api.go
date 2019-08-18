package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/convenient-utils/response"
	"github.com/sillyhatxu/docker-ui/config"
	"github.com/sillyhatxu/docker-ui/service"
	"github.com/sirupsen/logrus"
	"net/http"
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
		containerGroup.GET("", containerList)
		containerGroup.GET("/stats", stats)
		containerGroup.POST("/stop", stop)
		containerGroup.POST("/restart", restart)
	}
	serviceGroup := router.Group("/services")
	{
		serviceGroup.GET("", serviceList)
		serviceGroup.GET("/stats", stats)
		serviceGroup.POST("/stop", stop)
		serviceGroup.POST("/restart", restart)
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

func serviceList(context *gin.Context) {
	array, err := service.ServiceList()
	if err != nil {
		context.JSON(http.StatusOK, response.ServerError(nil, err.Error(), nil))
		return
	}
	context.JSON(http.StatusOK, response.ServerSuccess(array, nil))
}

func containerList(context *gin.Context) {

}

func list(context *gin.Context) {

}

func stats(context *gin.Context) {
	statsArray, err := service.DockerStats()
	if err != nil {
		context.JSON(http.StatusOK, response.ServerError(nil, err.Error(), nil))
		return
	}
	context.JSON(http.StatusOK, response.ServerSuccess(statsArray, nil))
}

func stop(context *gin.Context) {

}

func restart(context *gin.Context) {

}
