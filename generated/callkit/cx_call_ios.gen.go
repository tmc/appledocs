//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CXCall


// Returns a Boolean value that indicates whether a given call is equal to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCall/isEqualToCall:
func (c_ CXCall) IsEqualToCall(call ICXCall) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEqualToCall:"), call)
	return rv
}

// iOS-only properties

// A Boolean value that indicates whether the call has connected.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCall/hasConnected
func (c_ CXCall) HasConnected() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasConnected"))
	return rv
}

// A Boolean value that indicates whether the call has ended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCall/hasEnded
func (c_ CXCall) HasEnded() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasEnded"))
	return rv
}

// A Boolean value that indicates whether the call is on hold.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCall/isOnHold
func (c_ CXCall) OnHold() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("onHold"))
	return rv
}

// A Boolean value that indicates whether the call is outgoing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCall/isOutgoing
func (c_ CXCall) Outgoing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("outgoing"))
	return rv
}

// The unique identifier for the call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCall/uuid
func (c_ CXCall) UUID() foundation.UUID {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("UUID"))
	return rv
}





