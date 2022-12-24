package resourceHandler

import (
	databaseBusiness "homePage/backend/database/database_business"
	handlerError "homePage/backend/handler/handler_error"
	viewModel "homePage/backend/handler/view_model"
	"homePage/backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	router string
}

func (h *TokenHandler) Get(r *gin.Engine) error {
	return nil
}

func (h *TokenHandler) List(r *gin.Engine) error {
	return nil
}

func (h *TokenHandler) Update(r *gin.Engine) error {
	return nil
}

func (h *TokenHandler) Create(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		// middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		var user viewModel.UserView
		if err := g.ShouldBind(&user); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if len(user.Name) == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.UserNameEmpty.Error(),
			})
			return
		} else if ret, err := databaseBusiness.GetUserByName(user.Name); ret == nil || err != nil {
			if err != nil {
				g.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			} else if ret == nil {
				g.JSON(http.StatusBadRequest, gin.H{
					"message": handlerError.UserNotExist.Error(),
				})
				return
			} else if !util.CheckPasswordHash(user.Password, ret.Password) {
				g.JSON(http.StatusBadRequest, gin.H{
					"message": handlerError.PasswordNotCorrect.Error(),
				})
				return
			}
		} else if jwt, err := util.GenerateJWT(util.User{
			Username: ret.Name,
			UserId:   ret.ID,
		}); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			g.JSON(http.StatusOK, "Bearer "+jwt)
			return
		}
	}
	handlers = append(handlers, handler)
	r.POST(h.router, handlers...)
	return nil
}

func (h *TokenHandler) Delete(r *gin.Engine) error {
	return nil
}

func NewTokenHandler(router string) *TokenHandler {
	return &TokenHandler{
		router: router,
	}
}
