package dao

import (
	"homePage/backend/model"

	"gorm.io/gorm"
)

type UserManager struct {
}

func NewUserManager() *UserManager {
	return &UserManager{}
}

func (m *UserManager) GetByUsernameAndPassword(username, password string) (User *model.User, err error) {
	err = db.GetDB().Model(&model.User{}).Where("username = ? and password = ?", username, password).Find(&User).Error
	if err == nil && User.ID == 0 {
		err = gorm.ErrRecordNotFound
	}
	return
}

func (m *UserManager) Exists(username string, password string) (bool, error) {
	_, err := m.GetByUsernameAndPassword(username, password)
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func (m *UserManager) Get(newId uint) (User *model.User, err error) {
	User = &model.User{}
	err = db.GetDB().Model(&model.User{}).Where("ID = ?", newId).First(&User).Error
	return
}

func (m *UserManager) List(page, pageSize int) (User []*model.User, err error) {
	myDb := db.GetDB().Model(&model.User{}).Order("ID DESC").Offset(page * (pageSize - 1)).Limit(pageSize)
	if myDb.Error != nil {
		err = myDb.Error
		return
	}
	err = myDb.Find(&User).Error
	return
}

func (m *UserManager) Create(User *model.User) error {
	err := db.GetDB().Model(&model.User{}).Create(User).Error
	return err
}

func (m *UserManager) Update(User *model.User) error {
	err := db.GetDB().Model(&model.User{}).Where("id =?", User.ID).Updates(User).Error
	return err
}

func (m *UserManager) Delete(id string) error {
	err := db.GetDB().Model(&model.User{}).Delete(&model.User{}, id).Error
	return err
}
