package postgres

import (
	"fmt"
	"identity-enricher/internal/service"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

type FioRepository struct {
	db *sqlx.DB
}

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewFioRepository(config Config) (*FioRepository, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("user=%s dbname=%s host=%s port=%s password=%s sslmode=%s",
		config.Username, config.DBName, config.Host, config.Port, config.Password, config.SSLMode))

	if err != nil {
		return nil, fmt.Errorf("open sql connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return &FioRepository{db: db}, nil
}

func (r *FioRepository) Close() {
	r.db.Close()
}

type FioDataBaseDTO struct {
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic"`
	Age         int8   `json:"age"`
	Gender      string `json:"gender"`
	Nationality string `json:"nationality"`
}

func (r *FioRepository) Create(fio service.Fio) error {
	fioDB := mapServiceFioToFioDTO(fio)
	query := "INSERT INTO fio (name, surname, patronymic, age, gender, nationality) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("prepare statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(fioDB.Name, fioDB.Surname, fioDB.Patronymic, fioDB.Age, fioDB.Gender, fioDB.Nationality)
	if err != nil {
		return fmt.Errorf("execute create statement: %w", err)
	}

	return nil
}

func (r *FioRepository) GetById(id int64) (service.Fio, error) {
	query := "SELECT * FROM ticks WHERE id = ?"

	var fioDB FioDataBaseDTO
	err := r.db.Get(&fioDB, query, id)
	log.Debugf("================statement got data : %s , id: %d", fmt.Sprintf("%v", fioDB), id)
	if err != nil {
		return service.Fio{}, fmt.Errorf("execute get statement: %w", err)
	}

	fio := mapFioDTOToFio(fioDB)
	log.Debugf("================mapped DB tick to service tick : %s", fmt.Sprintf("%v", fio))
	return fio, nil
}


func mapServiceFioToFioDTO(fio service.Fio) FioDataBaseDTO {
	var fioDB FioDataBaseDTO
	fioDB.Name = fio.Name
	fioDB.Surname = fio.Surname
	fioDB.Patronymic = fio.Patronymic
	fioDB.Age = fio.Age
	fioDB.Gender = fio.Gender
	fioDB.Nationality = fio.Nationality
	return fioDB
}

func mapFioDTOToFio(fioDB FioDataBaseDTO) service.Fio {
	var fio service.Fio
	fio.Name = fioDB.Name
	fio.Surname = fioDB.Surname
	fio.Patronymic = fioDB.Patronymic
	fio.Age = fioDB.Age
	fio.Gender = fioDB.Gender
	fio.Nationality = fioDB.Nationality
	return fio
}
