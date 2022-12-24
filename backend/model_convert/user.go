package modelConvert

import (
	databaseModel "homePage/backend/database/database_model"
	viewModel "homePage/backend/handler/view_model"
	"homePage/backend/util"
)

func ViewModel2DatabaseModelUser(user *viewModel.UserView) *databaseModel.User {
	password, _ := util.HashPassword(user.Password)
	return &databaseModel.User{
		ID:       user.ID,
		Name:     user.Name,
		Password: password,
		RoleId:   int32(databaseModel.Role2RoleId[user.Role]),
		Extra:    user.Extra,
	}
}

func DatabaseModel2ViewModelUser(user *databaseModel.User) *viewModel.UserView {
	return &viewModel.UserView{
		ID:    user.ID,
		Name:  user.Name,
		Role:  string(databaseModel.RoleId2Role[user.RoleId]),
		Extra: user.Extra,
	}
}

func DatabaseModel2ViewModelUserList(user []*databaseModel.User) []*viewModel.UserView {
	result := make([]*viewModel.UserView, 0)
	if user == nil {
		return result
	}
	for _, v := range user {
		result = append(result, DatabaseModel2ViewModelUser(v))
	}
	return result
}
