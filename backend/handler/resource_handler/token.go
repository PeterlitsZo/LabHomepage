package resourceHandler

import (
	databaseBusiness "homePage/backend/database/database_business"
	handlerError "homePage/backend/handler/handler_error"
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
		var loginRequest struct {
			Name     string `json:"username"`
			Password string `json:"password"`
		}

		if err := g.ShouldBind(&loginRequest); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if len(loginRequest.Name) == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.UserNameEmpty.Error(),
			})
			return
		}

		user, err := databaseBusiness.GetUserByName(loginRequest.Name)
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		if user == nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.UserNotExist.Error(),
			})
			return
		}

		// Check username-password pair.
		if !util.CheckPasswordHash(loginRequest.Password, user.Password) {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.PasswordNotCorrect.Error(),
			})
			return
		}

		// Generate token to return.
		jwt, err := util.GenerateJWT(util.User{
			Username: user.Name,
			UserId:   user.ID,
		})
		if err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		g.JSON(http.StatusOK, gin.H{
			"token": jwt,
		})
		return
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
