package usecases

import (
	"nbhd/models/request"
	"nbhd/models/response"

	"errors"
)

func (controller Controller) UsersGet(req request.UsersGet) (response.UsersGet, error) {

	var res response.UsersGet

	if err := controller.validator.Process(req); err != nil {
		return res, errors.New("invalid params")
	}

	user, err := controller.db.GetUserById(req.Id)

	if err != nil {
		return res, errors.New("internal error")
	}

	if user.Id == 0 {
		return res, errors.New("invalid user id")
	}

	res = response.UsersGet{
		Id:   user.Id,
		Name: user.Name,
	}

	return res, nil

}
