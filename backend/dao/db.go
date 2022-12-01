package dao

import (
	"fmt"
	"homePage/backend/model"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Database struct
type Database struct {
	DB *gorm.DB
}

func (d *Database) AutoMigrate() error {
	if err := d.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.User{}); err != nil {
		return err
	}
	if err := d.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.News{}); err != nil {
		return err
	}
	if err := d.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Paper{}); err != nil {
		return err
	}
	if err := d.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Person{}); err != nil {
		return err
	}
	if err := d.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&model.Resource{}); err != nil {
		return err
	}
	return nil
}

func (d *Database) GetDB() *gorm.DB {
	return d.DB
}

var db Database
var Manager *DbManager

type DbManager struct {
	db              *Database
	NewsManager     *NewsManager
	PaperManager    *PaperManager
	PersonManager   *PersonManager
	ResourceManager *ResourceManager
	UserManager     *UserManager
}

func init() {
	db = NewDatabase()
	Manager = &DbManager{
		db:              &db,
		NewsManager:     NewNewsManager(),
		PaperManager:    NewPaperManager(),
		PersonManager:   NewPersonManager(),
		ResourceManager: NewResourceManager(),
		UserManager:     NewUserManager(),
	}
}

// NewDatabase : intializes and returns mysql db
func NewDatabase() Database {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS,
		HOST, DBNAME)
	fmt.Println(URL)
	var db *gorm.DB
	var err error
	for i := 0; i < 5; i++ {
		db, err = gorm.Open(
			mysql.Open(URL),
			&gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true,
			},
		)
		if err != nil {
			time.Sleep(1 * time.Second)
		} else {
			break
		}
	}
	if err != nil {
		panic("Failed to connect to database! " + USER + " " + PASS + " " + DBNAME + " " + HOST)
	}
	fmt.Println("Database connection established")
	ret := Database{
		DB: db,
	}
	if err := ret.AutoMigrate(); err != nil {
		panic(err)
	}
	return ret
}
