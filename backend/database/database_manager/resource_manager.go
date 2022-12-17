package databaseManager

import (
	databaseConnect "homePage/backend/database/database_connect"
	databaseError "homePage/backend/database/database_error"
	databaseModel "homePage/backend/database/database_model"

	"gorm.io/gorm"
)

type ResourceManager struct {
	dbConn *databaseConnect.DbConn
}

func (m *ResourceManager) RegisterDbConnection(dbConn *databaseConnect.DbConn) error {
	if dbConn == nil {
		return databaseError.DbConnectionIsNil
	}
	m.dbConn = dbConn
	return nil
}

func (m *ResourceManager) GetDbConnection() *databaseConnect.DbConn {
	return m.dbConn
}

func (m *ResourceManager) IsRegister() bool {
	return m.dbConn != nil
}

func (m *ResourceManager) GetByFilter(filter *databaseModel.ResourceFilter) ([]*databaseModel.Resource, error) {
	ret := make([]*databaseModel.Resource, 0)

	if !m.IsRegister() {
		return []*databaseModel.Resource{}, databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Model(&databaseModel.Resource{}); dbConn.Error != nil {
		return []*databaseModel.Resource{}, dbConn.Error
	} else if dbConn = dbConn.Where(filter); dbConn.Error != nil {
		return []*databaseModel.Resource{}, dbConn.Error
	} else if dbConn = dbConn.Find(&ret); dbConn.Error != nil && dbConn.Error != gorm.ErrRecordNotFound {
		return []*databaseModel.Resource{}, dbConn.Error
	} else {
		return ret, nil
	}
}

func (m *ResourceManager) DeleteByID(id uint) error {
	if !m.IsRegister() {
		return databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Delete(&databaseModel.Resource{}, id); dbConn.Error != nil {
		return dbConn.Error
	} else {
		return nil
	}
}

func (m *ResourceManager) Update(filter *databaseModel.ResourceFilter, resource *databaseModel.Resource) error {
	if !m.IsRegister() {
		return databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Model(filter); dbConn.Error != nil {
		return dbConn.Error
	} else if dbConn = dbConn.Updates(resource); dbConn.Error != nil {
		return dbConn.Error
	} else {
		return nil
	}
}

func (m *ResourceManager) Create(resource *databaseModel.Resource) error {
	if !m.IsRegister() {
		return databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Model(resource); dbConn.Error != nil {
		return dbConn.Error
	} else if dbConn = dbConn.Create(resource); dbConn.Error != nil {
		return dbConn.Error
	} else {
		return nil
	}
}

func NewResourceManager(dbConn *databaseConnect.DbConn) (*ResourceManager, error) {
	manager := new(ResourceManager)
	err := manager.RegisterDbConnection(dbConn)
	return manager, err
}
