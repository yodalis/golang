package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch, nil
}

// Consumindo mensagens através dos canais do Go
// Argumentos: Conexão com o Channel do RabbitMQ e mensagens recebidas pelo RabbitMQ armazenadas pelo channel
func Consume(ch *amqp.Channel, out chan amqp.Delivery, queue string) error {
	msgs, err := ch.Consume(
		queue,
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	// Qual fila queremos ler
	// Nome da aplicação que vai ta consumindo
	// Auto Hec - Recebe a mensagem e dar um baixa que a mensagem foi lida
	// Fila exclusiva = false
	// Local = false
	// Waiting = false
	// Table = nil
	)

	if err != nil {
		return err
	}

	// Consumindo mensagens recebidas

	for msg := range msgs {
		out <- msg // Mensagem recebida sendo armazenada no canal
	}

	return nil
}

func Publish(ch *amqp.Channel, body string, exName string) error {
	err := ch.Publish(
		exName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)
	if err != nil {
		return err
	}
	return nil
}
