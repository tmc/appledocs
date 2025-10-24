//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for CXSetGroupCallAction


// iOS-only properties

// The unique identifier of the call to be grouped with the call associated with the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetGroupCallAction/callUUIDToGroupWith
func (c_ CXSetGroupCallAction) CallUUIDToGroupWith() objc.IObject /* cross-framework: UUID */ {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("callUUIDToGroupWith"))
	return rv
}
func (c_ CXSetGroupCallAction) SetCallUUIDToGroupWith(value objc.IObject /* cross-framework: UUID */) {
	c_.ID.Send(objc.RegisterName("setCallUUIDToGroupWith:"), value)
}




