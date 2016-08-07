// Copyright 2016, Cong Ding. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: Cong Ding <dinggnu@gmail.com>

package stun

import (
	"log"
	"os"
)

// Logger is a simple logger specified for this STUN client.
type Logger struct {
	log.Logger
	debug bool
}

// NewLogger creates a default logger.
func NewLogger() *Logger {
	logger := &Logger{*log.New(os.Stdout, "", log.LstdFlags), false}
	return logger
}

// SetDebug sets the logger running in debug mode or not.
func (l *Logger) SetDebug(debug bool) {
	l.debug = debug
}

// Debug outputs the log in the format of log.Print.
func (l *Logger) Debug(v ...interface{}) {
	if l.debug {
		l.Print(v...)
	}
}

// Debugf outputs the log in the format of log.Printf.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.debug {
		l.Printf(format, v...)
	}
}

// Debugln outputs the log in the format of log.Println.
func (l *Logger) Debugln(v ...interface{}) {
	if l.debug {
		l.Println(v...)
	}
}