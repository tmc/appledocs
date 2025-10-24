// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASPasskeyAssertionCredential */


/* debug [class_header]: Header for ASPasskeyAssertionCredential */
// The class instance for the [PasskeyAssertionCredential] class.
var (
	PasskeyAssertionCredentialClass     _PasskeyAssertionCredentialClass
	PasskeyAssertionCredentialClassOnce sync.Once
)

func getPasskeyAssertionCredentialClass() _PasskeyAssertionCredentialClass {
	PasskeyAssertionCredentialClassOnce.Do(func() {
		PasskeyAssertionCredentialClass = _PasskeyAssertionCredentialClass{objc.GetClass("ASPasskeyAssertionCredential")}
	})
	return PasskeyAssertionCredentialClass
}

type _PasskeyAssertionCredentialClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PasskeyAssertionCredential */
// An interface definition for the [PasskeyAssertionCredential] class.
type IPasskeyAssertionCredential interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PasskeyAssertionCredential */
	// properties:
	AuthenticatorData() objc.IObject /* cross-framework: NSData */
	ClientDataHash() objc.IObject /* cross-framework: NSData */
	CredentialID() objc.IObject /* cross-framework: NSData */
	ExtensionOutput() IASPasskeyAssertionCredentialExtensionOutput
	SetExtensionOutput(value IASPasskeyAssertionCredentialExtensionOutput)
	RelyingParty() objc.IObject /* cross-framework: NSString */
	Signature() objc.IObject /* cross-framework: NSData */
	UserHandle() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PasskeyAssertionCredential */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PasskeyAssertionCredential */
// Alloc allocates a new instance without initialization.
func (pc _PasskeyAssertionCredentialClass) Alloc() PasskeyAssertionCredential {
	rv := objc.Send[PasskeyAssertionCredential](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PasskeyAssertionCredentialClass) New() PasskeyAssertionCredential {
	rv := objc.Send[PasskeyAssertionCredential](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasskeyAssertionCredential) Init() PasskeyAssertionCredential {
	rv := objc.Send[PasskeyAssertionCredential](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasskeyAssertionCredential) Autorelease() PasskeyAssertionCredential {
	rv := objc.Send[PasskeyAssertionCredential](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasskeyAssertionCredential creates a new PasskeyAssertionCredential instance.
func NewPasskeyAssertionCredential() PasskeyAssertionCredential {
	return getPasskeyAssertionCredentialClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PasskeyAssertionCredential */
// A passkey assertion credential.
//
// Create a passkey assertion credential to provide a response to a passkey authentication challenge from your credential provider extension. Call , passing your passkey assertion credential.


// A passkey assertion credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredential
type PasskeyAssertionCredential struct {
	objectivec.Object
}

// PasskeyAssertionCredentialFrom constructs a [PasskeyAssertionCredential] from an unsafe.Pointer.
//
// A passkey assertion credential.
func PasskeyAssertionCredentialFrom(ptr unsafe.Pointer) PasskeyAssertionCredential {
	return PasskeyAssertionCredential{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PasskeyAssertionCredential */

// Initializes a passkey assertion credential object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredential/init(userHandle:relyingParty:signature:clientDataHash:authenticatorData:credentialID:)
func NewPasskeyAssertionCredentialWithUserHandleRelyingPartySignatureClientDataHashAuthenticatorDataCredentialID(userHandle objc.IObject /* cross-framework: NSData */, relyingParty objc.IObject /* cross-framework: NSString */, signature objc.IObject /* cross-framework: NSData */, clientDataHash objc.IObject /* cross-framework: NSData */, authenticatorData objc.IObject /* cross-framework: NSData */, credentialID objc.IObject /* cross-framework: NSData */) PasskeyAssertionCredential {
	instance := getPasskeyAssertionCredentialClass().Alloc()
	rv := objc.Send[PasskeyAssertionCredential](instance.ID, objc.Sel("initWithUserHandle:relyingParty:signature:clientDataHash:authenticatorData:credentialID:"), userHandle, relyingParty, signature, clientDataHash, authenticatorData, credentialID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPasskeyAssertionCredentialWithUserHandleRelyingPartySignatureClientDataHashAuthenticatorDataCredentialID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredential/initWithUserHandle:relyingParty:signature:clientDataHash:authenticatorData:credentialID:extensionOutput:
func NewPasskeyAssertionCredentialWithUserHandleRelyingPartySignatureClientDataHashAuthenticatorDataCredentialIDExtensionOutput(userHandle objc.IObject /* cross-framework: NSData */, relyingParty objc.IObject /* cross-framework: NSString */, signature objc.IObject /* cross-framework: NSData */, clientDataHash objc.IObject /* cross-framework: NSData */, authenticatorData objc.IObject /* cross-framework: NSData */, credentialID objc.IObject /* cross-framework: NSData */, extensionOutput IASPasskeyAssertionCredentialExtensionOutput) PasskeyAssertionCredential {
	instance := getPasskeyAssertionCredentialClass().Alloc()
	rv := objc.Send[PasskeyAssertionCredential](instance.ID, objc.Sel("initWithUserHandle:relyingParty:signature:clientDataHash:authenticatorData:credentialID:extensionOutput:"), userHandle, relyingParty, signature, clientDataHash, authenticatorData, credentialID, extensionOutput)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPasskeyAssertionCredentialWithUserHandleRelyingPartySignatureClientDataHashAuthenticatorDataCredentialIDExtensionOutput */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PasskeyAssertionCredential */

// Creates and initializes a new passkey assertion credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredential/credentialWithUserHandle:relyingParty:signature:clientDataHash:authenticatorData:credentialID:
func (pc _PasskeyAssertionCredentialClass) CredentialWithUserHandleRelyingPartySignatureClientDataHashAuthenticatorDataCredentialID(userHandle objc.IObject /* cross-framework: NSData */, relyingParty objc.IObject /* cross-framework: NSString */, signature objc.IObject /* cross-framework: NSData */, clientDataHash objc.IObject /* cross-framework: NSData */, authenticatorData objc.IObject /* cross-framework: NSData */, credentialID objc.IObject /* cross-framework: NSData */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("credentialWithUserHandle:relyingParty:signature:clientDataHash:authenticatorData:credentialID:"), userHandle, relyingParty, signature, clientDataHash, authenticatorData, credentialID)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CredentialWithUserHandleRelyingPartySignatureClientDataHashAuthenticatorDataCredentialID) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PasskeyAssertionCredential */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PasskeyAssertionCredential */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PasskeyAssertionCredential */

// The authenticator data of the application that created this passkey assertion credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredential/authenticatorData
func (p_ PasskeyAssertionCredential) AuthenticatorData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("authenticatorData"))
	return rv
}/* debug [instance_properties/getter]: authenticatorData */


// A hash of the client data for this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredential/clientDataHash
func (p_ PasskeyAssertionCredential) ClientDataHash() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("clientDataHash"))
	return rv
}/* debug [instance_properties/getter]: clientDataHash */


// The identifier for this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredential/credentialID
func (p_ PasskeyAssertionCredential) CredentialID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("credentialID"))
	return rv
}/* debug [instance_properties/getter]: credentialID */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredential/extensionOutput-46ib3
func (p_ PasskeyAssertionCredential) ExtensionOutput() IASPasskeyAssertionCredentialExtensionOutput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionOutput](p_.ID, objc.Sel("extensionOutput"))
	return rv
}/* debug [instance_properties/getter]: extensionOutput */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredential/extensionOutput-46ib3
func (p_ PasskeyAssertionCredential) SetExtensionOutput(value IASPasskeyAssertionCredentialExtensionOutput) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setExtensionOutput:"), value)
}/* debug [instance_properties/setter]: extensionOutput */


// The relying party associated with this passkey.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredential/relyingParty
func (p_ PasskeyAssertionCredential) RelyingParty() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("relyingParty"))
	return rv
}/* debug [instance_properties/getter]: relyingParty */


// The cryptographic signature of this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredential/signature
func (p_ PasskeyAssertionCredential) Signature() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("signature"))
	return rv
}/* debug [instance_properties/getter]: signature */


// The user handle of this passkey.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredential/userHandle
func (p_ PasskeyAssertionCredential) UserHandle() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("userHandle"))
	return rv
}/* debug [instance_properties/getter]: userHandle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASPasskeyAssertionCredential */


