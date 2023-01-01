package databaseBusiness

import (
	databaseManager "homePage/backend/database/database_manager"
	databaseModel "homePage/backend/database/database_model"
	myError "homePage/backend/my_error"
)

func GetPaperByID(id uint) (*databaseModel.Paper, error) {
	if ret, err := databaseManager.PaperManage.GetByFilter(&databaseModel.PaperFilter{
		ID: id,
	}); err != nil {
		return nil, err
	} else if ret == nil || len(ret) == 0 {
		return nil, nil
	} else {
		return ret[0], nil
	}
}

func ListPaper() ([]*databaseModel.Paper, error) {
	if ret, err := databaseManager.PaperManage.GetByFilter(&databaseModel.PaperFilter{}); err != nil {
		return []*databaseModel.Paper{}, err
	} else {
		return ret, nil
	}
}

func UpdatePaperByID(id uint, paper *databaseModel.Paper) error {
	if paper == nil {
		return myError.PointerIsNil
	} else {
		paper.ID = 0
		if err := databaseManager.PaperManage.Update(&databaseModel.PaperFilter{
			ID: id,
		}, paper); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func GetPaperByTitle(title string) (*databaseModel.Paper, error) {
	if ret, err := databaseManager.PaperManage.GetByFilter(&databaseModel.PaperFilter{
		Title: title,
	}); err != nil {
		return nil, err
	} else if ret == nil || len(ret) == 0 {
		return nil, nil
	} else {
		return ret[0], nil
	}
}

func CreatePaper(paper *databaseModel.Paper) error {
	if paper == nil {
		return myError.PointerIsNil
	} else {
		paper.ID = 0
		if err := databaseManager.PaperManage.Create(paper); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func DeletePaperByID(id uint) error {
	if err := databaseManager.PaperManage.DeleteByID(id); err != nil {
		return err
	} else {
		return nil
	}
}
