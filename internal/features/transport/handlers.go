package http

import (
	er "Board_of_issuses/internal/core"
	dn "Board_of_issuses/internal/core/domains"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type Service interface {
	Registration(ctx context.Context, user *dn.User) (string, error)
	Authorization(ctx context.Context, user *dn.User) (string, error)
}

type UserHandler struct {
	serv Service
}

func NewUserHandler(src Service) *UserHandler {
	return &UserHandler{
		serv: src,
	}
}

func HandleBase(w http.ResponseWriter, r *http.Request) {
	///главная страница с ридми
}

func (h *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user User

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(httpRequestBody, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user.Login == "" || user.Password == "" || user.Email == "" || user.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user.Created_at = time.Now()

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	token, err := h.serv.Registration(ctx, user.ToServiceUser())
	if err != nil {

		var responseErr []byte

		switch {
		case er.IsError(err, "USER_HAVE_REGISTER"):
			var appErr *er.ErrorApp = err.(*er.ErrorApp)
			responseErr, _ = json.Marshal(appErr.Message)
			w.WriteHeader(http.StatusConflict)

		case er.IsError(err, "PASSWORD_IS_SHORT"), er.IsError(err, "PASSWORD_IS_lONG"):
			var appErr *er.ErrorApp = err.(*er.ErrorApp)
			responseErr, _ = json.Marshal(appErr.Message)
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		if len(responseErr) != 0 {
			w.Write(responseErr)
		}

		return
	}

	accesToken := &UserResponse{
		AccessToken: token,
		TokenType:   "Bearer",
	}

	resp, err := json.Marshal(accesToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(resp)

}

func (h *UserHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(httpRequestBody, &user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user.Password == "" || (user.Email == "" && user.Login == "") {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	token, err := h.serv.Authorization(ctx, user.ToServiceUser())

	if err != nil {

		var responseErr []byte

		switch {
		case er.IsError(err, "INVALID_PASSWORD"):
			var appErr *er.ErrorApp = err.(*er.ErrorApp)
			responseErr, _ = json.Marshal(appErr.Message)
			w.WriteHeader(http.StatusBadRequest)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		if len(responseErr) != 0 {
			w.Write(responseErr)
		}

		return
	}

	accesToken := &UserResponse{
		AccessToken: token,
		TokenType:   "Bearer",
	}

	resp, err := json.Marshal(accesToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}
