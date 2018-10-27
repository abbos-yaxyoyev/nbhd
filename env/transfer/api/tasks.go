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

func (api API) TasksGet(w http.ResponseWriter, r *http.Request) {

	var req request.TasksGet

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.TasksGet(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}

func (api API) TasksList(w http.ResponseWriter, r *http.Request) {

	var req request.TasksList

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.TasksList(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}

func (api API) TasksDelete(w http.ResponseWriter, r *http.Request) {

	var req request.TasksDelete

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.TasksDelete(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}

func (api API) TasksPerformanceRequest(w http.ResponseWriter, r *http.Request) {

	var req request.TasksPerformanceRequest

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.TasksPerformanceRequest(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}

func (api API) TasksPerformanceCancel(w http.ResponseWriter, r *http.Request) {

	var req request.TasksPerformanceCancel

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.TasksPerformanceCancel(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}

func (api API) TasksPerformerAccept(w http.ResponseWriter, r *http.Request) {

	var req request.TasksPerformerAccept

	err := ProcessRequest(r, &req, api.checkHash, api.hashSalt)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	res, err := api.controller.TasksPerformerAccept(req)

	if err != nil {
		ReturnErrorResponse(w, err.Error())
		return
	}

	ReturnSuccessResponse(w, res)

}
