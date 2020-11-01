package user

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"Token"`
}

type RegisterRequest struct {
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
}

type Details struct {
	Email          string `json:"email" db:"email"`
	FirstName      string `json:"firstName" db:"first_name"`
	LastName       string `json:"lastName" db:"last_name"`
	ProfilePicture string `json:"profilePicture" db:"profile_picture"`
}

type UnsuccessfulResponse struct {
	Error string `json:"error"`
}
