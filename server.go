package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/Gin-Gonic-Web-Framework/controller"
	"github.com/kwamekyeimonies/Gin-Gonic-Web-Framework/middlewares"
	"github.com/kwamekyeimonies/Gin-Gonic-Web-Framework/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	// server := gin.Default()
	setupLogOutput()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth()) //Using the BasicAuth Authorization
	// server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAPIKeyAuth())//Using the API-Key Authorization
	// server.Use(gin.Logger())
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
		// ctx.JSON(http.StatusOK, videoController.Save(ctx))
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error": err.Error(),
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"Message": "Video Validation Successfull",
			})
		}
	})

	server.Run(":" + port)
}
