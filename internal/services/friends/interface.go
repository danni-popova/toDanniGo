package friends

import "net/http"

type Service interface {
	// List all friends - pending, accepted and invited
	List(w http.ResponseWriter, r *http.Request)

	// Modify friendship - accept, reject, block
	Modify(w http.ResponseWriter, r *http.Request)

	// Send friendship request - existing user or email
	Send(w http.ResponseWriter, r *http.Request)
}
