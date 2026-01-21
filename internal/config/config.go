package config

type Config struct {
	DBURL string
}

func Load() *Config {
	return &Config{}
}
