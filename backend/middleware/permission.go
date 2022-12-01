package middleware

import (
	"errors"
	"homePage/backend/dao"
	"homePage/backend/model"
	"homePage/backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewPermissionsMiddleware(CheckFunc func(user *model.User) bool) gin.HandlerFunc {
	handler := func(g *gin.Context) {
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
		} else if user, err := util.ParseJWT(jwt); err != nil {
			g.AbortWithError(http.StatusInternalServerError, err)
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else if uid, err := util.String2Uint(user.UserId); err != nil {
			g.AbortWithError(http.StatusInternalServerError, err)
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		} else if usr, err := dao.Manager.UserManager.Get(uid); err != nil {
			g.AbortWithError(http.StatusInternalServerError, err)
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else if !CheckFunc(usr) {
			g.AbortWithError(http.StatusUnauthorized, errors.New("权限不够"))
			g.JSON(http.StatusUnauthorized, gin.H{
				"message": "权限不够",
			})
			return
		} else {
			return
		}
	}
	return handler
}
