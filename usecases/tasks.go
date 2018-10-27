package usecases

import (
	"errors"
	"sort"

	"nbhd/config/app"
	"nbhd/models"
	"nbhd/models/request"
	"nbhd/models/response"
	"nbhd/tools/datetime"
	"nbhd/tools/distance"
	"nbhd/tools/numbers"
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

func (controller Controller) TasksList(req request.TasksList) (response.TasksList, error) {

	var res response.TasksList

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

	userLat := numbers.TruncateFloat(user.Location[0], 2)
	userLong := numbers.TruncateFloat(user.Location[1], 2)

	searchArea := [4]float64{
		userLat - 0.01, userLat + 0.01,
		userLong - 0.01, userLong + 0.01,
	}

	tasks, err := controller.db.ListTasks(searchArea)

	if err != nil {
		return res, errors.New("internal error")
	}

	sort.Slice(tasks, func(i, j int) bool {
		return distance.Calculate(user.Location, tasks[i].Location) > distance.Calculate(user.Location, tasks[j].Location)
	})

	resTasks := make([]response.TasksListTask, 0, len(tasks))

	for _, task := range tasks {

		resTask := response.TasksListTask{
			Id:            task.Id,
			Title:         task.Title,
			Time:          task.Time,
			Encouragement: task.Encouragement,
		}

		taskDistance := distance.Calculate(user.Location, task.Location)

		switch {
		case taskDistance <= 500:
			resTask.Distance = app.Distance500
		case taskDistance <= 1000:
			resTask.Distance = app.Distance1000
		case taskDistance <= 2000:
			resTask.Distance = app.Distance2000
		case taskDistance <= 5000:
			resTask.Distance = app.Distance5000
		case taskDistance > 5000:
			resTask.Distance = app.DistanceFar
		}

		taskCreator, err := controller.db.GetUser(task.Creator)

		if err != nil {
			return res, errors.New("internal error")
		}

		resTask.Creator = response.TasksListCreator{
			Id:    taskCreator.Id,
			Name:  taskCreator.Name,
			Photo: taskCreator.Photo,
		}

		resTasks = append(resTasks, resTask)

	}

	res = response.TasksList{
		Tasks: resTasks,
	}

	return res, nil

}
