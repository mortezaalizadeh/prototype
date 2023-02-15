package outbox

import (
	"time"
)

type OutboxConfig struct {
	MaxRetryCount int           `yaml:"maxRetryCount" env:"CONNECTILLY_OUTBOX_MAXRETRYCOUNT"`
	RetryDelay    time.Duration `yaml:"retryDelay" env:"CONNECTILLY_OUTBOX_RETRYDELAY"`
}
