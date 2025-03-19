package configx

import (
	"portto/pkg/loggerx"

	"github.com/spf13/viper"
)

func init() {
	viper.AutomaticEnv()

	viper.MustBindEnv("host", "HOST")
	viper.MustBindEnv("port", "PORT")
	viper.MustBindEnv("db.host", "DB_HOST")
	viper.MustBindEnv("db.port", "DB_PORT")
	viper.MustBindEnv("db.user", "DB_USER")
	viper.MustBindEnv("db.password", "DB_PASSWORD")
	viper.MustBindEnv("db.name", "DB_NAME")
	viper.MustBindEnv("otel.target", "OTEL_TARGET")
}

type Application struct {
	Verbose bool   `mapstructure:"verbose" json:"verbose" yaml:"verbose"`
	Output  string `mapstructure:"output" json:"output" yaml:"output"`

	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`

	Database struct {
		Host     string `mapstructure:"host" json:"host" yaml:"host"`
		Port     int    `mapstructure:"port" json:"port" yaml:"port"`
		User     string `mapstructure:"user" json:"user" yaml:"user"`
		Password string `mapstructure:"password" json:"password" yaml:"password"`
		Name     string `mapstructure:"name" json:"name" yaml:"name"`
	} `mapstructure:"db" json:"db" yaml:"db"`

	OTel struct {
		Target string `mapstructure:"target" json:"target" yaml:"target"`
	} `mapstructure:"otel" json:"otel" yaml:"otel"`
}

func (x *Application) SetupLogger() error {
	level := "info"
	if x.Verbose {
		level = "debug"
	}

	format := "text"
	if x.Output == "json" {
		format = "json"
	}

	// Initialize the logger with the specified format and level
	_, err := loggerx.NewSlogLogger(loggerx.WithFormat(format), loggerx.WithLevel(level))
	if err != nil {
		return err
	}

	return nil
}
