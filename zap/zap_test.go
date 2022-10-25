package zap

import (
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestNone(t *testing.T) {
	firstStep("www.google.com")
}

func firstStep(url string) {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
