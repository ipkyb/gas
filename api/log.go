package api

import (
	"fmt"
	"log"
	"os"
	"path"
)

// The default flag for the Logger package-level.
const LogDefaultFlag = log.Ldate | log.Ltime | log.Lshortfile

// The default Logger configuration for the package-level.
var Log = NewLog(LogConfig{
	Info: LogTypeConfig{
		Flag:     LogDefaultFlag,
		Filepath: "logs/info.log",
	},
	Warn: LogTypeConfig{
		Flag:     LogDefaultFlag,
		Filepath: "logs/warn.log",
	},
	Debug: LogTypeConfig{
		Flag:     LogDefaultFlag,
		Filepath: "logs/debug.log",
	},
	Error: LogTypeConfig{
		Flag:     LogDefaultFlag,
		Filepath: "logs/error.log",
	},
})

// Logger represents instance for append-only logging stored as file. Using
// deferred approach to open files and load states only when 'log' method is
// invoked, this is to increase resource efficiency. The underlying
// implementation is served by built-in log package.
type Logger struct {
	config LogConfig
}

// LogTypeConfig is a configuration for specific type of logs. The underlying
// implementation is handled by built-in log package in go.
type LogTypeConfig struct {
	Flag     int    // Flag passed when creating new log.Logger.
	Prefix   string // Prefix passed when creating new log.Logger.
	Filepath string // Destination to store log output.

	logger *log.Logger // This logger object is deferred on instantiation.
}

// LogConfig is a configuration for various type of logs determined by the
// package. Currently there are four (4) types of log and subject to change.
type LogConfig struct {
	Info  LogTypeConfig
	Debug LogTypeConfig
	Warn  LogTypeConfig
	Error LogTypeConfig
}

// Creates a new Logger instance using the given config.
func NewLog(config LogConfig) *Logger {
	return &Logger{config}
}

// open will returns a new log.Logger based on the configuration provided by
// the current instance. It will try to open the given filepath for writing
// access, or panic.
func (i *LogTypeConfig) open() *log.Logger {
	dir := path.Dir(i.Filepath)
	dperm := os.FileMode(0755)
	err := os.MkdirAll(dir, dperm)
	if err != nil {
		panic(err)
	}

	perm := os.FileMode(0666)
	flag := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile(i.Filepath, flag, perm)
	if err != nil {
		panic(err)
	}

	return log.New(file, i.Prefix, i.Flag)
}

// append is supposed to be called from helper exported functions. This set
// the call depth to 3, making the file whom calls the Helper function is
// correctly identified and written to the log output.
func (i *LogTypeConfig) append(v ...any) {
	if i.logger == nil {
		i.logger = i.open()
	}
	s := fmt.Sprintln(v...)
	i.logger.Output(3, s)
}

// Log using Info to provide information about the state runtime of the
// application.
func (i *Logger) Info(v ...any) {
	i.config.Info.append(v...)
}

// Log using Debug to provide information only for debugging or development
// purposes.
func (i *Logger) Debug(v ...any) {
	i.config.Debug.append(v...)
}

// Log using Warn is to inform about abnormal activity in the application that
// might leads to an error in the future.
func (i *Logger) Warn(v ...any) {
	i.config.Warn.append(v...)
}

// Log using Error to signal a failure or error during the runtime of the
// application.
//
// Error means:
//   - Request or operation could not be served or completed.
//   - Application doesn't know how workaround current issue.
func (i *Logger) Error(v ...any) {
	i.config.Error.append(v...)
}
