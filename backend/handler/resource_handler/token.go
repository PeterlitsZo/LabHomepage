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
			g.JSON(http.StatusBadRequest, err)
			return
		} else if len(user.Name) == 0 {
			g.JSON(http.StatusBadRequest, handlerError.UserNameEmpty)
			return
		} else if ret, err := databaseBusiness.GetUserByName(user.Name); ret == nil || err != nil {
			if err != nil {
				g.JSON(http.StatusInternalServerError, err)
				return
			} else if ret == nil {
				g.JSON(http.StatusBadRequest, handlerError.UserNotExist)
				return
			}
		} else if jwt, err := util.GenerateJWT(util.User{
			Username: ret.Name,
			UserId:   ret.ID,
		}); err != nil {
			g.JSON(http.StatusInternalServerError, err)
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
