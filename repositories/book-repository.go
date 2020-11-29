package repositories

import (
	"github.com/ydhnwb/go_restful_api/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//BookRepository is an interface
type BookRepository interface {
	Insert(video entities.Book)
	Update(video entities.Book)
	Delete(video entities.Book)
	All() []entities.Book
	FindByID(bookID uint64) entities.Book
	CloseDatabaseConnection()
}

type database struct {
	connection *gorm.DB
}

//NewBookRepository creates a new instance of VideoRepository
func NewBookRepository() BookRepository {
	dsn := "root:yudhanewbie@tcp(127.0.0.1:3306)/go_rest_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to database")
	}
	db.AutoMigrate(&entities.Book{}, &entities.Author{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDatabaseConnection() {
	// err := db.connection.Close()
	dbSQL, err := db.connection.DB()
	if err != nil {
		panic("Failed when close a connection from database")
	}
	// defer dbSql.Close()
	dbSQL.Close()
}

func (db *database) Insert(book entities.Book) {
	db.connection.Create(&book)

}

func (db *database) Update(book entities.Book) {
	db.connection.Save(&book)
}

func (db *database) Delete(book entities.Book) {
	db.connection.Delete(&book)
}

func (db *database) All() []entities.Book {
	var books []entities.Book
	println(books)
	db.connection.Set("gorm:auto_preload", true).Find(&books)
	return books
}

func (db *database) FindByID(bookID uint64) entities.Book {
	var book entities.Book
	db.connection.Find(&book)
	return book
}
