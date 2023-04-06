package server

import (
	"learn/api/initer"
	"learn/utils/uToken"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (server *Server) getToken(ctx *gin.Context) {
	payload := uToken.Payload{
		Id:   90101293123,
		Name: "some jakk ji",
	}

	enc, err := uToken.Encrypt(initer.PasetoKey, payload, 2*time.Hour)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, enc)
}

type reqEncryption struct {
	Enc string `json:"enc" binding:"required"`
}

func (server *Server) getTokenDecryption(ctx *gin.Context) {
	var req reqEncryption
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	dec, err := uToken.Decrypt(initer.PasetoKey, req.Enc)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, dec)
}
