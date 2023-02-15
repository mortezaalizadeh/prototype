package postgres

type PostgresConfig struct {
	ConnectionString string `yaml:"connectionString" env:"CONNECTILLY_POSTGRES_CONNECTIONSTRING"`
	MaxOpenConns     int    `yaml:"maxOpenConns" env:"CONNECTILLY_POSTGRES_MAXOPENCONNS"`
}
