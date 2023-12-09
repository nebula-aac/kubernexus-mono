package nexuslogger

import (
	"context"
	"fmt"
	"time"

	gormlogger "gorm.io/gorm/logger"
)

type databaseWrapper struct {
	enabled bool
	base    *wrapper
}

// Error implements logger.Interface
func (gl *databaseWrapper) Error(ctx context.Context, msg string, data ...interface{}) {
	gl.base.logger.WithContext(ctx).Error(fmt.Sprintf(msg, data...))
}

// Info implements logger.Interface
func (gl *databaseWrapper) Info(ctx context.Context, msg string, data ...interface{}) {
	gl.base.logger.WithContext(ctx).Info(fmt.Sprintf(msg, data...))
}

// LogMode implements logger.Interface
func (gl *databaseWrapper) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return gl
}

// Trace implements logger.Interface
func (gl *databaseWrapper) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	elapsed := time.Since(begin)

	gl.base.logger.WithContext(ctx).Debugf("[SQL] %s [RowsAffected] %d [ElapsedTime] %v", sql, rows, elapsed)
	if err != nil {
		gl.base.logger.WithContext(ctx).Error(err)
	}
}

// Warn implements logger.Interface
func (gl *databaseWrapper) Warn(ctx context.Context, msg string, data ...interface{}) {
	gl.base.logger.WithContext(ctx).Warn(fmt.Sprintf(msg, data...))
}

func (gl *wrapper) DatabaseLogger() gormlogger.Interface {
	return &databaseWrapper{
		enabled: true,
		base:    gl,
	}
}
