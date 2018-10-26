package usecases

import (
	"errors"
	"nbhd/config/app"
	"nbhd/models"
	"nbhd/models/request"
	"nbhd/models/response"
	"nbhd/tools/datetime"
	"nbhd/tools/distance"
	"nbhd/tools/uuid"
)

func (controller Controller) TasksCreate(req request.TasksCreate) (response.TasksCreate, error) {

	var res response.TasksCreate

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

	task := models.Task{
		Id:            uuid.Generate(),
		Title:         req.Title,
		Category:      req.Category,
		Location:      req.Location,
		Description:   req.Description,
		Time:          req.Time,
		Creator:       user.Id,
		Performer:     uuid.Default,
		Encouragement: req.Encouragement,
		Pay:           req.Pay,
		Created:       datetime.Generate(),
	}

	err = controller.db.CreateTask(task)

	if err != nil {
		return res, errors.New("internal error")
	}

	res = response.TasksCreate{
		Id: task.Id,
	}

	return res, nil

}

func (controller Controller) TasksGet(req request.TasksGet) (response.TasksGet, error) {

	var res response.TasksGet

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

	task, err := controller.db.GetTask(req.Id)

	if err != nil {
		return res, errors.New("internal error")
	}

	if task.Id == "" {
		return res, errors.New("invalid task id")
	}

	taskCreator, err := controller.db.GetUser(task.Creator)

	if err != nil {
		return res, errors.New("internal error")
	}

	res = response.TasksGet{
		Id:            task.Id,
		Title:         task.Title,
		Category:      task.Category,
		Description:   task.Description,
		Time:          task.Time,
		Encouragement: task.Encouragement,
		Pay:           task.Pay,
		Created:       task.Created,
	}

	taskDistance := distance.Calculate(user.Location, task.Location)

	switch {
	case taskDistance <= 500:
		res.Distance = app.Distance500
	case taskDistance <= 1000:
		res.Distance = app.Distance1000
	case taskDistance <= 2000:
		res.Distance = app.Distance2000
	case taskDistance <= 5000:
		res.Distance = app.Distance5000
	case taskDistance > 5000:
		res.Distance = app.DistanceFar
	}

	res.Creator = response.TasksGetCreator{
		Id:    taskCreator.Id,
		Name:  taskCreator.Name,
		Photo: taskCreator.Photo,
	}

	if task.Performer == user.Id {
		res.Location = task.Location
		res.Creator.Phone = user.Phone
	}

	return res, nil

}
