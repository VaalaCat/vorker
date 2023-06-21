package models

import (
	"vorker/utils/database"
	"vorker/utils/secret"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

const (
	ErrInvalidParams = "invalid params"
)

type User struct {
	gorm.Model
	UserName string `json:"user_name" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
	Status   int    `json:"status"`
	Role     string `json:"role"`
}

func init() {
	db := database.GetDB()
	if err := db.AutoMigrate(&User{}); err != nil {
		logrus.Panic(err)
	}
	database.CloseDB(db)
}

func (u *User) TableName() string {
	return "users"
}

func CreateUser(user *User) error {
	if hashedPass, err := secret.HashPassword(user.Password); err != nil {
		return err
	} else {
		user.Password = hashedPass
	}

	return database.GetDB().Create(user).Error
}

func AdminGetUserNumber() (int64, error) {
	var count int64
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Model(&User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetUserByUserID(userID uint) (*User, error) {
	var user User
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Where(&User{
		Model: gorm.Model{ID: userID},
	}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByUserName(userName string) (*User, error) {
	var user User
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Where(&User{
		UserName: userName,
	}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Where(&User{
		Email: email,
	}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUser(userID uint, user *User) error {
	if hashedPass, err := secret.HashPassword(user.Password); err != nil {
		return err
	} else {
		user.Password = hashedPass
	}
	db := database.GetDB()
	defer database.CloseDB(db)
	return db.Model(&User{
		Model: gorm.Model{ID: userID},
	}).Updates(user).Error
}

func DeleteUser(userID uint) error {
	db := database.GetDB()
	defer database.CloseDB(db)
	return db.Delete(&User{
		Model: gorm.Model{ID: userID},
	}).Error
}

func ListUsers(page, pageSize int) ([]*User, error) {
	var users []*User
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func CountUsers() (int64, error) {
	var count int64
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Model(&User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func GetUserByUserNameAndPassword(userName, password string) (*User, error) {
	var user User
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Where(&User{
		UserName: userName,
		Password: password,
	}).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CheckUserPassword(userNameOrEmail, password string) (bool, error) {
	var user User
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Where(&User{
		UserName: userNameOrEmail,
	}).Or(&User{
		Email: userNameOrEmail}).First(&user).Error; err != nil {
		return false, err
	}
	return secret.CheckPasswordHash(password, user.Password), nil
}

func CheckUserNameAndEmail(userName, email string) error {
	var user User
	db := database.GetDB()
	defer database.CloseDB(db)
	if err := db.Where(&User{
		UserName: userName,
	}).Or(&User{
		Email: email,
	}).First(&user).Error; err != nil {
		return err
	}
	return nil
}
