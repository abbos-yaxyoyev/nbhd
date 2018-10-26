package usecases

import (
	"errors"
	"nbhd/models"
	"nbhd/models/request"
	"nbhd/models/response"
	"nbhd/tools/datetime"
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

	res.Id = task.Id

	return res, nil

}
