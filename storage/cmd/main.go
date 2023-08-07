package main

import (
	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"storage/pkg/broker"
)

func main() {
	broker.InitRabbitMQConnection("amqp://admin:admin@rabbitmq:5672/")

	go broker.ConsumeMessages("user_creation_queue")

	//db, err := database2.NewPostgresDB(database2.Config{
	//	Host:     viper.GetString("db.HOST"),
	//	Port:     viper.GetString("db.POSTGRES_PORT"),
	//	User:     viper.GetString("db.POSTGRES_USER"),
	//	Password: viper.GetString("db.POSTGRES_PASSWORD"),
	//	DBName:   viper.GetString("db.POSTGRES_DB"),
	//	SSLMode:  viper.GetString("db.SSL_MODE"),
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//repo := database2.NewRepository(db)
	//_ = repo
}
