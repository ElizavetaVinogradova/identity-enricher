package main

import (
	"fmt"
	"identity-enricher/cmd"

	"github.com/spf13/viper"
)

func main() {
	if err := cmd.InitConfig(); err != nil {
		panic(fmt.Sprintf("Couldn't initialize the config file: %s", err))
	}
	cmd.SetupLogging()

	//создаем и закрываем брокера
	//broker := kafka.NewBrokerProcessor([]string{"localhost:9093"}, "fio", "fio_failed")
	//defer broker.Read()

	//создаем и закрываем репу

	viper.SetDefault("service.batchSize", 1)
	//создаем сервис и вызываем на нем процессинг
}
