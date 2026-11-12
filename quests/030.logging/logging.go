package logging

import "log/slog"

func ProcessSearch(logger *slog.Logger, reqID string, query string, userID int) {
	// Create a child logger with request_id attached.
	logger = logger.With("request_id", reqID)

	// Log that processing has started.
	logger.Debug("processing search request")

	// Validate the search query.
	if query == "" {
		logger.Warn("empty search query")
		return
	}

	// Validate the user ID.
	if userID < 0 {
		logger.Error("invalid user ID")
		return
	}

	// Log a successful search with grouped user information.
	logger.Info(
		"search executed",
		slog.String("query", query),
		slog.Group("user", slog.Int("id", userID)),
	)
}
