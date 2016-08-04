package Domain

import (
	"phalgo-sample/Api/User/Model"
	"errors"
)

func (this *Domain_User)GetUserInfo(id int) (Model.User, error) {

	user, err := this.Model.User.GetInfoById(id)
	if err != nil {
		return user, errors.New("UserInfo There is no")
	}
	return user, nil
}

func (this *Domain_User)GetUserList() ([]Model.User, error) {

	user, err := this.Model.User.GetList()
	if err != nil {
		return user, errors.New("UserInfo There is no")
	}
	return user, nil

}