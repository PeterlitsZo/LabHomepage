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

func NewPersonHandlerImpl() *PersonHandleImpl {
	return &PersonHandleImpl{
		Route: "api/v1/Person",
	}
}

type PersonHandleImpl struct {
	Route string
}

func (h *PersonHandleImpl) Get(r *gin.Engine) {
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
		} else if ret, err := dao.Manager.PersonManager.Get(id); err != nil {
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

func (h *PersonHandleImpl) List(r *gin.Engine) {
	handle := func(g *gin.Context) {
		type Person struct {
			Person []*model.Person `json:"Person"`
		}
		if persons, err := dao.Manager.PersonManager.List(0, 1000); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			g.JSON(http.StatusOK, Person{
				Person: persons,
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

func (h *PersonHandleImpl) Create(r *gin.Engine) {
	handle := func(g *gin.Context) {
		new := model.Person{}
		if err := g.ShouldBind(&new); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if err := dao.Manager.PersonManager.Create(&new); err != nil {
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

func (h *PersonHandleImpl) Delete(r *gin.Engine) {
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
		} else if err := dao.Manager.PersonManager.Delete(id); err != nil {
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

func (h *PersonHandleImpl) Update(r *gin.Engine) {
	handle := func(g *gin.Context) {
		var Person model.Person
		if err := g.ShouldBind(&Person); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if err := dao.Manager.PersonManager.Update(&Person); err != nil {
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
