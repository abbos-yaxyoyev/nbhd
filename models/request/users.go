package request

type UsersSignUp struct {
	Phone    string `json:"phone" validate:"required,phone"`
	Name     string `json:"name" validate:"required,min=4,max=75"`
	Password string `json:"password" validate:"required,password"`
}

type UsersSignIn struct {
	Phone    string `json:"phone" validate:"required,phone"`
	Password string `json:"password" validate:"required,password"`
}

type UsersSignOut struct {
	Token string `json:"token" validate:"required,uuid"`
}

type UsersGet struct {
	Token string `json:"token" validate:"required,uuid"`
	Id    string `json:"id" validate:"required,uuid"`
}

type UsersLocationUpdate struct {
	Token    string    `json:"token" validate:"required,uuid"`
	Location []float64 `json:"location" validate:"required,len=2"`
}
