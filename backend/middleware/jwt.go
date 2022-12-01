package middleware

import (
	"errors"
	"net/http"

	"homePage/backend/util"

	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware() gin.HandlerFunc {
	authMiddleware := func(g *gin.Context) {
		if token, ok := g.Request.Header["Authorization"]; !ok {
			g.AbortWithError(http.StatusUnauthorized, errors.New("Authorization is empty"))
			g.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization is empty",
			})
			return
		} else if len(token) == 0 {
			g.AbortWithError(http.StatusUnauthorized, errors.New("Authorization is empty"))
			g.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization is empty",
			})
			return
		} else if jwt := util.SplitJwt(token[0]); jwt == "" {
			g.AbortWithError(http.StatusUnauthorized, errors.New("jwt is invalid"))
			g.JSON(http.StatusUnauthorized, gin.H{
				"message": "jwt is invalid",
			})
			return
		} else if isValid, err := util.IsJwtValid(jwt); err != nil {
			g.AbortWithError(http.StatusInternalServerError, err)
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else if !isValid {
			g.AbortWithError(http.StatusUnauthorized, errors.New("jwt is invalid"))
			g.JSON(http.StatusUnauthorized, gin.H{
				"message": "jwt is invalid",
			})
			return
		} else {
			return
		}
	}
	return authMiddleware
}
