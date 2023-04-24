package dig

import (
	"encoding/json"
	"go.uber.org/dig"
	"log"
	"os"
	"testing"
)

func TestFirst(t *testing.T) {
	type Config struct {
		Prefix string
	}

	c := dig.New()
	var err error

	// Provide a Config object. This can fail to decode.
	err = c.Provide(func() (*Config, error) {
		// In a real program, the configuration will probably be read from a
		// file.
		var cfg Config
		err := json.Unmarshal([]byte(`{"prefix": "[foo] "}`), &cfg)
		return &cfg, err
	})
	if err != nil {
		log.Fatalf("Failed to provide config: %v", err)
	}

	// Provide a way to build the logger based on the configuration.
	err = c.Provide(func(cfg *Config) *log.Logger {
		return log.New(os.Stdout, cfg.Prefix, 0)
	})
	if err != nil {
		log.Fatalf("Failed to provide logger: %v", err)
	}

	// Invoke a function that requires the logger, which in turn builds the
	// Config first.
	err = c.Invoke(func(l *log.Logger) {
		l.Print("You've been invoked")
	})
	if err != nil {
		log.Fatalf("Failed to invoke: %v", err)
	}

}
