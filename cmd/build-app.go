package cmd

import (
	"identity-enricher/internal/apiserver"
	"identity-enricher/internal/repo/postgres"

	"github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.AddConfigPath("../../configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	return viper.ReadInConfig()
}

func SetupLogging() {
	logLevel := viper.GetString("logLevel")

	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warning":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
}

func BuildKafkaReaderConfig() kafka.ReaderConfig {
	return kafka.ReaderConfig{
		Brokers: []string{viper.GetString("kafka.brokerAddress")},
		Topic:   viper.GetString("kafka.readerTopic"),
	}
}

func BuildKafkaWriterConfig() kafka.WriterConfig {
	return kafka.WriterConfig{
		Brokers: []string{viper.GetString("kafka.brokerAddress")},
		Topic:   viper.GetString("kafka.writerTopic"),
	}
}

func BuildPostgreSqlConfig() postgres.Config {
	return postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	}
}

func GetAgeURl() string {
	return viper.GetString("enrichmentClient.ageUrl")
}
func GetGenderURl() string {
	return viper.GetString("enrichmentClient.genderUrl")
}
func GetNationalityURl() string {
	return viper.GetString("enrichmentClient.nationalityUrl")
}

func BuildApiServerConfig() apiserver.Config {
	return apiserver.Config{
		BindAddress: viper.GetString("apiserver.bindAddress"),
	}
}
