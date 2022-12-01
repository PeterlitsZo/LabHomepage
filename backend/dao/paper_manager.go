package dao

import "homePage/backend/model"

type PaperManager struct {
}

func NewPaperManager() *PaperManager {
	return &PaperManager{}
}

func (m *PaperManager) Get(paperId uint) (Paper *model.Paper, err error) {
	Paper = &model.Paper{}
	err = db.GetDB().Model(&model.Paper{}).Where("id = ?", paperId).First(Paper).Error
	return
}

func (m *PaperManager) List(page, pageSize int) (Paper []*model.Paper, err error) {
	myDb := db.GetDB().Model(&model.Paper{}).Order("ID DESC").Offset(page * (pageSize - 1)).Limit(pageSize)
	if myDb.Error != nil {
		err = myDb.Error
		return
	}
	err = myDb.Find(&Paper).Error
	return
}

func (m *PaperManager) Create(Paper *model.Paper) error {
	err := db.GetDB().Model(&model.Paper{}).Create(Paper).Error
	return err
}

func (m *PaperManager) Update(Paper *model.Paper) error {
	err := db.GetDB().Model(&model.Paper{}).Where("id =?", Paper.ID).Updates(Paper).Error
	return err
}

func (m *PaperManager) Delete(id uint) error {
	err := db.GetDB().Model(&model.Paper{}).Delete(&model.Paper{}, id).Error
	return err
}
