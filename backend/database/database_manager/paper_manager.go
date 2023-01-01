package databaseManager

import (
	databaseConnect "homePage/backend/database/database_connect"
	databaseError "homePage/backend/database/database_error"
	databaseModel "homePage/backend/database/database_model"

	"gorm.io/gorm"
)

type PaperManager struct {
	dbConn *databaseConnect.DbConn
}

func (m *PaperManager) RegisterDbConnection(dbConn *databaseConnect.DbConn) error {
	if dbConn == nil {
		return databaseError.DbConnectionIsNil
	}
	m.dbConn = dbConn
	return nil
}

func (m *PaperManager) GetDbConnection() *databaseConnect.DbConn {
	return m.dbConn
}

func (m *PaperManager) IsRegister() bool {
	return m.dbConn != nil
}

func (m *PaperManager) GetByFilter(filter *databaseModel.PaperFilter) ([]*databaseModel.Paper, error) {
	ret := make([]*databaseModel.Paper, 0)

	if !m.IsRegister() {
		return []*databaseModel.Paper{}, databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Model(&databaseModel.Paper{}); dbConn.Error != nil {
		return []*databaseModel.Paper{}, dbConn.Error
	} else if dbConn = dbConn.Where(filter); dbConn.Error != nil {
		return []*databaseModel.Paper{}, dbConn.Error
	} else if dbConn = dbConn.Find(&ret); dbConn.Error != nil && dbConn.Error != gorm.ErrRecordNotFound {
		return []*databaseModel.Paper{}, dbConn.Error
	} else {
		return ret, nil
	}
}

func (m *PaperManager) DeleteByID(id uint) error {
	if !m.IsRegister() {
		return databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Delete(&databaseModel.Paper{}, id); dbConn.Error != nil {
		return dbConn.Error
	} else {
		return nil
	}
}

func (m *PaperManager) Update(filter *databaseModel.PaperFilter, paper *databaseModel.Paper) error {
	if !m.IsRegister() {
		return databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Model(filter); dbConn.Error != nil {
		return dbConn.Error
	} else if dbConn = dbConn.Updates(paper); dbConn.Error != nil {
		return dbConn.Error
	} else {
		return nil
	}
}

func (m *PaperManager) Create(paper *databaseModel.Paper) error {
	if !m.IsRegister() {
		return databaseError.DbConnectionNotRegistered
	} else if dbConn := m.GetDbConnection().Model(paper); dbConn.Error != nil {
		return dbConn.Error
	} else if dbConn = dbConn.Create(paper); dbConn.Error != nil {
		return dbConn.Error
	} else {
		return nil
	}
}

func NewPaperManager(dbConn *databaseConnect.DbConn) (*PaperManager, error) {
	manager := new(PaperManager)
	err := manager.RegisterDbConnection(dbConn)
	return manager, err
}
