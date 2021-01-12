package tasks

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var tables = []string{
	"tasks.sql",
}

func setup(t *testing.T) (Repository, func()) {
	// Some setup goes here - create test env
	var env int

	// Setup logger - don't have that so far

	// Setup repository
	r := NewRepository()

	return r, func() {
		require.NoError(t, env.Close())
	}
}

func TestRepository_InsertTask(t *testing.T) {
	r, close := setup(t)
	defer close()

	tcs := []struct {
		name           string
		taskToInsert   Task
		expectedResult Task
		expectedErr    bool
	}{
		{
			name: "Golden Path",
			taskToInsert: Task{
				Project:     1,
				Creator:     1,
				Assignee:    1,
				Done:        false,
				Title:       "Task Title",
				Status:      "Todo",
				Deadline:    time.Now().UTC().String(), // How the heck do I handle this null.String??
				Description: "Task Description",
			},
			expectedResult: Task{},
			expectedErr:    false,
		},
		{
			name: "Missing title",
			taskToInsert: Task{
				Project:     1,
				Creator:     1,
				Assignee:    1,
				Done:        false,
				Title:       "",
				Status:      "Todo",
				Deadline:    time.Now().UTC().String(), // How the heck do I handle this null.String??
				Description: "Task Description",
			},
			expectedResult: Task{},
			expectedErr:    false,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			task, err := r.InsertTask(tc.taskToInsert)
			if tc.expectedErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tc.expectedResult, task)
		})
	}

}

func TestRepository_SelectTasksByProjectID(t *testing.T) {

}

func TestRepository_UpdateTask(t *testing.T) {

}

func TestRepository_DeleteTask(t *testing.T) {

}
