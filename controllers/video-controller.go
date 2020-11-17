package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_restful_api/entities"
	"github.com/ydhnwb/go_restful_api/services"
)

//VideoController is an interface of logic what can Video do
type VideoController interface {
	FindAll(ctx *gin.Context)
	Save(context *gin.Context)
}

type controller struct {
	service services.VideoService
}

//NewVideoController function in creating a new VideoController instance
func NewVideoController(service services.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll(ctx *gin.Context) {
	var videos []entities.Video = c.service.All()
	println(videos)
	response := entities.BuildResponse(true, "OK", videos)
	ctx.JSON(http.StatusOK, response)
}

func (c *controller) Save(ctx *gin.Context) {
	// var video entities.Video
	// ctx.BindJSON(&video)
	// c.service.Insert(video)
	println("Example save()...")
}
