// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [PasskeyAssertionCredential] class.
type IPasskeyAssertionCredential interface {
	objectivec.IObject
	// properties:
	AuthenticatorData() objc.IObject /* cross-framework: Data */
	SetAuthenticatorData(value objc.IObject /* cross-framework: Data */)
	ClientDataHash() objc.IObject /* cross-framework: Data */
	SetClientDataHash(value objc.IObject /* cross-framework: Data */)
	CredentialID() objc.IObject /* cross-framework: Data */
	SetCredentialID(value objc.IObject /* cross-framework: Data */)
	ExtensionOutput() IPasskeyAssertionCredentialExtensionOutput
	SetExtensionOutput(value IPasskeyAssertionCredentialExtensionOutput)
	RelyingParty() objc.IObject /* cross-framework: NSString */
	SetRelyingParty(value objc.IObject /* cross-framework: NSString */)
	Signature() objc.IObject /* cross-framework: Data */
	SetSignature(value objc.IObject /* cross-framework: Data */)
	UserHandle() objc.IObject /* cross-framework: Data */
	SetUserHandle(value objc.IObject /* cross-framework: Data */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (pc _PasskeyAssertionCredentialClass) Alloc() PasskeyAssertionCredential {
	rv := objc.Send[PasskeyAssertionCredential](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The authenticator data of the application that created this passkey assertion credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/authenticatordata
func (p_ PasskeyAssertionCredential) AuthenticatorData() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("authenticatorData"))
	return rv
}


// The authenticator data of the application that created this passkey assertion credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/authenticatordata
func (p_ PasskeyAssertionCredential) SetAuthenticatorData(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAuthenticatorData:"), value)
}


// A hash of the client data for this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/clientdatahash
func (p_ PasskeyAssertionCredential) ClientDataHash() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("clientDataHash"))
	return rv
}


// A hash of the client data for this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/clientdatahash
func (p_ PasskeyAssertionCredential) SetClientDataHash(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setClientDataHash:"), value)
}


// The identifier for this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/credentialid
func (p_ PasskeyAssertionCredential) CredentialID() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("credentialID"))
	return rv
}


// The identifier for this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/credentialid
func (p_ PasskeyAssertionCredential) SetCredentialID(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCredentialID:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/extensionoutput-7t6rn
func (p_ PasskeyAssertionCredential) ExtensionOutput() IPasskeyAssertionCredentialExtensionOutput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionOutput](p_.ID, objc.Sel("extensionOutput"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/extensionoutput-7t6rn
func (p_ PasskeyAssertionCredential) SetExtensionOutput(value IPasskeyAssertionCredentialExtensionOutput) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setExtensionOutput:"), value)
}


// The relying party associated with this passkey.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/relyingparty
func (p_ PasskeyAssertionCredential) RelyingParty() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("relyingParty"))
	return rv
}


// The relying party associated with this passkey.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/relyingparty
func (p_ PasskeyAssertionCredential) SetRelyingParty(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRelyingParty:"), value)
}


// The cryptographic signature of this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/signature
func (p_ PasskeyAssertionCredential) Signature() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("signature"))
	return rv
}


// The cryptographic signature of this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/signature
func (p_ PasskeyAssertionCredential) SetSignature(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSignature:"), value)
}


// The user handle of this passkey.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/userhandle
func (p_ PasskeyAssertionCredential) UserHandle() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("userHandle"))
	return rv
}


// The user handle of this passkey.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/userhandle
func (p_ PasskeyAssertionCredential) SetUserHandle(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserHandle:"), value)
}



