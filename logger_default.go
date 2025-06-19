package logging

import (
	"context"
	"os"
)

// Fatal calls on [Logger.Fatal] the default logger.
// Before using this function, please set the function option using [WithSetDefault]
func Fatal(msg string, args ...any) {
	logWithSkip(nil, DefaultLogger(), 3, LevelFatal, msg, args...)

	os.Exit(1)
}

// FatalContext calls [Logger.FatalContext] on the default logger.
// Before using this function, please set the function option using [WithSetDefault]
func FatalContext(ctx context.Context, msg string, args ...any) {
	logWithSkip(ctx, DefaultLogger(), 3, LevelFatal, msg, args...)

	os.Exit(1)
}
