package log

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSetDebug(t *testing.T) {
	Convey("Set Debug", t, func() {
		err := SetDebug()
		So(func() {}, ShouldNotPanic)
		So(err, ShouldBeNil)
	})
}
