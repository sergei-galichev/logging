package logging

import (
	"context"
	"log/slog"
)

// loggerKey is a private type used as unique context key to store logger instances.
// The empty struct{} ensures no collisions with other context keys.
type loggerKey struct{}

// ContextWithLogger embeds a logger into the context and returns the derived context.
// This enables consistent logging configuration propagation through call chains.
//
// Parameters:
//   - ctx: parent context
//   - logger: pointer to slog.Logger to store in context
//
// Returns:
//   - new context.Context containing the logger
//
// Example:
//
//	logger := slog.New(...)
//	ctx := ContextWithLogger(context.Background(), logger)
//	// Pass ctx through function calls
func ContextWithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// LoggerFromContext extracts a logger from the context.
// Returns the default logger (slog.Default()) if no logger is found in context.
// This provides a safe way to retrieve loggers that always returns a valid logger instance.
//
// Parameters:
//   - ctx: context containing the logger
//
// Returns:
//   - *slog.Logger from context or slog.Default() as fallback
//
// Example:
//
//	logger := LoggerFromContext(ctx)
//	logger.Info("processing request")
func LoggerFromContext(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value(loggerKey{}).(*slog.Logger); ok {
		return logger
	}
	return slog.Default()
}
