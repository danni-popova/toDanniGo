package friends

import (
	"encoding/json"
	"net/http"

	"github.com/danni-popova/todannigo/internal/repositories/friends"
	log "github.com/sirupsen/logrus"
)

type service struct {
	repo friends.Repository
}

func (s service) List(w http.ResponseWriter, r *http.Request) {
	log.Info("List friends was called")
	userID := r.Context().Value("user_id")
	var friendships []friends.Friendship
	friendships, err := s.repo.List(userID.(int))

	// Return an error and exit
	if err != nil {
		writeFailure(w, err.Error())
		return
	}

	// Marshal response
	marshalled, err := json.Marshal(friendships)
	if err != nil {
		log.Error(err)
	}

	// Write successful response
	err = writeSuccess(w, []byte(marshalled))
	if err != nil {
		log.Error(err)
	}
}

func (s service) Modify(w http.ResponseWriter, r *http.Request) {
	log.Info("Modify friendship was called")
	userID := r.Context().Value("user_id")
}

func (s service) Send(w http.ResponseWriter, r *http.Request) {
	log.Info("Send friend request was called")
	userID := r.Context().Value("user_id")
}

func NewService(repo friends.Repository) Service {
	return &service{
		repo: repo,
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

type UnsuccessfulResponse struct {
	Error string `json:"error"`
}
