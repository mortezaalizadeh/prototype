package configuration

import (
	"github.com/Connectilly/connectilly/src/shared/enterprise/database/postgres"
)

type Config struct {
	Postgres postgres.PostgresConfig `yaml:"postgres"`
}
