package kslog

import (
	"net"
	"net/url"

	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	es6 "github.com/olivere/elastic"
	es7 "github.com/olivere/elastic/v7"
	"github.com/orandin/lumberjackrus"
	"github.com/sirupsen/logrus"
	elog6 "gopkg.in/sohlich/elogrus.v3"
	elog7 "gopkg.in/sohlich/elogrus.v7"
)

func NewFileHook(filename string, maxAge, maxSize, maxBackups int, compress bool, formatter Formatter) (*FileHook, error) {
	hook, err := lumberjackrus.NewHook(
		&lumberjackrus.LogFile{
			Filename:   filename,
			MaxSize:    maxSize,
			MaxBackups: maxBackups,
			MaxAge:     maxAge,
			Compress:   compress,
			LocalTime:  true,
		}, logrus.TraceLevel, formatter, nil)
	if err != nil {
		return nil, err
	}
	return hook, nil
}

func NewLogstashHook(addr, appName string) (Hook, error) {
	uri, err := url.Parse(addr)
	if err != nil {
		return nil, err
	}
	conn, err := net.Dial(uri.Scheme, uri.Host)
	if err != nil {
		return nil, err
	}
	return logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{appNameKey: appName})), nil
}

func NewElasticSearch6Hook(host string, esAddrs []string, index string, level Level) (*Elastic6Hook, error) {
	client, err := es6.NewClient(es6.SetURL(esAddrs...), es6.SetSniff(false))
	if err != nil {
		return nil, err
	}
	hook, err := elog6.NewAsyncElasticHook(client, host, level, index)
	if err != nil {
		return nil, err
	}
	return hook, nil
}

func NewElasticSearch7Hook(host string, esAddrs []string, index string, level Level) (*Elastic7Hook, error) {
	client, err := es7.NewClient(es7.SetURL(esAddrs...), es7.SetSniff(false))
	if err != nil {
		return nil, err
	}
	hook, err := elog7.NewAsyncElasticHook(client, host, level, index)
	if err != nil {
		return nil, err
	}
	return hook, nil
}
