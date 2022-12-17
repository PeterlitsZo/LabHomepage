package databaseBusiness

import (
	databaseManager "homePage/backend/database/database_manager"
	databaseModel "homePage/backend/database/database_model"
	myError "homePage/backend/my_error"
)

func GetResourceByID(id uint) (*databaseModel.Resource, error) {
	if ret, err := databaseManager.ResourceManage.GetByFilter(&databaseModel.ResourceFilter{
		ID: id,
	}); err != nil {
		return nil, err
	} else if ret == nil || len(ret) == 0 {
		return nil, nil
	} else {
		return ret[0], nil
	}
}

func ListResource() ([]*databaseModel.Resource, error) {
	if ret, err := databaseManager.ResourceManage.GetByFilter(&databaseModel.ResourceFilter{}); err != nil {
		return []*databaseModel.Resource{}, err
	} else {
		return ret, nil
	}
}

func UpdateResourceByID(id uint, Resource *databaseModel.Resource) error {
	if Resource == nil {
		return myError.PointerIsNil
	} else {
		Resource.ID = 0
		if err := databaseManager.ResourceManage.Update(&databaseModel.ResourceFilter{
			ID: id,
		}, Resource); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func GetResourceByTitle(title string) (*databaseModel.Resource, error) {
	if ret, err := databaseManager.ResourceManage.GetByFilter(&databaseModel.ResourceFilter{
		Title: title,
	}); err != nil {
		return nil, err
	} else if ret == nil || len(ret) == 0 {
		return nil, nil
	} else {
		return ret[0], nil
	}
}

func CreateResource(Resource *databaseModel.Resource) error {
	if Resource == nil {
		return myError.PointerIsNil
	} else {
		Resource.ID = 0
		if err := databaseManager.ResourceManage.Create(Resource); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func DeleteResourceByID(id uint) error {
	if err := databaseManager.ResourceManage.DeleteByID(id); err != nil {
		return err
	} else {
		return nil
	}
}
