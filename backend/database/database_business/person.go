package databaseBusiness

import (
	databaseManager "homePage/backend/database/database_manager"
	databaseModel "homePage/backend/database/database_model"
	myError "homePage/backend/my_error"
)

func GetPersonByID(id uint) (*databaseModel.Person, error) {
	if ret, err := databaseManager.PersonManage.GetByFilter(&databaseModel.PersonFilter{
		ID: id,
	}); err != nil {
		return nil, err
	} else if ret == nil || len(ret) == 0 {
		return nil, nil
	} else {
		return ret[0], nil
	}
}

func ListPerson() ([]*databaseModel.Person, error) {
	if ret, err := databaseManager.PersonManage.GetByFilter(&databaseModel.PersonFilter{}); err != nil {
		return []*databaseModel.Person{}, err
	} else {
		return ret, nil
	}
}

func UpdatePersonByID(id uint, person *databaseModel.Person) error {
	if person == nil {
		return myError.PointerIsNil
	} else {
		person.ID = 0
		if err := databaseManager.PersonManage.Update(&databaseModel.PersonFilter{
			ID: id,
		}, person); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func GetPersonByName(name string) (*databaseModel.Person, error) {
	if ret, err := databaseManager.PersonManage.GetByFilter(&databaseModel.PersonFilter{
		Name: name,
	}); err != nil {
		return nil, err
	} else if ret == nil || len(ret) == 0 {
		return nil, nil
	} else {
		return ret[0], nil
	}
}

func CreatePerson(person *databaseModel.Person) error {
	if person == nil {
		return myError.PointerIsNil
	} else {
		person.ID = 0
		if err := databaseManager.PersonManage.Create(person); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func DeletePersonByID(id uint) error {
	if err := databaseManager.PersonManage.DeleteByID(id); err != nil {
		return err
	} else {
		return nil
	}
}
