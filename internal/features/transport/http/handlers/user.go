package handlers

import (
	er "Board_of_issuses/internal/core"
	dto "Board_of_issuses/internal/features/transport"
	respond "Board_of_issuses/internal/features/transport/http"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func (h *UserHandler) HandleCreateUser(w http.ResponseWriter, r *http.Request) {

	var user dto.User

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "read request body error")
		return
	}

	if err := json.Unmarshal(httpRequestBody, &user); err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "parse request data error")
		return
	}

	if user.Login == "" || user.Password == "" || user.Email == "" || user.Name == "" {
		respond.RespondWithError(w, http.StatusBadRequest, "bad user data")
		return
	}

	user.Created_at = time.Now()

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	token, err := h.serv.Registration(ctx, user.ToServiceUser())
	if err != nil {

		switch {

		case er.IsError(err, "USER_HAVE_REGISTER"):
			var appErr *er.ErrorApp = err.(*er.ErrorApp)
			respond.RespondWithError(w, http.StatusConflict, appErr.Message)

		case er.IsError(err, "PASSWORD_IS_SHORT"), er.IsError(err, "PASSWORD_IS_lONG"):
			var appErr *er.ErrorApp = err.(*er.ErrorApp)
			respond.RespondWithError(w, http.StatusBadRequest, appErr.Message)

		default:
			respond.RespondWithError(w, http.StatusInternalServerError, "error to registrate user")
		}

		return
	}

	accesToken := &dto.UserResponse{
		AccessToken: token,
		TokenType:   "Bearer",
	}

	resp, _ := json.Marshal(accesToken)
	respond.RespondWithJSON(w, http.StatusCreated, resp)

}

func (h *UserHandler) HandleLoginUser(w http.ResponseWriter, r *http.Request) {
	var user dto.User

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "read request body error")
		return
	}

	if err := json.Unmarshal(httpRequestBody, &user); err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "parse request data error")
		return
	}

	if user.Password == "" || (user.Email == "" && user.Login == "") {
		respond.RespondWithError(w, http.StatusBadRequest, "bad user data")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	token, err := h.serv.Authorization(ctx, user.ToServiceUser())

	if err != nil {

		switch {

		case er.IsError(err, "INVALID_PASSWORD"):

			var appErr *er.ErrorApp = err.(*er.ErrorApp)
			respond.RespondWithError(w, http.StatusBadRequest, appErr.Message)

		default:

			respond.RespondWithError(w, http.StatusInternalServerError, "error to registrate user")
		}

		return
	}

	accesToken := &dto.UserResponse{
		AccessToken: token,
		TokenType:   "Bearer",
	}

	resp, _ := json.Marshal(accesToken)

	respond.RespondWithJSON(w, http.StatusOK, resp)

}

func (h *UserHandler) HandleChangeUserName(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r)

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "read request body error")
		return
	}

	newName := &dto.UpdateNameRequest{}

	if err := json.Unmarshal(reqBody, newName); err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "parse request data error")
		return
	}

	if err := h.serv.ChangeUserName(r.Context(), newName.Name, userID); err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "change user name error")
		return
	}

	respond.RespondWithJSON(w, http.StatusOK, "user name had change")

}

func (h *UserHandler) HandleChangeUserEmail(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r)

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "read request body error")
		return
	}

	newEmail := &dto.UpdateEmailRequest{}

	if err := json.Unmarshal(reqBody, newEmail); err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "parse request data error")
		return
	}

	if err := h.serv.ChangeUserEmail(r.Context(), newEmail.Email, userID); err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "change user email error")
	}

	respond.RespondWithJSON(w, http.StatusOK, "user email had change")

}

func (h *UserHandler) HandleChangeUserPassword(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r)

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "read request body error")
		return
	}

	newPassword := &dto.UpdatePasswordRequest{}

	if err := json.Unmarshal(reqBody, newPassword); err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "parse request data error")
		return
	}

	err = h.serv.ChangeUserEmail(r.Context(), newPassword.Password, userID)
	if err != nil {

		switch {

		case er.IsError(err, "PASSWORD_IS_SHORT"), er.IsError(err, "PASSWORD_IS_lONG"):

			var appErr *er.ErrorApp = err.(*er.ErrorApp)
			respond.RespondWithError(w, http.StatusBadRequest, appErr.Message)

		default:
			respond.RespondWithError(w, http.StatusInternalServerError, "change user password error")

		}

		return
	}

	respond.RespondWithJSON(w, http.StatusOK, "user password had change")

}

func (h *UserHandler) HandleConnectUserToDesk(w http.ResponseWriter, r *http.Request) {
	userId := getUserIDFromContext(r)

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "read request body error")
		return
	}

	desk := &dto.ConnectUserToDeskRequest{}

	if err := json.Unmarshal(reqBody, desk); err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "parse request data error")
		return
	}

	if desk.ID <= 0 {
		respond.RespondWithError(w, http.StatusBadRequest, "desk_id can not be less than or equal to zero")
		return
	}

	if err := h.serv.ConnectUserToDesk(r.Context(), userId, desk.ID, desk.Password); err != nil {
		switch {

		case er.IsError(err, "PASSWORD_IS_SHORT"), er.IsError(err, "PASSWORD_IS_lONG"):
			var appErr *er.ErrorApp = err.(*er.ErrorApp)
			respond.RespondWithError(w, http.StatusBadRequest, appErr.Message)

		default:
			respond.RespondWithError(w, http.StatusInternalServerError, "error connect to desk ")
		}

		return
	}

	respond.RespondWithJSON(w, http.StatusCreated, "you have connected to desk")

}
