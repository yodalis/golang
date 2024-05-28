package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/yodalis/fcutils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgs, "minhafila")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false) //Indicando que a mensagem ja foi lida e não precisa ser inserida na fila de novo
	}
}
