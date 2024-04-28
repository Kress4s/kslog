package kslog

import (
	"io"
	"time"
)

var (
	instance *Logger
)

func init() {
	instance = New("kress4s log")
}

func SetAppName(appname string) {
	instance.SetAppName(appname)
}

func GetInstance() *Logger {
	return instance
}

func AddHook(hook Hook) {
	instance.AddHook(hook)
}

func SetFormatter(formatter Formatter) {
	instance.SetFormatter(formatter)
}

func SetLevel(level Level) {
	instance.SetLevel(level)
}

func SetOutput(output io.Writer) {
	instance.SetOutput(output)
}

func GetLevel() Level {
	return instance.GetLevel()
}

func WithFields(fields Fields) *Entry {
	return instance.WithField(appNameKey, instance.appname).WithFields(fields)
}

func WithField(key string, value interface{}) *Entry {
	return instance.WithField(appNameKey, instance.appname).WithField(key, value)
}

func WithError(err error) *Entry {
	return instance.WithField(appNameKey, instance.appname).WithError(err)
}

func WithTime(t time.Time) *Entry {
	return instance.WithField(appNameKey, instance.appname).WithTime(t)
}

func Log(level Level, args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Log(level, args...)
}

func Logf(level Level, format string, args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Logf(level, format, args...)
}

func Logln(level Level, args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Logln(level, args...)
}

func Print(args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Print(args...)
}

func Printf(format string, args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Printf(format, args...)
}

func Println(args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Println(args...)
}

func Trace(args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Trace(args...)
}

func Tracef(format string, args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Tracef(format, args...)
}

func Traceln(args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Traceln(args...)
}

func Debug(args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Debugf(format, args...)
}

func Debugln(args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Debugln(args...)
}

func Info(args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Info(args...)
}

func Infof(format string, args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Infof(format, args...)
}

func Infoln(args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Infoln(args...)
}

func Warn(args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Warnf(format, args...)
}

func Warnln(args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Warnln(args...)
}

func Error(args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Error(args...)
}

func Errorf(format string, args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Errorf(format, args...)
}

func Errorln(args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Errorln(args...)
}

func Panic(args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Panicf(format, args...)
}

func Panicln(args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Panicln(args...)
}

func Fatal(args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Panicf(format, args...)
}

func Fatalln(args ...interface{}) {
	instance.WithField(appNameKey, instance.appname).Fatalln(args...)
}
