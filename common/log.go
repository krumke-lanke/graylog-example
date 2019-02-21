package common

import (
	"fmt"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
)

type LogService struct {
	*logrus.Logger
}

type CtxLogger struct {
	*logrus.Entry
}

func NewLogger() (*LogService, error) {
	logger := LogService{}
	logger.Logger = logrus.New()
	logger.Out = os.Stdout
	logger.Formatter = &prefixed.TextFormatter{}
	return &logger, nil
}

func (ls *LogService) NewPrefix(prefix string) *CtxLogger {
	ctxLog := CtxLogger{
		Entry: ls.WithField("prefix", prefix),
	}
	return &ctxLog
}

func (ctxl *CtxLogger) NewPrefix(prefix string) *CtxLogger {
	ctxLog := CtxLogger{
		Entry: ctxl.WithField("prefix", prefix),
	}
	return &ctxLog
}

func (ctxl *CtxLogger) Print(v ...interface{}) {
	ctxl.Debug(v)
}

func (ctxl *CtxLogger) AddPrefix(prefix string) *CtxLogger {
	var newPrefix string
	if data, ok := ctxl.Data["prefix"]; !ok {
		newPrefix = prefix
	} else {
		newPrefix = fmt.Sprintf("%s.%s", data, prefix)
	}
	return &CtxLogger{
		Entry: ctxl.WithField("prefix", newPrefix),
	}
}
