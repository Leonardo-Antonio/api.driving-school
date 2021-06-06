package utils

import "os"

type config struct {
	MongoUri      string
	NameDataBase  string
	Email         string
	PasswordEmail string
}

func Config() *config {
	return &config{
		MongoUri:      os.Getenv("MONGO_URI"),
		NameDataBase:  os.Getenv("NAME_DATABASE"),
		Email:         os.Getenv("EMAIL"),
		PasswordEmail: os.Getenv("PASSWORD_EMAIL"),
	}
}
