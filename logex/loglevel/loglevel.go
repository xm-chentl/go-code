package loglevel

type Value uint

const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	Panic Value = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	Fatal
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	Error
	// WarnLevel level. Non-critical entries that deserve eyes.
	Warn
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	Info
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	Debug
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	Trace
)
