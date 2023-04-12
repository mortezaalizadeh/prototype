package clienteventtemplates

import (
	_ "embed"
)

//go:embed metadata.tmpl
var metadata string

//go:embed handler.tmpl
var handler string

//go:embed consumer.tmpl
var consumer string

//go:embed producer.tmpl
var producer string

//go:embed generate.tmpl
var generate string

func GetMetadata() string {
	return metadata
}

func GetHandler() string {
	return handler
}

func GetConsumer() string {
	return consumer
}

func GetProducer() string {
	return producer
}

func GetGenerate() string {
	return generate
}
