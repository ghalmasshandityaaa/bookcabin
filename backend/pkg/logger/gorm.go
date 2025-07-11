package logger

import (
	"context"
	"time"

	gormlogger "gorm.io/gorm/logger"

	"github.com/sirupsen/logrus"
)

// CustomGormLogger implements gorm.Logger interface with JSON formatting using logrus
type CustomGormLogger struct {
	level  gormlogger.LogLevel
	logger *logrus.Logger
}

func NewCustomGormLogger(log *logrus.Logger) *CustomGormLogger {
	return &CustomGormLogger{
		level:  gormlogger.Info,
		logger: log,
	}
}

func (l *CustomGormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	newLogger := *l
	newLogger.level = level
	return &newLogger
}

func (l *CustomGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= gormlogger.Info {
		l.logger.WithFields(logrus.Fields{
			"component": "gorm",
			"data":      data,
		}).Info(msg)
	}
}

func (l *CustomGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= gormlogger.Warn {
		l.logger.WithFields(logrus.Fields{
			"component": "gorm",
			"data":      data,
		}).Warn(msg)
	}
}

func (l *CustomGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= gormlogger.Error {
		l.logger.WithFields(logrus.Fields{
			"component": "gorm",
			"data":      data,
		}).Error(msg)
	}
}

func (l *CustomGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.level <= gormlogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	sql, rows := fc()

	fields := logrus.Fields{
		"component":     "gorm",
		"elapsed":       elapsed.String(),
		"rows_affected": rows,
		"sql":           sql,
	}

	if err != nil {
		fields["error"] = err.Error()
		l.logger.WithFields(fields).Error("gorm query error")
	} else if l.level >= gormlogger.Info {
		l.logger.WithFields(fields).Info("gorm query executed")
	}
}
