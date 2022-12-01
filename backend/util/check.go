package util

import "homePage/backend/dao"

func CheckPassword(username, password string) (bool, error) {
	if username == "" || password == "" {
		return false, nil
	}
	return dao.Manager.UserManager.Exists(username, password)
}
