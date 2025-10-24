//go:build darwin && ios

// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for DisplayLink


// iOS-only properties

// The number of frames that must pass before the display link notifies the target again.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/frameInterval
func (d_ DisplayLink) FrameInterval() int {
	rv := objc.Send[int](d_.ID, objc.Sel("frameInterval"))
	return rv
}
func (d_ DisplayLink) SetFrameInterval(value int) {
	d_.ID.Send(objc.RegisterName("setFrameInterval:"), value)
}

// A frequency your app prefers for frame updates, affecting how often the system invokes your delegate’s callback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/preferredFramesPerSecond
func (d_ DisplayLink) PreferredFramesPerSecond() int {
	rv := objc.Send[int](d_.ID, objc.Sel("preferredFramesPerSecond"))
	return rv
}
func (d_ DisplayLink) SetPreferredFramesPerSecond(value int) {
	d_.ID.Send(objc.RegisterName("setPreferredFramesPerSecond:"), value)
}




