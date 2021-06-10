package utils

import "os"

type config struct {
	MongoUri      string
	NameDataBase  string
	Port          string
	Email         string
	PasswordEmail string
	BaseUri       string
	ApiReniecRuc  string
	ApiReniecDni  string
	TokenReniec   string
}

func Config() *config {
	return &config{
		MongoUri:      os.Getenv("MONGO_URI"),
		NameDataBase:  os.Getenv("NAME_DATABASE"),
		Port:          ":" + os.Getenv("PORT"),
		Email:         os.Getenv("EMAIL"),
		PasswordEmail: os.Getenv("PASSWORD_EMAIL"),
		BaseUri:       os.Getenv("BASE_URI"),
		ApiReniecRuc:  os.Getenv("API_RENIEC_RUC"),
		ApiReniecDni:  os.Getenv("API_RENIEC_DNI"),
		TokenReniec:   os.Getenv("TOKEN_API_RENIC"),
	}
}
