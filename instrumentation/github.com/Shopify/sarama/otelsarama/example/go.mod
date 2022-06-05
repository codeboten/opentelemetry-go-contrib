module go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama/example

go 1.14

replace (
	go.opentelemetry.io/contrib => ../../../../../../
	go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama => ../
)

require (
	github.com/Shopify/sarama v1.34.0
	go.opentelemetry.io/contrib/instrumentation/github.com/Shopify/sarama/otelsarama v0.13.0
	go.opentelemetry.io/otel v0.13.0
	go.opentelemetry.io/otel/exporters/stdout v0.13.0
	go.opentelemetry.io/otel/sdk v0.13.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
)
