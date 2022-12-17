package databaseModel

type Person struct {
	ID      uint   `gorm:"primaryKey;index;column:id"`
	Name    string `gorm:"unique;uniqueIndex;column:name"`
	Content string `gorm:"column:content"`
	Extra   string `gorm:"column:extra"`
}

type PersonFilter Person
