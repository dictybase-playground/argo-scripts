package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v1"
)

func GetLogger(c *cli.Context) (*logrus.Entry, error) {
	e := new(logrus.Entry)
	log := logrus.New()
	log.Out = os.Stderr
	format := c.GlobalString("log-format")
	switch format {
	case "text":
		log.Formatter = &logrus.TextFormatter{
			TimestampFormat: "02/Jan/2006:15:04:05",
		}
	case "json":
		log.Formatter = &logrus.JSONFormatter{
			TimestampFormat: "02/Jan/2006:15:04:05",
		}
	default:
		return e, fmt.Errorf(
			"only json and text are supported %s log format is not supported",
			format,
		)
	}
	level := c.GlobalString("log-level")
	switch level {
	case "debug":
		log.Level = logrus.DebugLevel
	case "warn":
		log.Level = logrus.WarnLevel
	case "error":
		log.Level = logrus.ErrorLevel
	case "fatal":
		log.Level = logrus.FatalLevel
	case "panic":
		log.Level = logrus.PanicLevel
	default:
		return e, fmt.Errorf(
			"%s log level is not supported",
			level,
		)
	}
	return logrus.NewEntry(log), nil
}
