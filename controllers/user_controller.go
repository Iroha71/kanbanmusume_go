package controllers

import (
	"kanbanmusume_go/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (pc UserController) Index(c *gin.Context) {
	var user services.UserService
	p, err := user.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, p)
	}
}

func (pc UserController) Create(c *gin.Context) {
	// var user services.UserService
	// p, err := user.Create(c)
	// if err != nil {
	// 	c.AbortWithStatus(400)
	// } else {

	// }
}
