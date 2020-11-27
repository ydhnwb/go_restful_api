package repositories

import (
	"github.com/ydhnwb/go_restful_api/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//VideoRepository is an interface
type VideoRepository interface {
	Insert(video entities.Video)
	Update(video entities.Video)
	Delete(video entities.Video)
	All() []entities.Video
	CloseDatabaseConnection()
}

type database struct {
	connection *gorm.DB
}

//NewVideoRepository creates a new instance of VideoRepository
func NewVideoRepository() VideoRepository {
	dsn := "root:yudhanewbie@tcp(127.0.0.1:3306)/go_rest_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to database")
	}
	db.AutoMigrate(&entities.Video{}, &entities.Person{})
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

func (db *database) Insert(video entities.Video) {
	println("Insert a video")
	db.connection.Create(&video)

}

func (db *database) Update(video entities.Video) {
	println("Update a video")
	db.connection.Save(&video)
}

func (db *database) Delete(video entities.Video) {
	println("Delete a video")
	db.connection.Delete(&video)
}

func (db *database) All() []entities.Video {
	var videos []entities.Video
	db.connection.Set("gorm:auto_preload", true).Find(&videos)
	return videos
}
