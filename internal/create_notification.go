package internal

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/yugarinn/github-issues-notificator/core"
)


type CreateNotificationInput struct {
	RepositoryUri string
	Filters       NotificationFilters
	Email         string
}

type CreateNotificationResult struct {
	Success      bool
	Error        error
}

func CreateNotification(app *core.App, input CreateNotificationInput) CreateNotificationResult {
	ctx := context.Background()

	if repositoryExists(input.RepositoryUri) == false {
		return CreateNotificationResult{
			Success: false,
			Error: errors.New("provided_repository_not_found"),
		}
	}

	now := time.Now()

	collection := app.Database.Collection("notifications")

	notification := bson.M{
		"repositoryUri":    input.RepositoryUri,
		"email":            input.Email,
		"filters":          input.Filters,
		"confirmationCode": generateNotificationConfirmationCode(&input),
		"isConfirmed":      false,
		"lastCheckAt":      nil,
		"createdAt":        now,
		"updatedAt":        now,
	}

	_, err := collection.InsertOne(ctx, notification)

	return CreateNotificationResult{
		Success: err != nil,
		Error: err,
	}
}

func repositoryExists(repositoryUri string) bool {
	url := fmt.Sprintf("https://github.com/%s", repositoryUri)
	response, err := http.Get(url)

	if err != nil || response.StatusCode != 200 {
		return false
	}

	return true
}

func generateNotificationConfirmationCode(input *CreateNotificationInput) string {
	bytes := make([]byte, 16)
	rand.Read(bytes)

	baseString := hex.EncodeToString(bytes)
	baseCode := fmt.Sprintf("%s-%s-%s", input.RepositoryUri, input.Email, baseString)

	hash := sha256.New()
	hash.Write([]byte(baseCode))
	hashedBytes := hash.Sum(nil)

	confirmationCode := hex.EncodeToString(hashedBytes)

	return confirmationCode
}
