module go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho

go 1.14

replace (
	go.opentelemetry.io/contrib => ../../../../../
	go.opentelemetry.io/contrib/propagators => ../../../../../propagators
)

require (
	github.com/labstack/echo/v4 v4.6.0
	github.com/stretchr/testify v1.6.1
	go.opentelemetry.io/contrib v0.13.0
	go.opentelemetry.io/contrib/propagators v0.13.0
	go.opentelemetry.io/otel v0.13.0
)
