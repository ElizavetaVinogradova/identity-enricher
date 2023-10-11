package service

type FioRepository interface {
	Create(fio Fio) error
}

type Fio struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Age         int8  `json:"age"`
	Gender         string `json:"gender"`
	Nationality string `json:"nationality"`
	Error      string `json:"error"`
}
