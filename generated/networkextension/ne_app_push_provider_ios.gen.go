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

// Informs the manager about a push-to-talk message on the connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushProvider/reportPushToTalkMessage(userInfo:)
func (n_ NEAppPushProvider) ReportPushToTalkMessageWithUserInfo(userInfo objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("reportPushToTalkMessageWithUserInfo:"), userInfo)
}

// Indicates that the framework has started the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushProvider/start()
func (n_ NEAppPushProvider) Start() {
	objc.Send[objc.ID](n_.ID, objc.Sel("start"))
}

// Indicates that the framework needs to stop the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushProvider/stop(with:completionHandler:)
func (n_ NEAppPushProvider) StopWithReasonCompletionHandler(reason NEProviderStopReason, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("stopWithReason:completionHandler:"), reason, completionHandler)
}

// Tells the framework not to use the provider with an active Ethernet connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushProvider/unmatchEthernet()
func (n_ NEAppPushProvider) UnmatchEthernet() {
	objc.Send[objc.ID](n_.ID, objc.Sel("unmatchEthernet"))
}

// iOS-only properties

// A dictionary that contains current vendor-specific configuration parameters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppPushProvider/providerConfiguration
func (n_ NEAppPushProvider) ProviderConfiguration() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](n_.ID, objc.Sel("providerConfiguration"))
	return rv
}





