package handlers

import (
	er "Board_of_issuses/internal/core"
	dto "Board_of_issuses/internal/features/transport"
	respond "Board_of_issuses/internal/features/transport/http"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func (h *UserHandler) HandleCreateDesk(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r)

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "read request body error")
		return
	}

	desk := &dto.Desk{}

	if err := json.Unmarshal(reqBody, desk); err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "parse request data error")
		return
	}

	if len(desk.Name) < 3 {
		respond.RespondWithError(w, http.StatusBadRequest, "length name can not be less than 3 ")
		return
	}

	desk.OwnerId = userID
	desk.Created_at = time.Now()

	if err = h.serv.CreateDesk(r.Context(), desk.ToServiceDeskr()); err != nil {
		switch {

		case er.IsError(err, "PASSWORD_IS_SHORT"), er.IsError(err, "PASSWORD_IS_lONG"):
			var appErr *er.ErrorApp = err.(*er.ErrorApp)
			respond.RespondWithError(w, http.StatusBadRequest, appErr.Message)

		default:
			respond.RespondWithError(w, http.StatusInternalServerError, "error to create desk")
		}

		return
	}

	respond.RespondWithJSON(w, http.StatusCreated, "desk have been created")

}

func (h *UserHandler) HandleChangeDeskName(w http.ResponseWriter, r *http.Request) {
	userId := getUserIDFromContext(r)

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "read request body error")
		return
	}

	newName := &dto.UpdateDeskNameRequest{}

	if err := json.Unmarshal(reqBody, newName); err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "parse request data error")
		return
	}

	if newName.DeskID <= 0 {
		respond.RespondWithError(w, http.StatusBadRequest, "desk_id can not be less than or equal to zero")
		return
	}

	if len(newName.Name) < 3 {
		respond.RespondWithError(w, http.StatusBadRequest, "length name can not be less than 3 ")
		return
	}

	if err := h.serv.ChangeDeskName(r.Context(), newName.Name, newName.DeskID, userId); err != nil {
		switch {

		case er.IsError(err, "USER_IS_NOT_OWNER"):
			var appErr *er.ErrorApp = err.(*er.ErrorApp)
			respond.RespondWithError(w, http.StatusForbidden, appErr.Message)

		default:
			respond.RespondWithError(w, http.StatusInternalServerError, "error to change desk name")
		}

		return
	}

	respond.RespondWithJSON(w, http.StatusOK, "desk name had change")

}

func (h *UserHandler) HandleChangeDeskPassword(w http.ResponseWriter, r *http.Request) {
	userId := getUserIDFromContext(r)

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "read request body error")
		return
	}

	newPassword := &dto.UpdateDeskPasswordRequest{}

	if err := json.Unmarshal(reqBody, newPassword); err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "parse request data error")
		return
	}

	if newPassword.DeskID <= 0 {
		respond.RespondWithError(w, http.StatusBadRequest, "desk_id can not be less than or equal to zero")
		return
	}

	if err := h.serv.ChangeDeskPassword(r.Context(), newPassword.Password, newPassword.DeskID, userId); err != nil {
		switch {

		case er.IsError(err, "USER_IS_NOT_OWNER"):
			var appErr *er.ErrorApp = err.(*er.ErrorApp)
			respond.RespondWithError(w, http.StatusForbidden, appErr.Message)

		case er.IsError(err, "PASSWORD_IS_SHORT"), er.IsError(err, "PASSWORD_IS_lONG"):
			var appErr *er.ErrorApp = err.(*er.ErrorApp)
			respond.RespondWithError(w, http.StatusBadRequest, appErr.Message)

		default:
			respond.RespondWithError(w, http.StatusInternalServerError, "error to change desk name")
		}

		return
	}

	respond.RespondWithJSON(w, http.StatusOK, "desk password had change")

}

func (h *UserHandler) HandleChangeDeskOwner(w http.ResponseWriter, r *http.Request) {
	userId := getUserIDFromContext(r)

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "read request body error")
		return
	}

	newOwner := &dto.UpdateDeskOwnerRequest{}

	if err := json.Unmarshal(reqBody, newOwner); err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "parse request data error")
		return
	}

	if newOwner.DeskID <= 0 {
		respond.RespondWithError(w, http.StatusBadRequest, "desk_id can not be less than or equal to zero")
		return
	}

	if newOwner.ID <= 0 {
		respond.RespondWithError(w, http.StatusBadRequest, "owner id can not be less than or equal to zero")
		return
	}

	if err := h.serv.ChangeDeskOwner(r.Context(), newOwner.DeskID, userId, newOwner.ID); err != nil {
		switch {

		case er.IsError(err, "USER_IS_NOT_OWNER"):
			var appErr *er.ErrorApp = err.(*er.ErrorApp)
			respond.RespondWithError(w, http.StatusForbidden, appErr.Message)
		default:
			respond.RespondWithError(w, http.StatusInternalServerError, "error to change desk owner")
		}

		return
	}

	respond.RespondWithJSON(w, http.StatusOK, "desk owner had change")

}

func (h *UserHandler) HandleDeleteDesk(w http.ResponseWriter, r *http.Request) {
	userId := getUserIDFromContext(r)

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "read request body error")
		return
	}

	deskID := &dto.DeleteDeskRequest{}

	if err := json.Unmarshal(reqBody, deskID); err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "parse request data error")
		return
	}

	if deskID.ID <= 0 {
		respond.RespondWithError(w, http.StatusBadRequest, "desk_id can not be less than or equal to zero")
		return
	}

	if err := h.serv.DeleteDesk(r.Context(), deskID.ID, userId); err != nil {
		switch {

		case er.IsError(err, "USER_IS_NOT_OWNER"):
			var appErr *er.ErrorApp = err.(*er.ErrorApp)
			respond.RespondWithError(w, http.StatusForbidden, appErr.Message)

		default:
			respond.RespondWithError(w, http.StatusInternalServerError, "error to delete desk")
		}

		return
	}

	respond.RespondWithJSON(w, http.StatusOK, "desk had deleted")

}

func (h *UserHandler) HandleGetAllDesksId(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r)

	desksID, err := h.serv.GetAllDesks(r.Context(), userID)
	if err != nil {
		respond.RespondWithError(w, http.StatusInternalServerError, "error to get desks id")
	}

	respond.RespondWithJSON(w, http.StatusOK, desksID)
}
