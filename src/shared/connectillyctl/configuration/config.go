package configuration

import (
	"github.com/Connectilly/connectilly/src/shared/enterprise/database/postgres"
	"github.com/Connectilly/connectilly/src/shared/enterprise/messaging/pulsar"
)

type Config struct {
	Pulsar   pulsar.PulsarConfig     `yaml:"pulsar"`
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}
