package model

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Extra   string `json:"extra"`
}

type Paper struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Extra   string `json:"extra"`
}

type Person struct {
	gorm.Model
	Name    string `json:"name"`
	Content string `json:"content"`
	Extra   string `json:"extra"`
}

type Resource struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `json:"content"`
	Extra   string `json:"extra"`
}

const (
	Admin  string = "admin"
	Editor string = "editor"
	Viewer string = "viewer"
)

type User struct {
	gorm.Model
	Role     string `json:"role"`
	Username string `json:"username"`
	Password string `json:"password"`
	Extra    string `json:"extra"`
}
