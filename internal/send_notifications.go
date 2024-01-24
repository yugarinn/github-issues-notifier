package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"time"

	"github.com/yugarinn/github-issues-notifier/core"

	bolt "go.etcd.io/bbolt"
)


const BOLT_BUCKET_NAME = "listeners"

func SendNotifications(app *core.App, listeners []Listener) {
	databaseName := os.Getenv("LISTENERS_DATABASE_PATH")
	db, err := bolt.Open(databaseName, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for _, listener := range listeners {
		var listenerTracker ListenerTracker
		listenerTrackerExists := false

		err = db.View(func(tx *bolt.Tx) error {
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

		var lastCheckAt time.Time

		if listenerTrackerExists {
			lastCheckAt = listenerTracker.LastCheckAt
		} else {
			lastCheckAt = time.Now().Add(-24 * time.Hour)
		}

		issues, _ := retrieveIssuesFor(listener, lastCheckAt)
		notTrackedIssues := filterNotTrackedIssuesFor(db, issues, listener)

		for _, issue := range notTrackedIssues {
			if listenerTrackerExists && slices.Contains(listenerTracker.TrackedIssues, issue.ID) {
				continue
			}

			sendIssueEmailAlertTo(listener, issue)

			if !listenerTrackerExists {
				var trackedIssues []int64

				listenerTracker = ListenerTracker{
					LastCheckAt: time.Now(),
					TrackedIssues: trackedIssues,
				}
			}

			err = db.Update(func(tx *bolt.Tx) error {
				b, err := tx.CreateBucketIfNotExists([]byte(BOLT_BUCKET_NAME))
				if err != nil {
					return fmt.Errorf("create bucket: %s", err)
				}

				listenerTracker.LastCheckAt = time.Now()
				listenerTracker.TrackedIssues = append(listenerTracker.TrackedIssues, issue.ID)

				jsonListenerTracker, err := json.Marshal(listenerTracker)
				if err != nil {
					log.Fatal(err)
				}

				err = b.Put([]byte(listener.Name), jsonListenerTracker)
				if err != nil {
					return fmt.Errorf("write to bucket: %s", err)
				}

				return nil
			})

			if err != nil {
				log.Fatalf("db update error: %s", err)
			}
		}
	}
}
