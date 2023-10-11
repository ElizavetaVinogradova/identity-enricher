package main

import (
	"fmt"
	"identity-enricher/cmd"
	"identity-enricher/internal/broker/kafka"

	"github.com/spf13/viper"
)

func main() {
	if err := cmd.InitConfig(); err != nil {
		panic(fmt.Sprintf("Couldn't initialize the config file: %s", err))
	}
	cmd.SetupLogging()

	broker := kafka.NewBrokerReader([]string{"localhost:9093"}, "fio")
	defer broker.Close()

	//создаем и закрываем репу

	viper.SetDefault("service.batchSize", 1)
	//создаем сервис и вызываем на нем процессинг
}
