package logging

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

const (
	defaultLogLevel       = LevelDebug
	defaultAddSource      = false
	defaultAddShortSource = false
	defaultJSONFormat     = false
	defaultSetDefault     = false
)

type Options struct {
	LogLevel       Level
	AddSource      bool
	AddShortSource bool
	JSONFormat     bool
	SetDefault     bool
	ReplaceAttrs   map[string]string
}

type Option func(*Options)

func (o *Options) replaceAttr(groups []string, a slog.Attr) slog.Attr {
	switch a.Key {
	case slog.TimeKey:
		return slog.Attr{
			Key:   o.ReplaceAttrs[slog.TimeKey],
			Value: a.Value,
		}
	case slog.SourceKey:
		if o.AddShortSource {
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
					o.ReplaceAttrs[slog.SourceKey],
					fmt.Sprintf("%s:%d", src.File, src.Line),
				)
			}
		} else {
			return slog.Attr{
				Key:   o.ReplaceAttrs[slog.SourceKey],
				Value: a.Value,
			}
		}
	case slog.MessageKey:
		return slog.Attr{
			Key:   o.ReplaceAttrs[slog.MessageKey],
			Value: a.Value,
		}
	case slog.LevelKey:
		return slog.Attr{
			Key:   o.ReplaceAttrs[slog.LevelKey],
			Value: a.Value,
		}
	}

	return a
}

func NewLogger(opts ...Option) *slog.Logger {
	config := &Options{
		LogLevel:       defaultLogLevel,
		AddSource:      defaultAddSource,
		AddShortSource: defaultAddShortSource,
		JSONFormat:     defaultJSONFormat,
		SetDefault:     defaultSetDefault,
		ReplaceAttrs: func() map[string]string {
			return map[string]string{
				slog.TimeKey:    slog.TimeKey,
				slog.SourceKey:  slog.SourceKey,
				slog.MessageKey: slog.MessageKey,
				slog.LevelKey:   slog.LevelKey,
			}
		}(),
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
		handler = NewJSONHandler(w, handlerOpts)
	} else {
		handler = NewTextHandler(w, handlerOpts)
	}

	logger := slog.New(handler)

	if config.SetDefault {
		SetDefault(logger)
	}

	return logger
}

func WithLogLevel(level Level) Option {
	return func(o *Options) {
		o.LogLevel = level
	}
}

func WithSource(source bool) Option {
	return func(o *Options) {
		o.AddSource = source
	}
}

func WithShortSource(shortSource bool) Option {
	return func(o *Options) {
		o.AddShortSource = shortSource
	}
}

func WithJSONFormat(format bool) Option {
	return func(o *Options) {
		o.JSONFormat = format
	}
}

func WithSetDefault(setDefault bool) Option {
	return func(o *Options) {
		o.SetDefault = setDefault
	}
}

func WithReplaceDefaultKeyName(keyName, replaceKeyName string) Option {
	return func(o *Options) {
		if _, ok := o.ReplaceAttrs[keyName]; ok {
			o.ReplaceAttrs[keyName] = replaceKeyName
		}
	}
}

//func (l *Log) WithError(err error, msg string, args ...any) {
//	passArgs := make([]any, len(args)+2)
//
//	passArgs[0] = "error"
//	passArgs[1] = err
//
//	for i, arg := range args {
//		passArgs[i+2] = arg
//	}
//
//	l.original.Error(msg, passArgs...)
//}
