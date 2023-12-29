package internal

import (
	"time"
)


type Notification struct {
	ID             string `bson:"_id"`
	RepositoryUri  string `bson:"repositoryUri"`
	Filters        NotificationFilters
	Email          string `bson:"email"`
	NotifiedIssues []int64 `bson:"notifiedIssues"`
	LastCheckAt    time.Time `bson:"lastCheckAt"`
}

type NotificationFilters struct {
	Author		string
	Assignee	string
	Label 		string
	Title 		string
}

type GithubIssue struct {
	ID              int64 `json:"id"`
	Title 			string `json:"title"`
	RepositoryUrl 	string `json:"repository_url"`
	Url 			string `json:"html_url"`
}
