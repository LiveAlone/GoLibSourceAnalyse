package dig

import (
	"encoding/json"
	"fmt"
	"go.uber.org/dig"
	"log"
	"os"
	"testing"
)

type TestOpt struct {
}

func TestFirst(t *testing.T) {
	type Config struct {
		Prefix string
	}

	c := dig.New()
	var err error

	err = c.Provide(func() (*Config, error) {
		fmt.Println(".......stating Config")
		var cfg Config
		err := json.Unmarshal([]byte(`{"prefix": "[foo] "}`), &cfg)
		return &cfg, err
	})
	if err != nil {
		log.Fatalf("Failed to provide config: %v", err)
	}

	err = c.Provide(func(cfg *Config) *log.Logger {
		fmt.Println(".......stating logger")
		return log.New(os.Stdout, cfg.Prefix, 0)
	})
	if err != nil {
		log.Fatalf("Failed to provide logger: %v", err)
	}

	fmt.Println(".......stating invoke")
	for i := 0; i < 300; i++ {
		go func() {
			err = c.Invoke(func(l *log.Logger) {
				l.Print("You've been invoked")
			})
		}()
	}

	if err != nil {
		log.Fatalf("Failed to invoke: %v", err)
	}

}
