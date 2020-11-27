package services

import (
	"github.com/ydhnwb/go_restful_api/entities"
	"github.com/ydhnwb/go_restful_api/repositories"
)

//VideoService is an interface that contains a contract what can the service do
type VideoService interface {
	Insert(video entities.Video) entities.Video
	Update(video entities.Video) entities.Video
	Delete(video entities.Video)
	All() []entities.Video
}

type videoService struct {
	videoRepository repositories.VideoRepository
}

//NewVideoService method is instancing a VideoService
func NewVideoService(videoRep repositories.VideoRepository) VideoService {
	return &videoService{
		videoRepository: videoRep,
	}
}

func (service *videoService) Insert(video entities.Video) entities.Video {
	service.videoRepository.Insert(video)
	return video
}

func (service *videoService) Update(video entities.Video) entities.Video {
	service.videoRepository.Update(video)
	return video
}

func (service *videoService) Delete(video entities.Video) {
	service.videoRepository.Delete(video)
}

func (service *videoService) All() []entities.Video {
	return service.videoRepository.All()
}
