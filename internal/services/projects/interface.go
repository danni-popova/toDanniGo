package projects

import "net/http"

type Service interface {
	CreateProject(w http.ResponseWriter, r *http.Request)

	ListProjects(w http.ResponseWriter, r *http.Request)

	AddMember(w http.ResponseWriter, r *http.Request)
}
