package log

import (
	"fmt"
	"path/filepath"

	"github.com/arxanfintech/go-logger"
)

type Logger struct {
	module_name   string
	log_prefix    string
	cfgLevel      LogLevel
	logger        *logger.Logger
	rotate_writer *RotateWriter
}

func New(opts *Options) (log *Logger, err error) {
	moduleName := opts.ModuleName
	filename := fmt.Sprintf("%s.log", moduleName)
	log_file := filepath.Join(opts.LogPath, filename)

	rotate_writer, err := NewRotateWriter(log_file,
		opts.LogMaxSize,
		opts.LogRotateDaily)
	if err != nil {
		return nil, err
	}

	color := 0
	logger, err := logger.New(moduleName, color, rotate_writer)
	if err != nil {
		return nil, err
	}

	log_level, err := ParseLogLevel(opts.LogLevel, opts.Verbose)
	if err != nil {
		log_level = INFO
	}

	log = &Logger{
		module_name:   moduleName,
		log_prefix:    opts.LogPrefix,
		cfgLevel:      log_level,
		logger:        logger,
		rotate_writer: rotate_writer,
	}

	return log, nil
}

func (this *Logger) Close() (err error) {
	return this.rotate_writer.Close()
}

func (this *Logger) log(level LogLevel, format string, args ...interface{}) (err error) {
	if level < this.cfgLevel {
		return nil
	}

	msg := fmt.Sprintf(format, args...)
	fmt.Println(msg)

	s := fmt.Sprintf("%s: %s", level, msg)

	this.Output(3, s)

	return nil
}

func (this *Logger) Output(maxdepth int, s string) error {
	maxdepth += 2
	this.logger.Log("", maxdepth, s)
	return nil
}

func (this *Logger) Fatal(format string, args ...interface{}) (err error) {
	return this.log(FATAL, format, args...)
}

func (this *Logger) Error(format string, args ...interface{}) (err error) {
	return this.log(ERROR, format, args...)
}

func (this *Logger) Warn(format string, args ...interface{}) (err error) {
	return this.log(WARN, format, args...)
}

func (this *Logger) Info(format string, args ...interface{}) (err error) {
	return this.log(INFO, format, args...)
}

func (this *Logger) Debug(format string, args ...interface{}) (err error) {
	return this.log(DEBUG, format, args...)
}
