package handler

import (
	"homePage/backend/dao"
	"homePage/backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewLoginHandler() gin.HandlerFunc {
	return LoginHandler
}

func LoginHandler(g *gin.Context) {
	type Login struct {
		Username string `form:"username" json:"username" binding:"required"`
		Password string `form:"password" json:"password" binding:"required"`
	}
	login := Login{}
	if err := g.ShouldBind(&login); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else if login.Password == "" || login.Username == "" {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "args is nil",
		})
		return
	} else if user, err := dao.Manager.UserManager.GetByUsernameAndPassword(login.Username, login.Password); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	} else if jwt, err := util.GenerateJWT(util.User{
		Username: user.Username,
		UserId:   util.Uint2String(user.ID),
	}); err != nil {
		g.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		g.JSON(200, gin.H{
			"jwt": jwt,
			// "user": user,
		})
	}
}
