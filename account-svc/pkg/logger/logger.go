package logger

type Logger interface {
	Debug(msg string, fields ...any)

	Info(msg string, fields ...any)

	Warn(msg string, fields ...any)

	Error(msg string, fields ...any)

	Fatal(msg string, fields ...any)
}
