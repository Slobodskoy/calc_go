package config

type Config struct {
	Port int `env:"PORT,required"`
}
