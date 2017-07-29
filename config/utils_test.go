package config

import (
	"testing"
)

var (
	sampleConfig = `
ListenAddr = ":8080"
`
)

func TestParseConfig(t *testing.T) {
	cfg, err := ParseConfig(sampleConfig)
	if err != nil {
		t.Fatalf("error parsing config: %s", err)
	}

	if cfg.ListenAddr != ":8080" {
		t.Fatalf("expected listen addr :8080; received %s", cfg.ListenAddr)
	}
}
