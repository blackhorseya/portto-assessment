package loggerx

// Options is the logging options.
type Options struct {
	// Level is the log level. options: debug, info, warn, error, dpanic, panic, fatal (default: info)
	Level string `json:"level" yaml:"level" mapstructure:"level"`

	// Format is the log format. options: json, text (default: text)
	Format string `json:"format" yaml:"format" mapstructure:"format"`
}

// Logger is the logging interface.
type Logger interface {
	Debug(msg string, fields ...any)
	Info(msg string, fields ...any)
	Warn(msg string, fields ...any)
	Error(msg string, fields ...any)
}
