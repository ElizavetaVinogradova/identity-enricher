package service

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

type FioBrokerReader interface {
	Read() (Fio, error)
}

type FioService struct {
	fioBrokerReader FioBrokerReader
	fioRepository   FioRepository
	batchSize       int
}

func NewFioService(fioBrokerReader FioBrokerReader, fioRepository FioRepository, batchSize int) *FioService {
	return &FioService{
		fioBrokerReader: fioBrokerReader,
		fioRepository:   fioRepository,
		batchSize:       batchSize,
	}
}

func (s *FioService) Run(fio Fio) error {
	if !isFioValid(fio){
		return errors.New("received FIO is invalid")
	}



	//где-то между этими операциями тут:
	//1. валидация полученного сообщения
	//2. отправка обратно в кафку в случае ошибки
	//3. обогащение инфой от клиентов

	err := s.fioRepository.Create(fio)
	if err != nil {
		log.Errorf("Couldn't create fio: %s", err)
	}
return nil
}

func isFioValid(fio Fio) bool {
	if (fio.Name == "" || fio.Surname=="") {
		//todo Какие еще условия для проверки?
		return false
	}
	return true
}
