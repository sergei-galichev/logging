package logging

import "log/slog"

var (
	Debug        = slog.Debug
	DebugContext = slog.DebugContext
	Info         = slog.Info
	InfoContext  = slog.InfoContext
	Warn         = slog.Warn
	WarnContext  = slog.WarnContext
	Error        = slog.Error
	ErrorContext = slog.ErrorContext
	Log          = slog.Log
	LogAttrs     = slog.LogAttrs

	// DefaultLogger is the default logger
	DefaultLogger = slog.Default
)
