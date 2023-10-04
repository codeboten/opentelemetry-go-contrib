// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package config // import "go.opentelemetry.io/contrib/config"

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type validatable interface {
	Validate() error
}

func TestValidate(t *testing.T) {
	tests := []struct {
		name      string
		component validatable
		err       error
	}{
		{
			name:      "invalid-span-processor",
			component: &SpanProcessor{},
			err:       errUnsupportedSpanProcessor,
		},
		{
			name: "valid-span-processor-batch-invalid-exporter",
			component: &SpanProcessor{
				Batch: &BatchSpanProcessor{},
			},
			err: errInvalidExporter,
		},
		{
			name: "valid-span-processor-simple-invalid-exporter",
			component: &SpanProcessor{
				Simple: &SimpleSpanProcessor{},
			},
			err: errInvalidExporter,
		},
		{
			name: "valid-span-processor-batch-valid-otlp-exporter",
			component: &SpanProcessor{
				Batch: &BatchSpanProcessor{
					Exporter: SpanExporter{
						OTLP: &OTLP{},
					},
				},
			},
		},
		{
			name: "valid-span-processor-batch-valid-console-exporter",
			component: &SpanProcessor{
				Batch: &BatchSpanProcessor{
					Exporter: SpanExporter{
						Console: Console{},
					},
				},
			},
		},
		{
			name: "valid-span-processor-batch-valid-zipkin-exporter",
			component: &SpanProcessor{
				Batch: &BatchSpanProcessor{
					Exporter: SpanExporter{
						Zipkin: &Zipkin{},
					},
				},
			},
		},
		{
			name:      "invalid-metric-reader",
			component: &MetricReader{},
			err:       errUnsupportedMetricReader,
		},
		{
			name: "valid-metric-reader-invalid-pull-reader",
			component: &MetricReader{
				Pull: &PullMetricReader{},
			},
			err: errInvalidExporter,
		},
		{
			name: "valid-metric-reader-valid-pull-reader-valid-prometheus-exporter",
			component: &MetricReader{
				Pull: &PullMetricReader{
					Exporter: MetricExporter{
						Prometheus: &Prometheus{},
					},
				},
			},
		},
		{
			name: "valid-metric-reader-invalid-pull-periodic-valid-otlp-exporter",
			component: &MetricReader{
				Periodic: &PeriodicMetricReader{
					Exporter: MetricExporter{
						OTLP: &OTLPMetric{},
					},
				},
			},
		},
		{
			name: "valid-metric-reader-invalid-pull-periodic-valid-console-exporter",
			component: &MetricReader{
				Periodic: &PeriodicMetricReader{
					Exporter: MetricExporter{
						Console: Console{},
					},
				},
			},
		},
		{
			name: "valid-metric-reader-invalid-pull-periodic",
			component: &MetricReader{
				Periodic: &PeriodicMetricReader{},
			},
			err: errInvalidExporter,
		},
	}

	for _, tt := range tests {
		require.Equal(t, tt.err, tt.component.Validate())
	}
}
