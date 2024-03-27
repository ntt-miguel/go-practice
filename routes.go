package routes

import "github.com/gin-gonic/gin"

func Router(server *gin.Engine) {
	server.GET("/fundings", getFundings)
	server.GET("/fundings/:id", getFunding)
	server.POST("/fundings", createFunding)
	server.DELETE("/fundings/:id", deleteFunding)
	server.POST("/signup", signUp)
	server.POST("/login", login)
}
