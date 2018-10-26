package response

type TasksCreate struct {
	Id string `json:"id"`
}

type TasksGet struct {
	Id            string          `json:"id"`
	Title         string          `json:"title"`
	Category      int             `json:"category"`
	Location      []float64       `json:"location"`
	Distance      int             `json:"distance"`
	Description   string          `json:"description"`
	Time          int             `json:"time"`
	Creator       TasksGetCreator `json:"creator"`
	Encouragement int             `json:"encouragement"`
	Pay           float64         `json:"pay"`
	Created       string          `json:"created"`
}

type TasksGetCreator struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
	Phone string `json:"phone"`
}
