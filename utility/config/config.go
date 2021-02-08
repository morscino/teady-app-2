package config

type Config struct {
	DB Database
}

type Database struct {
	Host     string `envconfig:"TEADY_DB_HOST"`
	Name     string `envconfig:"TEADY_DB_NAME"`
	Dialect  string `envconfig:"TEADY_DB_DIALECT"`
	User     string `envconfig:"TEADY_DB_USER"`
	Password string `envconfig:"TEADY_DB_PASSWORD"`
	Port     string `envconfig:"TEADY_DB_PORT"`
	SSLMode  string `envconfig:"TEADY_DB_SSL_MODE"`
}
