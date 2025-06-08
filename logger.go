// Package logging provides a configurable logger implementation
// based on Go's standard log/slog package with additional features
// like source location shortening and attribute key renaming.

package logging

import (
	"fmt"
	"log/slog"
	"maps"
	"os"
	"path/filepath"
	"strings"
)

const (
	// Default configuration constants
	defaultLogLevel       = LevelDebug
	defaultAddSource      = false
	defaultAddShortSource = false
	defaultJSONFormat     = false
	defaultSetDefault     = false
)

var (
	// defaultReplaceAttrs defines the default attribute key replacements
	defaultReplaceAttrs = map[string]string{
		slog.TimeKey:    slog.TimeKey,
		slog.SourceKey:  slog.SourceKey,
		slog.MessageKey: slog.MessageKey,
		slog.LevelKey:   slog.LevelKey,
	}
)

// Options contains configuration for the logger
type Options struct {
	LogLevel       Level             // Minimum log level to output
	AddSource      bool              // Whether to add source file location
	AddShortSource bool              // Whether to shorten source file paths
	JSONFormat     bool              // Use JSON format instead of text
	SetDefault     bool              // Set this logger as the default
	ReplaceAttrs   map[string]string // Attribute key replacements
}

// Option defines a function type for configuring Options
type Option func(*Options)

// shortSourceAttr shortens the source file path in the attribute
// a: Original source attribute
// newKey: New key name for the attribute
// Returns: Modified attribute with shortened source path
func (o *Options) shortSourceAttr(a slog.Attr, newKey string) slog.Attr {
	if src, ok := a.Value.Any().(*slog.Source); ok {
		dir, file := filepath.Split(src.File)

		dirParts := strings.Split(filepath.ToSlash(filepath.Clean(dir)), "/")

		if len(dirParts) > 0 {
			shortDir := dirParts[len(dirParts)-1]
			if shortDir == "" && len(dirParts) > 1 {
				shortDir = dirParts[len(dirParts)-2]
			}

			src.File = filepath.Join(shortDir, file)
		}

		return slog.String(
			newKey,
			fmt.Sprintf("%s:%d", src.File, src.Line),
		)
	}

	return slog.Attr{
		Key:   newKey,
		Value: a.Value,
	}
}

// replaceAttr handles attribute key replacement and source shortening
// groups: Current attribute groups
// a: Original attribute
// Returns: Modified attribute
func (o *Options) replaceAttr(groups []string, a slog.Attr) slog.Attr {
	if newKey, ok := o.ReplaceAttrs[a.Key]; ok {
		if a.Key == slog.SourceKey && o.AddShortSource {
			return o.shortSourceAttr(a, newKey)
		}

		return slog.Attr{
			Key:   newKey,
			Value: a.Value,
		}
	}

	return a
}

// NewLogger creates a new configured logger instance
// opts: Variadic list of configuration options
// Returns: Configured slog.Logger instance
func NewLogger(opts ...Option) *slog.Logger {
	config := &Options{
		LogLevel:       defaultLogLevel,
		AddSource:      defaultAddSource,
		AddShortSource: defaultAddShortSource,
		JSONFormat:     defaultJSONFormat,
		SetDefault:     defaultSetDefault,
		ReplaceAttrs:   maps.Clone(defaultReplaceAttrs),
	}

	for _, opt := range opts {
		opt(config)
	}

	handlerOpts := &slog.HandlerOptions{
		Level:       config.LogLevel,
		AddSource:   config.AddSource || config.AddShortSource,
		ReplaceAttr: config.replaceAttr,
	}

	w := os.Stdout

	var handler slog.Handler

	if config.JSONFormat {
		handler = slog.NewJSONHandler(w, handlerOpts)
	} else {
		handler = slog.NewTextHandler(w, handlerOpts)
	}

	logger := slog.New(handler)

	if config.SetDefault {
		slog.SetDefault(logger)
	}

	return logger
}

// WithLogLevel sets the minimum log level
// level: Minimum log level to output
// Returns: Configuration option function
func WithLogLevel(level Level) Option {
	return func(o *Options) {
		o.LogLevel = level
	}
}

// WithSource enables/disables source location logging
// source: Whether to enable source location
// Returns: Configuration option function
func WithSource(source bool) Option {
	return func(o *Options) {
		o.AddSource = source
	}
}

// WithShortSource enables/disables shortened source paths
// shortSource: Whether to enable shortened paths
// Returns: Configuration option function
func WithShortSource(shortSource bool) Option {
	return func(o *Options) {
		o.AddShortSource = shortSource
	}
}

// WithJSONFormat sets the output format to JSON
// format: Whether to use JSON format
// Returns: Configuration option function
func WithJSONFormat(format bool) Option {
	return func(o *Options) {
		o.JSONFormat = format
	}
}

// WithSetDefault sets whether to make this logger default
// setDefault: Whether to set as default logger
// Returns: Configuration option function
func WithSetDefault(setDefault bool) Option {
	return func(o *Options) {
		o.SetDefault = setDefault
	}
}

// WithReplaceDefaultKeyName replaces a default attribute key name
// keyName: Original key name to replace
// replaceKeyName: New key name to use
// Returns: Configuration option function
func WithReplaceDefaultKeyName(keyName, replaceKeyName string) Option {
	return func(o *Options) {
		if _, ok := o.ReplaceAttrs[keyName]; ok {
			o.ReplaceAttrs[keyName] = replaceKeyName
		}
	}
}
