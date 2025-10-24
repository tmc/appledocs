//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for CXAnswerCallAction


// Reports the successful execution of the action at the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAnswerCallAction/fulfill(withDateConnected:)
func (c_ CXAnswerCallAction) FulfillWithDateConnected(dateConnected objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fulfillWithDateConnected:"), dateConnected)
}

// iOS-only properties





