package services

import "github.com/ydhnwb/go_restful_api/entities"

//VideoService is an interface that contains a contract what can the service do
type VideoService interface {
	Insert(entities.Video) entities.Video
	All() []entities.Video
}

type videoService struct {
	videos []entities.Video
}

//NewVideoService method is instancing a VideoService
func NewVideoService() VideoService {
	return &videoService{
		videos: []entities.Video{},
	}
}

func (service *videoService) Insert(video entities.Video) entities.Video {
	service.videos = append(service.videos, video)
	return video
}

func (service *videoService) All() []entities.Video {
	// var videos []entities.Video
	// return videos
	return service.videos
}
