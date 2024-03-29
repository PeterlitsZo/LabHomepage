package middleware

import (
	databaseBusiness "homePage/backend/database/database_business"
	databaseModel "homePage/backend/database/database_model"
	handlerError "homePage/backend/handler/handler_error"
	"homePage/backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewPermissionsMiddleware(CheckFunc func(user *databaseModel.User) bool) gin.HandlerFunc {
	handler := func(g *gin.Context) {
		if token, ok := g.Request.Header["Authorization"]; !ok {
			g.AbortWithError(http.StatusUnauthorized, handlerError.AuthorizationIsEmpty)
			g.JSON(http.StatusUnauthorized, gin.H{
				"message": handlerError.AuthorizationIsEmpty.Error(),
			})
			return
		} else if len(token) == 0 {
			g.AbortWithError(http.StatusUnauthorized, handlerError.AuthorizationIsEmpty)
			g.JSON(http.StatusUnauthorized, gin.H{
				"message": handlerError.AuthorizationIsEmpty.Error(),
			})
			return
		} else if jwt := util.GetBearerToken(token[0]); jwt == "" {
			g.AbortWithError(http.StatusUnauthorized, handlerError.TokenIsInvalid)
			g.JSON(http.StatusUnauthorized, gin.H{
				"message": handlerError.AuthorizationIsEmpty.Error(),
			})
			return
		} else if user, err := util.ParseJWT(jwt); err != nil {
			g.AbortWithError(http.StatusInternalServerError, err)
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else if usr, err := databaseBusiness.GetUserByID(user.UserId); err != nil {
			g.AbortWithError(http.StatusInternalServerError, err)
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else if usr == nil {
			g.AbortWithError(http.StatusUnauthorized, handlerError.UserNotExist)
			g.JSON(http.StatusUnauthorized, gin.H{
				"message": handlerError.UserNotExist.Error(),
			})
			return
		} else if !CheckFunc(usr) {
			g.AbortWithError(http.StatusUnauthorized, handlerError.PermissionDenied)
			g.JSON(http.StatusUnauthorized, gin.H{
				"message": handlerError.PermissionDenied.Error(),
			})
			return
		} else {
			return
		}
	}
	return handler
}
