//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for CXPlayDTMFCallAction


// iOS-only properties

// The digits tapped by the user into the in-call keypad or included in the dial string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXPlayDTMFCallAction/digits
func (c_ CXPlayDTMFCallAction) Digits() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("digits"))
	return rv
}
func (c_ CXPlayDTMFCallAction) SetDigits(value objc.IObject /* cross-framework: NSString */) {
	c_.ID.Send(objc.RegisterName("setDigits:"), value)
}

// The type of the call action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXPlayDTMFCallAction/type
func (c_ CXPlayDTMFCallAction) Type() CXPlayDTMFCallActionType {
	rv := objc.Send[CXPlayDTMFCallActionType](c_.ID, objc.Sel("type"))
	return rv
}
func (c_ CXPlayDTMFCallAction) SetType(value CXPlayDTMFCallActionType) {
	c_.ID.Send(objc.RegisterName("setType:"), value)
}




