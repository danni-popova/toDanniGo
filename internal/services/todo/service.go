package todo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/todannigo/internal/repositories/todo"
	"io/ioutil"
	"net/http"
	"strconv"
)

type service struct {
	repo todo.Repository
}

func NewService(repo todo.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetHttp(w http.ResponseWriter, r *http.Request) {
	var td todo.ToDo

	pathParams := mux.Vars(r)
	rID := pathParams["id"]

	// TODO: validate parameter because that screwed me the first time
	i, err := strconv.Atoi(rID)
	td, err = s.repo.Get(i)
	// Return an error and exit
	if err != nil {
		fmt.Println(err)
	}

	// IDK if this is the best/worst way to do it
	rtd := Response{
		ID:          td.ID,
		Title:       td.Title,
		Description: td.Description,
		Deadline:    td.Deadline,
		Done:        td.Done,
	}
	marshalled, err := json.Marshal(rtd)
	if err != nil {
		fmt.Println("Failed to marshal response")
	}

	err = writeResponse(w, []byte(marshalled))
	if err != nil {
		fmt.Println("Failed to write response")
	}
}

func (s *service) ListHttp(w http.ResponseWriter, r *http.Request) {
	var td []todo.ToDo
	td, err := s.repo.List()
	// Return an error and exit
	if err != nil {
		fmt.Println(err)
	}

	// Marshal response
	marshalled, err := json.Marshal(td)
	if err != nil {
		fmt.Println("Failed to marshal response")
	}

	err = writeResponse(w, []byte(marshalled))
	if err != nil {
		fmt.Println("Failed to write response")
	}
}

func (s *service) CreateHttp(w http.ResponseWriter, r *http.Request) {
	// TODO: validate parameter because that screwed me the first time
	var td todo.ToDo
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(reqBody, &td)
	if err != nil {
		fmt.Println(err)
	}

	err = s.repo.Create(td)
	// Return an error and exit
	if err != nil {
		fmt.Println(err)
	}

	// IDK if this is the best/worst way to do it
	rtd := Response{
		ID:          td.ID,
		Title:       td.Title,
		Description: td.Description,
		Deadline:    td.Deadline,
		Done:        td.Done,
	}
	marshalled, err := json.Marshal(rtd)
	if err != nil {
		fmt.Println("Failed to marshal response")
	}

	err = writeResponse(w, []byte(marshalled))
	if err != nil {
		fmt.Println("Failed to write response")
	}
}

func (s *service) UpdateHttp(w http.ResponseWriter, r *http.Request) {
	var td todo.ToDo

	pathParams := mux.Vars(r)
	rID := pathParams["id"]

	// TODO: validate parameter because that screwed me the first time
	i, err := strconv.Atoi(rID)
	td, err = s.repo.Get(i)
	// Return an error and exit
	if err != nil {
		fmt.Println(err)
	}

	// IDK if this is the best/worst way to do it
	rtd := Response{
		ID:          td.ID,
		Title:       td.Title,
		Description: td.Description,
		Deadline:    td.Deadline,
		Done:        td.Done,
	}
	marshalled, err := json.Marshal(rtd)
	if err != nil {
		fmt.Println("Failed to marshal response")
	}

	err = writeResponse(w, []byte(marshalled))
	if err != nil {
		fmt.Println("Failed to write response")
	}
}

func (s *service) DeleteHttp(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	rID := pathParams["id"]

	// TODO: validate parameter because that screwed me the first time
	i, err := strconv.Atoi(rID)
	err = s.repo.Delete(i)
	// Return an error and exit
	if err != nil {
		fmt.Println(err)
	}
}

func writeResponse(w http.ResponseWriter, r []byte) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(r)
	return err
}

func (s *service) Get(ctx context.Context, req *GetRequest) (*Response, error) {
	var td todo.ToDo
	td, err := s.repo.Get(req.ID)
	if err != nil {
		fmt.Println(err)
	}

	// IDK if this is the best/worst way to do it
	rtd := Response{
		ID:          td.ID,
		Title:       td.Title,
		Description: td.Description,
		Deadline:    td.Deadline,
		Done:        td.Done,
	}

	return &rtd, nil
}
func (s *service) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	var resp ListResponse

	return &resp, nil
}
func (s *service) Create(ctx context.Context, req *CreateRequest) (*Response, error) {
	var resp Response

	return &resp, nil
}
func (s *service) Update(ctx context.Context, req *UpdateRequest) (*Response, error) {
	var resp Response

	return &resp, nil
}
func (s *service) Delete(ctx context.Context, req *DeleteRequest) error {

	return nil
}
