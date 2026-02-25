package handlers

import (
	"log"
	"net/http"
)

func getUserIDFromContext(r *http.Request) int {
	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		log.Panicf("middleware context error")
	}

	return userID
}

func HandleBase(w http.ResponseWriter, r *http.Request) {
	///главная страница с ридми
}
