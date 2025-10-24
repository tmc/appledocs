//go:build darwin && ios

// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for RPBroadcastActivityViewController


// iOS-only properties

// The delegate for the broadcast activity view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastActivityViewController/delegate
func (r_ RPBroadcastActivityViewController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("delegate"))
	return rv
}
func (r_ RPBroadcastActivityViewController) SetDelegate(value objc.ID) {
	r_.ID.Send(objc.RegisterName("setDelegate:"), value)
}





