package log

import (
	"fmt"
	"strings"
)

// Level type
type Level uint32

const (
	// WarnLevel level.
	WarnLevel Level = iota
	// InfoLevel level.
	InfoLevel
	// DebugLevel level.
	DebugLevel
)

// ParseLevel converts from string to logapi Level.
func ParseLevel(lvl string) (Level, error) {
	switch strings.ToLower(lvl) {
	case "warn":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	}

	var l Level
	return l, fmt.Errorf("%q is not a valid level", lvl)
}
