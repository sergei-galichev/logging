package main

import (
	"github.com/sergei-galichev/logging"
	"log"
	"net/http"
)

func main() {
	logger := logging.NewLogger(
		logging.WithLogLevel(logging.LevelDebug),
		//logging.WithSource(true),
		logging.WithShortSource(true),
		logging.WithJSONFormat(true),
		logging.WithReplaceDefaultKeyName(logging.TimeKey, "timestamp"),
		logging.WithReplaceDefaultKeyName(logging.SourceKey, "caller"),
		logging.WithSetDefault(true),
	)

	logger.Debug("debug message", logging.String("key", "debug"))
	logger.Info("info message", logging.String("key", "info"))
	logger.Error("error message", logging.String("key", "error"))

	logger.Info("info message", logging.Bool("key", true))

	mux := http.NewServeMux()

	mux.HandleFunc(
		"/", func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "not found", http.StatusInternalServerError)

		},
	)

	server := &http.Server{
		Addr:     ":8060",
		Handler:  mux,
		ErrorLog: log.Default(),
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Error("error starting server", logging.Error(err))
	}
}
