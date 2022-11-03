package configs

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("data/log/logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.SetOutput(os.Stdout)
		logrus.Info("Failed to log to file, using default stderr")
	}

	logrus.SetFormatter(&logrus.JSONFormatter{
		// DisableTimestamp: true,
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyFunc:  "caller",
			logrus.FieldKeyMsg:   "message",
		},
	})
	if os.Getenv("ENVIRONMENT") == "live" {
		logrus.SetLevel(logrus.WarnLevel)
	} else {
		logrus.SetReportCaller(true)
		logrus.SetLevel(logrus.DebugLevel)
	}

}
