package config

import (
	"github.com/ydhnwb/go_restful_api/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SetupDatabaseConnection is create a connection to database when server boot up
func SetupDatabaseConnection() *gorm.DB {
	dsn := "root:yudhanewbie@tcp(127.0.0.1:3306)/go_rest_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}
	db.AutoMigrate(&entities.Book{}, &entities.User{})
	return db
}

//CloseDatabaseConnection will close connection to database
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed when close a connection from database")
	}
	// defer dbSql.Close()
	dbSQL.Close()
}
