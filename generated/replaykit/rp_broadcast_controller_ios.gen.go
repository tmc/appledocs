//go:build darwin && ios

// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for RPBroadcastController


// iOS-only properties

// The bundle ID for the selected broadcast service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPBroadcastController/broadcastExtensionBundleID
func (r_ RPBroadcastController) BroadcastExtensionBundleID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("broadcastExtensionBundleID"))
	return rv
}





