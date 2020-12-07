package services

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/ydhnwb/go_restful_api/dto"
	"github.com/ydhnwb/go_restful_api/entities"
	"github.com/ydhnwb/go_restful_api/repositories"
)

//BookService is an interface that contains a contract what can the service do
type BookService interface {
	Insert(book dto.BookCreateDTO) entities.Book
	Update(book entities.Book) entities.Book
	Delete(book entities.Book)
	All() []entities.Book
	FindByID(bookID uint64) entities.Book
}

type bookService struct {
	bookRepository repositories.BookRepository
}

//NewBookService method is instancing a BookService
func NewBookService(bookRep repositories.BookRepository) BookService {
	return &bookService{
		bookRepository: bookRep,
	}
}

func (service *bookService) Insert(book dto.BookCreateDTO) entities.Book {
	b := entities.Book{}
	err := smapping.FillStruct(&b, smapping.MapFields(&book))
	if err != nil {
		log.Fatalf("failed map: %v", err)
	}
	res := service.bookRepository.InsertBook(b)
	return res
}

func (service *bookService) Update(book entities.Book) entities.Book {
	service.bookRepository.UpdateBook(book)
	return book
}

func (service *bookService) Delete(book entities.Book) {
	service.bookRepository.DeleteBook(book)
}

func (service *bookService) All() []entities.Book {
	return service.bookRepository.AllBook()
}

func (service *bookService) FindByID(bookID uint64) entities.Book {
	return service.bookRepository.FindBookByID(bookID)
}
