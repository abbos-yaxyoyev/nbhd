package api

import (
	"nbhd/models/request"
	"net/http"
)

func (api API) UsersSignUp(w http.ResponseWriter, r *http.Request) {

	var req request.UsersSignUp

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.UsersSignUp(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}

func (api API) UsersGet(w http.ResponseWriter, r *http.Request) {

	var req request.UsersGet

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.UsersGet(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}
