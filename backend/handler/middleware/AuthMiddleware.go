package middleware

import (
	databaseBusiness "homePage/backend/database/database_business"
	handlerError "homePage/backend/handler/handler_error"
	"homePage/backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware() gin.HandlerFunc {
	authMiddleware := func(g *gin.Context) {
		if token, ok := g.Request.Header["Authorization"]; !ok {
			g.AbortWithError(http.StatusUnauthorized, handlerError.AuthorizationIsEmpty)
			g.JSON(http.StatusUnauthorized, handlerError.AuthorizationIsEmpty)
			return
		} else if len(token) == 0 {
			g.AbortWithError(http.StatusUnauthorized, handlerError.AuthorizationIsEmpty)
			g.JSON(http.StatusUnauthorized, handlerError.AuthorizationIsEmpty)
			return
		} else if jwt := util.GetBearerToken(token[0]); jwt == "" {
			g.AbortWithError(http.StatusUnauthorized, handlerError.TokenIsInvalid)
			g.JSON(http.StatusUnauthorized, handlerError.TokenIsInvalid)
			return
		} else if isValid, err := util.IsJwtValid(jwt); err != nil {
			g.AbortWithError(http.StatusInternalServerError, err)
			g.JSON(http.StatusInternalServerError, err)
			return
		} else if !isValid {
			g.AbortWithError(http.StatusUnauthorized, handlerError.JwtIsInvalid)
			g.JSON(http.StatusUnauthorized, handlerError.JwtIsInvalid)
			return
		} else if usr, err := util.ParseJWT(jwt); err != nil {
			g.AbortWithError(http.StatusInternalServerError, err)
			g.JSON(http.StatusInternalServerError, err)
			return
		} else if usr.UserId == 0 {
			g.AbortWithError(http.StatusUnauthorized, handlerError.UserIdIsInvalid)
			g.JSON(http.StatusUnauthorized, handlerError.UserIdIsInvalid)
			return
		} else if user, err := databaseBusiness.GetUserByID(usr.UserId); err != nil {
			g.AbortWithError(http.StatusInternalServerError, err)
			g.JSON(http.StatusInternalServerError, err)
			return
		} else if user.Name != usr.Username {
			g.AbortWithError(http.StatusUnauthorized, handlerError.UsernameIsInvalid)
			g.JSON(http.StatusUnauthorized, handlerError.UsernameIsInvalid)
			return
		} else {
			return
		}

	}
	return authMiddleware
}
