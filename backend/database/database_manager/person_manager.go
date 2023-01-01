package databaseManager

import (
	databaseConnect "homePage/backend/database/database_connect"
	databaseError "homePage/backend/database/database_error"
	databaseModel "homePage/backend/database/database_model"

	"gorm.io/gorm"
)

type PersonManager struct {
	dbConn *databaseConnect.DbConn
}

func (m *PersonManager) RegisterDbConnection(dbConn *databaseConnect.DbConn) error {
	if dbConn == nil {
		return databaseError.DbConnectionIsNil
	}
	m.dbConn = dbConn
	return nil
}

func (m *PersonManager) GetDbConnection() *databaseConnect.DbConn {
	return m.dbConn
}

func (m *PersonManager) IsRegister() bool {
	return m.dbConn != nil
}

func (m *PersonManager) GetByFilter(filter *databaseModel.PersonFilter) ([]*databaseModel.Person, error) {
	ret := make([]*databaseModel.Person, 0)

	if !m.IsRegister() {
		return []*databaseModel.Person{}, databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Model(&databaseModel.Person{}); dbConn.Error != nil {
		return []*databaseModel.Person{}, dbConn.Error
	} else if dbConn = dbConn.Where(filter); dbConn.Error != nil {
		return []*databaseModel.Person{}, dbConn.Error
	} else if dbConn = dbConn.Find(&ret); dbConn.Error != nil && dbConn.Error != gorm.ErrRecordNotFound {
		return []*databaseModel.Person{}, dbConn.Error
	} else {
		return ret, nil
	}
}

func (m *PersonManager) DeleteByID(id uint) error {
	if !m.IsRegister() {
		return databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Delete(&databaseModel.Person{}, id); dbConn.Error != nil {
		return dbConn.Error
	} else {
		return nil
	}
}

func (m *PersonManager) Update(filter *databaseModel.PersonFilter, person *databaseModel.Person) error {
	if !m.IsRegister() {
		return databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Model(filter); dbConn.Error != nil {
		return dbConn.Error
	} else if dbConn = dbConn.Updates(person); dbConn.Error != nil {
		return dbConn.Error
	} else {
		return nil
	}
}

func (m *PersonManager) Create(person *databaseModel.Person) error {
	if !m.IsRegister() {
		return databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Model(person); dbConn.Error != nil {
		return dbConn.Error
	} else if dbConn = dbConn.Create(person); dbConn.Error != nil {
		return dbConn.Error
	} else {
		return nil
	}
}

func NewPersonManager(dbConn *databaseConnect.DbConn) (*PersonManager, error) {
	manager := new(PersonManager)
	err := manager.RegisterDbConnection(dbConn)
	return manager, err
}
