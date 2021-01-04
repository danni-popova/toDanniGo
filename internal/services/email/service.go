package email

import "net/http"

type service struct {
}

func (s service) SendVerificationEmail(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (s service) SendPasswordResetEmail(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func NewService() Service {
	return &service{}
}
