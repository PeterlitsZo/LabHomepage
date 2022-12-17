package main

import (
	"fmt"
	"homePage/backend/database"
)

func Init() {
	fmt.Println("init")
	fmt.Println("init database")
	database.Init()
}
