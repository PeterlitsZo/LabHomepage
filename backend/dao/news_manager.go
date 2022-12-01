package dao

import "homePage/backend/model"

type NewsManager struct {
}

func NewNewsManager() *NewsManager {
	return &NewsManager{}
}

func (m *NewsManager) Get(newId uint) (news *model.News, err error) {
	news = &model.News{}
	err = db.GetDB().Model(&model.News{}).Where("id = ?", newId).First(news).Error
	return
}

func (m *NewsManager) List(page, pageSize int) (news []*model.News, err error) {
	myDb := db.GetDB().Model(&model.News{}).Order("ID DESC").Offset(page * (pageSize - 1)).Limit(pageSize)
	if myDb.Error != nil {
		err = myDb.Error
		return
	}
	err = myDb.Find(&news).Error
	return
}

func (m *NewsManager) Create(news *model.News) error {
	err := db.GetDB().Model(&model.News{}).Create(news).Error
	return err
}

func (m *NewsManager) Update(news *model.News) error {
	err := db.GetDB().Model(&model.News{}).Where("id =?", news.ID).Updates(news).Error
	return err
}

func (m *NewsManager) Delete(id uint) error {
	err := db.GetDB().Model(&model.News{}).Delete(&model.News{}, id).Error
	return err
}
