package apihandler

import (
	"homePage/backend/dao"
	"homePage/backend/middleware"
	"homePage/backend/model"
	"homePage/backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewNewsHandlerImpl() *NewsHandleImpl {
	return &NewsHandleImpl{
		Route: "api/v1/news",
	}
}

type NewsHandleImpl struct {
	Route string
}

func (h *NewsHandleImpl) Get(r *gin.Engine) {
	handle := func(g *gin.Context) {
		if id, err := util.String2Uint(g.Param("id")); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if id == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": gorm.ErrRecordNotFound.Error(),
			})
			return
		} else if ret, err := dao.Manager.NewsManager.Get(id); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			g.JSON(http.StatusOK, ret)
			return
		}
	}
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
		handle,
	}
	r.GET(h.Route+"/:id", handlers...)
	return
}

func (h *NewsHandleImpl) List(r *gin.Engine) {
	handle := func(g *gin.Context) {
		type News struct {
			News []*model.News `json:"news"`
		}
		if news, err := dao.Manager.NewsManager.List(0, 1000); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			g.JSON(http.StatusOK, News{
				News: news,
			})
			return
		}
	}
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
		handle,
	}
	r.GET(h.Route, handlers...)
	return
}

func (h *NewsHandleImpl) Create(r *gin.Engine) {
	handle := func(g *gin.Context) {
		new := model.News{}
		if err := g.ShouldBind(&new); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if err := dao.Manager.NewsManager.Create(&new); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			g.JSON(http.StatusOK, gin.H{
				"message": new,
			})
			return
		}
	}

	// handler链，顺序不可乱
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
		handle,
	}
	r.POST(h.Route, handlers...)
	return
}

func (h *NewsHandleImpl) Delete(r *gin.Engine) {
	handle := func(g *gin.Context) {
		if Id, err := util.String2Uint(g.Param("id")); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if Id == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": gorm.ErrRecordNotFound.Error(),
			})
			return
		} else if err := dao.Manager.NewsManager.Delete(Id); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			g.JSON(http.StatusOK, gin.H{})
			return
		}
	}
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
		middleware.NewPermissionsMiddleware(
			func(user *model.User) bool {
				return user.Role == model.Admin
			},
		),
		handle,
	}
	r.DELETE(h.Route+"/:id", handlers...)
	return
}

func (h *NewsHandleImpl) Update(r *gin.Engine) {
	handle := func(g *gin.Context) {
		var news model.News
		if err := g.ShouldBind(&news); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if err := dao.Manager.NewsManager.Update(&news); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			g.JSON(http.StatusOK, gin.H{})
			return
		}
	}
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
		middleware.NewPermissionsMiddleware(
			func(user *model.User) bool {
				return user.Role == model.Admin
			},
		),
		handle,
	}
	r.PUT(h.Route+"/:id", handlers...)
	return
}
