package handlers

import (
	er "Board_of_issuses/internal/core"
	dto "Board_of_issuses/internal/features/transport"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func (h *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user dto.User

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

	accesToken := &dto.UserResponse{
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

func (h *UserHandler) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user dto.User

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

	accesToken := &dto.UserResponse{
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

func (h *UserHandler) HandleChangeUserName(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r)

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newName := &dto.UpdateNameRequest{}

	if err := json.Unmarshal(reqBody, newName); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := h.serv.ChangeUserName(r.Context(), newName.Name, userID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h *UserHandler) HandleChangeUserEmail(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r)

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newEmail := &dto.UpdateEmailRequest{}

	if err := json.Unmarshal(reqBody, newEmail); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := h.serv.ChangeUserEmail(r.Context(), newEmail.Email, userID); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (h *UserHandler) HandleChangeUserPassword(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r)

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newPassword := &dto.UpdatePasswordRequest{}

	if err := json.Unmarshal(reqBody, newPassword); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.serv.ChangeUserEmail(r.Context(), newPassword.Password, userID)
	if err != nil {

		var responseErr []byte

		switch {

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

}

///connect to desk user
