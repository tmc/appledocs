// Code generated from Apple documentation for LocalAuthentication. DO NOT EDIT.

package localauthentication

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Context] class.
var (
	ContextClass     _ContextClass
	ContextClassOnce sync.Once
)

func getContextClass() _ContextClass {
	ContextClassOnce.Do(func() {
		ContextClass = _ContextClass{objc.GetClass("LAContext")}
	})
	return ContextClass
}

type _ContextClass struct {
	class objc.Class
}

// An interface definition for the [Context] class.
type IContext interface {
	objectivec.IObject
	CanEvaluatePolicyError(policy Policy, error_ unsafe.Pointer) bool
	EvaluateAccessControlOperationLocalizedReasonReply(accessControl unsafe.Pointer, operation IAccessControlOperation, localizedReason string, reply unsafe.Pointer)
	EvaluatePolicyLocalizedReasonReply(policy Policy, localizedReason string, reply unsafe.Pointer)
	Invalidate()
	IsCredentialSet(type_ CredentialType) bool
	SetCredentialType(credential foundation.IData, type_ CredentialType) bool
	BiometryType() BiometryType
	DomainState() LADomainState
	EvaluatedPolicyDomainState() foundation.NSData
	InteractionNotAllowed() bool
	SetInteractionNotAllowed(value bool)
	LocalizedCancelTitle() string
	SetLocalizedCancelTitle(value string)
	LocalizedFallbackTitle() string
	SetLocalizedFallbackTitle(value string)
	LocalizedReason() string
	SetLocalizedReason(value string)
	MaxBiometryFailures() foundation.Number
	SetMaxBiometryFailures(value foundation.INumber)
	TouchIDAuthenticationAllowableReuseDuration() foundation.TimeInterval
	SetTouchIDAuthenticationAllowableReuseDuration(value foundation.ITimeInterval)
	LATouchIDAuthenticationMaximumAllowableReuseDuration() unsafe.Pointer
}

// A mechanism for evaluating authentication policies and access controls.
//
// You use an authentication context to evaluate the user’s identity, either with biometrics like Touch ID or Face ID, or by supplying the device passcode. The context handles user interaction, and also interfaces to the Secure Enclave, the underlying hardware element that manages biometric data. You create and configure the context, and ask it to carry out the authentication. You then receive an asynchronous callback, which provides an indication of authentication success or failure, and an error instance that explains the reason for a failure, if any.


// A mechanism for evaluating authentication policies and access controls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext
type Context struct {
	objectivec.Object
}

// ContextFrom constructs a [Context] from an unsafe.Pointer.
//
// A mechanism for evaluating authentication policies and access controls.
func ContextFrom(ptr unsafe.Pointer) Context {
	return Context{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ContextClass) Alloc() Context {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ContextClass) New() Context {
	rv := objc.Send[Context](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Context) Init() Context {
	rv := objc.Send[Context](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Context) Autorelease() Context {
	rv := objc.Send[Context](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewContext creates a new Context instance.
func NewContext() Context {
	return getContextClass().New()
}



// Assesses whether authentication can proceed for a given policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/canEvaluatePolicy(_:error:)
func (c_ Context) CanEvaluatePolicyError(policy Policy, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("canEvaluatePolicy:error:"), policy, error_)
	return rv
}


// Evaluates an access control for a given operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/evaluateAccessControl(_:operation:localizedReason:reply:)
func (c_ Context) EvaluateAccessControlOperationLocalizedReasonReply(accessControl unsafe.Pointer, operation IAccessControlOperation, localizedReason string, reply unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("evaluateAccessControl:operation:localizedReason:reply:"), accessControl, operation, objc.String(localizedReason), reply)
}


// Evaluates the specified policy.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/evaluatePolicy(_:localizedReason:reply:)
func (c_ Context) EvaluatePolicyLocalizedReasonReply(policy Policy, localizedReason string, reply unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("evaluatePolicy:localizedReason:reply:"), policy, objc.String(localizedReason), reply)
}


// Invalidates the authentication context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/invalidate()
func (c_ Context) Invalidate() {
	objc.Send[objc.ID](c_.ID, objc.Sel("invalidate"))
}


// Returns a Boolean value indicating whether the specified credential type is set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/isCredentialSet(_:)
func (c_ Context) IsCredentialSet(type_ CredentialType) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isCredentialSet:"), type_)
	return rv
}


// Sets an application-provided credential to be used when evaluating authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/setCredential(_:type:)
func (c_ Context) SetCredentialType(credential foundation.IData, type_ CredentialType) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("setCredential:type:"), credential, type_)
	return rv
}


// The type of biometric authentication supported by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/biometryType
func (c_ Context) BiometryType() BiometryType {
	rv := objc.Send[BiometryType](c_.ID, objc.Sel("biometryType"))
	return rv
}


// Contains authentication domain state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/domainState
func (c_ Context) DomainState() LADomainState {
	rv := objc.Send[LADomainState](c_.ID, objc.Sel("domainState"))
	return rv
}


// The current state of the evaluated policy domain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/evaluatedPolicyDomainState
func (c_ Context) EvaluatedPolicyDomainState() foundation.NSData {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("evaluatedPolicyDomainState"))
	return rv
}


// A Boolean value indicating whether authentication can be interactive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/interactionNotAllowed
func (c_ Context) InteractionNotAllowed() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("interactionNotAllowed"))
	return rv
}


// A Boolean value indicating whether authentication can be interactive.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/interactionNotAllowed
func (c_ Context) SetInteractionNotAllowed(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setInteractionNotAllowed:"), value)
}


// The localized title for the cancel button in the dialog presented to the user during authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/localizedCancelTitle
func (c_ Context) LocalizedCancelTitle() string {
	rv := objc.Send[string](c_.ID, objc.Sel("localizedCancelTitle"))
	return rv
}


// The localized title for the cancel button in the dialog presented to the user during authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/localizedCancelTitle
func (c_ Context) SetLocalizedCancelTitle(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocalizedCancelTitle:"), objc.String(value))
}


// The localized title for the fallback button in the dialog presented to the user during authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/localizedFallbackTitle
func (c_ Context) LocalizedFallbackTitle() string {
	rv := objc.Send[string](c_.ID, objc.Sel("localizedFallbackTitle"))
	return rv
}


// The localized title for the fallback button in the dialog presented to the user during authentication.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/localizedFallbackTitle
func (c_ Context) SetLocalizedFallbackTitle(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocalizedFallbackTitle:"), objc.String(value))
}


// The localized explanation for authentication shown in the dialog presented to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/localizedReason
func (c_ Context) LocalizedReason() string {
	rv := objc.Send[string](c_.ID, objc.Sel("localizedReason"))
	return rv
}


// The localized explanation for authentication shown in the dialog presented to the user.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/localizedReason
func (c_ Context) SetLocalizedReason(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLocalizedReason:"), objc.String(value))
}


// The number of biometric authentication failures after which the context falls back to another mechanism.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/maxBiometryFailures
func (c_ Context) MaxBiometryFailures() foundation.Number {
	rv := objc.Send[foundation.Number](c_.ID, objc.Sel("maxBiometryFailures"))
	return rv
}


// The number of biometric authentication failures after which the context falls back to another mechanism.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/maxBiometryFailures
func (c_ Context) SetMaxBiometryFailures(value foundation.INumber) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMaxBiometryFailures:"), value)
}


// The duration for which Touch ID authentication reuse is allowable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/touchIDAuthenticationAllowableReuseDuration
func (c_ Context) TouchIDAuthenticationAllowableReuseDuration() foundation.TimeInterval {
	rv := objc.Send[foundation.TimeInterval](c_.ID, objc.Sel("touchIDAuthenticationAllowableReuseDuration"))
	return rv
}


// The duration for which Touch ID authentication reuse is allowable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/LocalAuthentication/LAContext/touchIDAuthenticationAllowableReuseDuration
func (c_ Context) SetTouchIDAuthenticationAllowableReuseDuration(value foundation.ITimeInterval) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTouchIDAuthenticationAllowableReuseDuration:"), value)
}


// The maximum allowable reuse duration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/localauthentication/latouchidauthenticationmaximumallowablereuseduration
func (c_ Context) LATouchIDAuthenticationMaximumAllowableReuseDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("LATouchIDAuthenticationMaximumAllowableReuseDuration"))
	return rv
}



