//go:build darwin && ios

// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// iOS-only methods for MutableURLRequest


// Binds a URL request to the network interface associated with the hotspot helper command instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableURLRequest/bind(to:)
func (m_ MutableURLRequest) BindToHotspotHelperCommand(command unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("bindToHotspotHelperCommand:"), command)
}

// iOS-only properties





