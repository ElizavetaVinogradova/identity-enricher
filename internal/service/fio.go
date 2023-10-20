package service

import (
	"fmt"
	"regexp"

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
	err := isFioValid(fio)
	if err != nil {
		return err
	}

	s.enrichWithAge(&fio)
	s.enrichWithGender(&fio)
	s.enrichWithNationality(&fio)

	err = s.fioRepository.Create(fio)
	if err != nil {
		log.Errorf("Couldn't create fio: %s", err)
	}
	return nil
}

func (s *FioService) enrichWithAge(fio *Fio) error {
	age, err := s.ageClient.FetchAge()
	if err != nil {
		return err
	}
	fio.Age = age
	return nil
}

func (s *FioService) enrichWithGender(fio *Fio) error {
	gender, err := s.genderClient.FetchGender()
	if err != nil {
		return err
	}
	fio.Gender = gender
	return nil
}

func (s *FioService) enrichWithNationality(fio *Fio) error {
	nationality, err := s.nationalityClient.FetchNationality()
	if err != nil {
		return err
	}
	fio.Nationality = nationality
	return nil
}

func isFioValid(fio Fio) error {
	nonAlphabetic := regexp.MustCompile("^[A-Za-z]+$")
	if fio.Name == "" || !nonAlphabetic.MatchString(fio.Name) {
		return fmt.Errorf("the name must contain only alpfabetic characters: %s", fio.Name)
	}
	if fio.Surname == "" || !nonAlphabetic.MatchString(fio.Surname) {
		return fmt.Errorf("the surname must contain only alpfabetic characters: %s", fio.Surname)
	}
	if fio.Patronymic != "" && !nonAlphabetic.MatchString(fio.Patronymic) {
		return fmt.Errorf("the patronymic can be empty or must contain only alpfabetic characters: %s", fio.Surname)
	}
	return nil
}
