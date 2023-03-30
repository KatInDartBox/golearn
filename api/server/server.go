package server

import (
	"api-test/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	store  *sqlc.Queries
}

func NewServer(store *sqlc.Queries) *Server {
	router := gin.Default()
	server := &Server{store: store}
	server.router = router

	routeAccount := router.Group("/account")
	routeAccount.POST("/", server.createAccount)
	routeAccount.GET("/:id", server.getAccount)
	routeAccount.GET("/", server.getAccountList)

	routeCustomeValidation := router.Group("/validation")
	routeCustomeValidation.POST("/", server.postBooking)
	routeCustomeValidation.GET("/defaultForm", server.getDefaultForm)
	routeCustomeValidation.POST("/defaultJson", server.getDefaultJson)

	router.GET("/", server.getHome)
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {

	return gin.H{
		"error": err.Error(),
	}
}
