// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// A passkey assertion credential.
//
// Create a passkey assertion credential to provide a response to a passkey authentication challenge from your credential provider extension. Call , passing your passkey assertion credential.
//
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


// A hash of the client data for this credential.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/clientdatahash
func (p_ PasskeyAssertionCredential) ClientDataHash() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("clientDataHash"))
	return rv
}


// SetClientDataHash sets the value of the clientDataHash property.
// A hash of the client data for this credential.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/clientdatahash
func (p_ PasskeyAssertionCredential) SetClientDataHash(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setClientDataHash:"), value)
}

// The relying party associated with this passkey.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/relyingparty
func (p_ PasskeyAssertionCredential) RelyingParty() string {
	rv := objc.Send[string](p_.ID, objc.Sel("relyingParty"))
	return rv
}


// SetRelyingParty sets the value of the relyingParty property.
// The relying party associated with this passkey.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/relyingparty
func (p_ PasskeyAssertionCredential) SetRelyingParty(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRelyingParty:"), objc.String(value))
}

// The authenticator data of the application that created this passkey assertion credential.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/authenticatordata
func (p_ PasskeyAssertionCredential) AuthenticatorData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("authenticatorData"))
	return rv
}


// SetAuthenticatorData sets the value of the authenticatorData property.
// The authenticator data of the application that created this passkey assertion credential.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/authenticatordata
func (p_ PasskeyAssertionCredential) SetAuthenticatorData(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAuthenticatorData:"), value)
}

// The user handle of this passkey.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/userhandle
func (p_ PasskeyAssertionCredential) UserHandle() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("userHandle"))
	return rv
}


// SetUserHandle sets the value of the userHandle property.
// The user handle of this passkey.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/userhandle
func (p_ PasskeyAssertionCredential) SetUserHandle(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserHandle:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/extensionoutput-7t6rn
func (p_ PasskeyAssertionCredential) ExtensionOutput() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("extensionOutput"))
	return rv
}


// SetExtensionOutput sets the value of the extensionOutput property.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/extensionoutput-7t6rn
func (p_ PasskeyAssertionCredential) SetExtensionOutput(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setExtensionOutput:"), value)
}

// The identifier for this credential.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/credentialid
func (p_ PasskeyAssertionCredential) CredentialID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("credentialID"))
	return rv
}


// SetCredentialID sets the value of the credentialID property.
// The identifier for this credential.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyassertioncredential/credentialid
func (p_ PasskeyAssertionCredential) SetCredentialID(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCredentialID:"), value)
}

// The cryptographic signature of this credential.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredential/signature
func (p_ PasskeyAssertionCredential) Signature() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("signature"))
	return rv
}



