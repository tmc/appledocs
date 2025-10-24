//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CXAction


// Reports the failed execution of the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAction/fail()
func (c_ CXAction) Fail() {
	objc.Send[objc.ID](c_.ID, objc.Sel("fail"))
}

// Reports the successful execution of the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAction/fulfill()
func (c_ CXAction) Fulfill() {
	objc.Send[objc.ID](c_.ID, objc.Sel("fulfill"))
}

// iOS-only properties

// A Boolean value that indicates whether the action has been performed by the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAction/isComplete
func (c_ CXAction) Complete() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("complete"))
	return rv
}

// The time after which the action cannot be completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAction/timeoutDate
func (c_ CXAction) TimeoutDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("timeoutDate"))
	return rv
}

// The unique identifier for the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAction/uuid
func (c_ CXAction) UUID() foundation.UUID {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("UUID"))
	return rv
}




