package postgres

import (
	"fmt"
	"identity-enricher/internal/service"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
	Id          int64  `db:"id"`
	Name        string `db:"name"`
	Surname     string `db:"surname"`
	Patronymic  string `db:"patronymic"`
	Age         int8   `db:"age"`
	Gender      string `db:"gender"`
	Nationality string `db:"nationality"`
}

func (r *FioRepository) Create(fio service.Fio) error {
	fioDB := mapServiceFioToFioDTO(fio)

	sqlStatement := `INSERT INTO fio (name, surname, patronymic, age, gender, nationality) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(sqlStatement, fioDB.Name, fioDB.Surname, fioDB.Patronymic, fioDB.Age, fioDB.Gender, fioDB.Nationality)
	if err != nil {
		return fmt.Errorf("execute create statement: %w", err)
	}

	return nil
}

func (r *FioRepository) GetFioById(id int64) (service.Fio, error) {
	query := "SELECT * FROM fio WHERE id = $1"

	var fioDB FioDataBaseDTO
	err := r.db.Get(&fioDB, query, id)
	if err != nil {
		return service.Fio{}, fmt.Errorf("execute get statement: %w", err)
	}

	fio := mapFioDTOToFio(fioDB)
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
