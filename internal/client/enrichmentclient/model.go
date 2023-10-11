package enrichmentclient

import (
	"encoding/json"
	"fmt"
	"identity-enricher/internal/service"
)

type FioClientDTO struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int8  `json:"age"` 
	Gender         string `json:"gender"`
	Nationality string `json:"nationality"`
}

func (dto *FioClientDTO) mapToServiceFio() service.Fio {
	var fio service.Fio
	fio.Name = dto.Name
	fio.Surname = dto.Surname
	fio.Patronymic = dto.Patronymic
	fio.Age = dto.Age
	fio.Gender = dto.Gender
	fio.Nationality = dto.Nationality
	return fio
}

type RequestMessage struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
}

func MarshalRequest(request RequestMessage) ([]byte, error) {
	requestMsg, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("marshall request to coinbase: %w", err)
	}
	return requestMsg, nil
}
