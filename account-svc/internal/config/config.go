package config

type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	Port        string `mapstructure:"PORT"`
}

func Load() *Config {
	// TODO WRITE CONFIG LOAD FUNCTION
}
