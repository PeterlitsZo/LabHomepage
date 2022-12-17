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

type ResourceHandler struct {
	router string
}

func (h *ResourceHandler) Get(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		if id, err := util.String2Uint(g.Param("id")); err != nil {
			g.JSON(http.StatusBadRequest, err)
			return
		} else if id == 0 {
			g.JSON(http.StatusBadRequest, handlerError.ResourceNotExist)
			return
		} else if dbResource, err := databaseBusiness.GetResourceByID(id); err != nil {
			g.JSON(http.StatusInternalServerError, err)
			return
		} else {
			ret := modelConvert.DatabaseModel2ViewModelResource(dbResource)
			g.JSON(http.StatusOK, ret)
			return
		}
	}
	handlers = append(handlers, handler)
	r.GET(h.router+"/:id", handlers...)
	return nil
}

func (h *ResourceHandler) List(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		type Resource struct {
			Resource []*viewModel.ResourceView `json:"resource"`
		}
		if resource, err := databaseBusiness.ListResource(); err != nil {
			g.JSON(http.StatusInternalServerError, err)
			return
		} else {
			g.JSON(http.StatusOK, Resource{
				Resource: modelConvert.DatabaseModel2ViewModelResourceList(resource),
			})
			return
		}
	}
	handlers = append(handlers, handler)
	r.GET(h.router, handlers...)
	return nil
}

func (h *ResourceHandler) Update(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		var resource viewModel.ResourceView
		if err := g.ShouldBind(&resource); err != nil {
			g.JSON(http.StatusBadRequest, err)
			return
		} else if len(resource.Title) == 0 {
			g.JSON(http.StatusBadRequest, handlerError.ResourceTitleEmpty)
			return
		} else if id, err := util.String2Uint(g.Param("id")); err != nil {
			g.JSON(http.StatusBadRequest, handlerError.ResourceNotExist)
			return
		} else if err = databaseBusiness.UpdateResourceByID(id, modelConvert.ViewModel2DatabaseModelResource(&resource)); err != nil {
			g.JSON(http.StatusInternalServerError, err)
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

func (h *ResourceHandler) Create(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		var resource viewModel.ResourceView
		if err := g.ShouldBind(&resource); err != nil {
			g.JSON(http.StatusBadRequest, err)
			return
		} else if len(resource.Title) == 0 {
			g.JSON(http.StatusBadRequest, handlerError.ResourceTitleEmpty)
			return
		} else if ret, err := databaseBusiness.GetResourceByTitle(resource.Title); ret == nil || err != nil {
			if err != nil {
				g.JSON(http.StatusInternalServerError, err)
				return
			} else if ret == nil {
				g.JSON(http.StatusBadRequest, handlerError.ResourceNotExist)
				return
			}
		} else if err = databaseBusiness.CreateResource(modelConvert.ViewModel2DatabaseModelResource(&resource)); err != nil {
			g.JSON(http.StatusInternalServerError, err)
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

func (h *ResourceHandler) Delete(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
		middleware.NewPermissionsMiddleware(func(user *databaseModel.User) bool {
			return user.IsAdmin()
		}),
	}
	handler := func(g *gin.Context) {
		if id, err := util.String2Uint(g.Param("id")); err != nil {
			g.JSON(http.StatusBadRequest, err)
			return
		} else if id == 0 {
			g.JSON(http.StatusBadRequest, handlerError.ResourceNotExist)
			return
		} else if resource, err := databaseBusiness.GetResourceByID(id); err != nil || resource == nil {
			if err != nil {
				g.JSON(http.StatusInternalServerError, err)
				return
			} else if resource == nil {
				g.JSON(http.StatusBadRequest, handlerError.ResourceNotExist)
				return
			}
		} else if err = databaseBusiness.DeleteResourceByID(id); err != nil {
			g.JSON(http.StatusInternalServerError, err)
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

func NewResourceHandler(router string) *ResourceHandler {
	return &ResourceHandler{
		router: router,
	}
}
