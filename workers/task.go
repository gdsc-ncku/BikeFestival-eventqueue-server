package workers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/hibiken/asynq"
	"github.com/line/line-bot-sdk-go/linebot"
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

	bot, err := linebot.New(
		os.Getenv("LINEBOT_CLIENT_CHANNEL_SECRET"),
		os.Getenv("LINEBOT_CLIENT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	message := linebot.NewTextMessage(fmt.Sprintf("Hello, Event %s is going to start within 30 minutes!!!", p.EventID))

	_, err = bot.PushMessage(p.UserID, message).Do()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
