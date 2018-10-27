package response

type UsersSignUp struct {
	Token string `json:"token"`
	Id    string `json:"id"`
}

type UsersSignIn struct {
	Token string `json:"token"`
	Id    string `json:"id"`
}

type UsersSignOut struct{}

type UsersGet struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Photo  string  `json:"photo"`
	Rating float64 `json:"rating"`
}

type UsersLocationUpdate struct{}
