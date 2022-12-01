package dao

import "homePage/backend/model"

type ResourceManager struct {
}

func NewResourceManager() *ResourceManager {
	return &ResourceManager{}
}

func (m *ResourceManager) Get(newId uint) (Resource *model.Resource, err error) {
	Resource = &model.Resource{}
	err = db.GetDB().Model(&model.Resource{}).Where("id = ?", newId).First(Resource).Error
	return
}

func (m *ResourceManager) List(page, pageSize int) (Resource []*model.Resource, err error) {
	myDb := db.GetDB().Model(&model.Resource{}).Order("ID DESC").Offset(page * (pageSize - 1)).Limit(pageSize)
	if myDb.Error != nil {
		err = myDb.Error
		return
	}
	err = myDb.Find(&Resource).Error
	return
}

func (m *ResourceManager) Create(Resource *model.Resource) error {
	err := db.GetDB().Model(&model.Resource{}).Create(Resource).Error
	return err
}

func (m *ResourceManager) Update(Resource *model.Resource) error {
	err := db.GetDB().Model(&model.Resource{}).Where("id =?", Resource.ID).Updates(Resource).Error
	return err
}

func (m *ResourceManager) Delete(id uint) error {
	err := db.GetDB().Model(&model.Resource{}).Delete(&model.Resource{}, id).Error
	return err
}
