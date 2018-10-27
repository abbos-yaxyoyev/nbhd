package response

type UsersSignUp struct {
	Token string `json:"token"`
	Id    string `json:"id"`
}

type UsersGet struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Photo string `json:"photo"`
}
