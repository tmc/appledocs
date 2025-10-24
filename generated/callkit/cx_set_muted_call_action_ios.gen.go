//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for CXSetMutedCallAction


// iOS-only properties

// A Boolean value that indicates whether the call is muted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetMutedCallAction/isMuted
func (c_ CXSetMutedCallAction) Muted() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("muted"))
	return rv
}
func (c_ CXSetMutedCallAction) SetMuted(value bool) {
	c_.ID.Send(objc.RegisterName("setMuted:"), value)
}




