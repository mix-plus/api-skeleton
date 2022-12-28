package controllers

import "github.com/gin-gonic/gin"

type HelloController struct{}

func (h *HelloController) SayHello(ctx *gin.Context) {
	ctx.String(200, "Hello MixPlus")
}
