module go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc/example

go 1.14

replace (
	go.opentelemetry.io/contrib => ../../../../../
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc => ../
)

require (
	github.com/golang/protobuf v1.5.2
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.13.0
	go.opentelemetry.io/otel v0.13.0
	go.opentelemetry.io/otel/exporters/stdout v0.13.0
	go.opentelemetry.io/otel/sdk v0.13.0
	golang.org/x/net v0.0.0-20220722155237-a158d28d115b
	google.golang.org/grpc v1.51.0
)
