package utils

import "os"

type config struct {
	MongoUri string `bson:"mongo_uri" json:"mongo_uri"`
	Port string `bson:"port" json:"port"`
	SecretKey string `bson:"secret_key" json:"secret_key"`
	DBName string `bson:"db_name" json:"db_name"`
}

func Config() *config {
	return &config{
		MongoUri: os.Getenv("MONGO_URI"),
		Port: ":" + os.Getenv("PORT"),
		SecretKey: os.Getenv("SECRET_KEY"),
		DBName: os.Getenv("DB_NAME"),
	}
}

