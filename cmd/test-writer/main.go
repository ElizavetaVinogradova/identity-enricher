package main

import (
	"context"
	"fmt"
	"identity-enricher/cmd"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

const kafkaTopic = "FIO"

var (
	russianNames      = []string{"Dmitriy", "Ivan", "Alexander", "Ekaterina", "Olga", "Mikhail", "Anna", "Sergei"}
	russianSurnames   = []string{"Ivanov", "Petrov", "Smirnov", "Kozlov", "Vasiliev", "Fedorov", "Morozov", "Kovalenko"}
	russianPatronymic = []string{"Ivanovich", "Petrovich", "Sergeevich", "Alexandrovich", "Mikhailovich", "Dmitrievna", "Andreevna"}

	invalidDataRatio = 0.3
)

func main() {
	fmt.Print("==========MAIN========== \n")

	cmd.SetupLogging()
	rand.Seed(time.Now().UnixNano())

	broker := "localhost:9093"
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{broker},
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	})

	defer writer.Close()

	for {
		data := generateRandomData()
		fmt.Print("==========Generate message for Kafka: ", data, "\n")
		message := kafka.Message{
			Key:   nil,
			Value: []byte(data),
		}
		err := writer.WriteMessages(context.Background(), message)
		fmt.Print("==========Sent message to Kafka: ", message, "\n")
		if err != nil {
			fmt.Printf("Error sending message to Kafka: %v\n", err)
		}
	}
}

func generateRandomData() string {
	var data string

	if rand.Float64() < invalidDataRatio {
		// Генерация невалидных данных (например, отсутствие имени или фамилии)
		data = `{"name": "!!", "surname": "   ", "patronymic": "5432"}`
	} else {
		// Генерация валидных данных
		name := russianNames[rand.Intn(len(russianNames))]
		surname := russianSurnames[rand.Intn(len(russianSurnames))]
		patronymic := russianPatronymic[rand.Intn(len(russianPatronymic))]
		data = fmt.Sprintf(`{"name": "%s", "surname": "%s", "patronymic": "%s"}`, name, surname, patronymic)
	}

	return data
}
