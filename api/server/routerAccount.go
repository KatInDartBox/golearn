package server

import (
	md "learn/api/middlewares"
)

func (s *Server) setAccountRoute() {
	router := s.router
	routeAccount := router.Group("/account")
	routeAccount.Use(md.AuthMiddleware(""))

	routeAccount.POST("/", s.createAccount)
	routeAccount.GET("/:id", md.Authorize, s.getAccount)

	routeAccount.GET("/", s.getAccountList)
}
