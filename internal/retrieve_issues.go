package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func retrieveIssuesFor(listener Listener, lastCheckAt time.Time) ([]GithubIssue, error) {
	var client = &http.Client{Timeout: 10 * time.Second}
	url := buildIssuesUrlFor(listener, lastCheckAt)
	resp, err := client.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Non-OK HTTP status: %s\n", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var issues []GithubIssue
	if err := json.Unmarshal(body, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}

func buildIssuesUrlFor(listener Listener, lastCheckAt time.Time) string {
	labels := listener.Filters.Labels
	issuesUrl := fmt.Sprintf("https://api.github.com/repos/%s/issues?labels=%s", listener.Repository, url.QueryEscape(labels))
	issuesUrl += "&since=" + url.QueryEscape(lastCheckAt.Format(time.RFC3339))

	return issuesUrl
}
