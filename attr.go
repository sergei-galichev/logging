// Package logging provides attribute constructors for structured logging.
// These functions wrap slog's attribute creation with additional
// convenience functions for pointer types and type conversions.
package logging

import (
	"log/slog"
)

// Bool creates a boolean logging attribute.
// key: The attribute key
// val: The boolean value
// Returns: A structured log attribute
func Bool(key string, val bool) Attr {
	return slog.Bool(key, val)
}

// BoolPtr creates a boolean attribute from a pointer.
// key: The attribute key
// val: Pointer to boolean value. If nil, the attribute will be logged as "nil"
// Returns: A structured log attribute
func BoolPtr(key string, val *bool) Attr {
	if val == nil {
		return slog.String(key, "nil")
	}

	return slog.Bool(key, *val)
}

// String creates a string logging attribute.
// key: The attribute key
// val: The string value
// Returns: A structured log attribute
func String(key string, val string) Attr {
	return slog.String(key, val)
}

// StringPtr creates a string attribute from a pointer.
// key: The attribute key
// val: Pointer to string value. If nil, the attribute will be logged as "nil"
// Returns: A structured log attribute
func StringPtr(key string, val *string) Attr {
	if val == nil {
		return slog.String(key, "nil")
	}

	return slog.String(key, *val)
}

// Int creates an integer logging attribute.
// key: The attribute key
// val: The integer value
// Returns: A structured log attribute
func Int(key string, val int) Attr {
	return slog.Int(key, val)
}

// IntPtr creates an integer attribute from a pointer.
// key: The attribute key
// val: Pointer to integer value. If nil, the attribute will be logged as "nil"
// Returns: A structured log attribute
func IntPtr(key string, val *int) Attr {
	if val == nil {
		return slog.String(key, "nil")
	}

	return slog.Int(key, *val)
}

// Int32 creates an attribute from an int32 value.
// key: The attribute key
// val: The int32 value
// Returns: A structured log attribute
func Int32(key string, val int32) Attr {
	return slog.Int(key, int(val))
}

// Int32Ptr creates an attribute from an int32 pointer.
// key: The attribute key
// val: Pointer to int32 value. If nil, the attribute will be logged as "nil"
// Returns: A structured log attribute
func Int32Ptr(key string, val *int32) Attr {
	if val == nil {
		return slog.String(key, "nil")
	}

	return slog.Int(key, int(*val))
}

// Int64 creates an int64 logging attribute.
// key: The attribute key
// val: The int64 value
// Returns: A structured log attribute
func Int64(key string, val int64) Attr {
	return slog.Int64(key, val)
}

// Int64Ptr creates an int64 attribute from a pointer.
// key: The attribute key
// val: Pointer to int64 value. If nil, the attribute will be logged as "nil"
// Returns: A structured log attribute
func Int64Ptr(key string, val *int64) Attr {
	if val == nil {
		return slog.String(key, "nil")
	}

	return slog.Int64(key, *val)
}

// Float32 creates an attribute from a float32 value.
// key: The attribute key
// val: The float32 value
// Returns: A structured log attribute
func Float32(key string, val float32) Attr {
	return slog.Float64(key, float64(val))
}

// Float32Ptr creates an attribute from a float32 pointer.
// key: The attribute key
// val: Pointer to float32 value. If nil, the attribute will be logged as "nil"
// Returns: A structured log attribute
func Float32Ptr(key string, val *float32) Attr {
	if val == nil {
		return slog.String(key, "nil")
	}

	return slog.Float64(key, float64(*val))
}

// Float64 creates a float64 logging attribute.
// key: The attribute key
// val: The float64 value
// Returns: A structured log attribute
func Float64(key string, val float64) Attr {
	return slog.Float64(key, val)
}

// Float64Ptr creates a float64 attribute from a pointer.
// key: The attribute key
// val: Pointer to float64 value. If nil, the attribute will be logged as "nil"
// Returns: A structured log attribute
func Float64Ptr(key string, val *float64) Attr {
	if val == nil {
		return slog.String(key, "nil")
	}

	return slog.Float64(key, *val)
}

// Error creates an error logging attribute.
// Handles nil errors by logging "nil" as the error value.
// err: The error to log
// Returns: A structured log attribute with key "error"
func Error(err error) Attr {
	if err == nil {
		return slog.String("error", "nil")
	}

	return slog.String("error", err.Error())
}

// Any creates an attribute for any value type.
// key: The attribute key
// val: The value of any type
// Returns: A structured log attribute
func Any(key string, val any) Attr {
	return slog.Any(key, val)
}

// Dict creates a group attribute containing multiple attributes.
// key: The group key
// val: Variadic list of attributes to include in the group
// Returns: A structured group attribute
func Dict(key string, val ...Attr) Attr {
	return slog.Group(key, val)
}
