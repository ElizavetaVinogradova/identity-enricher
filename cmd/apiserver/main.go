package main

import (
	"errors"
	"fmt"
	"identity-enricher/cmd"
	"identity-enricher/internal/apiserver"
	"identity-enricher/internal/client/enrichmentclient"
	"identity-enricher/internal/repo/postgres"
	"identity-enricher/internal/service"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.InitConfig(); err != nil {
		panic(fmt.Sprintf("Couldnt initialize the config file: %s", err))
	}
	cmd.SetupLogging()
	
	repository, err := postgres.NewFioRepository(cmd.BuildPostgreSqlConfig())
	if err != nil {
		panic(fmt.Sprintf("Couldnt create Repository: %s", err))
	}
	defer repository.Close()

	ageClient := enrichmentclient.NewAgeEnrichmentClient(cmd.GetAgeURl())
	genderClient := enrichmentclient.NewGenderEnrichmentClient(cmd.GetGenderURl())
	nationalityClient := enrichmentclient.NewNationalityEnrichmentClient(cmd.GetNationalityURl())

	tickService := service.NewFioService(ageClient, genderClient, nationalityClient, repository)
	
	server := apiserver.NewApiServer(cmd.BuildApiServerConfig(), *tickService)
	err = server.Start()

	if errors.Is(err, http.ErrServerClosed) {
		log.Error("server closed")
	} else if err != nil {
		log.Errorf("error starting server: %s", err)
		os.Exit(1)
	}
}