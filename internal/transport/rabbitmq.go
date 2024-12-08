package transport

import (
	"fin_data_processing/internal/config"
	"github.com/streadway/amqp"
	"log/slog"
)

type Rabbitmq struct {
	Chan  *amqp.Channel
	Queue amqp.Queue
}

func New() *Rabbitmq {
	return &Rabbitmq{}
}

func (rabbit *Rabbitmq) InitConn(cfg *config.Config) {
	// Установите соединение с RabbitMQ
	conn, err := amqp.Dial(cfg.GetRabbitDSN())
	if err != nil {
		slog.Error("Failed to connect to RabbitMQ: ", "error", err)
	}

	ch, err := conn.Channel()
	rabbit.Chan = ch

	if err != nil {
		slog.Error("Failed to open a channel: %s", "error", err)
	}
}

func (rabbit *Rabbitmq) ConnClose() {
	rabbit.Chan.Close()
}

func (rabbit *Rabbitmq) DeclareQueue(name string) {
	// Объявите очередь, в которую будете отправлять сообщения
	queue, err := rabbit.Chan.QueueDeclare(
		name,  // имя очереди
		true,  // durable постоянная очередь
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // аргументы
	)
	rabbit.Queue = queue
	if err != nil {
		slog.Error("Failed to declare a queue: ", "", err)
	}
}

func (rabbit *Rabbitmq) SendMsg(data []byte) {
	err := rabbit.Chan.Publish(
		"",                // обменник
		rabbit.Queue.Name, // ключ маршрутизации (имя очереди)
		false,             // обязательное
		false,             // немедленное
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, // сохранять сообщение
			ContentType:  "text/plain",
			Body:         data,
		})
	if err != nil {
		slog.Error("Failed to publish a message: %s", "error", err)
	}
}
