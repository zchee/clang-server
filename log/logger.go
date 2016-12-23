// Copyright 2016 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import (
	"fmt"
	"os"

	"github.com/uber-go/zap"
)

type stdLogger struct {
	write func(string, ...zap.Field)
	panic func(string, ...zap.Field)
	fatal func(string, ...zap.Field)
	debug func(string, ...zap.Field)
}

var logger stdLogger

func init() {
	l := zap.New(zap.NewTextEncoder(zap.TextNoTime()))
	ld := zap.New(zap.NewTextEncoder(zap.TextNoTime()),
		zap.AddCaller(),
	)

	logger = stdLogger{
		panic: l.Panic,
		fatal: l.Fatal,
		write: l.Info,
		debug: ld.Debug,
	}

	switch os.Getenv("CLANG_SERVER_DEBUG") {
	case "1":
		// nothing to do
	default:
		ld.With(zap.Skip())
	}
}

// Print calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Print.
func Print(args ...interface{}) {
	logger.write(fmt.Sprint(args...))
}

// Printf calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Printf.
func Printf(format string, args ...interface{}) {
	logger.write(fmt.Sprintf(format, args...))
}

// Println calls Output to print to the standard logger. Arguments are handled in the manner of fmt.Println.
func Println(args ...interface{}) {
	logger.write(fmt.Sprint(args...))
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(args ...interface{}) {
	logger.panic(fmt.Sprint(args...))
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, args ...interface{}) {
	logger.panic(fmt.Sprintf(format, args...))
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(args ...interface{}) {
	logger.panic(fmt.Sprint(args...))
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(args ...interface{}) {
	logger.fatal(fmt.Sprint(args...))
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, args ...interface{}) {
	logger.fatal(fmt.Sprintf(format, args...))
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(args ...interface{}) {
	logger.fatal(fmt.Sprint(args...))
}

// Debug calls Output to print to the standard logger if set debug level. Arguments are handled in the manner of fmt.Print.
func Debug(args ...interface{}) {
	logger.debug(fmt.Sprint(args...))
}

// Debugf calls Output to print to the standard logger if set debug level. Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, args ...interface{}) {
	logger.debug(fmt.Sprint(args...))
}

// Debugln calls Output to print to the standard logger if set debug level. Arguments are handled in the manner of fmt.Println.
func Debugln(args ...interface{}) {
	logger.debug(fmt.Sprint(args...))
}
