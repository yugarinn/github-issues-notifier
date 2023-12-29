package http


type NotificationCreationRequest struct {
	RepositoryUri string 	`json:"repositoryUri"`
	Email         string 	`json:"email"`
	Filters       Filters	`json:"filters"`
}

type Filters struct {
	Author		string `json:"author"`
	Assignee	string `json:"assignee"`
	Label 		string `json:"label"`
	Title 		string `json:"title"`
}
