package kafka

import (
	"encoding/json"
	"fmt"
)

type FioKafkaDTO struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Error      string `json:"error"`
}

func MarshalMessage(fio FioKafkaDTO) ([]byte, error) {
	message, err := json.Marshal(&fio)
	if err != nil {
		return nil, fmt.Errorf("marshall message to Kafka: %w", err)
	}
	return message, nil
}

func UnmarshalMessage(message []byte) (FioKafkaDTO, error) {
	var fio FioKafkaDTO
	err := json.Unmarshal(message, &fio)
	if err != nil {
		return FioKafkaDTO{}, fmt.Errorf("unmarshall message to Kafka: %w", err)
	}
	return fio, nil
}
