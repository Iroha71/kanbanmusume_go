package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := make_router()

	router.Run(":3000")
}

func make_router() *gin.Engine {
	router := gin.Default()

	return router
}
