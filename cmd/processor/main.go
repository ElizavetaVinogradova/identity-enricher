package main

import (
	"fmt"
	"identity-enricher/cmd"
	"identity-enricher/internal/broker/kafka"
	"identity-enricher/internal/client/enrichmentclient"
	"identity-enricher/internal/repo/postgres"
	"identity-enricher/internal/service"

	kafka_go "github.com/segmentio/kafka-go"

	_ "github.com/lib/pq"
)

func main() {
	if err := cmd.InitConfig(); err != nil {
		panic(fmt.Sprintf("Couldnt initialize the config file: %s", err))
	}
	cmd.SetupLogging()

	ageClient := enrichmentclient.NewAgeEnrichmentClient(cmd.GetAgeURl())
	genderClient := enrichmentclient.NewGenderEnrichmentClient(cmd.GetGenderURl())
	nationalityClient := enrichmentclient.NewNationalityEnrichmentClient(cmd.GetNationalityURl())
	repository, error := postgres.NewFioRepository(cmd.BuildPostgreSqlConfig())

	if error != nil {
		panic(fmt.Sprintf("Couldnt create Repository: %s", error))
	}
	fioService := service.NewFioService(ageClient, genderClient, nationalityClient, repository)

	reader := kafka_go.NewReader(cmd.BuildKafkaReaderConfig())
	writer := kafka_go.NewWriter(cmd.BuildKafkaWriterConfig())

	kafka.NewBrokerProcessor(reader, writer, *fioService).Read()
	// defer broker.Close()
}
