package internal

import (
	"encoding/json"
	"slices"

	"go.etcd.io/bbolt"
	bolt "go.etcd.io/bbolt"
)


func filterNotTrackedIssuesFor(db *bbolt.DB, issues []GithubIssue, listener Listener) []GithubIssue {
	var filteredIssues []GithubIssue

	var listenerTracker ListenerTracker
	listenerTrackerExists := false

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BOLT_BUCKET_NAME))
		if b == nil {
			return nil
		}

		storedListenerTracker := b.Get([]byte(listener.Name))
		if storedListenerTracker == nil {
			return nil
		}

		if err := json.Unmarshal(storedListenerTracker, &listenerTracker); err != nil {
			return nil
		}

		listenerTrackerExists = true

		return nil
	})

	if !listenerTrackerExists {
		return issues
	}

	for _, issue := range issues {
		if !slices.Contains(listenerTracker.TrackedIssues, issue.ID) {
			filteredIssues = append(filteredIssues, issue)
		}
	}

	return filteredIssues
}
