//go:build darwin && ios

// Code generated from Apple documentation for AVKit. DO NOT EDIT.

package avkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CaptureEventInteraction


// iOS-only properties

// A Boolean value that indicates whether this capture event interaction is in an enabled state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVKit/AVCaptureEventInteraction/isEnabled
func (c_ CaptureEventInteraction) Enabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("enabled"))
	return rv
}
func (c_ CaptureEventInteraction) SetEnabled(value bool) {
	c_.ID.Send(objc.RegisterName("setEnabled:"), value)
}




