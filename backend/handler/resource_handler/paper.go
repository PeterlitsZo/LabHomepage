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

type PaperHandler struct {
	router string
}

func (h *PaperHandler) Get(r *gin.Engine) error {
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
				"message": handlerError.PaperNotExist.Error(),
			})
			return
		} else if dbPaper, err := databaseBusiness.GetPaperByID(id); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			ret := modelConvert.DatabaseModel2ViewModelPaper(dbPaper)
			g.JSON(http.StatusOK, ret)
			return
		}
	}
	handlers = append(handlers, handler)
	r.GET(h.router+"/:id", handlers...)
	return nil
}

func (h *PaperHandler) List(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		// middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		type Paper struct {
			Paper []*viewModel.PaperView `json:"papers"`
		}
		if paper, err := databaseBusiness.ListPaper(); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			g.JSON(http.StatusOK, Paper{
				Paper: modelConvert.DatabaseModel2ViewModelPaperList(paper),
			})
			return
		}
	}
	handlers = append(handlers, handler)
	r.GET(h.router, handlers...)
	return nil
}

func (h *PaperHandler) Update(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		var paper viewModel.PaperView
		if err := g.ShouldBind(&paper); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if len(paper.Title) == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.PaperTitleEmpty.Error(),
			})
			return
		} else if id, err := util.String2Uint(g.Param("id")); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.PaperNotExist.Error(),
			})
			return
		} else if err = databaseBusiness.UpdatePaperByID(id, modelConvert.ViewModel2DatabaseModelPaper(&paper)); err != nil {
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

func (h *PaperHandler) Create(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		var paper viewModel.PaperView
		if err := g.ShouldBind(&paper); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if len(paper.Title) == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.PaperTitleEmpty.Error(),
			})
			return
		} else if ret, err := databaseBusiness.GetPaperByTitle(paper.Title); ret != nil || err != nil {
			if err != nil {
				g.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			} else if ret != nil {
				g.JSON(http.StatusBadRequest, gin.H{
					"message": handlerError.PaperAlreadyExist.Error(),
				})
				return
			}
		} else if err = databaseBusiness.CreatePaper(modelConvert.ViewModel2DatabaseModelPaper(&paper)); err != nil {
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

func (h *PaperHandler) Delete(r *gin.Engine) error {
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
				"message": handlerError.PaperNotExist.Error(),
			})
			return
		} else if paper, err := databaseBusiness.GetPaperByID(id); err != nil || paper == nil {
			if err != nil {
				g.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			} else if paper == nil {
				g.JSON(http.StatusBadRequest, gin.H{
					"message": handlerError.PaperNotExist.Error(),
				})
				return
			}
		} else if err = databaseBusiness.DeletePaperByID(id); err != nil {
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

func NewPaperHandler(router string) *PaperHandler {
	return &PaperHandler{
		router: router,
	}
}
