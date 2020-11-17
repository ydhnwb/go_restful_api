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
	response := entities.BuildResponse(true, "OK", videos)
	ctx.JSON(http.StatusOK, response)
}

func (c *controller) Save(ctx *gin.Context) {
	var video entities.Video
	err := ctx.BindJSON(&video)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to precess your data", err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		c.service.Insert(video)
		response := entities.BuildResponse(true, "OK", video)
		ctx.JSON(http.StatusCreated, response)
	}

}
