package user

type RegisterRequest struct {
	Email     string `json:"email" db:"email"`
	Password  string `json:"password" db:"password"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
}

type Details struct {
	Email          string `json:"email" db:"email"`
	FirstName      string `json:"first_name" db:"first_name"`
	LastName       string `json:"last_name" db:"last_name"`
	ProfilePicture string `json:"profile_picture" db:"profile_picture"`
}
