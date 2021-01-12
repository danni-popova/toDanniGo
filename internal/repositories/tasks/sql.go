package tasks

import (
	"gopkg.in/guregu/null.v3"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

type Task struct {
	gorm.Model
	Project     int         `json:"project"     db:"project"`
	Creator     int         `json:"creator"     db:"creator"`
	Assignee    int         `json:"assignee"     db:"assignee"`
	Done        bool        `json:"done"   db:"done"`
	Title       string      `json:"title"  db:"title"`
	Status      string      `json:"status"  db:"status"`
	CompletedAt null.String `json:"completedAt" db:"completed_at"`
	Deadline    null.String `json:"deadline" db:"deadline"`
	Description null.String `json:"description" db:"description"`
}

func (r *repository) InsertTask(task Task) (Task, error) {
	err := r.db.Create(&task).Error
	return task, err
}

func (r *repository) SelectTasksByProjectID(projectID int) (tasks []Task, err error) {
	err = r.db.Where("project=?", projectID).Find(&tasks).Error
	return tasks, err
}

func (r *repository) UpdateTask(task Task) (Task, error) {
	// Note: for some stupid reason, update, unlike the rest of the calls from the GORM
	// doesn't update the task after the query with the data from the DB.
	// So I instead perform a query to get the updated object.
	err := r.db.Model(&task).Updates(&task).Error
	r.db.First(&task, task.ID)
	return task, err
}

func (r *repository) DeleteTask(id int) error {
	err := r.db.Delete(&Task{}, id).Error
	return err
}
