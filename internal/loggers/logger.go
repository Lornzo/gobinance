package loggers

type ConcreteLogger interface {
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

type Logger struct {
	concreteLogger ConcreteLogger
}

func (l *Logger) SetConcreteLogger(concreteLogger ConcreteLogger) {
	l.concreteLogger = concreteLogger
}

func (l *Logger) Infof(format string, args ...interface{}) {
	if l.concreteLogger == nil {
		return
	}
	l.concreteLogger.Infof(format, args...)
}

func (l *Logger) Debuf(format string, args ...interface{}) {
	if l.concreteLogger == nil {
		return
	}
	l.concreteLogger.Debugf(format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.concreteLogger == nil {
		return
	}
	l.concreteLogger.Errorf(format, args...)
}
