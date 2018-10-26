package api

import (
	"nbhd/models/request"
	"net/http"
)

func (api API) TasksCreate(w http.ResponseWriter, r *http.Request) {

	var req request.TasksCreate

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.TasksCreate(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}
