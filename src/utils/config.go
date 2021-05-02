package utils

import "os"

type config struct {
	UrlMain string
	MongoUri string
	Port string
	SecretKeyToken string
	SecretKeyPassword string
	DBName string
}

func Config() *config {
	return &config{
		UrlMain: os.Getenv("URL_MAIN"),
		MongoUri: os.Getenv("MONGO_URI"),
		Port: ":" + os.Getenv("PORT"),
		SecretKeyToken: os.Getenv("SECRET_KEY_TOKEN"),
		SecretKeyPassword: os.Getenv("SECRET_KEY_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
	}
}

