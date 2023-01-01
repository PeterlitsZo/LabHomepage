package databaseConnect

import (
	"fmt"
	databaseModel "homePage/backend/database/database_model"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConn struct {
	*gorm.DB
}

func (d *DbConn) AutoMigrate() error {
	if err := d.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&databaseModel.User{}); err != nil {
		return err
	} else if err := d.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&databaseModel.News{}); err != nil {
		return err
	} else if err := d.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&databaseModel.Paper{}); err != nil {
		return err
	} else if err := d.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&databaseModel.Person{}); err != nil {
		return err
	} else if err := d.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&databaseModel.Resource{}); err != nil {
		return err
	} else {
		return nil
	}
}

func NewDbConn() DbConn {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		USER, PASS, HOST, DBNAME)
	fmt.Println(URL)
	var db *gorm.DB
	var err error
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(
			mysql.Open(URL),
			&gorm.Config{
				DisableForeignKeyConstraintWhenMigrating: true,
			},
		)
		if err != nil {
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
	if err != nil {
		panic("Failed to connect to database! " + USER + " " + PASS + " " + DBNAME + " " + HOST)
	}
	fmt.Println("Database connection established")
	ret := DbConn{
		DB: db,
	}
	if err := ret.AutoMigrate(); err != nil {
		panic(err)
	}
	return ret
}
