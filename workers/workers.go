package workers

import (
	"log"

	"github.com/hibiken/asynq"
)

func StartServe() {
	// Fixme : Addr should be in the form of "redis-container-name:port"
	// and it should be read from a configuration file or env variable
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "localhost:6379"},
		asynq.Config{Concurrency: 10},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(TypeEventReminder, HandleEventTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal(err)
	}
}
