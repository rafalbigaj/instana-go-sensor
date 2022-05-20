// (c) Copyright IBM Corp. 2021
// (c) Copyright Instana Inc. 2016

package instana

import (
	ot "github.com/opentracing/opentracing-go"
	"time"
)

type SpanState struct {
	Service     string             `json:"service,omitempty"`
	Operation   string             `json:"op"`
	Start       time.Time          `json:"start,omitempty"`
	Duration    time.Duration      `json:"dur,omitempty"`
	Correlation EUMCorrelationData `json:"correlation,omitempty"`
	Tags        ot.Tags            `json:"tags,omitempty"`
	Logs        []ot.LogRecord     `json:"logs,omitempty"`
	ErrorCount  int                `json:"errorCount,omitempty"`
}

type TrackerWithRestore interface {
	RestoreSpan(sc ot.SpanContext, operationName string, opts ...ot.StartSpanOption) ot.Span
	RestoreSpanWithOptions(sc ot.SpanContext, operationName string, opts ot.StartSpanOptions) ot.Span
}

func (r *tracerS) RestoreSpan(sc ot.SpanContext, operationName string, opts ...ot.StartSpanOption) ot.Span {
	sso := ot.StartSpanOptions{}
	for _, o := range opts {
		o.Apply(&sso)
	}

	return r.RestoreSpanWithOptions(sc, operationName, sso)
}

func (r *tracerS) RestoreSpanWithOptions(sc ot.SpanContext, operationName string, opts ot.StartSpanOptions) ot.Span {
	startTime := opts.StartTime
	if startTime.IsZero() {
		startTime = time.Now()
	}

	ctx := sc.(SpanContext)

	return &spanS{
		context:     ctx,
		tracer:      r,
		Service:     sensor.options.Service,
		Operation:   operationName,
		Start:       startTime,
		Duration:    -1,
		Correlation: ctx.Correlation,
		Tags:        opts.Tags,
	}
}
