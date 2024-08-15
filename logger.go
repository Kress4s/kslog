package kslog

import (
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	PanicLevel  = logrus.PanicLevel
	FatalLevel  = logrus.FatalLevel
	ErrorLevel  = logrus.ErrorLevel
	WarnLevel   = logrus.WarnLevel
	InfoLevel   = logrus.InfoLevel
	DebugLevel  = logrus.DebugLevel
	TranceLevel = logrus.TraceLevel
)

const appNameKey = "app_name"

type Logger struct {
	appname string
	logger  *logrus.Logger
}

func New(appname string) *Logger {
	return &Logger{
		appname: appname,
		logger:  logrus.New(),
	}
}

func (l *Logger) SetAppName(appname string) {
	l.appname = appname
}

func (l *Logger) AddHook(hook Hook) {
	l.logger.Hooks.Add(hook)
}

func (l *Logger) SetFormatter(formatter Formatter) {
	l.logger.SetFormatter(formatter)
}

func (l *Logger) SetLevel(level Level) {
	l.logger.SetLevel(level)
}

func (l *Logger) SetOutput(output io.Writer) {
	l.logger.SetOutput(output)
}

func (l *Logger) GetLevel() Level {
	return l.logger.GetLevel()
}

func (l *Logger) WithFields(fields Fields) *Entry {
	return l.logger.WithField(appNameKey, l.appname).WithFields(fields)
}

func (l *Logger) WithField(key string, value interface{}) *Entry {
	return l.logger.WithField(appNameKey, l.appname).WithField(key, value)
}

func (l *Logger) WithError(err error) *Entry {
	return l.logger.WithField(appNameKey, l.appname).WithError(err)
}

func (l *Logger) WithTime(t time.Time) *Entry {
	return l.logger.WithField(appNameKey, l.appname).WithTime(t)
}

func (l *Logger) Log(level Level, args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Log(level, args...)
}

func (l *Logger) Logf(level Level, format string, args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Logf(level, format, args...)
}

func (l *Logger) Logln(level Level, args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Logln(level, args...)
}

func (l *Logger) Print(args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Print(args...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Printf(format, args...)
}

func (l *Logger) Println(args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Println(args...)
}

func (l *Logger) Trace(args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Trace(args...)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Tracef(format, args...)
}

func (l *Logger) Traceln(args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Traceln(args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Debug(args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Debugf(format, args...)
}

func (l *Logger) Debugln(args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Debugln(args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Info(args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Infof(format, args...)
}

func (l *Logger) Infoln(args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Infoln(args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Warn(args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Warnf(format, args...)
}

func (l *Logger) Warnln(args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Warnln(args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Error(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Errorf(format, args...)
}

func (l *Logger) Errorln(args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Errorln(args...)
}

func (l *Logger) Panic(args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Panic(args...)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Panicf(format, args...)
}

func (l *Logger) Panicln(args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Panicln(args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Fatal(args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Panicf(format, args...)
}

func (l *Logger) Fatalln(args ...interface{}) {
	l.logger.WithField(appNameKey, l.appname).Fatalln(args...)
}
