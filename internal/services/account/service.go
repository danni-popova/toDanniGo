package account

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/danni-popova/todannigo/internal/pkg/response"
	"github.com/danni-popova/todannigo/internal/pkg/token"
	"github.com/danni-popova/todannigo/internal/repositories/account"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo account.Repository
}

func NewService(repo account.Repository) Service {
	return &service{
		repo: repo,
	}
}

// Authenticate servers the http endpoint that returns a JWT token
// when given valid credentials
func (s service) Authenticate(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	var loginRequest LoginRequest
	err = json.Unmarshal(reqBody, &loginRequest)
	if err != nil {
		log.Error(err)
	}

	// Check if both email and password are provided
	if loginRequest.Email == "" || loginRequest.Password == "" {
		response.WriteFailure(w, "password and email cannot be blank")
		return
	}

	authDetails, err := s.repo.SelectAuthDetails(loginRequest.Email)
	if err != nil {
		log.Error(err)
		response.WriteFailure(w, "incorrect username or password")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(authDetails.Password), []byte(loginRequest.Password))
	if err != nil {
		log.Error(err)
		response.WriteFailure(w, "wrong credentials")
		return
	}

	token := token.Generate(authDetails)
	responseString := fmt.Sprintf("{ \"token\" : \"%s\" }", token)
	err = response.WriteSuccess(w, []byte(responseString))
	if err != nil {
		log.Error(err)
	}
}

// Register serves the http endpoint that creates a new user
func (s service) Register(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	var registerRequest RegisterRequest
	err = json.Unmarshal(reqBody, &registerRequest)
	if err != nil {
		log.Error(err)
	}

	// TODO: validate user details
	pass, err := bcrypt.GenerateFromPassword([]byte(registerRequest.Password), token.Cost)
	if err != nil {
		log.Error(err)
		return
	}

	// Call insert user data
	usrData := account.AccountData{
		FirstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
		Role:      registerRequest.Role,
		Email:     registerRequest.Email,
		Password:  string(pass),
	}
	err = s.repo.InsertAccountData(usrData)
	if err != nil {
		log.Error(err)
	}

	//TODO: Trigger email verification process
	err = response.WriteSuccess(w, []byte(""))
	if err != nil {
		log.Error(err)
	}
}

// GetAccountDetails serves the http endpoint that returns
// publicly available info for a user given an ID
func (s service) GetAccountDetails(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	id := pathParams["id"]
	userID, err := strconv.Atoi(id)

	var accData account.AccountData
	accData, err = s.repo.SelectAccountDetails(userID)
	if err != nil {
		log.Error(err)
		response.WriteFailure(w, "user does not exist")
		return
	}

	marshalled, err := json.Marshal(accData)
	if err != nil {
		log.Error(err)
	}

	err = response.WriteSuccess(w, marshalled)
	if err != nil {
		log.Error(err)
	}
}
