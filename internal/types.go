package internal

import (
	"time"
)


type ListenersWrapper struct {
    Listeners []Listener `yaml:"listeners"`
}

type Listener struct {
    Name          string            `yaml:"name"`
    EmailTo       string            `yaml:"email_to"`
    Repository    string            `yaml:"repository"`
    IsActive      bool              `yaml:"is_active"`
    Filters       ListenerFilters `yaml:"filters"`
}

type ListenerFilters struct {
	Labels string   `yaml:"labels"`
	Assignee string `yaml:"assignee"`
	Author string   `yaml:"author"`
}

type ListenerTracker struct {
	LastCheckAt   time.Time
	TrackedIssues []int64
}

type GithubIssue struct {
	ID              int64 `json:"id"`
	Title 			string `json:"title"`
	RepositoryUrl 	string `json:"repository_url"`
	Url 			string `json:"html_url"`
}
