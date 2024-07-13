// Package glog implements powerful and easy-to-use levelled logging functionality.
package glog

import (
	"github.com/ilylx/gconv/internal/cmdenv"
	"github.com/ilylx/gconv/internal/os/grpool"
)

var (
	// Default logger object, for package method usage.
	logger = New()

	// Goroutine pool for async logging output.
	// It uses only one asynchronize worker to ensure log sequence.
	asyncPool = grpool.New(1)

	// defaultDebug enables debug level or not in default,
	// which can be configured using command option or system environment.
	defaultDebug = true
)

func init() {
	defaultDebug = cmdenv.Get("gf.glog.debug", true).Bool()
	SetDebug(defaultDebug)
}

// Default returns the default logger.
func DefaultLogger() *Logger {
	return logger
}

// SetDefaultLogger sets the default logger for package glog.
// Note that there might be concurrent safety issue if calls this function
// in different goroutines.
func SetDefaultLogger(l *Logger) {
	logger = l
}
