module go.opentelemetry.io/opentelemetry-go-contrib/instrumentation/github.com/emicklei/go-restful/otelrestful/example

go 1.14

replace (
	go.opentelemetry.io/contrib => ../../../../../../
	go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful => ../
	go.opentelemetry.io/contrib/propagators => ../../../../../../propagators
)

require (
	github.com/DataDog/sketches-go v0.0.1 // indirect
	github.com/emicklei/go-restful/v3 v3.3.1
	github.com/google/gofuzz v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/emicklei/go-restful/otelrestful v0.13.0
	go.opentelemetry.io/otel v0.18.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout v0.18.0
	go.opentelemetry.io/otel/sdk v0.18.0
)
