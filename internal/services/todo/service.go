package todo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/danni-popova/todannigo/internal/repositories/todo"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type service struct {
	repo todo.Repository
}

func NewService(repo todo.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateHttp(w http.ResponseWriter, r *http.Request) {
	var td todo.ToDo

	// Read request body and save into a todo
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	err = json.Unmarshal(reqBody, &td)
	if err != nil {
		log.Error(err)
	}

	// Add the user ID from the request context
	userID := r.Context().Value("user_id")
	td.UserID = userID.(int)

	var ctd todo.ToDo
	ctd, err = s.repo.Create(td)
	// Return an error and exit
	if err != nil {
		writeFailure(w, err.Error())
		return
	}

	// IDK if this is the best/worst way to do it
	rtd := Response{
		ID:          ctd.ID,
		Title:       ctd.Title,
		Description: ctd.Description,
		Deadline:    ctd.Deadline.String,
		Done:        ctd.Done,
		CreatedAt:   ctd.CreatedAt,
	}
	marshalled, err := json.Marshal(rtd)
	if err != nil {
		log.Error(err)
	}

	err = writeSuccess(w, []byte(marshalled))
	if err != nil {
		log.Error(err)
	}
}

func (s *service) GetHttp(w http.ResponseWriter, r *http.Request) {
	var td todo.ToDo
	pathParams := mux.Vars(r)
	reqToDoID := pathParams["id"]
	i, err := strconv.Atoi(reqToDoID)
	usrID := r.Context().Value("user_id")

	td, err = s.repo.Get(i, usrID.(int))
	// Return an error and exit
	if err != nil {
		writeFailure(w, err.Error())
		return
	}

	// IDK if this is the best/worst way to do it
	rtd := Response{
		ID:          td.ID,
		Title:       td.Title,
		Description: td.Description,
		Deadline:    td.Deadline.String,
		Done:        td.Done,
		CreatedAt:   td.CreatedAt,
	}
	marshalled, err := json.Marshal(rtd)
	if err != nil {
		log.Error(err)
	}

	err = writeSuccess(w, marshalled)
	if err != nil {
		log.Error(err)
	}
}

func (s *service) ListHttp(w http.ResponseWriter, r *http.Request) {
	log.Info("List was called")
	userID := r.Context().Value("user_id")
	var td []todo.ToDo
	td, err := s.repo.List(userID.(int))

	// Return an error and exit
	if err != nil {
		writeFailure(w, err.Error())
		return
	}

	// Marshal response
	marshalled, err := json.Marshal(td)
	if err != nil {
		log.Error(err)
	}

	err = writeSuccess(w, []byte(marshalled))
	if err != nil {
		log.Error(err)
	}
}

func (s *service) UpdateHttp(w http.ResponseWriter, r *http.Request) {
	log.Info("Update was called")

	pathParams := mux.Vars(r)
	rID := pathParams["id"]
	i, err := strconv.Atoi(rID)

	// Insert the userID from the validated token
	userID := r.Context().Value("user_id")

	// Update the value in the DB
	err = s.repo.Update(i, userID.(int))

	// Write response
	marshalled, err := json.Marshal(Response{
		Done: true,
	})

	if err != nil {
		log.Error(err)
	}

	err = writeSuccess(w, []byte(marshalled))
	if err != nil {
		log.Error(err)
	}

}

func (s *service) DeleteHttp(w http.ResponseWriter, r *http.Request) {
	log.Info("Delete was called")

	pathParams := mux.Vars(r)
	rID := pathParams["id"]

	// TODO: validate parameter because that screwed me the first time
	i, err := strconv.Atoi(rID)
	userID := r.Context().Value("user_id")
	err = s.repo.Delete(i, userID.(int))
	// Return an error and exit
	if err != nil {
		writeFailure(w, err.Error())
		return
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
