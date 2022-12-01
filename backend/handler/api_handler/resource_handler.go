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

func NewResourceHandlerImpl() *ResourceHandleImpl {
	return &ResourceHandleImpl{
		Route: "api/v1/Resource",
	}
}

type ResourceHandleImpl struct {
	Route string
}

func (h *ResourceHandleImpl) Get(r *gin.Engine) {
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
		} else if ret, err := dao.Manager.ResourceManager.Get(id); err != nil {
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

func (h *ResourceHandleImpl) List(r *gin.Engine) {
	handle := func(g *gin.Context) {
		type Resource struct {
			Resource []*model.Resource `json:"Resource"`
		}
		if resources, err := dao.Manager.ResourceManager.List(0, 1000); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			g.JSON(http.StatusOK, Resource{
				Resource: resources,
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

func (h *ResourceHandleImpl) Create(r *gin.Engine) {
	handle := func(g *gin.Context) {
		new := model.Resource{}
		if err := g.ShouldBind(&new); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if err := dao.Manager.ResourceManager.Create(&new); err != nil {
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
		handle,
	}
	r.POST(h.Route, handlers...)
	return
}

func (h *ResourceHandleImpl) Delete(r *gin.Engine) {
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
		} else if err := dao.Manager.ResourceManager.Delete(id); err != nil {
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

func (h *ResourceHandleImpl) Update(r *gin.Engine) {
	handle := func(g *gin.Context) {
		var Resource model.Resource
		if err := g.ShouldBind(&Resource); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if err := dao.Manager.ResourceManager.Update(&Resource); err != nil {
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
