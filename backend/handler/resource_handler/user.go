package resourceHandler

import (
	databaseBusiness "homePage/backend/database/database_business"
	databaseModel "homePage/backend/database/database_model"
	handlerError "homePage/backend/handler/handler_error"
	"homePage/backend/handler/middleware"
	viewModel "homePage/backend/handler/view_model"
	modelConvert "homePage/backend/model_convert"
	"homePage/backend/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	router string
}

func (h *UserHandler) Get(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		if id, err := util.String2Uint(g.Param("id")); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if id == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.UserNotExist.Error(),
			})
			return
		} else if dbUser, err := databaseBusiness.GetUserByID(id); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			ret := modelConvert.DatabaseModel2ViewModelUser(dbUser)
			g.JSON(http.StatusOK, ret)
			return
		}
	}
	handlers = append(handlers, handler)
	r.GET(h.router+"/:id", handlers...)
	return nil
}

func (h *UserHandler) List(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		type User struct {
			User []*viewModel.UserView `json:"user"`
		}
		if user, err := databaseBusiness.ListUser(); err != nil {
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			g.JSON(http.StatusOK, User{
				User: modelConvert.DatabaseModel2ViewModelUserList(user),
			})
			return
		}
	}
	handlers = append(handlers, handler)
	r.GET(h.router, handlers...)
	return nil
}

func (h *UserHandler) Update(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		var User viewModel.UserView
		if err := g.ShouldBind(&User); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else if len(User.Name) == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.UserNameEmpty.Error(),
			})
			return
		} else if id, err := util.String2Uint(g.Param("id")); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.UserNotExist.Error(),
			})
			return
		} else if err = databaseBusiness.UpdateUserByID(id, modelConvert.ViewModel2DatabaseModelUser(&User)); err != nil {
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

func (h *UserHandler) Create(r *gin.Engine) error {
	handlers := []gin.HandlerFunc{
		// middleware.NewAuthMiddleware(),
	}
	handler := func(g *gin.Context) {
		var User viewModel.UserView
		if err := g.ShouldBind(&User); err != nil {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			log.Println("User 格式不正确")
			return
		} else if len(User.Name) == 0 {
			g.JSON(http.StatusBadRequest, gin.H{
				"message": handlerError.UserNameEmpty.Error(),
			})
			log.Println("username 为空")
			return
		} else if ret, err := databaseBusiness.GetUserByName(User.Name); ret != nil || err != nil {
			if err != nil {
				log.Fatalln("数据库查询用户出错")
				g.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			} else if ret != nil {
				log.Println("用户名已存在")
				g.JSON(http.StatusBadRequest, gin.H{
					"message": handlerError.UserAlreadyExist.Error(),
				})
				return
			}
		} else if err = databaseBusiness.CreateUser(modelConvert.ViewModel2DatabaseModelUser(&User)); err != nil {
			log.Fatalln("数据库创建用户失败")
			g.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			log.Println("创建用户成功")
			g.JSON(http.StatusOK, nil)
			return
		}
	}
	handlers = append(handlers, handler)
	r.POST(h.router, handlers...)
	return nil
}

func (h *UserHandler) Delete(r *gin.Engine) error {
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
				"message": handlerError.UserNotExist.Error(),
			})
			return
		} else if User, err := databaseBusiness.GetUserByID(id); err != nil || User == nil {
			if err != nil {
				g.JSON(http.StatusInternalServerError, gin.H{
					"message": err.Error(),
				})
				return
			} else if User == nil {
				g.JSON(http.StatusBadRequest, gin.H{
					"message": handlerError.UserNotExist.Error(),
				})
				return
			}
		} else if err = databaseBusiness.DeleteUserByID(id); err != nil {
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

func NewUserHandler(router string) *UserHandler {
	return &UserHandler{
		router: router,
	}
}
