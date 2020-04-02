package opentracing

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfigureTracing(t *testing.T) {
	Convey("Test tracing:", t, func() {
		jaegerurl := "google.com"
		Convey("right url tracing closer", func() {
			_, _, err := ConfigureTracing(jaegerurl)
			So(func() {}, ShouldNotPanic)
			So(err, ShouldBeNil)
		})
		Convey("wrong url tracing", func() {
			jaegerurl := "http://opentracing.toto.svc.local"
			_, closer, err := ConfigureTracing(jaegerurl)
			So(closer, ShouldBeEmpty)
			So(err.Error(), ShouldResemble,
				"Unable to launch opentracing: wrong caracters set in jaeger url, do not set http:// or port")
		})
		Convey("right IP url tracing closer", func() {
			jaegerurl := "127.0.0.1"
			_, _, err := ConfigureTracing(jaegerurl)
			So(func() {}, ShouldNotPanic)
			So(err, ShouldBeNil)
		})
		Convey("wrong url resolution for tracing", func() {
			jaegerurl := "todfgdfdsffdg.com"
			_, _, err := ConfigureTracing(jaegerurl)
			So(err, ShouldNotBeNil)
			So(err.Error(), ShouldResemble, "Unable to launch opentracing: could not resolv DNS jaeger tracer")
		})
	})
}
