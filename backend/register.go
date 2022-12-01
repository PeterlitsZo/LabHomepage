package main

import (
	"homePage/backend/handler"

	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {
	handler.InitHandler()
	for _, handler := range handler.ApiHandlerImplMap {
		handler.Get(r)
		handler.List(r)
		handler.Create(r)
		handler.Delete(r)
		handler.Update(r)
	}
	r.POST("api/v1/token", handler.NewLoginHandler())
	r.POST("api/v1/register", handler.NewRegisterHandler())
}
