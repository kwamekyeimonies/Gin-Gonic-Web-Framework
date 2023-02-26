package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kwamekyeimonies/Gin-Gonic-Web-Framework/model"
	"github.com/kwamekyeimonies/Gin-Gonic-Web-Framework/service"
)

type VideoController interface {
	FindAll() []model.Video
	Save(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()

	// validate.RegisterValidation("is-cool",validators.ValidateCoo lTitle)

	return &controller{
		service: service,
	}
}

func (ctx *controller) FindAll() []model.Video {
	return service.New().FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var video model.Video
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	c.service.Save(video)

	return nil
}
