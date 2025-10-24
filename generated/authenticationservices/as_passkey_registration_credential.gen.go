// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASPasskeyRegistrationCredential */


/* debug [class_header]: Header for ASPasskeyRegistrationCredential */
// The class instance for the [PasskeyRegistrationCredential] class.
var (
	PasskeyRegistrationCredentialClass     _PasskeyRegistrationCredentialClass
	PasskeyRegistrationCredentialClassOnce sync.Once
)

func getPasskeyRegistrationCredentialClass() _PasskeyRegistrationCredentialClass {
	PasskeyRegistrationCredentialClassOnce.Do(func() {
		PasskeyRegistrationCredentialClass = _PasskeyRegistrationCredentialClass{objc.GetClass("ASPasskeyRegistrationCredential")}
	})
	return PasskeyRegistrationCredentialClass
}

type _PasskeyRegistrationCredentialClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PasskeyRegistrationCredential */
// An interface definition for the [PasskeyRegistrationCredential] class.
type IPasskeyRegistrationCredential interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PasskeyRegistrationCredential */
	// properties:
	AttestationObject() objc.IObject /* cross-framework: NSData */
	ClientDataHash() objc.IObject /* cross-framework: NSData */
	CredentialID() objc.IObject /* cross-framework: NSData */
	ExtensionOutput() IASPasskeyRegistrationCredentialExtensionOutput
	SetExtensionOutput(value IASPasskeyRegistrationCredentialExtensionOutput)
	RelyingParty() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PasskeyRegistrationCredential */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PasskeyRegistrationCredential */
// Alloc allocates a new instance without initialization.
func (pc _PasskeyRegistrationCredentialClass) Alloc() PasskeyRegistrationCredential {
	rv := objc.Send[PasskeyRegistrationCredential](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PasskeyRegistrationCredentialClass) New() PasskeyRegistrationCredential {
	rv := objc.Send[PasskeyRegistrationCredential](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasskeyRegistrationCredential) Init() PasskeyRegistrationCredential {
	rv := objc.Send[PasskeyRegistrationCredential](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasskeyRegistrationCredential) Autorelease() PasskeyRegistrationCredential {
	rv := objc.Send[PasskeyRegistrationCredential](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasskeyRegistrationCredential creates a new PasskeyRegistrationCredential instance.
func NewPasskeyRegistrationCredential() PasskeyRegistrationCredential {
	return getPasskeyRegistrationCredentialClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PasskeyRegistrationCredential */
// A passkey registration credential.
//
// Create a passkey registration credential to provide a response to a passkey registration request from your credential provider extension. Call , passing your passkey registration credential.


// A passkey registration credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredential
type PasskeyRegistrationCredential struct {
	objectivec.Object
}

// PasskeyRegistrationCredentialFrom constructs a [PasskeyRegistrationCredential] from an unsafe.Pointer.
//
// A passkey registration credential.
func PasskeyRegistrationCredentialFrom(ptr unsafe.Pointer) PasskeyRegistrationCredential {
	return PasskeyRegistrationCredential{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PasskeyRegistrationCredential */

// Initializes a passkey registration credential object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredential/init(relyingParty:clientDataHash:credentialID:attestationObject:)
func NewPasskeyRegistrationCredentialWithRelyingPartyClientDataHashCredentialIDAttestationObject(relyingParty objc.IObject /* cross-framework: NSString */, clientDataHash objc.IObject /* cross-framework: NSData */, credentialID objc.IObject /* cross-framework: NSData */, attestationObject objc.IObject /* cross-framework: NSData */) PasskeyRegistrationCredential {
	instance := getPasskeyRegistrationCredentialClass().Alloc()
	rv := objc.Send[PasskeyRegistrationCredential](instance.ID, objc.Sel("initWithRelyingParty:clientDataHash:credentialID:attestationObject:"), relyingParty, clientDataHash, credentialID, attestationObject)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPasskeyRegistrationCredentialWithRelyingPartyClientDataHashCredentialIDAttestationObject */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredential/initWithRelyingParty:clientDataHash:credentialID:attestationObject:extensionOutput:
func NewPasskeyRegistrationCredentialWithRelyingPartyClientDataHashCredentialIDAttestationObjectExtensionOutput(relyingParty objc.IObject /* cross-framework: NSString */, clientDataHash objc.IObject /* cross-framework: NSData */, credentialID objc.IObject /* cross-framework: NSData */, attestationObject objc.IObject /* cross-framework: NSData */, extensionOutput IASPasskeyRegistrationCredentialExtensionOutput) PasskeyRegistrationCredential {
	instance := getPasskeyRegistrationCredentialClass().Alloc()
	rv := objc.Send[PasskeyRegistrationCredential](instance.ID, objc.Sel("initWithRelyingParty:clientDataHash:credentialID:attestationObject:extensionOutput:"), relyingParty, clientDataHash, credentialID, attestationObject, extensionOutput)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPasskeyRegistrationCredentialWithRelyingPartyClientDataHashCredentialIDAttestationObjectExtensionOutput */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PasskeyRegistrationCredential */

// Creates and initializes a new passkey registration credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredential/credentialWithRelyingParty:clientDataHash:credentialID:attestationObject:
func (pc _PasskeyRegistrationCredentialClass) CredentialWithRelyingPartyClientDataHashCredentialIDAttestationObject(relyingParty objc.IObject /* cross-framework: NSString */, clientDataHash objc.IObject /* cross-framework: NSData */, credentialID objc.IObject /* cross-framework: NSData */, attestationObject objc.IObject /* cross-framework: NSData */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("credentialWithRelyingParty:clientDataHash:credentialID:attestationObject:"), relyingParty, clientDataHash, credentialID, attestationObject)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CredentialWithRelyingPartyClientDataHashCredentialIDAttestationObject) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PasskeyRegistrationCredential */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PasskeyRegistrationCredential */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PasskeyRegistrationCredential */

// The attestation object for this passkey.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredential/attestationObject
func (p_ PasskeyRegistrationCredential) AttestationObject() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("attestationObject"))
	return rv
}/* debug [instance_properties/getter]: attestationObject */


// A hash of the client data for this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredential/clientDataHash
func (p_ PasskeyRegistrationCredential) ClientDataHash() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("clientDataHash"))
	return rv
}/* debug [instance_properties/getter]: clientDataHash */


// The identifier for this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredential/credentialID
func (p_ PasskeyRegistrationCredential) CredentialID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("credentialID"))
	return rv
}/* debug [instance_properties/getter]: credentialID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredential/extensionOutput-95gvu
func (p_ PasskeyRegistrationCredential) ExtensionOutput() IASPasskeyRegistrationCredentialExtensionOutput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionOutput](p_.ID, objc.Sel("extensionOutput"))
	return rv
}/* debug [instance_properties/getter]: extensionOutput */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredential/extensionOutput-95gvu
func (p_ PasskeyRegistrationCredential) SetExtensionOutput(value IASPasskeyRegistrationCredentialExtensionOutput) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setExtensionOutput:"), value)
}/* debug [instance_properties/setter]: extensionOutput */


// The relying party associated with this passkey.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredential/relyingParty
func (p_ PasskeyRegistrationCredential) RelyingParty() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("relyingParty"))
	return rv
}/* debug [instance_properties/getter]: relyingParty */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASPasskeyRegistrationCredential */


