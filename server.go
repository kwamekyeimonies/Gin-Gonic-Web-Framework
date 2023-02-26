package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/Gin-Gonic-Web-Framework/controller"
	"github.com/kwamekyeimonies/Gin-Gonic-Web-Framework/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {

	server := gin.Default()
	port := "6666"

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "Server Running",
		})
	})

	server.GET("/videos", func(ctx *gin.Context) {
		// videoController.FindAll()
		ctx.JSON(http.StatusOK, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.Save(ctx))
	})

	server.Run(":" + port)
}
