package resourceHandler

import (
	databaseBusiness "homePage/backend/database/database_business"
	databaseModel "homePage/backend/database/database_model"
	handlerError "homePage/backend/handler/handler_error"
	"homePage/backend/handler/middleware"
	viewModel "homePage/backend/handler/view_model"
	modelConvert "homePage/backend/model_convert"
	"homePage/backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NewsHandler struct {
	router string
}

func (h *NewsHandler) Get(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		// middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		if id, err := util.String2Uint(g.Param("id")); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if id == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.NewsNotExist.Error(),
			})
			return
		} else if dbNews, err := databaseBusiness.GetNewsByID(id); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			ret := modelConvert.DatabaseModel2ViewModelNews(dbNews)
			g.JSON(http.StatusOK, ret)
			return
		}
	}
	handlers = append(handlers, handler)
	r.GET(h.router+"/:id", handlers...)
	return nil
}

func (h *NewsHandler) List(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		// middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		type News struct {
			News []*viewModel.NewsView `json:"news"`
		}
		if news, err := databaseBusiness.ListNews(); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			g.JSON(http.StatusOK, News{
				News: modelConvert.DatabaseModel2ViewModelNewsList(news),
			})
			return
		}
	}
	handlers = append(handlers, handler)
	r.GET(h.router, handlers...)
	return nil
}

func (h *NewsHandler) Update(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		var news viewModel.NewsView
		if err := g.ShouldBind(&news); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if len(news.Title) == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.NewsTitleEmpty.Error(),
			})
			return
		} else if id, err := util.String2Uint(g.Param("id")); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.NewsNotExist.Error(),
			})
			return
		} else if err = databaseBusiness.UpdateNewsByID(id, modelConvert.ViewModel2DatabaseModelNews(&news)); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			g.JSON(http.StatusOK, nil)
			return
		}
	}
	handlers = append(handlers, handler)
	r.PUT(h.router+"/:id", handlers...)
	return nil
}

func (h *NewsHandler) Create(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		var news viewModel.NewsView
		if err := g.ShouldBind(&news); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if len(news.Title) == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.NewsTitleEmpty.Error(),
			})
			return
		} else if ret, err := databaseBusiness.GetNewsByTitle(news.Title); ret != nil || err != nil {
			if err != nil {
				g.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			} else if ret != nil {
				g.JSON(http.StatusBadRequest, gin.H{
					"message": handlerError.NewsAlreadyExist.Error(),
				})
				return
			}
		} else if err = databaseBusiness.CreateNews(modelConvert.ViewModel2DatabaseModelNews(&news)); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			g.JSON(http.StatusOK, nil)
			return
		}
	}
	handlers = append(handlers, handler)
	r.POST(h.router, handlers...)
	return nil
}

func (h *NewsHandler) Delete(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
		middleware.NewPermissionsMiddleware(func(user *databaseModel.User) bool {
			return user.IsAdmin()
		}),
	}
	handler := func(g *gin.Context) {
		if id, err := util.String2Uint(g.Param("id")); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if id == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.NewsNotExist.Error(),
			})
			return
		} else if news, err := databaseBusiness.GetNewsByID(id); err != nil || news == nil {
			if err != nil {
				g.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			} else if news == nil {
				g.JSON(http.StatusBadRequest, gin.H{
					"message": handlerError.NewsNotExist.Error(),
				})
				return
			}
		} else if err = databaseBusiness.DeleteNewsByID(id); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			g.JSON(http.StatusOK, nil)
			return
		}
	}
	handlers = append(handlers, handler)
	r.DELETE(h.router+"/:id", handlers...)
	return nil
}

func NewNewsHandler(router string) *NewsHandler {
	return &NewsHandler{
		router: router,
	}
}
