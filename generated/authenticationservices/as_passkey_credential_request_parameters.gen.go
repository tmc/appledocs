// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASPasskeyCredentialRequestParameters */


/* debug [class_header]: Header for ASPasskeyCredentialRequestParameters */
// The class instance for the [PasskeyCredentialRequestParameters] class.
var (
	PasskeyCredentialRequestParametersClass     _PasskeyCredentialRequestParametersClass
	PasskeyCredentialRequestParametersClassOnce sync.Once
)

func getPasskeyCredentialRequestParametersClass() _PasskeyCredentialRequestParametersClass {
	PasskeyCredentialRequestParametersClassOnce.Do(func() {
		PasskeyCredentialRequestParametersClass = _PasskeyCredentialRequestParametersClass{objc.GetClass("ASPasskeyCredentialRequestParameters")}
	})
	return PasskeyCredentialRequestParametersClass
}

type _PasskeyCredentialRequestParametersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PasskeyCredentialRequestParameters */
// An interface definition for the [PasskeyCredentialRequestParameters] class.
type IPasskeyCredentialRequestParameters interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PasskeyCredentialRequestParameters */
	// properties:
	AllowedCredentials() []foundation.Data
	ClientDataHash() objc.IObject /* cross-framework: NSData */
	ExtensionInput() IASPasskeyAssertionCredentialExtensionInput
	RelyingPartyIdentifier() objc.IObject /* cross-framework: NSString */
	UserVerificationPreference() AuthorizationPublicKeyCredentialUserVerificationPreference /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PasskeyCredentialRequestParameters */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PasskeyCredentialRequestParameters */
// Alloc allocates a new instance without initialization.
func (pc _PasskeyCredentialRequestParametersClass) Alloc() PasskeyCredentialRequestParameters {
	rv := objc.Send[PasskeyCredentialRequestParameters](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PasskeyCredentialRequestParametersClass) New() PasskeyCredentialRequestParameters {
	rv := objc.Send[PasskeyCredentialRequestParameters](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasskeyCredentialRequestParameters) Init() PasskeyCredentialRequestParameters {
	rv := objc.Send[PasskeyCredentialRequestParameters](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasskeyCredentialRequestParameters) Autorelease() PasskeyCredentialRequestParameters {
	rv := objc.Send[PasskeyCredentialRequestParameters](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasskeyCredentialRequestParameters creates a new PasskeyCredentialRequestParameters instance.
func NewPasskeyCredentialRequestParameters() PasskeyCredentialRequestParameters {
	return getPasskeyCredentialRequestParametersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PasskeyCredentialRequestParameters */
// A class that represents information about a passkey credential request.
//
// The system creates instances of this class to handle active passkey requests, and passes them to your extension by calling . Use the properties of the given request parameters object, along with the passkey credential that the person chooses, to construct a passkey credential response that you return to the system using .


// A class that represents information about a passkey credential request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequestParameters
type PasskeyCredentialRequestParameters struct {
	objectivec.Object
}

// PasskeyCredentialRequestParametersFrom constructs a [PasskeyCredentialRequestParameters] from an unsafe.Pointer.
//
// A class that represents information about a passkey credential request.
func PasskeyCredentialRequestParametersFrom(ptr unsafe.Pointer) PasskeyCredentialRequestParameters {
	return PasskeyCredentialRequestParameters{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PasskeyCredentialRequestParameters *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PasskeyCredentialRequestParameters */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PasskeyCredentialRequestParameters */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PasskeyCredentialRequestParameters */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PasskeyCredentialRequestParameters */

// A list of passkey credentials that the relying party accepts for this challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequestParameters/allowedCredentials
func (p_ PasskeyCredentialRequestParameters) AllowedCredentials() []foundation.Data {
	rv := objc.Send[[]foundation.Data](p_.ID, objc.Sel("allowedCredentials"))
	return rv
}/* debug [instance_properties/getter]: allowedCredentials */


// The client data that you sign as part of the response.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequestParameters/clientDataHash
func (p_ PasskeyCredentialRequestParameters) ClientDataHash() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("clientDataHash"))
	return rv
}/* debug [instance_properties/getter]: clientDataHash */


// Inputs for WebAuthn extensions used for passkey assertion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequestParameters/extensionInput-5ny4g
func (p_ PasskeyCredentialRequestParameters) ExtensionInput() IASPasskeyAssertionCredentialExtensionInput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionInput](p_.ID, objc.Sel("extensionInput"))
	return rv
}/* debug [instance_properties/getter]: extensionInput */


// The relying party that issues the challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequestParameters/relyingPartyIdentifier
func (p_ PasskeyCredentialRequestParameters) RelyingPartyIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("relyingPartyIdentifier"))
	return rv
}/* debug [instance_properties/getter]: relyingPartyIdentifier */


// The relying party’s preference for user verification.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyCredentialRequestParameters/userVerificationPreference
func (p_ PasskeyCredentialRequestParameters) UserVerificationPreference() AuthorizationPublicKeyCredentialUserVerificationPreference /* typedef */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("userVerificationPreference"))
	return rv
}/* debug [instance_properties/getter]: userVerificationPreference */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASPasskeyCredentialRequestParameters */



