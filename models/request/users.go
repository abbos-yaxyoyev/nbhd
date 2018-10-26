package request

type UsersGet struct {
	Token string `json:"token" validate:"required,uuid"`
	Id    string `json:"id" validate:"required,uuid"`
}
