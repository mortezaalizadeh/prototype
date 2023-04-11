package generators

//go:generate mockgen -source=client_event_consumer_generator_service.go -destination=mock/client_event_consumer_generator_service_gen.go
//go:generate mockgen -source=client_event_generate_generator_service.go -destination=mock/client_event_generate_generator_service_gen.go
//go:generate mockgen -source=client_event_handler_generator_service.go -destination=mock/client_event_handler_generator_service_gen.go
//go:generate mockgen -source=client_event_metadata_generator_service.go -destination=mock/client_event_metadata_generator_service_gen.go
//go:generate mockgen -source=client_event_producer_generator_service.go -destination=mock/client_event_producer_generator_service_gen.go
//go:generate mockgen -source=client_event_schema_generator_service.go -destination=mock/client_event_schema_generator_service_gen.go
