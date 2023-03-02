package log

import "io"

// WithLogOutput returns a LoggerOption that sets the output for the logger. The
// default is os.Stderr.
func WithLogOutput(w io.Writer) LoggerOption {
	return func(l *logger) {
		l.w = w
	}
}

// WithLogTimeFunction returns a LoggerOption that sets the time function for the
// logger. The default is time.Now.
func WithLogTimeFunction(f TimeFunction) LoggerOption {
	return func(l *logger) {
		l.timeFunc = f
	}
}

// WithLogTimeFormat returns a LoggerOption that sets the time format for the
// logger. The default is "2006/01/02 15:04:05".
func WithLogTimeFormat(format string) LoggerOption {
	return func(l *logger) {
		l.timeFormat = format
	}
}

// WithLogLevel returns a LoggerOption that sets the level for the logger. The
// default is InfoLevel.
func WithLogLevel(level Level) LoggerOption {
	return func(l *logger) {
		l.level = int32(level)
	}
}

// WithLogPrefix returns a LoggerOption that sets the prefix for the logger.
func WithLogPrefix(prefix string) LoggerOption {
	return func(l *logger) {
		l.prefix = prefix
	}
}

// WithLogTimestamp returns a LoggerOption that enables timestamps for the logger.
func WithLogTimestamp() LoggerOption {
	return func(l *logger) {
		l.timestamp = true
	}
}

// WithLogCaller returns a LoggerOption that enables caller for the logger.
func WithLogCaller() LoggerOption {
	return func(l *logger) {
		l.caller = true
	}
}

// WithLogFields returns a LoggerOption that sets the fields for the logger.
func WithLogFields(keyvals ...interface{}) LoggerOption {
	return func(l *logger) {
		l.keyvals = keyvals
	}
}

// WithLogFormatter returns a LoggerOption that sets the formatter for the logger.
func WithLogFormatter(f Formatter) LoggerOption {
	return func(l *logger) {
		l.formatter = f
	}
}
