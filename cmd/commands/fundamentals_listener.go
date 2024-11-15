package commands

import (
	"fin_data_processing/internal/config"
	"github.com/streadway/amqp"
	"log"
)

func RunFundamentalsListener(cfg *config.Config) {
	conn, err := amqp.Dial(config.RabbitDsn)
	if err != nil {
		log.Fatalf("Ошибка подключения: %s", err)
	}
	defer conn.Close()

	// Создание канала
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Ошибка создания канала: %s", err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		cfg.RabbitQueueFundamentals, // имя очереди
		"",                          // имя потребителя
		true,                        // авто-активировать
		false,                       // не эксклюзивная
		false,                       // не ожидать
		false,                       // не ждать
		nil,                         // дополнительные параметры
	)
	if err != nil {
		log.Fatalf("Ошибка при получении сообщений: %s", err)
	}

	// Обработка сообщений
	go func() {
		for d := range msgs {
			log.Printf("Получено сообщение: %s", d.Body)
		}
	}()

	// Блокировка главной горутины
	select {}

}
