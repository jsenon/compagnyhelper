// Package opentracing ...
package opentracing

import (
	"io"
	"net"
	"strings"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	myjaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/zipkin"
)

const nbrFlush time.Duration = 1

// ConfigureTracing configures OpenTracing client
func ConfigureTracing(jaeger string) (tracer opentracing.Tracer, closer io.Closer, err error) {
	var localagenthostport string

	if strings.Contains(jaeger, "http") || strings.Contains(jaeger, ":5") ||
		strings.Contains(jaeger, ":") || strings.Contains(jaeger, "/") {
		errmsg := errors.New("wrong caracters set in jaeger url, do not set http:// or port")
		return tracer, closer, errors.Wrapf(errmsg, "Unable to launch opentracing")
	}

	if net.ParseIP(jaeger) == nil {
		localAgentHostPortIP, err2 := net.LookupHost(jaeger)
		if err2 != nil {
			errmsg := errors.New("could not resolv DNS jaeger tracer")
			return tracer, closer, errors.Wrapf(errmsg, "Unable to launch opentracing")
		}

		localagenthostport = localAgentHostPortIP[0] + ":5775"
	} else {
		localagenthostport = jaeger + ":5775"
	}

	sender, err := myjaeger.NewUDPTransport(localagenthostport, 0)
	if err != nil {
		return tracer, closer, err
	}

	//B3
	zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	injector := myjaeger.TracerOptions.Injector(opentracing.HTTPHeaders, zipkinPropagator)
	extractor := myjaeger.TracerOptions.Extractor(opentracing.HTTPHeaders, zipkinPropagator)
	// Zipkin shares span ID between client and server spans; it must be enabled via the following option.
	zipkinSharedRPCSpan := myjaeger.TracerOptions.ZipkinSharedRPCSpan(true)

	tracer, closer = myjaeger.NewTracer(
		"compagnyhelper",
		myjaeger.NewConstSampler(true),
		myjaeger.NewRemoteReporter(
			sender,
			myjaeger.ReporterOptions.BufferFlushInterval(nbrFlush*time.Second)),
		injector,
		extractor,
		zipkinSharedRPCSpan,
	)
	// Important in order to reuse tracer, to extract b3
	opentracing.SetGlobalTracer(tracer)

	return tracer, closer, nil
}
