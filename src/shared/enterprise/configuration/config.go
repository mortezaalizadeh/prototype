package configuration

type AppConfig struct {
	ListeningInterface string `yaml:"listeningInterface" env:"CONNECTILLY_APP_LISTENINGINTERFACE"`
	Source             string `yaml:"source" env:"CONNECTILLY_APP_SOURCE"`
}
