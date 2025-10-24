//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for CXCallAction


// iOS-only properties

// The unique identifier for the call associated with the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallAction/callUUID
func (c_ CXCallAction) CallUUID() foundation.UUID {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("callUUID"))
	return rv
}




