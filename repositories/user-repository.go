package repositories

import (
	"log"

	"github.com/ydhnwb/go_restful_api/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

//UserRepository is
type UserRepository interface {
	InsertUser(user entities.User) entities.User
	UpdateUser(user entities.User)
	DeleteUser(user entities.User)
	VerifyCredential(email string, password string) bool
	IsDuplicateEmail(user entities.User) bool
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

func (db *userConnection) InsertUser(user entities.User) entities.User {
	user.Password = hashAndSalt([]byte(user.Password))
	db.connection.Create(&user)
	var createdUser entities.User
	db.connection.Where("email = ?", user.Email).Take(&createdUser)
	return createdUser
}

func (db *userConnection) UpdateUser(user entities.User) {
	user.Password = hashAndSalt([]byte(user.Password))
	println("hashed when update")
	println(user.Password)
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

func (db *userConnection) VerifyCredential(email string, password string) bool {
	var user entities.User
	err := db.connection.Where("email = ?", email).Take(&user).Error
	if err == nil {
		comparedPassword := comparePasswords(user.Password, []byte(password))
		if user.Email == email && comparedPassword {
			return true
		}
		return false
	}
	return false
}

func (db *userConnection) IsDuplicateEmail(user entities.User) bool {
	err := db.connection.Where("email = ?", user.Email).Take(&user).Error
	if err == nil {
		return false
	}
	return true
	// errors.Is(err, gorm.ErrRecordNotFound)
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func (db *userConnection) FindByEmail(email string) entities.User {
	var user entities.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	println("hashed password " + hashedPwd)
	println(plainPwd)
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
