package logrusex

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/xm-chentl/go-code/logex"
	"github.com/xm-chentl/go-code/logex/loglevel"
)

var (
	once         sync.Once
	mappingLevel = make(map[loglevel.Value]logrus.Level)
)

type loggerImpl struct {
	logrus  *logrus.Logger
	title   string
	dataMap map[string]string
}

func (l *loggerImpl) Title(format string, args ...interface{}) logex.ILogger {
	l.title = format
	if len(args) > 0 {
		l.title = fmt.Sprintf(l.title, args...)
	}

	return l
}

func (l *loggerImpl) Label(name string, format string, args ...interface{}) logex.ILogger {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args...)
	}
	l.dataMap[name] = format

	return l
}

func (l *loggerImpl) LabelBy(args ...interface{}) {
	key := "detail"
	if len(args) == 0 {
		l.dataMap[key] = "[]"
		return
	}
	for index := range args {
		value := args[index]
		valueRt := reflect.ValueOf(value)
		if valueRt.Kind() == reflect.Ptr {
			args[index] = reflect.ValueOf(value).Elem().Interface()
		}
	}
	argsBytes, _ := json.Marshal(args)
	l.dataMap[key] = string(argsBytes)
}

func (l *loggerImpl) GetLog() logex.ILog {
	return &logImpl{
		logrus: l.logrus,
	}
}

func (l *loggerImpl) Debug(err error) {
	l.writer(func(entry *logrus.Entry) {
		entry.Debug(err)
	})
}

func (l *loggerImpl) Error(err error) {
	l.writer(func(entry *logrus.Entry) {
		entry.Error(err)
	})
}

func (l *loggerImpl) Fatal(err error) {
	l.writer(func(entry *logrus.Entry) {
		entry.Fatal(err)
	})
}

func (l *loggerImpl) Info(err error) {
	l.writer(func(entry *logrus.Entry) {
		entry.Info(err)
	})
}

func (l *loggerImpl) Trace(err error) {
	l.writer(func(entry *logrus.Entry) {
		entry.Trace(err)
	})
}

func (l *loggerImpl) Warn(err error) {
	l.writer(func(entry *logrus.Entry) {
		entry.Warn(err)
	})
}

func (l *loggerImpl) Panic(err error) {
	l.writer(func(entry *logrus.Entry) {
		entry.Panic(err)
	})
}

func (l *loggerImpl) writer(f func(entry *logrus.Entry)) {
	defer l.reset()

	fields := logrus.Fields{}
	if l.title != "" {
		fields["title"] = l.title
	}
	for k, v := range l.dataMap {
		fields[k] = v
	}
	if len(fields) > 0 {
		entry := l.logrus.WithFields(fields)
		if f != nil {
			f(entry)
		}
	}

	return
}

func (l *loggerImpl) reset() {
	l.dataMap = make(map[string]string)
}

type logImpl struct {
	logrus *logrus.Logger
}

func (l logImpl) Logger() logex.ILogger {
	return &loggerImpl{
		logrus:  l.logrus,
		dataMap: make(map[string]string),
	}
}

func New(level loglevel.Value, writers ...io.Writer) logex.ILog {
	inst := logrus.New()
	if len(writers) > 0 {
		inst.SetOutput(io.MultiWriter())
	}
	inst.SetLevel(mappingLevel[level])
	inst.SetFormatter(&logrus.JSONFormatter{})
	inst.SetReportCaller(false)
	return &logImpl{
		logrus: inst,
	}
}

func init() {
	once.Do(func() {
		mappingLevel[loglevel.Panic] = logrus.PanicLevel
		mappingLevel[loglevel.Error] = logrus.ErrorLevel
		mappingLevel[loglevel.Fatal] = logrus.FatalLevel
		mappingLevel[loglevel.Debug] = logrus.DebugLevel
		mappingLevel[loglevel.Info] = logrus.InfoLevel
		mappingLevel[loglevel.Trace] = logrus.TraceLevel
		mappingLevel[loglevel.Warn] = logrus.WarnLevel
	})
}
