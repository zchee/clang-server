// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import (
	"context"

	"github.com/uber-go/zap"
)

var logger zap.Logger

func init() {
	logger = zap.New(
		zap.NewTextEncoder(
			zap.TextNoTime(),
		),
		zap.AddCaller(),
		// zap.AddStacks(zap.InfoLevel),
		// zap.Fields(zap.Int("pid", os.Getpid()),
		// zap.String("exe", filepath.Base(os.Args[0]))),
	)
}

type correlationIdType int

const (
	requestIdKey correlationIdType = iota
	sessionIdKey
)

// WithRqId returns a context which knows its request ID
func WithRqId(ctx context.Context, rqId string) context.Context {
	return context.WithValue(ctx, requestIdKey, rqId)
}

// WithSessionId returns a context which knows its session ID
func WithSessionId(ctx context.Context, sessionId string) context.Context {
	return context.WithValue(ctx, sessionIdKey, sessionId)
}

// Logger returns a zap logger with as much context as possible
func Logger(ctx context.Context) zap.Logger {
	newLogger := logger
	if ctx != nil {
		if ctxRqId, ok := ctx.Value(requestIdKey).(string); ok {
			newLogger = newLogger.With(zap.String("rqId", ctxRqId))
		}
		if ctxSessionId, ok := ctx.Value(sessionIdKey).(string); ok {
			newLogger = newLogger.With(zap.String("sessionId", ctxSessionId))
		}
	}
	return newLogger
}

// Implements zap standard logger methods.

func Check(lvl zap.Level, msg string) *zap.CheckedMessage {
	return logger.Check(lvl, msg)
}

func Log(lvl zap.Level, msg string, fields ...zap.Field) {
	logger.Log(lvl, msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}
func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
	panic(msg)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
