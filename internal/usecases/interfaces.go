package usecases

type (
	LoggerInterface interface {
		Infof(format string, args ...interface{})
		Errorf(format string, args ...interface{})
		Debugf(format string, args ...interface{})

		Info(args ...interface{})
		Error(args ...interface{})
		Debug(args ...interface{})
	}
)
