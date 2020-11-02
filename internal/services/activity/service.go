package activity

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/danni-popova/todannigo/internal/repositories/activity"
	log "github.com/sirupsen/logrus"
)

type service struct {
	repo activity.Repository
}

func NewService(repo activity.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) ListActions(w http.ResponseWriter, r *http.Request) {
	// Retrieve user from request token
	userID := r.Context().Value("user_id")

	var actions []activity.Action
	actions, err := s.repo.List(userID.(int))

	// Return an error and exit
	if err != nil {
		writeFailure(w, err.Error())
		return
	}

	// Marshal response
	marshalled, err := json.Marshal(actions)
	if err != nil {
		log.Error(err)
	}

	err = writeSuccess(w, []byte(marshalled))
	if err != nil {
		log.Error(err)
	}
}

func (s *service) RecordAction(w http.ResponseWriter, r *http.Request) {
	// Retrieve user from request token
	userID := r.Context().Value("user_id")

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}
	log.Info(reqBody)

	var action activity.Action
	err = json.Unmarshal(reqBody, &action)
	if err != nil {
		log.Error(err)
	}
	action.UserID = userID.(int)

	err = s.repo.Add(action)
	if err != nil {
		log.Error(err)
	}

}

func writeSuccess(w http.ResponseWriter, r []byte) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(r)
	return err
}

func writeFailure(w http.ResponseWriter, e string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	errMessage := UnsuccessfulResponse{Error: e}
	marshalled, err := json.Marshal(errMessage)
	if err != nil {
		log.Error(err)
	}
	_, err = w.Write(marshalled)
	return err
}
