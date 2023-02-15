package pulsar

type PulsarConfig struct {
	HttpUrl                    string                           `yaml:"httpUrl" env:"CONNECTILLY_PULSAR_HTTPURL"`
	Url                        string                           `yaml:"url" env:"CONNECTILLY_PULSAR_URL"`
	ProducerName               string                           `yaml:"producerName" env:"CONNECTILLY_PULSAR_PRODUCERNAME"`
	SubscriptionName           string                           `yaml:"subscriptionName" env:"CONNECTILLY_PULSAR_SUBSCRIPTIONNAME"`
	OperationTimeout           string                           `yaml:"operationTimeout" env:"CONNECTILLY_PULSAR_OPERATIONTIMEOUT"`
	ConnectionTimeout          string                           `yaml:"connectionTimeout" env:"CONNECTILLY_PULSAR_CONNECTIONTIMEOUT"`
	EnableAuthenticationOAuth2 bool                             `yaml:"enableAuthenticationOAuth2" env:"CONNECTILLY_PULSAR_ENABLEAUTHENTICATIONOAUTH2"`
	AuthenticationOAuth2       PulsarAuthenticationOAuth2Config `yaml:"authenticationOAuth2,omitempty"`
}

type PulsarAuthenticationOAuth2Config struct {
	Type       string `yaml:"type" env:"CONNECTILLY_PULSAR_AUTHENTICATION_TYPE"`
	IssuerUrl  string `yaml:"issuerUrl" env:"CONNECTILLY_PULSAR_AUTHENTICATION_ISSUERURL"`
	Audience   string `yaml:"audience" env:"CONNECTILLY_PULSAR_AUTHENTICATION_AUDIENCE"`
	PrivateKey string `yaml:"privateKey" env:"CONNECTILLY_PULSAR_AUTHENTICATION_PRIVATEKEY"`
	ClientId   string `yaml:"clientId" env:"CONNECTILLY_PULSAR_AUTHENTICATION_CLIENTID"`
}
