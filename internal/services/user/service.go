package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/danni-popova/todannigo/internal/services/claims"

	"github.com/danni-popova/todannigo/internal/repositories/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
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

func (s *service) Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register request received")
	var usr user.User

	// Read new user details from request body
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	// Unmarshall body
	err = json.Unmarshal(reqBody, &usr)
	if err != nil {
		log.Error(err)
	}

	//TODO: Validate user details

	pass, err := bcrypt.GenerateFromPassword([]byte(usr.Password), 14)
	if err != nil {
		log.Error(err)
		return
	}
	usr.Password = string(pass)

	// Insert user into database
	// TODO: Convert email to lower case before inserting
	err = s.repo.Create(usr)
	if err != nil {
		log.Error(err)
		return
	}

	response := LoginResponse{Token: "valid-token"}
	marshalled, err := json.Marshal(response)
	if err != nil {
		log.Error(err)
		return
	}

	err = writeResponse(w, marshalled)
	if err != nil {
		log.Error(err)
	}
}

func (s *service) Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login request received")

	var lr LoginRequest
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error(err)
	}

	err = json.Unmarshal(reqBody, &lr)
	if err != nil {
		log.Error(err)
	}

	fmt.Println(lr)

	if lr.Email == "" || lr.Password == "" {
		writeFailure(w, "password and email cannot be blank")
	}

	usr, err := s.repo.GetByEmail(lr.Email)
	if err != nil {
		log.Error(err)
		writeFailure(w, "incorrect username or password")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(lr.Password))
	if err != nil {
		log.Error(err)
	}

	token := generateToken(usr)

	// Create response to write
	response := LoginResponse{Token: token}
	marshalled, err := json.Marshal(response)
	if err != nil {
		log.Error(err)
	}

	// Return token
	err = writeResponse(w, marshalled)
	if err != nil {
		log.Error(err)
	}
}

func (s *service) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUser request received")
	// Read ID from request
	pathParams := mux.Vars(r)
	id := pathParams["id"]
	rID, err := strconv.Atoi(id)

	// TODO: validate parameter
	var usr user.User
	usr, err = s.repo.GetByID(rID)
	if err != nil {
		log.Error(err)
	}

	// Map details
	uD := Details{
		FirstName:      usr.FirstName,
		LastName:       usr.LastName,
		Email:          usr.Email,
		ProfilePicture: usr.ProfilePicture,
	}

	marshalled, err := json.Marshal(uD)
	if err != nil {
		log.Error(err)
	}

	err = writeResponse(w, marshalled)
	if err != nil {
		log.Error(err)
	}
}

func (s *service) ResetPassword(w http.ResponseWriter, r *http.Request) {
	//TODO: To be done when the form is created
}

func writeResponse(w http.ResponseWriter, r []byte) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(r)
	return err
}

func generateToken(u user.User) string {
	// Create the Claims
	userInfoClaims := claims.UserClaims{
		UserID:         u.UserID,
		Email:          u.Email,
		ProfilePicture: u.ProfilePicture,
	}

	clms := claims.ToDanniClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    "todanni-user-service",
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
		UserInfo: userInfoClaims,
	}

	// Generate the Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clms)
	ss, err := token.SignedString([]byte(claims.HmacSampleSecret))
	if err != nil {
		log.Error(err)
	}
	return ss
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
