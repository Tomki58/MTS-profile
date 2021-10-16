package api

import (
	"MTS/profile/httpserver/serializer"
	sendinglist "MTS/profile/sending_list"
	"net/http"

	"github.com/go-chi/chi"
)

func (a *App) createSendingList(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "user_id")
	if userID != "" {
		a.listMap[userID] = sendinglist.New()
	}
	w.Write(serializer.SerializeResponseJSON("User's list created"))
}
