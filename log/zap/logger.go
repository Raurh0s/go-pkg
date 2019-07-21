// MIT License
//
// Copyright (c) 2019 Thibault NORMAND
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package zap

import (
	"go.uber.org/zap"

	"go.zenithar.org/pkg/log"
)

// logger delegates all calls to the underlying zap.Logger
type logger struct {
	logger *zap.Logger
}

// Debug logs an debug msg with fields
func (l logger) Debug(msg string, fields ...log.Field) {
	l.logger.Debug(msg, zfields(fields)...)
}

// Info logs an info msg with fields
func (l logger) Info(msg string, fields ...log.Field) {
	l.logger.Info(msg, zfields(fields)...)
}

// Error logs an error msg with fields
func (l logger) Error(msg string, fields ...log.Field) {
	l.logger.Error(msg, zfields(fields)...)
}

// Warn logs a warning with fields
func (l logger) Warn(msg string, fields ...log.Field) {
	l.logger.Warn(msg, zfields(fields)...)
}

// Fatal logs a fatal error msg with fields
func (l logger) Fatal(msg string, fields ...log.Field) {
	l.logger.Fatal(msg, zfields(fields)...)
}

// With creates a child logger, and optionally adds some context fields to that logger.
func (l logger) With(fields ...log.Field) log.Logger {
	return &logger{logger: l.logger.With(zfields(fields)...)}
}
