//go:build darwin && ios

// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for Call


// iOS-only properties

// A unique identifier for the cellular call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCall/callID
func (c_ Call) CallID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("callID"))
	return rv
}

// The state of the cellular call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCall/callState
func (c_ Call) CallState() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("callState"))
	return rv
}





