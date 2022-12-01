package handler

import (
	apihandler "homePage/backend/handler/api_handler"

	"github.com/gin-gonic/gin"
)

type HandlerImpl interface {
	Get(r *gin.Engine)
	List(r *gin.Engine)
	Delete(r *gin.Engine)
	Update(r *gin.Engine)
	Create(r *gin.Engine)
}

var ApiHandlerImplMap = []HandlerImpl{}

func InitHandler() {
	ApiHandlerImplMap = append(ApiHandlerImplMap, apihandler.NewNewsHandlerImpl())
	ApiHandlerImplMap = append(ApiHandlerImplMap, apihandler.NewPaperHandlerImpl())
	ApiHandlerImplMap = append(ApiHandlerImplMap, apihandler.NewPersonHandlerImpl())
	ApiHandlerImplMap = append(ApiHandlerImplMap, apihandler.NewResourceHandlerImpl())
}
