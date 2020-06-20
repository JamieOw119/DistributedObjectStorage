package heartbeat

import (
	"lib/rabbitmq"
	"os"
	"time"
)

func StartHeartbeat() {
	q := rabbitmq.New(os.Getenv("RABBITMQ_SERVER_API"))
	defer q.Close()
	for {
		q.Publish("apiServers", os.Getenv("LOCATE_ADDRESS"))
		time.Sleep(5 * time.Second)
	}
}
