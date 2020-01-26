package loggerconfig

import (
	"github.com/herb-go/logger"
)

type Config map[string]string

func (c *Config) ApplyTo(l *logger.Logger, name string) error {
	url := (*c)[name]
	return logger.InitLogger(l, url)
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
