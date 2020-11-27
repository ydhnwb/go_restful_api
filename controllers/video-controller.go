package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/go_restful_api/entities"
	"github.com/ydhnwb/go_restful_api/services"
)

//VideoController is an interface of logic what can Video do
type VideoController interface {
	All(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
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

func (c *controller) All(context *gin.Context) {
	var videos []entities.Video = c.service.All()
	response := entities.BuildResponse(true, "OK", videos)
	context.JSON(http.StatusOK, response)
}

func (c *controller) Insert(context *gin.Context) {
	var video entities.Video
	err := context.ShouldBindJSON(&video)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to precess your data", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
	} else {
		c.service.Insert(video)
		response := entities.BuildResponse(true, "OK", video)
		context.JSON(http.StatusCreated, response)
	}

}

func (c *controller) Update(context *gin.Context) {
	var video entities.Video
	err := context.ShouldBindJSON(&video)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to precess your data", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to find your id", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	video.ID = id
	c.service.Update(video)
	response := entities.BuildResponse(true, "OK", video)
	context.JSON(http.StatusOK, response)
}

func (c *controller) Delete(context *gin.Context) {
	var video entities.Video
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := entities.BuildErrorResponse("Failed to find your id", err.Error(), nil)
		context.JSON(http.StatusBadRequest, response)
		return
	}
	video.ID = id
	c.service.Delete(video)
	response := entities.BuildResponse(true, "Deleted", nil)
	context.JSON(http.StatusOK, response)
}
