//go:build darwin && ios

// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for FallDetectionManager


// Requests authorization to receive notifications about fall detection events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionManager/requestAuthorization(handler:)
func (f_ FallDetectionManager) RequestAuthorizationWithHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("requestAuthorizationWithHandler:"), handler)
}

// iOS-only properties

// The authorization status for receiving fall detection event notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionManager/authorizationStatus
func (f_ FallDetectionManager) AuthorizationStatus() AuthorizationStatus {
	rv := objc.Send[AuthorizationStatus](f_.ID, objc.Sel("authorizationStatus"))
	return rv
}

// A delegate that can receive notifications about fall detection events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionManager/delegate
func (f_ FallDetectionManager) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("delegate"))
	return rv
}
func (f_ FallDetectionManager) SetDelegate(value unsafe.Pointer) {
	f_.ID.Send(objc.RegisterName("setDelegate:"), value)
}





