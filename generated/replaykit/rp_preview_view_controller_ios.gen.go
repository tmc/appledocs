//go:build darwin && ios

// Code generated from Apple documentation for ReplayKit. DO NOT EDIT.

package replaykit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// iOS-only methods for RPPreviewViewController


// iOS-only properties

// The type of screen that appears when the view is presented.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ReplayKit/RPPreviewViewController/mode
func (r_ RPPreviewViewController) Mode() RPPreviewViewControllerMode {
	rv := objc.Send[RPPreviewViewControllerMode](r_.ID, objc.Sel("mode"))
	return rv
}
func (r_ RPPreviewViewController) SetMode(value RPPreviewViewControllerMode) {
	r_.ID.Send(objc.RegisterName("setMode:"), value)
}





