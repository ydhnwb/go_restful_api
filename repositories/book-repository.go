package repositories

import (
	"github.com/ydhnwb/go_restful_api/entities"
	"gorm.io/gorm"
)

//BookRepository is an interface
type BookRepository interface {
	InsertBook(book entities.Book) entities.Book
	UpdateBook(book entities.Book) entities.Book
	DeleteBook(book entities.Book)
	AllBook() []entities.Book
	FindBookByID(bookID uint64) entities.Book
}

type bookConnection struct {
	connection *gorm.DB
}

//NewBookRepository creates a new instance of BookRepository
func NewBookRepository(dbConnection *gorm.DB) BookRepository {
	return &bookConnection{
		connection: dbConnection,
	}
}

func (db *bookConnection) InsertBook(book entities.Book) entities.Book {
	db.connection.Save(&book)
	db.connection.Preload("User").Find(&book)
	return book
}

func (db *bookConnection) UpdateBook(book entities.Book) entities.Book {
	db.connection.Save(&book)
	db.connection.Preload("User").Find(&book)
	return book
}

func (db *bookConnection) DeleteBook(book entities.Book) {
	db.connection.Delete(&book)
}

func (db *bookConnection) AllBook() []entities.Book {
	var books []entities.Book
	db.connection.Preload("User").Find(&books)
	return books
}

func (db *bookConnection) FindBookByID(bookID uint64) entities.Book {
	var book entities.Book
	db.connection.Preload("User").Find(&book, bookID)
	return book
}
