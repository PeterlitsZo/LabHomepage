package modelConvert

import (
	databaseModel "homePage/backend/database/database_model"
	viewModel "homePage/backend/handler/view_model"
)

func ViewModel2DatabaseModelPaper(paper *viewModel.PaperView) *databaseModel.Paper {
	return &databaseModel.Paper{
		ID:      paper.ID,
		Title:   paper.Title,
		Content: paper.Content,
		Extra:   paper.Extra,
	}
}

func DatabaseModel2ViewModelPaper(paper *databaseModel.Paper) *viewModel.PaperView {
	return &viewModel.PaperView{
		ID:      paper.ID,
		Title:   paper.Title,
		Content: paper.Content,
		Extra:   paper.Extra,
	}
}

func DatabaseModel2ViewModelPaperList(paper []*databaseModel.Paper) []*viewModel.PaperView {
	result := make([]*viewModel.PaperView, 0)
	if paper == nil {
		return result
	}
	for _, v := range paper {
		result = append(result, DatabaseModel2ViewModelPaper(v))
	}
	return result
}
