//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for CXSetHeldCallAction


// iOS-only properties

// A Boolean value that indicates whether the call is placed on hold.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetHeldCallAction/isOnHold
func (c_ CXSetHeldCallAction) OnHold() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("onHold"))
	return rv
}
func (c_ CXSetHeldCallAction) SetOnHold(value bool) {
	c_.ID.Send(objc.RegisterName("setOnHold:"), value)
}




