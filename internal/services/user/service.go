package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/danni-popova/todannigo/internal/repositories/user"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo user.Repository
}

func NewService(repo user.Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) Login(w http.ResponseWriter, r *http.Request) {
	var lr LoginRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(reqBody, &lr)
	if err != nil {
		fmt.Println("failed to unmarshal")
		fmt.Println(err)
	}

	// TODO: validate request body contents
	expPass, err := s.repo.GetPassword(lr.Email)
	if err != nil {
		fmt.Println("")
		fmt.Println(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(expPass), []byte(lr.Password))
	if err != nil {
		fmt.Println("")
		fmt.Println(err)
	}

	// Create response to write
	response := LoginResponse{Token: "valid-token"}
	marshalled, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}

	// Return token
	err = writeResponse(w, marshalled)
	if err != nil {
		fmt.Println(err)
	}
}

func (s *service) Register(w http.ResponseWriter, r *http.Request) {
	var user user.User

	// Read new user details from request body
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Unmarshall body
	err = json.Unmarshal(reqBody, &user)
	if err != nil {
		fmt.Println(err)
	}

	//TODO: Validate user details

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		fmt.Println("error generating from password")
		fmt.Println(err)
	}
	user.Password = string(pass)

	// Insert user into database
	// TODO: Convert email to lower case before inserting
	err = s.repo.Create(user)
	if err != nil {
		fmt.Println("error creating user:")
		fmt.Println(err)
	}

	response := LoginResponse{Token: "valid-token"}
	marshalled, err := json.Marshal(response)
	if err != nil {
		fmt.Println("Failed to marshal response")
	}

	err = writeResponse(w, marshalled)
	if err != nil {
		fmt.Println("Failed to write response")
	}
}

func writeResponse(w http.ResponseWriter, r []byte) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(r)
	return err
}
