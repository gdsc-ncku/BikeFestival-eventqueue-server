package workers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
)

// Task payload for any event notification related tasks.
type eventNotificationPayload struct {
	UserID  string
	EventID string
}

const (
	TypeEventReminder = "reminder"
)

func HandleEventTask(ctx context.Context, t *asynq.Task) error {
	var p eventNotificationPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}

	// Fixme : call line message unicast or broadcasting API
	log.Printf(" [*] Send %s Event Notification to User %s", p.EventID, p.UserID)
	return nil
}
