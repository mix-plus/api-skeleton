package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mix-plus/core/validator"
	"net/http"
)

func Load(router *gin.Engine) {
	// register validate
	binding.Validator = new(validator.DefaultValidator)

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})
}
