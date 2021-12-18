package controllers

import (
	"github.com/fadhiilrachman/e-bpm/graph/model"
	utils "github.com/fadhiilrachman/e-bpm/utils"
)

type LoginData struct {
	Role     string `json:"id"`
	Username string `json:"username"`
}

func IsUsernameExist(username string) (*LoginData, error) {
	var user model.User
	err := dbConnect.Model(&user).Where("username=?", username).Select()

	if err != nil {
		panic(err)
	}

	data := &LoginData{
		Role:     user.Role,
		Username: user.Username,
	}

	return data, nil
}

func AuthUser(username, password string) (*LoginData, bool) {
	var user model.User
	err := dbConnect.Model(&user).Where("username=?", username).Select()

	if err != nil {
		panic(err)
	}

	check := utils.CheckPassword(password, user.Password)

	_, updateErr := dbConnect.Model(&user).
		Set("last_login_at=?", utils.TimeNow()).
		Where("id=?", user.ID).
		Update()

	if updateErr != nil {
		panic(err)
	}

	data := &LoginData{
		Role:     user.Role,
		Username: user.Username,
	}

	return data, check
}