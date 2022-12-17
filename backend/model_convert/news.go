package modelConvert

import (
	databaseModel "homePage/backend/database/database_model"
	viewModel "homePage/backend/handler/view_model"
)

func DatabaseModel2ViewModelNews(news *databaseModel.News) *viewModel.NewsView {
	return &viewModel.NewsView{
		ID:      news.ID,
		Title:   news.Title,
		Content: news.Content,
		Extra:   news.Extra,
	}
}

func ViewModel2DatabaseModelNews(news *viewModel.NewsView) *databaseModel.News {
	return &databaseModel.News{
		ID:      news.ID,
		Title:   news.Title,
		Content: news.Content,
		Extra:   news.Extra,
	}
}

func DatabaseModel2ViewModelNewsList(news []*databaseModel.News) []*viewModel.NewsView {
	result := make([]*viewModel.NewsView, 0)
	if news == nil {
		return result
	}
	for _, v := range news {
		result = append(result, DatabaseModel2ViewModelNews(v))
	}
	return result
}
