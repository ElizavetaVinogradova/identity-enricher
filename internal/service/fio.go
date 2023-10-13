package service

import (
	"errors"

	log "github.com/sirupsen/logrus"
)

type FioService struct {
	ageClient         AgeClient
	genderClient      GenderClient
	nationalityClient NationalityClient
	fioRepository     FioRepository
}

func NewFioService(ageClient AgeClient, genderClient GenderClient, nationalityClient NationalityClient, fioRepository FioRepository) *FioService {
	return &FioService{
		ageClient:         ageClient,
		genderClient:      genderClient,
		nationalityClient: nationalityClient, 
		fioRepository:     fioRepository,
	}
}

func (s *FioService) StoreFio(fio Fio) error {
	if !isFioValid(fio) { //1. валидация полученного сообщения
		return errors.New("received FIO is invalid") //2. отправка обратно в кафку в случае ошибки
	}

	//3. обогащение инфой от клиентов

	err := s.fioRepository.Create(fio)
	if err != nil {
		log.Errorf("Couldn't create fio: %s", err)
	}
	return nil
}

func isFioValid(fio Fio) bool {
	if fio.Name == "" || fio.Surname == "" {
		//todo Какие еще условия для проверки?
		return false
	}
	return true
}

func (s *FioService) enrichWithAge(fio *Fio) error {
	age, err := s.ageClient.FetchAge()//валидировать пришедшие данные?
	if err != nil {
		return err 
	}
	fio.Age = age
	return nil
}

func (s *FioService) enrichWithGender(fio *Fio) error {
	gender, err := s.genderClient.FetchGender()//валидировать пришедшие данные?
	if err != nil {
		return err 
	}
	fio.Gender = gender
	return nil
}

func (s *FioService) enrichWithNationality(fio *Fio) error {
	nationality, err := s.nationalityClient.FetchNationality() //валидировать пришедшие данные?
	if err != nil {
		return err
	}
	fio.Nationality = nationality
	return nil
}
