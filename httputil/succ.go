package httputil

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewSucc example
func NewSucc(ctx *gin.Context, data interface{}) {
	su := HTTPSucc{
		Code:    http.StatusOK,
		Message: "success",
		Data:    data,
	}
	ctx.JSON(http.StatusOK, su)
}

// HTTPSucc example
type HTTPSucc struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}
