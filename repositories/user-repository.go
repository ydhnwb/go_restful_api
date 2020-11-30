package repositories

import (
	"github.com/ydhnwb/go_restful_api/entities"
	"gorm.io/gorm"
)

//UserRepository is
type UserRepository interface {
	InsertUser(user entities.User)
	UpdateUser(user entities.User)
	DeleteUser(user entities.User)
	ProfileUser(token string) entities.User
}

type userConnection struct {
	connection *gorm.DB
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user entities.User) {
	db.connection.Create(&user)

}

func (db *userConnection) UpdateUser(user entities.User) {
	db.connection.Save(&user)
}

func (db *userConnection) DeleteUser(user entities.User) {
	db.connection.Delete(&user)

}
func (db *userConnection) ProfileUser(token string) entities.User {
	var user entities.User
	db.connection.Find(&user, token)
	return user
}
