// Copyright 2017 The clang-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package log

import (
	"fmt"
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
)

var (
	logger      = log.New(os.Stderr, "", log.Lshortfile)
	debug       = os.Getenv("CLANG_SERVER_DEBUG")
	debugLogger = log.New(os.Stderr, "DEBUG: ", log.Lshortfile)
)

// Print calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Print.
func Print(v ...interface{}) {
	logger.Output(2, fmt.Sprint(v...))
}

// Printf calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Printf(format string, v ...interface{}) {
	logger.Output(2, fmt.Sprintf(format, v...))
}

// Println calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Println.
func Println(v ...interface{}) {
	logger.Output(2, fmt.Sprintln(v...))
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	logger.Output(2, "FATAL: "+fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	logger.Output(2, "FATAL: "+fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func Fatalln(v ...interface{}) {
	logger.Output(2, "FATAL: "+fmt.Sprintln(v...))
	os.Exit(1)
}

// Panic is equivalent to Print() followed by a call to panic().
func Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	logger.Output(2, "PANIC: "+s)
	panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	logger.Output(2, "PANIC: "+s)
	panic(s)
}

// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	logger.Output(2, "PANIC: "+s)
	panic(s)
}

// Debug calls Output to print to the standard logger if debug is true.
// Arguments are handled in the manner of fmt.Print.
func Debug(v ...interface{}) {
	if len(debug) == 0 {
		return
	}
	debugLogger.Output(2, fmt.Sprint(v...))
}

// Debugf calls Output to print to the standard logger if debug is true.
// Arguments are handled in the manner of fmt.Printf.
func Debugf(format string, v ...interface{}) {
	if len(debug) == 0 {
		return
	}
	debugLogger.Output(2, fmt.Sprintf(format, v...))
}

// Debugln calls Output to print to the standard logger if debug is true.
// Arguments are handled in the manner of fmt.Println.
func Debugln(format string, v ...interface{}) {
	if len(debug) == 0 {
		return
	}
	debugLogger.Output(2, fmt.Sprintln(v...))
}

func Dump(v ...interface{}) {
	spew.Dump(v...)
}
