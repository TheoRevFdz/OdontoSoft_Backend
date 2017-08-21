package config

import (
	"fmt"

	"github.com/eduardogpg/gonv"
)

// DatabaseConfig estructura con atributos de configuracion
type DatabaseConfig struct {
	Username string
	Password string
	Host     string
	Port     int
	Databse  string
}

var database *DatabaseConfig

func init() {
	database = &DatabaseConfig{}
	database.Username = gonv.GetStringEnv("USER", "theo")
	database.Password = gonv.GetStringEnv("PASSWORD", "ambu")
	database.Host = gonv.GetStringEnv("HOST", "localhost")
	database.Port = gonv.GetIntEnv("PORT", 5432)
	database.Databse = gonv.GetStringEnv("DATABASE", "odontosoft")
}

// GetUrlDatabase obtiene la cadena de coneccion con la db
func GetUrlDatabase() string {
	return database.url()
}

func (this *DatabaseConfig) url() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", this.Username, this.Password, this.Host, this.Port, this.Databse)
}
