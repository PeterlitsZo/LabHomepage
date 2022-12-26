package databaseManager

import (
	databaseConnect "homePage/backend/database/database_connect"
	databaseError "homePage/backend/database/database_error"
	databaseModel "homePage/backend/database/database_model"

	"gorm.io/gorm"
)

type UserManager struct {
	dbConn *databaseConnect.DbConn
}

func (m *UserManager) RegisterDbConnection(dbConn *databaseConnect.DbConn) error {
	if dbConn == nil {
		return databaseError.DbConnectionIsNil
	}
	m.dbConn = dbConn
	return nil
}

func (m *UserManager) GetDbConnection() *databaseConnect.DbConn {
	return m.dbConn
}

func (m *UserManager) IsRegister() bool {
	return m.dbConn != nil
}

func (m *UserManager) GetByFilter(filter *databaseModel.UserFilter) ([]*databaseModel.User, error) {
	ret := make([]*databaseModel.User, 0)

	if !m.IsRegister() {
		return []*databaseModel.User{}, databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Model(&databaseModel.User{}); dbConn.Error != nil {
		return []*databaseModel.User{}, dbConn.Error
	} else if dbConn = dbConn.Where(filter); dbConn.Error != nil {
		return []*databaseModel.User{}, dbConn.Error
	} else if dbConn = dbConn.Find(&ret); dbConn.Error != nil && dbConn.Error != gorm.ErrRecordNotFound {
		return []*databaseModel.User{}, dbConn.Error
	} else {
		return ret, nil
	}
}

func (m *UserManager) DeleteByID(id uint) error {
	if !m.IsRegister() {
		return databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Delete(&databaseModel.User{}, id); dbConn.Error != nil {
		return dbConn.Error
	} else {
		return nil
	}
}

func (m *UserManager) Update(filter *databaseModel.UserFilter, user *databaseModel.User) error {
	if !m.IsRegister() {
		return databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Model(filter); dbConn.Error != nil {
		return dbConn.Error
	} else if dbConn = dbConn.Model(&databaseModel.User{}).Updates(user); dbConn.Error != nil {
		return dbConn.Error
	} else {
		return nil
	}
}

func (m *UserManager) Create(user *databaseModel.User) error {
	if !m.IsRegister() {
		return databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Model(user); dbConn.Error != nil {
		return dbConn.Error
	} else if dbConn = dbConn.Create(user); dbConn.Error != nil {
		return dbConn.Error
	} else {
		return nil
	}
}

func NewUserManager(dbConn *databaseConnect.DbConn) (*UserManager, error) {
	manager := new(UserManager)
	err := manager.RegisterDbConnection(dbConn)
	return manager, err
}
