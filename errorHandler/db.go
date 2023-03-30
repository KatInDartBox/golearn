package errorHandler

import (
	"github.com/gin-gonic/gin"
)

func BadRequest(err error, c *gin.Context){
	DBerror(err,c,400)
}
func Unauthorized(err error, c *gin.Context){
	DBerror(err,c,401)
}
func Forbidden(err error, c *gin.Context){
	DBerror(err,c,403)
}
func NotFound(err error, c *gin.Context){
	DBerror(err,c,404)
}
func NotAcceptable(err error, c *gin.Context){
	DBerror(err,c,406)
}

func DBerror(err error, c *gin.Context,code int) {
	if err != nil {
		c.Status(code)
	}
}
