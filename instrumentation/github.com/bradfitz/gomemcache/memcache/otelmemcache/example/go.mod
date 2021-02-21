module go.opentelemetry.io/opentelemetry-go-contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache/example

go 1.14

replace (
	go.opentelemetry.io/contrib => ../../../../../../../
	go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache => ../
)

require (
	github.com/DataDog/sketches-go v0.0.1 // indirect
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b
	github.com/google/gofuzz v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/bradfitz/gomemcache/memcache/otelmemcache v0.13.0
	go.opentelemetry.io/otel v0.17.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout v0.17.0
	go.opentelemetry.io/otel/sdk v0.17.0
)
