package kafka

import (
	"context"
	"fmt"
	"identity-enricher/internal/service"

	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type BrokerReader struct {
	reader *kafka.Reader
}

func NewBrokerReader(brokerAddresses []string, topic string) *BrokerReader {
	config := kafka.ReaderConfig{
		Brokers: brokerAddresses,
		Topic:   topic,
	}

	reader := kafka.NewReader(config)

	return &BrokerReader{reader: reader}
}

func (b *BrokerReader) Close() {
	b.reader.Close()
}

func (b *BrokerReader) Read() (service.Fio, error) {
	//todo В случае некорректного сообщения, обогатить его причиной ошибки (нет обязательного поля, некорректный формат...) и отправить в очередь кафки FIO_FAILED
	message, err := b.reader.ReadMessage(context.Background())
	if err != nil {
		return service.Fio{}, fmt.Errorf("read message from Kafka: %w", err)
	}

	fioDTO, err := UnmarshalMessage(message.Value)
	if err != nil {
		return service.Fio{}, err
	}
	log.Debugf("Fio from Kafka Unmarshalled: %s", fmt.Sprintf("%v", fioDTO))

	return mapKafkaDTOToServiceFio(fioDTO), nil
}

func mapKafkaDTOToServiceFio(dto FioKafkaDTO) service.Fio {
	var fio service.Fio
	fio.Name = dto.Name
	fio.Surname = dto.Surname
	fio.Patronymic = dto.Patronymic
	fio.Error = dto.Error
	return fio
}
