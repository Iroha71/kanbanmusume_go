package services

import (
	"kanbanmusume_go/db"
	"kanbanmusume_go/models"

	"github.com/gin-gonic/gin"
)

type UserService struct{}

type User models.User

func (_ UserService) GetAll() ([]User, error) {
	db := db.GetDB()
	var user []User
	if err := db.Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (_ UserService) Create(c *gin.Context) (User, error) {
	db := db.GetDB()
	var user User
	if err := c.BindJSON(&user); err != nil {
		return user, err
	}
	if err := db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
