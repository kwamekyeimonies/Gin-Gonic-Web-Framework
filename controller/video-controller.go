package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kwamekyeimonies/Gin-Gonic-Web-Framework/model"
	"github.com/kwamekyeimonies/Gin-Gonic-Web-Framework/service"
)


type VideoController interface{
	FindAll() []model.Video
	Save(ctx *gin.Context) model.Video
}

type controller struct{
	service service.VideoService
}

func New(service service.VideoService) VideoController{
	return &controller{
		service: service,
	}
}

func (ctx *controller) FindAll() []model.Video{
	return service.New().FindAll()
}

func (c *controller) Save(ctx *gin.Context) model.Video{
	var video model.Video
	ctx.BindJSON(&video)
	c.service.Save(video)

	return video
}