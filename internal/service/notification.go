package service

import (
	"encoding/json"
	"fin_data_processing/internal/entities"
	"fin_data_processing/internal/log"
	"fin_data_processing/internal/transport"
)

func SendEmailNotificationMessage(target entities.TargetUser, rabbit *transport.Rabbitmq, queueName string) {
	data, err := json.Marshal(target)
	if err != nil {
		return
	}
	err = rabbit.DeclareQueue(queueName, true)
	if err != nil {
		log.Error("Ошибка оюъявлине очереди: ", err)
	}

	rabbit.SendMsg(data)
}
