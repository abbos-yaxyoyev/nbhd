package usecases

import (
	"errors"
	"nbhd/models/request"
	"nbhd/models/response"
)

func (controller Controller) UsersGet(req request.UsersGet) (response.UsersGet, error) {

	var res response.UsersGet

	if err := controller.validator.Process(req); err != nil {
		return res, errors.New("invalid params")
	}

	session, err := controller.db.GetSession(req.Token)

	if err != nil {
		return res, errors.New("internal error")
	}

	if session.Id == "" {
		return res, errors.New("invalid session id")
	}

	user, err := controller.db.GetUser(session.User)

	if err != nil {
		return res, errors.New("internal error")
	}

	person, err := controller.db.GetUser(req.Id)

	if err != nil {
		return res, errors.New("internal error")
	}

	if person.Id == "" {
		return res, errors.New("invalid user id")
	}

	res = response.UsersGet{
		Id:    user.Id,
		Name:  user.Name,
		Photo: user.Photo,
	}

	return res, nil

}
