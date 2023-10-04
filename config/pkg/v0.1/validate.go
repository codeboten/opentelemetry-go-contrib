// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package config // import "go.opentelemetry.io/contrib/config"

import (
	"fmt"
)

var (
	errUnsupportedSpanProcessor = fmt.Errorf("unsupported span processor type")
	errInvalidExporter          = fmt.Errorf("invalid exporter configuration")
	errUnsupportedMetricReader  = fmt.Errorf("unsupported metric reader type")
)

// Validate checks for a valid span processor to be configured for the SpanProcessor.
func (sp *SpanProcessor) Validate() error {
	if sp.Batch != nil {
		return sp.Batch.Exporter.Validate()
	}
	if sp.Simple != nil {
		return sp.Simple.Exporter.Validate()
	}
	return errUnsupportedSpanProcessor
}

// Validate checks for valid exporters to be configured for the SpanExporter.
func (se *SpanExporter) Validate() error {
	if se.Console == nil && se.OTLP == nil && se.Zipkin == nil {
		return errInvalidExporter
	}
	return nil
}

// Validate checks for metric readers to be configured for the MetricReader.
func (mr *MetricReader) Validate() error {
	if mr.Pull != nil {
		return mr.Pull.Validate()
	}
	if mr.Periodic != nil {
		return mr.Periodic.Validate()
	}

	return errUnsupportedMetricReader
}

// Validate checks for valid exporters to be configured for the PullMetricReader.
func (pmr *PullMetricReader) Validate() error {
	if pmr.Exporter.Prometheus == nil {
		return errInvalidExporter
	}
	return nil
}

// Validate checks for valid exporters to be configured for the PeriodicMetricReader.
func (pmr *PeriodicMetricReader) Validate() error {
	if pmr.Exporter.OTLP == nil && pmr.Exporter.Console == nil {
		return errInvalidExporter
	}
	return nil
}
