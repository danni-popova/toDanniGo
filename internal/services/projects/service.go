package projects

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/danni-popova/todannigo/internal/repositories/projects"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type service struct {
	repo projects.Repository
}

func NewService(repo projects.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) CreateProject(w http.ResponseWriter, r *http.Request) {
	var proj projects.Project
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	err = json.Unmarshal(reqBody, &proj)
	if err != nil {
		log.Error(err)
	}
	userID := r.Context().Value("user_id")
	proj.Creator = userID.(int)

	// Call sql create
	proj, err = s.repo.Create(proj)
	if err != nil {
		writeFailure(w, err.Error())
		return
	}

	// Call sql add member
	// TODO: this can cause issues if one is successful but the other one isn't
	// maybe make it a transaction?
	err = s.repo.AddMember(proj.ID, userID.(int))
	if err != nil {
		writeFailure(w, err.Error())
	}

	marshalled, err := json.Marshal(proj)
	if err != nil {
		log.Error(err)
	}

	err = writeSuccess(w, []byte(marshalled))
	if err != nil {
		log.Error(err)
	}
}

func (s *service) ListProjects(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id")
	var projcts []projects.Project
	log.Info("user ID is %d", userID.(int))
	projcts, err := s.repo.List(userID.(int))

	if err != nil {
		writeFailure(w, err.Error())
		return
	}

	marshalled, err := json.Marshal(projcts)
	if err != nil {
		log.Error(err)
	}

	err = writeSuccess(w, []byte(marshalled))
	if err != nil {
		log.Error(err)
	}
}

func (s *service) AddMember(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	reqProjID := pathParams["id"]
	projID, err := strconv.Atoi(reqProjID)

	// Read request body and save into a tasks
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	var memberID int
	err = json.Unmarshal(reqBody, &memberID)
	if err != nil {
		log.Error(err)
	}

	err = s.repo.AddMember(projID, memberID)
	if err != nil {
		writeFailure(w, err.Error())
	}

	err = writeSuccess(w, []byte("member added"))
	if err != nil {
		log.Error(err)
	}
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

func writeSuccess(w http.ResponseWriter, r []byte) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(r)
	return err
}
