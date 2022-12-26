package main

import (
	"fmt"
	"homePage/backend/database"
	databaseBusiness "homePage/backend/database/database_business"
	viewModel "homePage/backend/handler/view_model"
	modelConvert "homePage/backend/model_convert"
)

func Init() {
	fmt.Println("init")
	fmt.Println("init database")
	database.Init()
	fmt.Println("database init complete")
	fmt.Println("\nadmin acount init")
	users, err := databaseBusiness.ListUser()
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		err := databaseBusiness.CreateUser(modelConvert.ViewModel2DatabaseModelUser(&viewModel.UserView{
			Name:     "admin",
			Password: "admin",
			Role:     "admin",
		}))
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("admin acount init complete")
}
