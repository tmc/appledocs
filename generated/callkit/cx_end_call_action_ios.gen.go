//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for CXEndCallAction


// Reports the successful execution of the action at the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXEndCallAction/fulfill(withDateEnded:)
func (c_ CXEndCallAction) FulfillWithDateEnded(dateEnded objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fulfillWithDateEnded:"), dateEnded)
}

// iOS-only properties





