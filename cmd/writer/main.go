package main

import (
	"fmt"
	"identity-enricher/cmd"

	"github.com/spf13/viper"
)

func main() {
	if err := cmd.InitConfig(); err != nil {
		panic(fmt.Sprintf("Couldnt initialize the config file: %s", err))
	}
	cmd.SetupLogging()

	//создаем и закрываем клиентов

	// broker := kafka.NewBrokerWriter([]string{"localhost:9093"}, "ticks")
	// defer broker.Close()

	viper.SetDefault("service.batchSize", 1)
	//создаем сервис и начинаем процессинг
}
