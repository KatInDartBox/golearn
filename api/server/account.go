package server

import (
	"api-test/db/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Holder  string `json:"holder" binding:"required"`
	Balance int64  `json:"balance" binding:"required,oneof=0 15"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := sqlc.CreateAccountParams{
		Holder:  req.Holder,
		Balance: req.Balance,
	}
	acc, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, acc)
}

type getAccountRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	acc, err := server.store.GetAccount(ctx, req.Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, acc)
}

type getAccountsRequest struct {
	LastId    int64 `form:"lastId"`
	PageLimit int64 `form:"pageLimit" binding:"required,min=5,max=50"`
}

func (server *Server) getAccountList(ctx *gin.Context) {
	req := getAccountsRequest{
		LastId: 0,
	}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}
	arg := sqlc.GetAccountListParams{
		LastID:    req.LastId,
		PageLimit: req.PageLimit,
	}
	acc, err := server.store.GetAccountList(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, acc)
}
