package loggerconfig

import (
	"github.com/herb-go/logger"
)

type LoggerConfig struct {
	URI    string
	Config func(v interface{}) error `config:", lazyload"`
}
type Config struct {
	Loggers map[string]LoggerConfig
	Formats map[string]string
}

func (c *Config) ApplyTo(l *logger.Logger, name string) error {
	lc := c.Loggers[name]
	return logger.InitLogger(l, lc.URI, lc.Config)
}

func (c *Config) SetFormatter(l *logger.FormatLogger, name string) error {
	f := c.Formats[name]
	l.SetFormatter(logger.ReplacementFormater(f))
	return nil
}
func (c *Config) ApplyToBuiltinLoggers() error {
	var err error
	builtinLoggers := map[string]*logger.Logger{
		"Fatal":   logger.FatalLogger,
		"Panic":   logger.PanicLogger,
		"Error":   logger.ErrorLogger,
		"Print":   logger.PrintLogger,
		"Warning": logger.WarningLogger,
		"Info":    logger.InfoLogger,
		"Trace":   logger.TraceLogger,
		"Debug":   logger.DebugLogger,
	}
	for k := range builtinLoggers {
		err = c.ApplyTo(builtinLoggers[k], k)
		if err != nil {
			return err
		}
	}
	return nil
}
