package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yugarinn/github-issues-notificator/core"
	"github.com/yugarinn/github-issues-notificator/internal"
)


func createNotificationHandler(app *core.App, w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var notificationCreationRequest NotificationCreationRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&notificationCreationRequest)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if notificationCreationRequest.RepositoryUri == "" || notificationCreationRequest.Email == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	filters := internal.NotificationFilters{
		Author:		notificationCreationRequest.Filters.Author,
		Assignee:	notificationCreationRequest.Filters.Assignee,
		Label:		notificationCreationRequest.Filters.Label,
		Title:		notificationCreationRequest.Filters.Title,
	}

	input := internal.CreateNotificationInput{
		RepositoryUri:	notificationCreationRequest.RepositoryUri,
		Email: 			notificationCreationRequest.Email,
		Filters:		filters,

	}
	notificationCreationResult := internal.CreateNotification(app, input)

	if notificationCreationResult.Error != nil {
		http.Error(w, fmt.Sprintf("Error on notification creation: %s.", notificationCreationResult.Error), http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
}
