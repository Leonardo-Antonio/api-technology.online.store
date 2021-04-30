package utils

import "os"

type config struct {
	UrlMain string
	MongoUri string
	Port string
	SecretKey string
	DBName string
}

func Config() *config {
	return &config{
		UrlMain: os.Getenv("URL_MAIN"),
		MongoUri: os.Getenv("MONGO_URI"),
		Port: ":" + os.Getenv("PORT"),
		SecretKey: os.Getenv("SECRET_KEY"),
		DBName: os.Getenv("DB_NAME"),
	}
}

