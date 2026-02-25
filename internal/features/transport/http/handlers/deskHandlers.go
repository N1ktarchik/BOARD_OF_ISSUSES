package handlers

import (
	er "Board_of_issuses/internal/core"
	dto "Board_of_issuses/internal/features/transport"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func (h *UserHandler) HandleCreateDesk(w http.ResponseWriter, r *http.Request) {
	userID := getUserIDFromContext(r)

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	desk := &dto.Desk{}

	if err := json.Unmarshal(reqBody, desk); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if desk.Name == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	desk.OwnerId = userID
	desk.Created_at = time.Now()

	err = h.serv.CreateDesk(r.Context(), desk.ToServiceDeskr())
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
