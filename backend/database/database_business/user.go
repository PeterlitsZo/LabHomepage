package databaseBusiness

import (
	databaseManager "homePage/backend/database/database_manager"
	databaseModel "homePage/backend/database/database_model"
	myError "homePage/backend/my_error"
)

func GetUserByID(uid uint) (*databaseModel.User, error) {
	if ret, err := databaseManager.UserManage.GetByFilter(&databaseModel.UserFilter{
		ID: uid,
	}); err != nil {
		return nil, err
	} else if ret == nil || len(ret) == 0 {
		return nil, nil
	} else {
		return ret[0], nil
	}
}

func ListUser() ([]*databaseModel.User, error) {
	if ret, err := databaseManager.UserManage.GetByFilter(&databaseModel.UserFilter{}); err != nil {
		return []*databaseModel.User{}, err
	} else {
		return ret, nil
	}
}

func UpdateUserByID(id uint, User *databaseModel.User) error {
	if User == nil {
		return myError.PointerIsNil
	} else {
		User.ID = 0
		if err := databaseManager.UserManage.Update(&databaseModel.UserFilter{
			ID: id,
		}, User); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func GetUserByName(name string) (*databaseModel.User, error) {
	if ret, err := databaseManager.UserManage.GetByFilter(&databaseModel.UserFilter{
		Name: name,
	}); err != nil {
		return nil, err
	} else if ret == nil || len(ret) == 0 {
		return nil, nil
	} else {
		return ret[0], nil
	}
}

func CreateUser(User *databaseModel.User) error {
	if User == nil {
		return myError.PointerIsNil
	} else {
		User.ID = 0
		if err := databaseManager.UserManage.Create(User); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func DeleteUserByID(id uint) error {
	if err := databaseManager.UserManage.DeleteByID(id); err != nil {
		return err
	} else {
		return nil
	}
}
