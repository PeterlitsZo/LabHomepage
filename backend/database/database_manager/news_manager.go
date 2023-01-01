package databaseManager

import (
	databaseConnect "homePage/backend/database/database_connect"
	databaseError "homePage/backend/database/database_error"
	databaseModel "homePage/backend/database/database_model"

	"gorm.io/gorm"
)

type NewsManager struct {
	dbConn *databaseConnect.DbConn
}

func (m *NewsManager) RegisterDbConnection(dbConn *databaseConnect.DbConn) error {
	if dbConn == nil {
		return databaseError.DbConnectionIsNil
	}
	m.dbConn = dbConn
	return nil
}

func (m *NewsManager) GetDbConnection() *databaseConnect.DbConn {
	return m.dbConn
}

func (m *NewsManager) IsRegister() bool {
	return m.dbConn != nil
}

func (m *NewsManager) GetByFilter(filter *databaseModel.NewsFilter) ([]*databaseModel.News, error) {
	ret := make([]*databaseModel.News, 0)

	if !m.IsRegister() {
		return []*databaseModel.News{}, databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Model(&databaseModel.News{}); dbConn.Error != nil {
		return []*databaseModel.News{}, dbConn.Error
	} else if dbConn = dbConn.Where(filter).Order("id"); dbConn.Error != nil {
		return []*databaseModel.News{}, dbConn.Error
	} else if dbConn = dbConn.Find(&ret); dbConn.Error != nil && dbConn.Error != gorm.ErrRecordNotFound {
		return []*databaseModel.News{}, dbConn.Error
	} else {
		return ret, nil
	}
}

func (m *NewsManager) DeleteByID(id uint) error {
	if !m.IsRegister() {
		return databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Delete(&databaseModel.News{}, id); dbConn.Error != nil {
		return dbConn.Error
	} else {
		return nil
	}
}

func (m *NewsManager) Update(filter *databaseModel.NewsFilter, news *databaseModel.News) error {
	if !m.IsRegister() {
		return databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Model(filter); dbConn.Error != nil {
		return dbConn.Error
	} else if dbConn = dbConn.Updates(news); dbConn.Error != nil {
		return dbConn.Error
	} else {
		return nil
	}
}

func (m *NewsManager) Create(news *databaseModel.News) error {
	if !m.IsRegister() {
		return databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Model(news); dbConn.Error != nil {
		return dbConn.Error
	} else if dbConn = dbConn.Create(news); dbConn.Error != nil {
		return dbConn.Error
	} else {
		return nil
	}
}

func NewNewsManager(dbConn *databaseConnect.DbConn) (*NewsManager, error) {
	manager := new(NewsManager)
	err := manager.RegisterDbConnection(dbConn)
	return manager, err
}
