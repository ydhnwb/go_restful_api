package repositories

import (
	"github.com/ydhnwb/go_restful_api/dto"
	"github.com/ydhnwb/go_restful_api/entities"
	"gorm.io/gorm"
)

//UserRepository is
type UserRepository interface {
	InsertUser(user dto.UserCreateDTO) entities.User
	UpdateUser(user dto.UserUpdateDTO)
	DeleteUser(user entities.User)
	VerifyCredential(email string, password string) (tx *gorm.DB)
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entities.User
	ProfileUser(userID string) entities.User
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

func (db *userConnection) InsertUser(user dto.UserCreateDTO) entities.User {
	db.connection.Save(&user)
	var createdUser entities.User
	db.connection.Where("email = ?", user.Email).Take(&createdUser)
	return createdUser
}

func (db *userConnection) UpdateUser(user dto.UserUpdateDTO) {
	db.connection.Save(&user)
}

func (db *userConnection) DeleteUser(user entities.User) {
	db.connection.Delete(&user)

}
func (db *userConnection) ProfileUser(userID string) entities.User {
	var user entities.User
	db.connection.Find(&user, userID)
	return user
}

func (db *userConnection) VerifyCredential(email string, password string) (tx *gorm.DB) {
	var user entities.User
	return db.connection.Where("email = ?").Take(&user)
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entities.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) entities.User {
	var user entities.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}
