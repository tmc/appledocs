//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CXHandle


// Returns a Boolean value that indicates whether a given handle is equal to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle/isEqualToHandle:
func (c_ CXHandle) IsEqualToHandle(handle ICXHandle) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEqualToHandle:"), handle)
	return rv
}

// iOS-only properties

// The type of the handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle/type
func (c_ CXHandle) Type() CXHandleType {
	rv := objc.Send[CXHandleType](c_.ID, objc.Sel("type"))
	return rv
}

// The value of the handle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXHandle/value
func (c_ CXHandle) Value() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("value"))
	return rv
}




