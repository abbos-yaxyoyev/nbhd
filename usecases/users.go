package usecases

import (
	"errors"
	"nbhd/models"
	"nbhd/models/request"
	"nbhd/models/response"
	"nbhd/tools/passwords"
	"nbhd/tools/uuid"
)

func (controller Controller) UsersSignUp(req request.UsersSignUp) (response.UsersSignUp, error) {

	var res response.UsersSignUp

	if err := controller.validator.Process(req); err != nil {
		return res, errors.New("invalid params")
	}

	user, err := controller.db.GetUserByPhone(req.Phone)

	if err != nil {
		return res, errors.New("internal error")
	}

	if user.Id != "" {
		return res, errors.New("phone is already used")
	}

	user = models.User{
		Id:       uuid.Generate(),
		Name:     req.Name,
		Phone:    req.Phone,
		Password: passwords.Encode(req.Password),
	}

	err = controller.db.StoreUser(user)

	if err != nil {
		return res, errors.New("internal error")
	}

	session := models.Session{
		Id:   uuid.Generate(),
		User: user.Id,
	}

	err = controller.db.StoreSession(session)

	if err != nil {
		return res, errors.New("internal error")
	}

	res = response.UsersSignUp{
		Token: session.Id,
		Id:    user.Id,
	}

	return res, nil

}

func (controller Controller) UsersSignIn(req request.UsersSignIn) (response.UsersSignIn, error) {

	var res response.UsersSignIn

	if err := controller.validator.Process(req); err != nil {
		return res, errors.New("invalid params")
	}

	user, err := controller.db.GetUserByPhone(req.Phone)

	if err != nil {
		return res, errors.New("internal error")
	}

	if user.Id == "" {
		return res, errors.New("invalid phone")
	}

	if !passwords.Validate(req.Password, user.Password) {
		return res, errors.New("invalid password")
	}

	session := models.Session{
		Id:   uuid.Generate(),
		User: user.Id,
	}

	err = controller.db.StoreSession(session)

	if err != nil {
		return res, errors.New("internal error")
	}

	res = response.UsersSignIn{
		Token: session.Id,
		Id:    user.Id,
	}

	return res, nil

}

func (controller Controller) UsersSignOut(req request.UsersSignOut) (response.UsersSignOut, error) {

	var res response.UsersSignOut

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

	err = controller.db.DeleteSession(session)

	if err != nil {
		return res, errors.New("internal error")
	}

	return res, nil

}

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
