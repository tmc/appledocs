// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASPasskeyCredentialRequest */


/* debug [class_header]: Header for ASPasskeyCredentialRequest */
// The class instance for the [PasskeyCredentialRequest] class.
var (
	PasskeyCredentialRequestClass     _PasskeyCredentialRequestClass
	PasskeyCredentialRequestClassOnce sync.Once
)

func getPasskeyCredentialRequestClass() _PasskeyCredentialRequestClass {
	PasskeyCredentialRequestClassOnce.Do(func() {
		PasskeyCredentialRequestClass = _PasskeyCredentialRequestClass{objc.GetClass("ASPasskeyCredentialRequest")}
	})
	return PasskeyCredentialRequestClass
}

type _PasskeyCredentialRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PasskeyCredentialRequest */
// An interface definition for the [PasskeyCredentialRequest] class.
type IPasskeyCredentialRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PasskeyCredentialRequest */
	// properties:
	AssertionExtensionInput() IASPasskeyAssertionCredentialExtensionInput
	ClientDataHash() objc.IObject /* cross-framework: NSData */
	ExcludedCredentials() []AuthorizationPlatformPublicKeyCredentialDescriptor
	RegistrationExtensionInput() IASPasskeyRegistrationCredentialExtensionInput
	SupportedAlgorithms() []foundation.Number
	UserVerificationPreference() AuthorizationPublicKeyCredentialUserVerificationPreference /* typedef */
	SetUserVerificationPreference(value AuthorizationPublicKeyCredentialUserVerificationPreference /* typedef */)
	ExtensionInput() PasskeyCredentialExtensionInput /* not a class type */
	SetExtensionInput(value PasskeyCredentialExtensionInput /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PasskeyCredentialRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PasskeyCredentialRequest */
// Alloc allocates a new instance without initialization.
func (pc _PasskeyCredentialRequestClass) Alloc() PasskeyCredentialRequest {
	rv := objc.Send[PasskeyCredentialRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PasskeyCredentialRequestClass) New() PasskeyCredentialRequest {
	rv := objc.Send[PasskeyCredentialRequest](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasskeyCredentialRequest) Init() PasskeyCredentialRequest {
	rv := objc.Send[PasskeyCredentialRequest](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasskeyCredentialRequest) Autorelease() PasskeyCredentialRequest {
	rv := objc.Send[PasskeyCredentialRequest](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasskeyCredentialRequest creates a new PasskeyCredentialRequest instance.
func NewPasskeyCredentialRequest() PasskeyCredentialRequest {
	return getPasskeyCredentialRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PasskeyCredentialRequest */
// A class that represents a request to supply a passkey credential.


// A class that represents a request to supply a passkey credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequest
type PasskeyCredentialRequest struct {
	objectivec.Object
}

// PasskeyCredentialRequestFrom constructs a [PasskeyCredentialRequest] from an unsafe.Pointer.
//
// A class that represents a request to supply a passkey credential.
func PasskeyCredentialRequestFrom(ptr unsafe.Pointer) PasskeyCredentialRequest {
	return PasskeyCredentialRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PasskeyCredentialRequest */

// Initializes a passkey credential request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequest/initWithCredentialIdentity:clientDataHash:userVerificationPreference:supportedAlgorithms:
func NewPasskeyCredentialRequestWithCredentialIdentityClientDataHashUserVerificationPreferenceSupportedAlgorithms(credentialIdentity IASPasskeyCredentialIdentity, clientDataHash objc.IObject /* cross-framework: NSData */, userVerificationPreference AuthorizationPublicKeyCredentialUserVerificationPreference /* typedef */, supportedAlgorithms []foundation.Number) PasskeyCredentialRequest {
	instance := getPasskeyCredentialRequestClass().Alloc()
	rv := objc.Send[PasskeyCredentialRequest](instance.ID, objc.Sel("initWithCredentialIdentity:clientDataHash:userVerificationPreference:supportedAlgorithms:"), credentialIdentity, clientDataHash, userVerificationPreference, supportedAlgorithms)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPasskeyCredentialRequestWithCredentialIdentityClientDataHashUserVerificationPreferenceSupportedAlgorithms */


// Initializes an instance of ASPasskeyCredentialRequest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequest/initWithCredentialIdentity:clientDataHash:userVerificationPreference:supportedAlgorithms:assertionExtensionInput:
func NewPasskeyCredentialRequestWithCredentialIdentityClientDataHashUserVerificationPreferenceSupportedAlgorithmsAssertionExtensionInput(credentialIdentity IASPasskeyCredentialIdentity, clientDataHash objc.IObject /* cross-framework: NSData */, userVerificationPreference AuthorizationPublicKeyCredentialUserVerificationPreference /* typedef */, supportedAlgorithms []foundation.Number, assertionExtensionInput IASPasskeyAssertionCredentialExtensionInput) PasskeyCredentialRequest {
	instance := getPasskeyCredentialRequestClass().Alloc()
	rv := objc.Send[PasskeyCredentialRequest](instance.ID, objc.Sel("initWithCredentialIdentity:clientDataHash:userVerificationPreference:supportedAlgorithms:assertionExtensionInput:"), credentialIdentity, clientDataHash, userVerificationPreference, supportedAlgorithms, assertionExtensionInput)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPasskeyCredentialRequestWithCredentialIdentityClientDataHashUserVerificationPreferenceSupportedAlgorithmsAssertionExtensionInput */


// Initializes an instance of ASPasskeyCredentialRequest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequest/initWithCredentialIdentity:clientDataHash:userVerificationPreference:supportedAlgorithms:registrationExtensionInput:
func NewPasskeyCredentialRequestWithCredentialIdentityClientDataHashUserVerificationPreferenceSupportedAlgorithmsRegistrationExtensionInput(credentialIdentity IASPasskeyCredentialIdentity, clientDataHash objc.IObject /* cross-framework: NSData */, userVerificationPreference AuthorizationPublicKeyCredentialUserVerificationPreference /* typedef */, supportedAlgorithms []foundation.Number, registrationExtensionInput IASPasskeyRegistrationCredentialExtensionInput) PasskeyCredentialRequest {
	instance := getPasskeyCredentialRequestClass().Alloc()
	rv := objc.Send[PasskeyCredentialRequest](instance.ID, objc.Sel("initWithCredentialIdentity:clientDataHash:userVerificationPreference:supportedAlgorithms:registrationExtensionInput:"), credentialIdentity, clientDataHash, userVerificationPreference, supportedAlgorithms, registrationExtensionInput)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPasskeyCredentialRequestWithCredentialIdentityClientDataHashUserVerificationPreferenceSupportedAlgorithmsRegistrationExtensionInput */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PasskeyCredentialRequest */

// Initializes a passkey credential request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequest/init(credentialIdentity:clientDataHash:userVerificationPreference:supportedAlgorithms:)-1jihy
func (pc _PasskeyCredentialRequestClass) RequestWithCredentialIdentityClientDataHashUserVerificationPreferenceSupportedAlgorithms(credentialIdentity IASPasskeyCredentialIdentity, clientDataHash objc.IObject /* cross-framework: NSData */, userVerificationPreference AuthorizationPublicKeyCredentialUserVerificationPreference /* typedef */, supportedAlgorithms []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("requestWithCredentialIdentity:clientDataHash:userVerificationPreference:supportedAlgorithms:"), credentialIdentity, clientDataHash, userVerificationPreference, supportedAlgorithms)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RequestWithCredentialIdentityClientDataHashUserVerificationPreferenceSupportedAlgorithms) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PasskeyCredentialRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PasskeyCredentialRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PasskeyCredentialRequest */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequest/assertionExtensionInput
func (p_ PasskeyCredentialRequest) AssertionExtensionInput() IASPasskeyAssertionCredentialExtensionInput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionInput](p_.ID, objc.Sel("assertionExtensionInput"))
	return rv
}/* debug [instance_properties/getter]: assertionExtensionInput */


// The hash of the client data for this assertion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequest/clientDataHash
func (p_ PasskeyCredentialRequest) ClientDataHash() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("clientDataHash"))
	return rv
}/* debug [instance_properties/getter]: clientDataHash */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequest/excludedCredentials
func (p_ PasskeyCredentialRequest) ExcludedCredentials() []AuthorizationPlatformPublicKeyCredentialDescriptor {
	rv := objc.Send[[]AuthorizationPlatformPublicKeyCredentialDescriptor](p_.ID, objc.Sel("excludedCredentials"))
	return rv
}/* debug [instance_properties/getter]: excludedCredentials */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequest/registrationExtensionInput
func (p_ PasskeyCredentialRequest) RegistrationExtensionInput() IASPasskeyRegistrationCredentialExtensionInput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionInput](p_.ID, objc.Sel("registrationExtensionInput"))
	return rv
}/* debug [instance_properties/getter]: registrationExtensionInput */


// A list of cryptographic signature algorithms that the relying party supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequest/supportedAlgorithms-27z68
func (p_ PasskeyCredentialRequest) SupportedAlgorithms() []foundation.Number {
	rv := objc.Send[[]foundation.Number](p_.ID, objc.Sel("supportedAlgorithms"))
	return rv
}/* debug [instance_properties/getter]: supportedAlgorithms */


// The relying party’s user verification preference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequest/userVerificationPreference
func (p_ PasskeyCredentialRequest) UserVerificationPreference() AuthorizationPublicKeyCredentialUserVerificationPreference /* typedef */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("userVerificationPreference"))
	return rv
}/* debug [instance_properties/getter]: userVerificationPreference */


// The relying party’s user verification preference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequest/userVerificationPreference
func (p_ PasskeyCredentialRequest) SetUserVerificationPreference(value AuthorizationPublicKeyCredentialUserVerificationPreference /* typedef */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserVerificationPreference:"), value)
}/* debug [instance_properties/setter]: userVerificationPreference */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeycredentialrequest/extensioninput
func (p_ PasskeyCredentialRequest) ExtensionInput() PasskeyCredentialExtensionInput /* not a class type */ {
	rv := objc.Send[PasskeyCredentialExtensionInput](p_.ID, objc.Sel("extensionInput"))
	return rv
}/* debug [instance_properties/getter]: extensionInput */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeycredentialrequest/extensioninput
func (p_ PasskeyCredentialRequest) SetExtensionInput(value PasskeyCredentialExtensionInput /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setExtensionInput:"), value)
}/* debug [instance_properties/setter]: extensionInput */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASPasskeyCredentialRequest */


