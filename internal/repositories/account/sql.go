package account

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/guregu/null.v3"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

type AccountData struct {
	ID             int         `json:"id" db:"id"`
	Password       string      `json:"-" db:"password"`
	Email          string      `json:"-" db:"email"`
	FirstName      string      `json:"firstName" db:"first_name"`
	LastName       string      `json:"lastName" db:"last_name"`
	Role           string      `json:"role" db:"job_role"`
	CreatedAt      time.Time   `json:"createdAt" db:"created_at"`
	UpdatedAt      null.String `json:"-" db:"updated_at"`
	DeletedAt      null.String `json:"-" db:"deleted_at"`
	ProfilePicture null.String `json:"profilePicture" db:"profile_picture"`
}

type AuthDetails struct {
	ID       int    `json:"id" db:"id"`
	Password string `json:"password" db:"password"`
	Email    string `json:"email" db:"email"`
}

func (r *repository) InsertAccountData(acc AccountData) error {
	_, err := r.db.NamedExec(`INSERT INTO account_data(first_name, last_name, job_role, email, password) 
							   VALUES (:first_name, :last_name, :job_role, :email, :password)`, &acc)
	return err
}

func (r *repository) SelectAuthDetails(email string) (authDetails AuthDetails, err error) {
	err = r.db.QueryRow(`SELECT id, password FROM account_data WHERE email=$1`, email).
		Scan(&authDetails.ID, &authDetails.Password)
	return
}

func (r *repository) SelectAccountDetails(id int) (acc AccountData, err error) {
	err = r.db.QueryRowx(`SELECT * FROM account_data WHERE id = $1`, id).StructScan(&acc)
	return
}
