package resourceHandler

import "github.com/gin-gonic/gin"

type Handler interface {
	Get(r *gin.Engine) error
	List(r *gin.Engine) error
	Create(r *gin.Engine) error
	Update(r *gin.Engine) error
	Delete(r *gin.Engine) error
}

var Handlers = []interface{}{
	NewNewsHandler("/api/v1/news"),
	NewPaperHandler("/api/v1/papers"),
	NewPersonHandler("/api/v1/people"),
	NewResourceHandler("/api/v1/resources"),
	NewTokenHandler("/api/v1/login"),
	NewUserHandler("/api/v1/users"),
}
