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

type PersonHandler struct {
	router string
}

func (h *PersonHandler) Get(r *gin.Engine) error {
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
				"message": handlerError.PersonNotExist.Error(),
			})
			return
		} else if dbPerson, err := databaseBusiness.GetPersonByID(id); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			ret := modelConvert.DatabaseModel2ViewModelPerson(dbPerson)
			g.JSON(http.StatusOK, ret)
			return
		}
	}
	handlers = append(handlers, handler)
	r.GET(h.router+"/:id", handlers...)
	return nil
}

func (h *PersonHandler) List(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		// middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		type Person struct {
			Person []*viewModel.PersonView `json:"person"`
		}
		if person, err := databaseBusiness.ListPerson(); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			g.JSON(http.StatusOK, Person{
				Person: modelConvert.DatabaseModel2ViewModelPersonList(person),
			})
			return
		}
	}
	handlers = append(handlers, handler)
	r.GET(h.router, handlers...)
	return nil
}

func (h *PersonHandler) Update(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		var Person viewModel.PersonView
		if err := g.ShouldBind(&Person); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if len(Person.Name) == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.PersonNameEmpty.Error(),
			})
			return
		} else if id, err := util.String2Uint(g.Param("id")); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.PersonNotExist.Error(),
			})
			return
		} else if err = databaseBusiness.UpdatePersonByID(id, modelConvert.ViewModel2DatabaseModelPerson(&Person)); err != nil {
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

func (h *PersonHandler) Create(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		var Person viewModel.PersonView
		if err := g.ShouldBind(&Person); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if len(Person.Name) == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.PersonNameEmpty.Error(),
			})
			return
		} else if ret, err := databaseBusiness.GetPersonByName(Person.Name); ret != nil || err != nil {
			if err != nil {
				g.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			} else if ret != nil {
				g.JSON(http.StatusBadRequest, gin.H{
					"message": handlerError.PersonAlreadyExist.Error(),
				})
				return
			}
		} else if err = databaseBusiness.CreatePerson(modelConvert.ViewModel2DatabaseModelPerson(&Person)); err != nil {
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

func (h *PersonHandler) Delete(r *gin.Engine) error {
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
				"message": handlerError.PersonNotExist.Error(),
			})
			return
		} else if Person, err := databaseBusiness.GetPersonByID(id); err != nil || Person == nil {
			if err != nil {
				g.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			} else if Person == nil {
				g.JSON(http.StatusBadRequest, gin.H{
					"message": handlerError.PersonNotExist.Error(),
				})
				return
			}
		} else if err = databaseBusiness.DeletePersonByID(id); err != nil {
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

func NewPersonHandler(router string) *PersonHandler {
	return &PersonHandler{
		router: router,
	}
}
