package middlewares

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"learn/api/initer"
	"learn/utils/uToken"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenType string) gin.HandlerFunc {
	if tokenType == "" {
		tokenType = authorizationTypeBearer
	}
	// fmt.Println("token type out: ", tokenType)
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		// fmt.Println("token type inside: ", tokenType)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		authorizationType := strings.TrimSpace(strings.ToLower(fields[0]))
		if authorizationType != tokenType {

			err := fmt.Errorf("unsupported authorization type %s", authorizationType)

			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := uToken.Decrypt(initer.PasetoKey, accessToken)
		if err != nil {
			fmt.Println("f0", fields[0])
			fmt.Println("f1", fields[1])
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

type getAccountRequest struct {
	Id int64 `uri:"id" binding:"required,min=1"`
}

func Authorize(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(uToken.Payload)
	if authPayload.Id == 0 || authPayload.Id != req.Id {
		err := fmt.Errorf("not authorize req id: %v", req.Id)

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, errResponse(err))
		return
	}
	ctx.Next()
}

func errResponse(err error) gin.H {

	return gin.H{
		"error": err.Error(),
	}
}
