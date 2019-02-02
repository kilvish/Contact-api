package users

import (
	"encoding/json"
	"net/http"
	"strings"

	serr "github.com/anil/contacts-api/error"
	"github.com/gorilla/mux"
)

type addUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (aur *addUserRequest) validate() error {
	if strings.TrimSpace(aur.Username) == "" {
		return serr.ErrUserNameMissingAddUserRequest
	}

	if strings.TrimSpace(aur.Email) == "" {
		return serr.ErrEmailMissingAddUserRequest
	}
	return nil
}

func (aur *addUserRequest) load(r *http.Request) error {
	if r == nil {
		return serr.ErrInvalidRequest
	} else if r.Body == nil {
		return serr.ErrInvalidOrMissingBody
	}
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(aur)
	if err != nil {
		return serr.ErrDecodingAddUserRequestBody
	}
	return nil
}

type getUserRequest struct {
	Email *string `json:"email,omitempty"`
	ID    *string
}

func (gur *getUserRequest) validate() error {
	if gur.Email == nil && gur.ID == nil {
		return serr.ErrInvalidGetUserRequest
	}
	return nil
}

func (gur *getUserRequest) load(r *http.Request) error {
	if r == nil {
		return serr.ErrInvalidRequest
	}

	for query, value := range r.URL.Query() {
		if query == "email" {
			gur.Email = &value[0]
		}
	}
	userID := mux.Vars(r)["user_id"]
	if userID != "" {
		gur.ID = &userID
	}
	return nil
}
