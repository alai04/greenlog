package api

import (
	"net/http"

	"github.com/alai04/greenlog/httputil"
	"github.com/alai04/greenlog/log"
	"github.com/gin-gonic/gin"
)

// @Description ping example
// @Accept json
// @Produce  json
// @Success 200 {string} string	"pong"
// @Failure 400 {string} string "ok"
// @Failure 404 {string} string "ok"
// @Failure 500 {string} string "ok"
// @Router /ping [get]
func Pong(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

// @Description post log message
// @Accept  json
// @Produce  json
// @Param   log     body    log.LogRec     true        "Log Record"
// @Success 200 {string} string	"ok"
// @Router /log [post]
func LogPost(ctx *gin.Context) {
	var rec log.LogRec
	if err := ctx.ShouldBindJSON(&rec); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err := log.Log2File(&rec)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	httputil.NewSucc(ctx, nil)
}
