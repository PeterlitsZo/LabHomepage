package databaseBusiness

import (
	databaseManager "homePage/backend/database/database_manager"
	databaseModel "homePage/backend/database/database_model"
	myError "homePage/backend/my_error"
)

func GetNewsByID(id uint) (*databaseModel.News, error) {
	if ret, err := databaseManager.NewsManage.GetByFilter(&databaseModel.NewsFilter{
		ID: id,
	}); err != nil {
		return nil, err
	} else if ret == nil || len(ret) == 0 {
		return nil, nil
	} else {
		return ret[0], nil
	}
}

func GetNewsByTitle(title string) (*databaseModel.News, error) {
	if ret, err := databaseManager.NewsManage.GetByFilter(&databaseModel.NewsFilter{
		Title: title,
	}); err != nil {
		return nil, err
	} else if ret == nil || len(ret) == 0 {
		return nil, nil
	} else {
		return ret[0], nil
	}
}

func ListNews() ([]*databaseModel.News, error) {
	if ret, err := databaseManager.NewsManage.GetByFilter(&databaseModel.NewsFilter{}); err != nil {
		return []*databaseModel.News{}, err
	} else {
		return ret, nil
	}
}

func UpdateNewsByID(id uint, news *databaseModel.News) error {
	if news == nil {
		return myError.PointerIsNil
	} else {
		news.ID = 0
		if err := databaseManager.NewsManage.Update(&databaseModel.NewsFilter{
			ID: id,
		}, news); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func CreateNews(news *databaseModel.News) error {
	if news == nil {
		return myError.PointerIsNil
	} else {
		news.ID = 0
		if err := databaseManager.NewsManage.Create(news); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func DeleteNewsByID(id uint) error {
	if err := databaseManager.NewsManage.DeleteByID(id); err != nil {
		return err
	} else {
		return nil
	}
}
