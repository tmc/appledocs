//go:build darwin && ios

// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for CXProvider


// Invalidates the provider and completes all active calls with an error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/invalidate()
func (c_ CXProvider) Invalidate() {
	objc.Send[objc.ID](c_.ID, objc.Sel("invalidate"))
}

// Returns all call actions in any pending transactions of the specified class for the specified call identifier that are incomplete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/pendingCallActions(of:withCall:)
func (c_ CXProvider) PendingCallActionsOfClassWithCallUUID(callActionClass objc.Class, callUUID foundation.UUID) []CXCallAction {
	rv := objc.Send[[]CXCallAction](c_.ID, objc.Sel("pendingCallActionsOfClass:withCallUUID:"), callActionClass, callUUID)
	return rv
}

// Reports to the provider that a call with the specified identifier ended at a given date for a particular reason.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/reportCall(with:endedAt:reason:)
func (c_ CXProvider) ReportCallWithUUIDEndedAtDateReason(UUID foundation.UUID, dateEnded objc.IObject /* cross-framework: NSDate */, endedReason CXCallEndedReason) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reportCallWithUUID:endedAtDate:reason:"), UUID, dateEnded, endedReason)
}

// Reports to the provider that an active call updated its information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/reportCall(with:updated:)
func (c_ CXProvider) ReportCallWithUUIDUpdated(UUID foundation.UUID, update ICXCallUpdate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reportCallWithUUID:updated:"), UUID, update)
}

// Reports a new incoming call with the specified unique identifier to the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/reportNewIncomingCall(with:update:completion:)
func (c_ CXProvider) ReportNewIncomingCallWithUUIDUpdateCompletion(UUID foundation.UUID, update ICXCallUpdate, completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reportNewIncomingCallWithUUID:update:completion:"), UUID, update, completion)
}

// Reports to the provider that an outgoing call with the specified unique identifier finished connecting at a particular time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/reportOutgoingCall(with:connectedAt:)
func (c_ CXProvider) ReportOutgoingCallWithUUIDConnectedAtDate(UUID foundation.UUID, dateConnected objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reportOutgoingCallWithUUID:connectedAtDate:"), UUID, dateConnected)
}

// Reports to the provider that an outgoing call with the specified unique identifier started connecting at a particular time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/reportOutgoingCall(with:startedConnectingAt:)
func (c_ CXProvider) ReportOutgoingCallWithUUIDStartedConnectingAtDate(UUID foundation.UUID, dateStartedConnecting objc.IObject /* cross-framework: NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reportOutgoingCallWithUUID:startedConnectingAtDate:"), UUID, dateStartedConnecting)
}

// Sets a provider delegate, specifying an optional queue on which to execute delegate methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/setDelegate(_:queue:)
func (c_ CXProvider) SetDelegateQueue(delegate unsafe.Pointer, queue unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:queue:"), delegate, queue)
}

// iOS-only properties

// The configuration of the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/configuration
func (c_ CXProvider) Configuration() ICXProviderConfiguration {
	rv := objc.Send[CXProviderConfiguration](c_.ID, objc.Sel("configuration"))
	return rv
}
func (c_ CXProvider) SetConfiguration(value ICXProviderConfiguration) {
	c_.ID.Send(objc.RegisterName("setConfiguration:"), value)
}

// Incomplete transactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/pendingTransactions
func (c_ CXProvider) PendingTransactions() []CXTransaction {
	rv := objc.Send[[]CXTransaction](c_.ID, objc.Sel("pendingTransactions"))
	return rv
}




