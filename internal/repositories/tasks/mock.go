package tasks

import "github.com/stretchr/testify/mock"

type RepositoryMock struct {
	mock.Mock
}

func (r *RepositoryMock) InsertTask(task Task) (Task, error) {
	args := r.Called(task)
	return args.Get(0).(Task), args.Error(1)
}

func (r *RepositoryMock) SelectTasksByProjectID(projectID int) (tasks []Task, err error) {
	args := r.Called(projectID)
	return args.Get(0).([]Task), args.Error(1)
}

func (r *RepositoryMock) UpdateTask(task Task) (Task, error) {
	args := r.Called(task)
	return args.Get(0).(Task), args.Error(1)
}

func (r *RepositoryMock) DeleteTask(id int) error {
	args := r.Called(id)
	return args.Error(1)
}
