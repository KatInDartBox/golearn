package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type bookingRequest struct {
	CheckIn  time.Time `json:"checkIn" binding:"required,bookabledate" `
	CheckOut time.Time `json:"checkOut" binding:"required,gtfield=CheckIn,bookabledate" `
}

/*
bookabledate is not a default validation field
we init this function first before validator/gin could process it
*/

func (server *Server) postBooking(ctx *gin.Context) {
	var req bookingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, req)

}

type sFormDefault struct {
	Age  int    `form:"age"`
	Name string `form:"name" binding:"required"`
}

func (server *Server) getDefaultForm(ctx *gin.Context) {
	// var req sFormDefault
	req := sFormDefault{
		Age: 18,
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, req)
}

type sDefaultJson struct {
	Age  int    `json:"age"`
	Name string `json:"name" binding:"required"`
}

func (server *Server) getDefaultJson(ctx *gin.Context) {
	req := sDefaultJson{
		Age: 24,
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, req)
}
