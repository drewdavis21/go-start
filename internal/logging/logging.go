// Package logging includes formatting for the logrus package.
package logging

import (
	log "github.com/sirupsen/logrus"
)

// init sets up the package for use.
func init() {
	formatter := &log.TextFormatter {
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp: true,
		ForceColors: true,
	}
	log.SetFormatter(formatter)
}
