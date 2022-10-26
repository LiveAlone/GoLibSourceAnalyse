package zap

import (
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestNone(t *testing.T) {
	firstStep("www.google.com")
}

func TestNamespace(t *testing.T) {
	logger := zap.NewExample()
	defer logger.Sync()

	logger.With(
		zap.Namespace("metrics"),
		zap.Int("counter", 1),
	).Info("tracked some metrics")
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
