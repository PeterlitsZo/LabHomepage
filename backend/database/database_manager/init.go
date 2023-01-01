package databaseManager

import databaseConnect "homePage/backend/database/database_connect"

var (
	NewsManage     *NewsManager
	UserManage     *UserManager
	PaperManage    *PaperManager
	ResourceManage *ResourceManager
	PersonManage   *PersonManager
)

func Init(dbConn *databaseConnect.DbConn) {
	var err error
	if NewsManage, err = NewNewsManager(dbConn); err != nil {
		panic(err)
	} else if UserManage, err = NewUserManager(dbConn); err != nil {
		panic(err)
	} else if PaperManage, err = NewPaperManager(dbConn); err != nil {
		panic(err)
	} else if ResourceManage, err = NewResourceManager(dbConn); err != nil {
		panic(err)
	} else if PersonManage, err = NewPersonManager(dbConn); err != nil {
		panic(err)
	} else {
		return
	}
}
