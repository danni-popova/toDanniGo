package projects

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

type Project struct {
	ID          int            `json:"id"     db:"id"`
	Creator     int            `json:"creator"  db:"creator"`
	IsDefault   bool           `json:"isDefault"   db:"is_default"`
	Title       string         `json:"title"  db:"title"`
	Description string         `json:"description,omitempty" db:"description"`
	Logo        string         `json:"logo,omitempty" db:"logo"`
	CreatedAt   time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt   sql.NullString `json:"updatedAt,omitempty" db:"updated_at"`
	DeletedAt   sql.NullString `json:"deletedAt,omitempty" db:"deleted_at"`
}

func (r *repository) Create(userID int, project Project) (Project, error) {
	var proj Project
	result, err := r.db.NamedQuery(
		`INSERT INTO projects(title, description, creator, is_default, logo) 
				VALUES (:title, :description, :creator, :isDefault, :logo) RETURNING *`,
		&project)
	if err != nil {
		return proj, err
	}

	result.Next()
	err = result.StructScan(&proj)
	return proj, err
}

func (r *repository) List(userID int) (projects []Project, err error) {
	err = r.db.Select(&projects,
		`SELECT * 
				FROM projects 
				INNER JOIN project_members pm ON projects.id = pm.project_id 
				WHERE pm.member_id=$1`, userID)
	return projects, err
}

func (r *repository) AddMember(projID, userID int) error {
	_, err := r.db.Exec(
		`INSERT INTO project_members VALUES (:project_id, :member_id)`,
		projID, userID)

	return err
}
