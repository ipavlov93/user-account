package config

type Config struct {
	DBConfig DBConfig
}

type DBConfig struct {
	Host         string
	Port         int
	Username     string
	Password     string
	DatabaseName string
}
