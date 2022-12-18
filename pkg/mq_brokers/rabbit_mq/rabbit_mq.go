package rabbit_mq

import (
	"github.com/streadway/amqp"
	"log"
	"main/config"
	"os"
)

func NewRabbitMqClient(cfg *config.Config) (rmq *amqp.Connection, err error) {
	log.Printf("dialing %q", os.Getenv("MQ_RABBIT_URI"))
	rmq, err = amqp.Dial(os.Getenv("MQ_RABBIT_URI"))
	if err != nil {
		panic(err)
	}
	log.Println("connect to RMQ is successfully")

	return rmq, err
}
