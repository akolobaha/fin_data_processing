package service

import (
	"encoding/json"
	"fin_data_processing/internal/entities"
	"fin_data_processing/internal/transport"
)

func SendNotificationMessage(target entities.TargetUser, rabbit *transport.Rabbitmq, queueName string) {
	data, err := json.Marshal(target)
	if err != nil {
		return
	}
	rabbit.DeclareQueue(queueName)

	rabbit.SendMsg(data)
}
