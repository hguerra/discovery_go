package logging

type catchLogger interface {
	Panic(args ...any)
}

func Catch(err error, logger catchLogger) {
	if err != nil {
		logger.Panic(err.Error())
	}
}
