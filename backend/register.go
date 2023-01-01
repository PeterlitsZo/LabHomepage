package main

import (
	resourceHandler "homePage/backend/handler/resource_handler"

	"github.com/gin-gonic/gin"
)

func register(r *gin.Engine) {
	for _, val := range resourceHandler.Handlers {
		handler := val.(resourceHandler.Handler)
		handler.Get(r)
		handler.List(r)
		handler.Create(r)
		handler.Delete(r)
		handler.Update(r)
	}
}
