package dao

import "homePage/backend/model"

type PersonManager struct {
}

func NewPersonManager() *PersonManager {
	return &PersonManager{}
}

func (m *PersonManager) Get(newId uint) (Person *model.Person, err error) {
	Person = &model.Person{}
	err = db.GetDB().Model(&model.Person{}).Where("id = ?", newId).First(Person).Error
	return
}

func (m *PersonManager) List(page, pageSize int) (Person []*model.Person, err error) {
	myDb := db.GetDB().Model(&model.Person{}).Order("ID DESC").Offset(page * (pageSize - 1)).Limit(pageSize)
	if myDb.Error != nil {
		err = myDb.Error
		return
	}
	err = myDb.Find(&Person).Error
	return
}

func (m *PersonManager) Create(Person *model.Person) error {
	err := db.GetDB().Model(&model.Person{}).Create(Person).Error
	return err
}

func (m *PersonManager) Update(Person *model.Person) error {
	err := db.GetDB().Model(&model.Person{}).Where("id =?", Person.ID).Updates(Person).Error
	return err
}

func (m *PersonManager) Delete(id uint) error {
	err := db.GetDB().Model(&model.Person{}).Delete(&model.Person{}, id).Error
	return err
}
