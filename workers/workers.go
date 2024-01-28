package workers

import (
	"fmt"
	"log"
	"os"

	"github.com/hibiken/asynq"
)

func StartServe() {
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%d", os.Getenv("APP_REDIS_HOST"), os.Getenv("APP_REDIS_PORT"))},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeEventReminder, HandleEventTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
