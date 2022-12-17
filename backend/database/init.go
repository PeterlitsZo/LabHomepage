package database

import (
	databaseConnect "homePage/backend/database/database_connect"
	databaseManager "homePage/backend/database/database_manager"
)

func Init() {
	dbConn := databaseConnect.NewDbConn()
	databaseManager.Init(&dbConn)
}
