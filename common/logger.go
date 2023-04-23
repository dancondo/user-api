package common

import (
	"github.com/sirupsen/logrus"
)

var Log = logrus.WithFields(logrus.Fields{
	"app": GetEnv("APP_NAME"),
	"env": GetEnv("ENV"),
})

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyLevel: "severity",
		},
	})

	logrus.RegisterExitHandler(func() {
		logrus.Info("application will stop probably due to a os signal")
	})

	ll := GetEnv("LOG_LEVEL")
	l, err := logrus.ParseLevel(ll)
	if err != nil {
		Log.WithError(err).Errorf("error parsing log level %s", ll)
		return
	}

	logrus.SetLevel(l)
}
