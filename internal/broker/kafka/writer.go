package kafka

import (
	"context"
	"fmt"
	"identity-enricher/internal/service"

	log "github.com/sirupsen/logrus"

	kafka "github.com/segmentio/kafka-go"
)

type BrokerWriter struct {
	writer *kafka.Writer
}

func NewBrokerWriter(brokerAddresses []string, topic string) *BrokerWriter {
	config := kafka.WriterConfig{
		Brokers: brokerAddresses,
		Topic:   topic,
	}

	writer := kafka.NewWriter(config)
	return &BrokerWriter{writer: writer}
}

func (b *BrokerWriter) Close() {
	b.writer.Close()
}

func (b *BrokerWriter) Write(fio service.Fio) error {
	//todo В случае некорректного сообщения, обогатить его причиной ошибки (нет обязательного поля, некорректный формат...) и отправить в очередь кафки FIO_FAILED
	byteKey := []byte("fio_failed")
	byteValue, err := MarshalMessage(mapServiceFioToKafkaDTO(fio))
	if err != nil {
		return err
	}
	log.Debugf("marshaled message for Kafka: %s", byteValue)

	message := kafka.Message{
		Key:   byteKey,
		Value: byteValue,
	}

	err = b.writer.WriteMessages(context.Background(), message)
	if err != nil {
		return fmt.Errorf("write message to Kafka: %w", err)
	}
	return nil
}

func mapServiceFioToKafkaDTO(fio service.Fio) FioKafkaDTO {
	var kafkaFio FioKafkaDTO
	kafkaFio.Name = fio.Name
	kafkaFio.Surname = fio.Surname
	kafkaFio.Patronymic = fio.Patronymic

	log.Debugf("Service Fio mapped to kafkaDTO: %s", fmt.Sprintf("%v", kafkaFio))
	return kafkaFio
}
