package kslog

import (
	"github.com/orandin/lumberjackrus"
	"github.com/sirupsen/logrus"
	elog6 "gopkg.in/sohlich/elogrus.v3"
	elog7 "gopkg.in/sohlich/elogrus.v7"
)

type (
	Level         = logrus.Level
	Fields        = logrus.Fields
	Entry         = logrus.Entry
	Hook          = logrus.Hook
	Elastic6Hook  = elog6.ElasticHook
	Elastic7Hook  = elog7.ElasticHook
	FileHook      = lumberjackrus.Hook
	Formatter     = logrus.Formatter
	JSONFormatter = logrus.JSONFormatter
)

var PrettyJSONFormatter = JSONFormatter{
	DisableHTMLEscape: true,
	// json indent
	PrettyPrint: true,
}
