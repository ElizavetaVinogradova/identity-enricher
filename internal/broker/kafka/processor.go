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

func NewBrokerProcessor(brokerAddresses []string, readFromTopic string, writeToTopic string, service service.FioService) *BrokerProcessor {
	//todo конфиги и инициализацию ридера с райтером перенести в мейн, а тут в методе получать сразу готовых ридера и райтера
	readerConfig := kafka.ReaderConfig{
		Brokers: brokerAddresses,
		Topic:   readFromTopic,
	}
	reader := kafka.NewReader(readerConfig)

	writerConfig := kafka.WriterConfig{
		Brokers: brokerAddresses,
		Topic:   writeToTopic,
	}
	writer := kafka.NewWriter(writerConfig)

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
		log.Debugf("Fio from Kafka Unmarshalled: %s", fmt.Sprintf("%v", fioDTO))
		fio := mapKafkaDTOToServiceFio(fioDTO)
		error := b.service.Run(fio)
		if error != nil {
			fio.Error = err.Error()
			b.writeCorruptedFioToKafka(fioDTO)
		}

	}
}

func (b *BrokerProcessor) writeCorruptedFioToKafka(fioDTO FioKafkaDTO) error {
	byteValue, err := MarshalMessage(fioDTO)
	if err != nil {
		return err
	}
	log.Debugf("marshaled message for Kafka: %s", byteValue)

	message := kafka.Message{
		Value: byteValue,
	}

	err = b.writer.WriteMessages(context.Background(), message)
	if err != nil {
		return fmt.Errorf("write message to Kafka: %w", err)
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
