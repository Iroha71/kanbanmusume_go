package main

import (
	"kanbanmusume_go/controllers"
	"kanbanmusume_go/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	router := make_router()
	router.Run(":3000")
	db.Close()
}

func make_router() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		ctrl := controllers.UserController{}
		user.GET("", ctrl.Index)
	}

	return router
}
