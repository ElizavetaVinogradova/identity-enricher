package kafka

import (
	"context"
	"fmt"
	"identity-enricher/internal/service"

	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type BrokerProcessor struct {
	reader  *kafka.Reader
	writer  *kafka.Writer
	service service.FioService
}

func NewBrokerProcessor(reader *kafka.Reader, writer *kafka.Writer, service service.FioService) *BrokerProcessor {
	return &BrokerProcessor{reader: reader, writer: writer, service: service}
}

func (b *BrokerProcessor) Read() {
	for {
		message, err := b.reader.ReadMessage(context.Background())
		if err != nil {
			log.Errorf("Couldn't read fio from Kafka: %s", err)
		}


		fioDTO, err := UnmarshalMessage(message.Value)
		if err != nil {
			log.Errorf("Couldn't Unmarshal fio from Kafka: %s", err)
		}

		fio := mapKafkaDTOToServiceFio(fioDTO)
		err = b.service.StoreFio(fio)
		if err != nil {
			fioDTO.Error = err.Error()
			b.writeCorruptedFioToKafka(fioDTO)
		}

	}
}

func (b *BrokerProcessor) writeCorruptedFioToKafka(fioDTO FioKafkaDTO) error {
	byteValue, err := MarshalMessage(fioDTO)
	if err != nil {
		return err
	}
	log.Debugf("marshaled corrupted message for Kafka: %s", byteValue)

	message := kafka.Message{
		Value: byteValue,
	}

	err = b.writer.WriteMessages(context.Background(), message)
	if err != nil {
		return fmt.Errorf("write corrupted message to Kafka: %w", err)
	}
	return nil
}

func mapKafkaDTOToServiceFio(dto FioKafkaDTO) service.Fio {
	var fio service.Fio
	fio.Name = dto.Name
	fio.Surname = dto.Surname
	fio.Patronymic = dto.Patronymic
	fio.Error = dto.Error
	return fio
}
