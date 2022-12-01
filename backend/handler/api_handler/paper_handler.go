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

func NewPaperHandlerImpl() *PaperHandleImpl {
	return &PaperHandleImpl{
		Route: "api/v1/Paper",
	}
}

type PaperHandleImpl struct {
	Route string
}

func (h *PaperHandleImpl) Get(r *gin.Engine) {
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
		} else if ret, err := dao.Manager.PaperManager.Get(id); err != nil {
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

func (h *PaperHandleImpl) List(r *gin.Engine) {
	handle := func(g *gin.Context) {
		type Paper struct {
			Paper []*model.Paper `json:"Paper"`
		}
		if papers, err := dao.Manager.PaperManager.List(0, 1000); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			g.JSON(http.StatusOK, Paper{
				Paper: papers,
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

func (h *PaperHandleImpl) Create(r *gin.Engine) {
	handle := func(g *gin.Context) {
		new := model.Paper{}
		if err := g.ShouldBind(&new); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if err := dao.Manager.PaperManager.Create(&new); err != nil {
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

func (h *PaperHandleImpl) Delete(r *gin.Engine) {
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
		} else if err = dao.Manager.PaperManager.Delete(Id); err != nil {
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

func (h *PaperHandleImpl) Update(r *gin.Engine) {
	handle := func(g *gin.Context) {
		var Paper model.Paper
		if err := g.ShouldBind(&Paper); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if err = dao.Manager.PaperManager.Update(&Paper); err != nil {
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
