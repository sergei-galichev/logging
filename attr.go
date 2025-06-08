// Package logging provides attribute constructors for structured logging.
// These functions wrap slog's attribute creation with additional
// convenience functions for pointer types and type conversions.
package logging

import (
	"log/slog"
)

// BoolPtr creates a boolean attribute from a pointer.
// key: The attribute key
// val: Pointer to boolean value. If nil, the attribute will be logged as "nil"
// Returns: A structured log attribute
func BoolPtr(key string, val *bool) slog.Attr {
	if val == nil {
		return slog.String(key, "nil")
	}

	return slog.Bool(key, *val)
}

// StringPtr creates a string attribute from a pointer.
// key: The attribute key
// val: Pointer to string value. If nil, the attribute will be logged as "nil"
// Returns: A structured log attribute
func StringPtr(key string, val *string) slog.Attr {
	if val == nil {
		return slog.String(key, "nil")
	}

	return slog.String(key, *val)
}

// IntPtr creates an integer attribute from a pointer.
// key: The attribute key
// val: Pointer to integer value. If nil, the attribute will be logged as "nil"
// Returns: A structured log attribute
func IntPtr(key string, val *int) slog.Attr {
	if val == nil {
		return slog.String(key, "nil")
	}

	return slog.Int(key, *val)
}

// Int32 creates an attribute from an int32 value.
// key: The attribute key
// val: The int32 value
// Returns: A structured log attribute
func Int32(key string, val int32) slog.Attr {
	return slog.Int(key, int(val))
}

// Int32Ptr creates an attribute from an int32 pointer.
// key: The attribute key
// val: Pointer to int32 value. If nil, the attribute will be logged as "nil"
// Returns: A structured log attribute
func Int32Ptr(key string, val *int32) slog.Attr {
	if val == nil {
		return slog.String(key, "nil")
	}

	return slog.Int(key, int(*val))
}

// Int64Ptr creates an int64 attribute from a pointer.
// key: The attribute key
// val: Pointer to int64 value. If nil, the attribute will be logged as "nil"
// Returns: A structured log attribute
func Int64Ptr(key string, val *int64) slog.Attr {
	if val == nil {
		return slog.String(key, "nil")
	}

	return slog.Int64(key, *val)
}

// Uint creates an attribute from a uint value.
// key: The attribute key
// val: The uint value
// Returns: A structured log attribute
func Uint(key string, val uint) slog.Attr {
	return slog.Uint64(key, uint64(val))
}

// Uint32 creates an attribute from a uint32 value.
// key: The attribute key
// val: The uint32 value
// Returns: A structured log attribute
func Uint32(key string, val uint32) slog.Attr {
	return slog.Uint64(key, uint64(val))
}

// Uint32Ptr creates an attribute from a uint32 pointer.
// key: The attribute key
// val: Pointer to uint32 value. If nil, the attribute will be logged as "nil"
// Returns: A structured log attribute
func Uint32Ptr(key string, val *uint32) slog.Attr {
	if val == nil {
		return slog.String(key, "nil")
	}

	return slog.Uint64(key, uint64(*val))
}

// Float32 creates an attribute from a float32 value.
// key: The attribute key
// val: The float32 value
// Returns: A structured log attribute
func Float32(key string, val float32) slog.Attr {
	return slog.Float64(key, float64(val))
}

// Float32Ptr creates an attribute from a float32 pointer.
// key: The attribute key
// val: Pointer to float32 value. If nil, the attribute will be logged as "nil"
// Returns: A structured log attribute
func Float32Ptr(key string, val *float32) slog.Attr {
	if val == nil {
		return slog.String(key, "nil")
	}

	return slog.Float64(key, float64(*val))
}

// Float64Ptr creates a float64 attribute from a pointer.
// key: The attribute key
// val: Pointer to float64 value. If nil, the attribute will be logged as "nil"
// Returns: A structured log attribute
func Float64Ptr(key string, val *float64) slog.Attr {
	if val == nil {
		return slog.String(key, "nil")
	}

	return slog.Float64(key, *val)
}

// Error creates an error logging attribute.
// Handles nil errors by logging "nil" as the error value.
// err: The error to log
// Returns: A structured log attribute with key "error"
func Error(err error) slog.Attr {
	if err == nil {
		return slog.String("error", "nil")
	}

	return slog.String("error", err.Error())
}
