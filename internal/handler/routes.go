package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/mix-plus/api-skeleton/internal/handler/controllers"
	"github.com/mix-plus/core/validator"
)

func Load(router *gin.Engine) {
	// register validate
	binding.Validator = new(validator.DefaultValidator)

	hello := controllers.HelloController{}
	router.GET("/", hello.SayHello)
}
