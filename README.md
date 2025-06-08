# `logging` Package - Powerful Structured Logging Made Simple

This package provides a modern interface for structured logging built on Go's standard `log/slog`, with powerful enhancements that will make your code cleaner and your logs more useful.

## 🔥 Key Features

- **Full compatibility** with standard `slog`
- **Convenient attribute constructors** for all types
- **Pointer support** - automatic nil value handling
- **Flexible output formatting** (text/JSON)
- **Shortened source paths** for compact logs
- **Key renaming** in log output
- **Simple configuration** through functional options

## 🚀 Quick Start

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/sergei-galichev/logging"
)

func main() {
	// Create a configured logger
	logger := logging.NewLogger(
		logging.WithLogLevel(logging.LevelDebug),
		logging.WithShortSource(true),
		logging.WithJSONFormat(true),
	)

	// Start logging!
	logger.Info("Application starting",
		logging.String("version", "1.0.0"),
		logging.Time("started_at", time.Now()),
	)

	userID := 42
	logger.Debug("User authenticated",
		logging.Int("user_id", userID),
		logging.BoolPtr("is_admin", nil), // automatically logged as "nil"
	)

	err := fmt.Errorf("connection timeout")
	logger.Error("Request failed",
		logging.Error(err),
	)
}
```

## ⚙️ Logger Configuration

### Log Levels
```go
logging.NewLogger(
	logging.WithLogLevel(logging.LevelDebug), // Debug, Info, Warn or Error
)
```

### Output Format
```go
// Text format (default)
logging.NewLogger(logging.WithJSONFormat(false))

// JSON format
logging.NewLogger(logging.WithJSONFormat(true))
```

### Source Information
```go
// Full file path
logging.NewLogger(logging.WithSource(true))

// Short path (only last directory)
logging.NewLogger(logging.WithShortSource(true))
```

### Renaming Standard Keys
```go
logging.NewLogger(
	logging.WithReplaceDefaultKeyName(logging.TimeKey, "timestamp"),
	logging.WithReplaceDefaultKeyName(logging.LevelKey, "severity"),
)
```

## 🛠 Creating Attributes

The package provides convenient constructors for all types:

### Basic Types
```go
logging.String("name", "Alice")
logging.Int("age", 30)
logging.Bool("active", true)
logging.Float64("score", 4.5)
logging.Time("created_at", time.Now())
logging.Duration("latency", 150*time.Millisecond)
logging.Any("custom_data", map[string]interface{}{"key": "value"})
```

### Pointer Support
```go
var name *string
logging.StringPtr("name", name) // will log as "nil"

age := 30
logging.IntPtr("age", &age)
```

### Special Types
```go
logging.Int32("count", int32(10))
logging.Uint32("flags", uint32(0xFF))
logging.Float32("temp", float32(36.6))
```

## 🏆 Best Practices

1. **Use context** - add request IDs, user identifiers and other useful data to logs
2. **Configure format** - use JSON for production, text format for development
3. **Optimize paths** - enable WithShortSource for more readable logs
4. **Handle nil values** - use *Ptr functions for safe pointer handling

## 📜 Sample Output

**Text Format:**
```
2023-10-05T14:23:18.123Z INFO src/api/user.go:42 User authenticated user_id=42 is_admin=nil
```

**JSON Format:**
```json
{
  "time": "2023-10-05T14:23:18.123Z",
  "level": "INFO",
  "source": "api/user.go:42",
  "msg": "User authenticated",
  "user_id": 42,
  "is_admin": "nil"
}
```

## 📦 Installation

```bash
go get github.com/sergei-galichev/logging
```

Now your logs will be both informative and beautiful! 🎉