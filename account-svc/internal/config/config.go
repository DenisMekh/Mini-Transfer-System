package config

import (
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Database DBConfig     `mapstructure:"database" validate:"required"`
	Server   ServerConfig `mapstructure:"server" validate:"required"`
	Log      LogConfig    `mapstructure:"log" validate:"required"`
}
type DBConfig struct {
	Host              string        `mapstructure:"host" validate:"required"`
	Port              int           `mapstructure:"port" validate:"required"`
	User              string        `mapstructure:"user" validate:"required"`
	Password          string        `mapstructure:"password" validate:"required"`
	Dbname            string        `mapstructure:"dbname" validate:"required"`
	SSLMode           string        `mapstructure:"sslmode" validate:"required"`
	MaxOpenConns      int           `mapstructure:"max_open_conns" validate:"required,gt=0"`
	MaxIdleConns      int           `mapstructure:"max_idle_conns" validate:"required,gt=0"`
	MaxConnLifetime   time.Duration `mapstructure:"max_conn_lifetime"`
	MaxConnIdleTime   time.Duration `mapstructure:"max_conn_idle_time"`
	HealthCheckPeriod time.Duration `mapstructure:"health_check_period"`
}
type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
type LogConfig struct {
	Level    string `mapstructure:"level" validate:"required"`
	Format   string `mapstructure:"format" validate:"required,oneof=json text"`
	Mode     string `mapstructure:"mode" validate:"required,oneof=debug production development"`
	FilePath string `mapstructure:"file_path"`
	Encoding string `mapstructure:"encoding"`
}

func Load(configPath string) (*Config, error) {
	var err error
	v := viper.New()
	v.SetDefault("server.host", "localhost")
	v.SetDefault("server.port", 8080)
	v.SetDefault("server.read_timeout", "5s")
	v.SetDefault("server.write_timeout", "5s")
	// database defaults
	v.SetDefault("database.dsn", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	v.SetDefault("database.host", "localhost")
	v.SetDefault("database.port", 5432)
	v.SetDefault("database.user", "postgres")
	v.SetDefault("database.password", "postgres")
	v.SetDefault("database.dbname", "postgres")
	v.SetDefault("database.sslmode", "disable")
	v.SetDefault("database.max_open_conns", 25)
	v.SetDefault("database.max_idle_conns", 5)
	v.SetDefault("database.max_conn_lifetime", "1h")
	v.SetDefault("database.max_conn_idle_time", "30m")
	v.SetDefault("database.health_check_period", "1m")
	// log defaults
	v.SetDefault("log.level", "info")
	v.SetDefault("log.format", "json")
	v.SetDefault("log.mode", "production")
	v.SetDefault("log.encoding", "console")
	v.AutomaticEnv()

	if configPath != "" {
		v.SetConfigFile(configPath)
		if err = v.ReadInConfig(); err != nil {
			log.Printf("Error reading config file, %s", err)
		}
	}
	var config Config
	if err = v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct, %v", err)
	}

	validate := validator.New()
	if err = validate.Struct(&config); err != nil {
		return nil, fmt.Errorf("config validation failed, %v", err)
	}
	return &config, nil
}
