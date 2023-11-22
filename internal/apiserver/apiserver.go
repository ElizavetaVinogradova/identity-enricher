package apiserver

import (
	"fmt"
	"identity-enricher/internal/service"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

type ApiServer struct {
	fioService service.FioService
	config      Config
}

type Config struct {
	BindAddress string
}

func NewApiServer(config Config, fioService service.FioService) *ApiServer {
	return &ApiServer{
		config:      config,
		fioService: fioService,
	}
}

func (s *ApiServer) Start() error {
	log.Info("Start server")
	mux := http.NewServeMux()
	mux.HandleFunc("/fio", s.getFioById)
	err := http.ListenAndServe(s.config.BindAddress, mux)
	return err
}

func (s *ApiServer) getFioById(w http.ResponseWriter, r *http.Request) {
	log.Info("got /fio request")

	id := r.URL.Query().Get("id")
	if id == "" {
		log.Errorf("Missing id parameter, %d", http.StatusBadRequest)
		return
	}

	idValue, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		log.Errorf("id parsing, %d", http.StatusInternalServerError)
		return
	}

	fio, err := s.fioService.GetFioById(idValue)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting fio by id: %s", err), http.StatusInternalServerError)
		return
	}
	fioAPI := mapFioToFioAPI(fio)
	jsonResponse, err := MarshalRequest(fioAPI)
	if err != nil {
		http.Error(w, "Error marshalling", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

//todo what about map FioAPI to Fio???
func mapFioToFioAPI(fio service.Fio) FioAPI {
	var fioAPI FioAPI
	fioAPI.Id = fio.Id
	fioAPI.Name = fio.Name
	fioAPI.Surname = fio.Surname
	fioAPI.Patronymic = fio.Patronymic
	fioAPI.Age = fio.Age
	fioAPI.Gender = fio.Gender
	fioAPI.Nationality = fio.Nationality
	return fioAPI
}

