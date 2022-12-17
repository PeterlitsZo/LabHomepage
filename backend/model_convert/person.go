package modelConvert

import (
	databaseModel "homePage/backend/database/database_model"
	viewModel "homePage/backend/handler/view_model"
)

func DatabaseModel2ViewModelPerson(person *databaseModel.Person) *viewModel.PersonView {
	return &viewModel.PersonView{
		ID:      person.ID,
		Name:    person.Name,
		Content: person.Content,
		Extra:   person.Extra,
	}
}

func ViewModel2DatabaseModelPerson(person *viewModel.PersonView) *databaseModel.Person {
	return &databaseModel.Person{
		ID:      person.ID,
		Name:    person.Name,
		Content: person.Content,
		Extra:   person.Extra,
	}
}

func DatabaseModel2ViewModelPersonList(person []*databaseModel.Person) []*viewModel.PersonView {
	result := make([]*viewModel.PersonView, 0)
	if person == nil {
		return result
	}
	for _, v := range person {
		result = append(result, DatabaseModel2ViewModelPerson(v))
	}
	return result
}
