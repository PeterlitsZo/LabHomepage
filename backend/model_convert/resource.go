package modelConvert

import (
	databaseModel "homePage/backend/database/database_model"
	viewModel "homePage/backend/handler/view_model"
)

func ViewModel2DatabaseModelResource(resource *viewModel.ResourceView) *databaseModel.Resource {
	return &databaseModel.Resource{
		ID:      resource.ID,
		Title:   resource.Title,
		Content: resource.Content,
		Extra:   resource.Extra,
	}
}

func DatabaseModel2ViewModelResource(resource *databaseModel.Resource) *viewModel.ResourceView {
	return &viewModel.ResourceView{
		ID:      resource.ID,
		Title:   resource.Title,
		Content: resource.Content,
		Extra:   resource.Extra,
	}
}

func DatabaseModel2ViewModelResourceList(resource []*databaseModel.Resource) []*viewModel.ResourceView {
	result := make([]*viewModel.ResourceView, 0)
	if resource == nil {
		return result
	}
	for _, v := range resource {
		result = append(result, DatabaseModel2ViewModelResource(v))
	}
	return result
}
