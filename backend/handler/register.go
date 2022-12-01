package handler

import (
	"homePage/backend/dao"
	"homePage/backend/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRegisterHandler() gin.HandlerFunc {
	return RegisterHandler
}

func RegisterHandler(g *gin.Context) {
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
	} else if exist, err := dao.Manager.UserManager.Exists(login.Username, login.Password); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	} else if exist {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "用户名或密码已存在",
		})
		return
	} else if err := dao.Manager.UserManager.Create(&model.User{
		Username: login.Username,
		Password: login.Password,
	}); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		g.JSON(200, gin.H{})
	}
}
