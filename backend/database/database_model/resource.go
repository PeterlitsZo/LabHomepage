package databaseModel

type Resource struct {
	ID      uint   `gorm:"primaryKey;index;column:id"`
	Title   string `gorm:"unique;uniqueIndex;not null;column:title"`
	Content string `gorm:"column:content"`
	Extra   string `gorm:"column:extra"`
}

type ResourceFilter Resource
