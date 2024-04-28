package kslog

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestNewLoggerSetAppName(t *testing.T) {
	logger := New("default log")
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&PrettyJSONFormatter)
	logger.Infoln("hello log")
	logger.SetAppName("hello, kress4s")
	logger.Infoln("hello, kress4s!")
}

// func TestNewLoggerInfo(t *testing.T) {
// 	logger := New("default log")
// 	logger.Infoln("hello log")
// }
