package userrepo

import (
	usersModel "anna/models/users"
	"crypto/md5"
	"encoding/hex"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"name"`
	Password string `gorm:"password"`
	Email    string `gorm:"email"`
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (repo *UserRepo) RegisterUser(name, email, password string) error {
	userModal := usersModel.User{
		Name:     name,
		Password: password,
		Email:    email,
	}

	if err := repo.db.Create(&userModal).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) LoginUser(name, email, password string) (bool, error) {
	var loginUser usersModel.User
	res := repo.db.Where("email = ?", email).First(&loginUser)
	if res.Error != nil || res.RowsAffected != 1 {
		return false, res.Error
	}
	passBytes := []byte(password)
	hashPassword := md5.Sum(passBytes)
	hashStringPass := hex.EncodeToString(hashPassword[:])

	if loginUser.Password != hashStringPass {
		return false, nil
	}
	return true, nil

}
