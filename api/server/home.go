package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (server *Server) getHome(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "hello world",
	})
}
