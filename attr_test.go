package logging

import (
	"log/slog"
	"math"
	"testing"
)

func TestBoolPtr(t *testing.T) {
	tests := []struct {
		name string
		val  *bool
		want slog.Attr
	}{
		{"nil pointer", nil, slog.String("key", "nil")},
		{"true value", boolPtr(true), slog.Bool("key", true)},
		{"false value", boolPtr(false), slog.Bool("key", false)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BoolPtr("key", tt.val)
			if got.Key != tt.want.Key || got.Value.String() != tt.want.Value.String() {
				t.Errorf("BoolPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringPtr(t *testing.T) {
	tests := []struct {
		name string
		val  *string
		want slog.Attr
	}{
		{"nil pointer", nil, slog.String("key", "nil")},
		{"non-empty string", stringPtr("test"), slog.String("key", "test")},
		{"empty string", stringPtr(""), slog.String("key", "")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringPtr("key", tt.val)
			if got.Key != tt.want.Key || got.Value.String() != tt.want.Value.String() {
				t.Errorf("StringPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntPtr(t *testing.T) {
	tests := []struct {
		name string
		val  *int
		want slog.Attr
	}{
		{"nil pointer", nil, slog.String("key", "nil")},
		{"positive value", intPtr(42), slog.Int("key", 42)},
		{"zero value", intPtr(0), slog.Int("key", 0)},
		{"negative value", intPtr(-1), slog.Int("key", -1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IntPtr("key", tt.val)
			if got.Key != tt.want.Key || got.Value.String() != tt.want.Value.String() {
				t.Errorf("IntPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32(t *testing.T) {
	tests := []struct {
		name string
		val  int32
		want slog.Attr
	}{
		{"positive value", 42, slog.Int("key", 42)},
		{"zero value", 0, slog.Int("key", 0)},
		{"negative value", -1, slog.Int("key", -1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int32("key", tt.val)
			if got.Key != tt.want.Key || got.Value.String() != tt.want.Value.String() {
				t.Errorf("Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Ptr(t *testing.T) {
	tests := []struct {
		name string
		val  *int32
		want slog.Attr
	}{
		{"nil pointer", nil, slog.String("key", "nil")},
		{"positive value", int32Ptr(42), slog.Int("key", 42)},
		{"zero value", int32Ptr(0), slog.Int("key", 0)},
		{"negative value", int32Ptr(-1), slog.Int("key", -1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int32Ptr("key", tt.val)
			if got.Key != tt.want.Key || got.Value.String() != tt.want.Value.String() {
				t.Errorf("Int32Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Ptr(t *testing.T) {
	tests := []struct {
		name string
		val  *int64
		want slog.Attr
	}{
		{"nil pointer", nil, slog.String("key", "nil")},
		{"positive value", int64Ptr(42), slog.Int64("key", 42)},
		{"zero value", int64Ptr(0), slog.Int64("key", 0)},
		{"negative value", int64Ptr(-1), slog.Int64("key", -1)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Int64Ptr("key", tt.val)
			if got.Key != tt.want.Key || got.Value.String() != tt.want.Value.String() {
				t.Errorf("Int64Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint(t *testing.T) {
	tests := []struct {
		name string
		val  uint
		want slog.Attr
	}{
		{"positive value", 42, slog.Uint64("key", 42)},
		{"zero value", 0, slog.Uint64("key", 0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint("key", tt.val)
			if got.Key != tt.want.Key || got.Value.String() != tt.want.Value.String() {
				t.Errorf("Uint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32(t *testing.T) {
	tests := []struct {
		name string
		val  uint32
		want slog.Attr
	}{
		{"positive value", 42, slog.Uint64("key", 42)},
		{"zero value", 0, slog.Uint64("key", 0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint32("key", tt.val)
			if got.Key != tt.want.Key || got.Value.String() != tt.want.Value.String() {
				t.Errorf("Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUint32Ptr(t *testing.T) {
	tests := []struct {
		name string
		val  *uint32
		want slog.Attr
	}{
		{"nil pointer", nil, slog.String("key", "nil")},
		{"positive value", uint32Ptr(42), slog.Uint64("key", 42)},
		{"zero value", uint32Ptr(0), slog.Uint64("key", 0)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Uint32Ptr("key", tt.val)
			if got.Key != tt.want.Key || got.Value.String() != tt.want.Value.String() {
				t.Errorf("Uint32Ptr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFloat32(t *testing.T) {
	tests := []struct {
		name string
		val  float32
		want float64
	}{
		{"positive value", 3.14, 3.14},
		{"zero value", 0, 0},
		{"negative value", -1.23, -1.23},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float32("key", tt.val)
			// Compare float values directly instead of string representations
			if got.Key != "key" {
				t.Errorf("Float32() key = %v, want 'key'", got.Key)
			}
			if got.Value.Kind() != slog.KindFloat64 {
				t.Errorf("Float32() value kind = %v, want %v", got.Value.Kind(), slog.KindFloat64)
			}
			if math.Abs(got.Value.Float64()-tt.want) > 1e-6 {
				t.Errorf("Float32() value = %v, want %v", got.Value.Float64(), tt.want)
			}
		})
	}
}

func TestFloat32Ptr(t *testing.T) {
	tests := []struct {
		name string
		val  *float32
		want float64
	}{
		{"nil pointer", nil, 0}, // The actual value doesn't matter for nil case
		{"positive value", float32Ptr(3.14), 3.14},
		{"zero value", float32Ptr(0), 0},
		{"negative value", float32Ptr(-1.23), -1.23},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float32Ptr("key", tt.val)
			if tt.val == nil {
				if got.Key != "key" || got.Value.String() != "nil" {
					t.Errorf("Float32Ptr() = %v, want key=nil", got)
				}
				return
			}

			if got.Key != "key" {
				t.Errorf("Float32Ptr() key = %v, want 'key'", got.Key)
			}
			if got.Value.Kind() != slog.KindFloat64 {
				t.Errorf("Float32Ptr() value kind = %v, want %v", got.Value.Kind(), slog.KindFloat64)
			}
			if math.Abs(got.Value.Float64()-tt.want) > 1e-6 {
				t.Errorf("Float32Ptr() value = %v, want %v", got.Value.Float64(), tt.want)
			}
		})
	}
}

func TestFloat64Ptr(t *testing.T) {
	tests := []struct {
		name string
		val  *float64
		want float64
	}{
		{"nil pointer", nil, 0}, // The actual value doesn't matter for nil case
		{"positive value", float64Ptr(3.14), 3.14},
		{"zero value", float64Ptr(0), 0},
		{"negative value", float64Ptr(-1.23), -1.23},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float64Ptr("key", tt.val)
			if tt.val == nil {
				if got.Key != "key" || got.Value.String() != "nil" {
					t.Errorf("Float64Ptr() = %v, want key=nil", got)
				}
				return
			}

			if got.Key != "key" {
				t.Errorf("Float64Ptr() key = %v, want 'key'", got.Key)
			}
			if got.Value.Kind() != slog.KindFloat64 {
				t.Errorf("Float64Ptr() value kind = %v, want %v", got.Value.Kind(), slog.KindFloat64)
			}
			if math.Abs(got.Value.Float64()-tt.want) > 1e-6 {
				t.Errorf("Float64Ptr() value = %v, want %v", got.Value.Float64(), tt.want)
			}
		})
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want slog.Attr
	}{
		{"nil error", nil, slog.String("error", "nil")},
		{"non-nil error", errorMock("test error"), slog.String("error", "test error")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Error(tt.err)
			if got.Key != tt.want.Key || got.Value.String() != tt.want.Value.String() {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper functions to create pointers for testing
func boolPtr(b bool) *bool          { return &b }
func stringPtr(s string) *string    { return &s }
func intPtr(i int) *int             { return &i }
func int32Ptr(i int32) *int32       { return &i }
func int64Ptr(i int64) *int64       { return &i }
func uint32Ptr(i uint32) *uint32    { return &i }
func float32Ptr(f float32) *float32 { return &f }
func float64Ptr(f float64) *float64 { return &f }

// errorMock is a simple error implementation for testing
type errorMock string

func (e errorMock) Error() string { return string(e) }
