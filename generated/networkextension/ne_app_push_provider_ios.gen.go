//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// iOS-only methods for NEAppPushProvider


// Indicates a periodic status check from the framework to the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushProvider/handleTimerEvent()
func (n_ NEAppPushProvider) HandleTimerEvent() {
	objc.Send[objc.ID](n_.ID, objc.Sel("handleTimerEvent"))
}

// Informs the manager about an incoming call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushProvider/reportIncomingCall(userInfo:)
func (n_ NEAppPushProvider) ReportIncomingCallWithUserInfo(userInfo objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("reportIncomingCallWithUserInfo:"), userInfo)
}

// iOS-only properties





