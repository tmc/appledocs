// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CXProvider] class.
var (
	CXProviderClass     _CXProviderClass
	CXProviderClassOnce sync.Once
)

func getCXProviderClass() _CXProviderClass {
	CXProviderClassOnce.Do(func() {
		CXProviderClass = _CXProviderClass{objc.GetClass("CXProvider")}
	})
	return CXProviderClass
}

type _CXProviderClass struct {
	class objc.Class
}

// An interface definition for the [CXProvider] class.
type ICXProvider interface {
	objectivec.IObject
	Invalidate()
	PendingCallActionsOfClassWithCallUUID(callActionClass objc.Class, callUUID unsafe.Pointer) []CXCallAction
	ReportCallWithUUIDEndedAtDateReason(UUID unsafe.Pointer, dateEnded unsafe.Pointer, endedReason unsafe.Pointer)
	ReportCallWithUUIDUpdated(UUID unsafe.Pointer, update unsafe.Pointer)
	ReportNewIncomingCallWithUUIDUpdateCompletion(UUID unsafe.Pointer, update unsafe.Pointer, completion unsafe.Pointer)
	ReportOutgoingCallWithUUIDConnectedAtDate(UUID unsafe.Pointer, dateConnected unsafe.Pointer)
	ReportOutgoingCallWithUUIDStartedConnectingAtDate(UUID unsafe.Pointer, dateStartedConnecting unsafe.Pointer)
	SetDelegateQueue(delegate objc.ID, queue unsafe.Pointer)
}

// An object that represents a telephony provider.
//
// A object is responsible for reporting out-of-band notifications that occur to the system. A VoIP app should create only one instance of and store it for use globally. A object is initialized with a object to specify the behavior and capabilities of calls. Each provider can specify an object conforming to the protocol to respond to events, such as the call starting, the call being put on hold, or the provider’s audio session being activated.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider
type CXProvider struct {
	objectivec.Object
}

// CXProviderFrom constructs a [CXProvider] from an unsafe.Pointer.
//
// An object that represents a telephony provider.
func CXProviderFrom(ptr unsafe.Pointer) CXProvider {
	return CXProvider{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CXProviderClass) Alloc() CXProvider {
	rv := objc.Send[CXProvider](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXProviderClass) New() CXProvider {
	rv := objc.Send[CXProvider](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXProvider) Init() CXProvider {
	rv := objc.Send[CXProvider](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXProvider) Autorelease() CXProvider {
	rv := objc.Send[CXProvider](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXProvider creates a new CXProvider instance.
func NewCXProvider() CXProvider {
	return getCXProviderClass().New()
}




// Initializes a new provider with the specified configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/init(configuration:)
func NewCXProviderWithConfiguration(configuration unsafe.Pointer) CXProvider {
	instance := getCXProviderClass().Alloc()
	rv := objc.Send[CXProvider](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	rv.Autorelease()
	return rv
}


// Reports a new incoming call after your notification service extension decrypts a VoIP call request.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/reportNewIncomingVoIPPushPayload(_:completion:)
func (cc _CXProviderClass) ReportNewIncomingVoIPPushPayloadCompletion(dictionaryPayload objc.ID, completion unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("reportNewIncomingVoIPPushPayload:completion:"), dictionaryPayload, completion)
}

// Invalidates the provider and completes all active calls with an error.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/invalidate()
func (c_ CXProvider) Invalidate() {
	objc.Send[objc.ID](c_.ID, objc.Sel("invalidate"))
}

// Returns all call actions in any pending transactions of the specified class for the specified call identifier that are incomplete.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/pendingCallActions(of:withCall:)
func (c_ CXProvider) PendingCallActionsOfClassWithCallUUID(callActionClass objc.Class, callUUID unsafe.Pointer) []CXCallAction {
	rv := objc.Send[[]CXCallAction](c_.ID, objc.Sel("pendingCallActionsOfClass:withCallUUID:"), callActionClass, callUUID)
	return rv
}

// Reports to the provider that a call with the specified identifier ended at a given date for a particular reason.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/reportCall(with:endedAt:reason:)
func (c_ CXProvider) ReportCallWithUUIDEndedAtDateReason(UUID unsafe.Pointer, dateEnded unsafe.Pointer, endedReason unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reportCallWithUUID:endedAtDate:reason:"), UUID, dateEnded, endedReason)
}

// Reports to the provider that an active call updated its information.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/reportCall(with:updated:)
func (c_ CXProvider) ReportCallWithUUIDUpdated(UUID unsafe.Pointer, update unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reportCallWithUUID:updated:"), UUID, update)
}

// Reports a new incoming call with the specified unique identifier to the provider.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/reportNewIncomingCall(with:update:completion:)
func (c_ CXProvider) ReportNewIncomingCallWithUUIDUpdateCompletion(UUID unsafe.Pointer, update unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reportNewIncomingCallWithUUID:update:completion:"), UUID, update, completion)
}

// Reports to the provider that an outgoing call with the specified unique identifier finished connecting at a particular time.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/reportOutgoingCall(with:connectedAt:)
func (c_ CXProvider) ReportOutgoingCallWithUUIDConnectedAtDate(UUID unsafe.Pointer, dateConnected unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reportOutgoingCallWithUUID:connectedAtDate:"), UUID, dateConnected)
}

// Reports to the provider that an outgoing call with the specified unique identifier started connecting at a particular time.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/reportOutgoingCall(with:startedConnectingAt:)
func (c_ CXProvider) ReportOutgoingCallWithUUIDStartedConnectingAtDate(UUID unsafe.Pointer, dateStartedConnecting unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("reportOutgoingCallWithUUID:startedConnectingAtDate:"), UUID, dateStartedConnecting)
}

// Sets a provider delegate, specifying an optional queue on which to execute delegate methods.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/setDelegate(_:queue:)
func (c_ CXProvider) SetDelegateQueue(delegate objc.ID, queue unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:queue:"), delegate, queue)
}

// The domain for CallKit errors.
//
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxerrordomain
func (c_ CXProvider) CXErrorDomain() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CXErrorDomain"))
	return rv
}

// The domain for errors that occur during incoming calls.
//
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxerrordomainincomingcall
func (c_ CXProvider) CXErrorDomainIncomingCall() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CXErrorDomainIncomingCall"))
	return rv
}

// The configuration of the provider.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/configuration
func (c_ CXProvider) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("configuration"))
	return rv
}


// SetConfiguration sets the value of the configuration property.
// The configuration of the provider.

//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/configuration
func (c_ CXProvider) SetConfiguration(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConfiguration:"), value)
}

// Incomplete transactions.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXProvider/pendingTransactions
func (c_ CXProvider) PendingTransactions() []CXTransaction {
	rv := objc.Send[[]CXTransaction](c_.ID, objc.Sel("pendingTransactions"))
	return rv
}


