// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [PasskeyRegistrationCredential] class.
type IPasskeyRegistrationCredential interface {
	objectivec.IObject
}

// A passkey registration credential.
//
// Create a passkey registration credential to provide a response to a passkey registration request from your credential provider extension. Call , passing your passkey registration credential.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PasskeyRegistrationCredentialClass) Alloc() PasskeyRegistrationCredential {
	rv := objc.Send[PasskeyRegistrationCredential](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// The attestation object for this passkey.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyregistrationcredential/attestationobject
func (p_ PasskeyRegistrationCredential) AttestationObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("attestationObject"))
	return rv
}


// SetAttestationObject sets the value of the attestationObject property.
// The attestation object for this passkey.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyregistrationcredential/attestationobject
func (p_ PasskeyRegistrationCredential) SetAttestationObject(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setAttestationObject:"), value)
}

// A hash of the client data for this credential.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyregistrationcredential/clientdatahash
func (p_ PasskeyRegistrationCredential) ClientDataHash() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("clientDataHash"))
	return rv
}


// SetClientDataHash sets the value of the clientDataHash property.
// A hash of the client data for this credential.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyregistrationcredential/clientdatahash
func (p_ PasskeyRegistrationCredential) SetClientDataHash(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setClientDataHash:"), value)
}

// The identifier for this credential.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyregistrationcredential/credentialid
func (p_ PasskeyRegistrationCredential) CredentialID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("credentialID"))
	return rv
}


// SetCredentialID sets the value of the credentialID property.
// The identifier for this credential.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyregistrationcredential/credentialid
func (p_ PasskeyRegistrationCredential) SetCredentialID(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCredentialID:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyregistrationcredential/extensionoutput-2lf9m
func (p_ PasskeyRegistrationCredential) ExtensionOutput() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("extensionOutput"))
	return rv
}


// SetExtensionOutput sets the value of the extensionOutput property.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyregistrationcredential/extensionoutput-2lf9m
func (p_ PasskeyRegistrationCredential) SetExtensionOutput(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setExtensionOutput:"), value)
}

// The relying party associated with this passkey.
//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyregistrationcredential/relyingparty
func (p_ PasskeyRegistrationCredential) RelyingParty() string {
	rv := objc.Send[string](p_.ID, objc.Sel("relyingParty"))
	return rv
}


// SetRelyingParty sets the value of the relyingParty property.
// The relying party associated with this passkey.

//
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeyregistrationcredential/relyingparty
func (p_ PasskeyRegistrationCredential) SetRelyingParty(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setRelyingParty:"), objc.String(value))
}



