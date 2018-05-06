package util

import (
	"errors"
	"strings"

	"github.com/sirupsen/logrus"
)

// Options holds the configuration to setup the logger
type Options struct {
	LogLevel  string
	LogFields logrus.Fields
	LogFormat string
}

// LogHandle is our wrapper around logrus.Entry
type LogHandle struct {
	*logrus.Entry
}

// NewLogger creates new logrus logger for a package
func NewLogger(o Options) (*LogHandle, error) {

	switch strings.ToLower(o.LogLevel) {
	case "warn":
		logrus.SetLevel(logrus.WarnLevel)
	case "info":
		logrus.SetLevel(logrus.InfoLevel)
	case "debug":
		logrus.SetLevel(logrus.DebugLevel)
	case "":
		return nil, errors.New("Option cannot be empty")
	default:
		return nil, errors.New("Unknown LogLevel, expected Debug,Info or Warn")
	}

	switch strings.ToLower(o.LogFormat) {
	case "text":
		logrus.SetFormatter(&logrus.TextFormatter{})
	case "json":
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case "":
		return nil, errors.New("Option cannot be empty")
	default:
		return nil, errors.New("Unknown format, expected text,json")
	}

	logger := &LogHandle{logrus.StandardLogger().WithFields(o.LogFields)}

	return logger, nil

}

// SetFormat sets the output format of the logs
// either text or json
func (lh *LogHandle) SetFormat(format string) error {
	if format != "" {
		switch strings.ToLower(format) {
		case "text":
			lh.Logger.Formatter = &logrus.TextFormatter{}
			return nil
		case "json":
			lh.Logger.Formatter = &logrus.JSONFormatter{}
			return nil
		default:
			return errors.New("Unknown format, expected text,json")
		}
	}
	return errors.New("Unknown format, expected text,json")
}

// SetLevel sets the log level: warn, info or debug
func (lh *LogHandle) SetLevel(level string) error {
	if level != "" {
		switch strings.ToLower(level) {
		case "warn":
			lh.Level = logrus.WarnLevel
			return nil
		case "info":
			lh.Level = logrus.InfoLevel
			return nil
		case "debug":
			lh.Level = logrus.DebugLevel
			return nil
		default:
			return errors.New("Unknown LogLevel, expected Debug,Info or Warn")
		}
	}
	return errors.New("Unknown LogLevel, expected Debug,Info or Warn")
}
