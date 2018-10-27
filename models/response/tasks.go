package response

type TasksCreate struct {
	Id string `json:"id"`
}

type TasksGet struct {
	Id            string              `json:"id"`
	Title         string              `json:"title"`
	Category      int                 `json:"category"`
	Location      []float64           `json:"location"`
	Distance      int                 `json:"distance"`
	Description   string              `json:"description"`
	Time          int                 `json:"time"`
	Creator       TasksGetCreator     `json:"creator"`
	Encouragement int                 `json:"encouragement"`
	Pay           float64             `json:"pay"`
	Created       string              `json:"created"`
	Performers    []TasksGetPerformer `json:"performers"`
}

type TasksGetPerformer struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Photo  string  `json:"photo"`
	Rating float64 `json:"rating"`
}

type TasksGetCreator struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Photo  string  `json:"photo"`
	Phone  string  `json:"phone"`
	Rating float64 `json:"rating"`
}

type TasksList struct {
	Tasks []TasksListTask `json:"tasks"`
}

type TasksListTask struct {
	Id            string           `json:"id"`
	Title         string           `json:"title"`
	Distance      int              `json:"distance"`
	Time          int              `json:"time"`
	Creator       TasksListCreator `json:"creator"`
	Encouragement int              `json:"encouragement"`
}

type TasksListCreator struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Photo  string  `json:"photo"`
	Rating float64 `json:"rating"`
}

type TasksDelete struct{}

type TasksPerformanceRequest struct{}

type TasksPerformanceCancel struct{}

type TasksPerformerAccept struct{}

type TasksPerformerDecline struct{}

type TasksRate struct{}
