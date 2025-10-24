//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CXCallObserver


// Sets a call observer delegate, specifying an optional queue on which to execute delegate methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallObserver/setDelegate(_:queue:)
func (c_ CXCallObserver) SetDelegateQueue(delegate unsafe.Pointer, queue unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:queue:"), delegate, queue)
}

// iOS-only properties

// Returns the active calls of the telephony provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallObserver/calls
func (c_ CXCallObserver) Calls() []CXCall {
	rv := objc.Send[[]CXCall](c_.ID, objc.Sel("calls"))
	return rv
}





