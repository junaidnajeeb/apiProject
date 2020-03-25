package utils

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {

	// Log as JSON instead of the default ASCII formatter.

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//log.SetOutput(os.Stdout)

	// Only log the info severity or above.
	//log.SetLevel(log.DebugLevel)

	//log.SetReportCaller(true)

	log.Formatter = new(logrus.TextFormatter)

	log.Formatter.(*logrus.TextFormatter).FullTimestamp = true
	log.Formatter.(*logrus.TextFormatter).PadLevelText = true
	log.Formatter.(*logrus.TextFormatter).ForceColors = true
	log.Level = logrus.DebugLevel

}

func GetLogger() *logrus.Logger {
	return log
}
func LoggerTrace(message string) {
	log.Trace(message)
}

func LoggerInfo(message string) {
	log.Info(message)
}

func LoggerDebug(message string) {
	log.Debug(message)
}

func LoggerWarn(message string) {
	log.Warn(message)
}

func LoggerError(message string) {
	log.Error(message)
}

func LoggerFatal(message string) {
	log.Fatal(message)
}
