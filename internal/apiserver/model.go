package apiserver

import (
	"encoding/json"
	"fmt"
)

type FioAPI struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int8   `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

func MarshalRequest(request FioAPI) ([]byte, error) {
	requestMsg, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("marshall request to API: %w", err)
	}
	return requestMsg, nil
}

func UnmarshalResponse(message []byte) (FioAPI, error) {
	var fioAPI FioAPI
	err := json.Unmarshal(message, &fioAPI)
	if err != nil {
		return FioAPI{}, fmt.Errorf("unmarshall request to coinbase: %w", err)
	}
	return fioAPI, nil
}
