package services

import (
	"github.com/ydhnwb/go_restful_api/entities"
	"github.com/ydhnwb/go_restful_api/repositories"
)

//BookService is an interface that contains a contract what can the service do
type BookService interface {
	Insert(book entities.Book) entities.Book
	Update(book entities.Book) entities.Book
	Delete(book entities.Book)
	All() []entities.Book
	FindByID(bookID uint64) entities.Book
}

type bookService struct {
	bookRepository repositories.BookRepository
}

//NewBookService method is instancing a VideoService
func NewBookService(bookRep repositories.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRep,
	}
}

func (service *bookService) Insert(video entities.Book) entities.Book {
	service.bookRepository.Insert(video)
	return video
}

func (service *bookService) Update(video entities.Book) entities.Book {
	service.bookRepository.Update(video)
	return video
}

func (service *bookService) Delete(video entities.Book) {
	service.bookRepository.Delete(video)
}

func (service *bookService) All() []entities.Book {
	return service.bookRepository.All()
}

func (service *bookService) FindByID(bookID uint64) entities.Book {
	return service.bookRepository.FindByID(bookID)
}
