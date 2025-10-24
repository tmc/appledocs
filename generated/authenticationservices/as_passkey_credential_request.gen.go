// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [PasskeyCredentialRequest] class.
type IPasskeyCredentialRequest interface {
	objectivec.IObject
	// properties:
	ClientDataHash() objc.IObject /* cross-framework: Data */
	SetClientDataHash(value objc.IObject /* cross-framework: Data */)
	ExcludedCredentials() AuthorizationPlatformPublicKeyCredentialDescriptor /* not a class type */
	SetExcludedCredentials(value AuthorizationPlatformPublicKeyCredentialDescriptor /* not a class type */)
	ExtensionInput() PasskeyCredentialExtensionInput /* not a class type */
	SetExtensionInput(value PasskeyCredentialExtensionInput /* not a class type */)
	SupportedAlgorithms() COSEAlgorithmIdentifier /* not a class type */
	SetSupportedAlgorithms(value COSEAlgorithmIdentifier /* not a class type */)
	UserVerificationPreference() AuthorizationPublicKeyCredentialUserVerificationPreference /* not a class type */
	SetUserVerificationPreference(value AuthorizationPublicKeyCredentialUserVerificationPreference /* not a class type */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (pc _PasskeyCredentialRequestClass) Alloc() PasskeyCredentialRequest {
	rv := objc.Send[PasskeyCredentialRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The hash of the client data for this assertion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeycredentialrequest/clientdatahash
func (p_ PasskeyCredentialRequest) ClientDataHash() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("clientDataHash"))
	return rv
}


// The hash of the client data for this assertion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeycredentialrequest/clientdatahash
func (p_ PasskeyCredentialRequest) SetClientDataHash(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setClientDataHash:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeycredentialrequest/excludedcredentials
func (p_ PasskeyCredentialRequest) ExcludedCredentials() AuthorizationPlatformPublicKeyCredentialDescriptor /* not a class type */ {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialDescriptor](p_.ID, objc.Sel("excludedCredentials"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeycredentialrequest/excludedcredentials
func (p_ PasskeyCredentialRequest) SetExcludedCredentials(value AuthorizationPlatformPublicKeyCredentialDescriptor /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setExcludedCredentials:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeycredentialrequest/extensioninput
func (p_ PasskeyCredentialRequest) ExtensionInput() PasskeyCredentialExtensionInput /* not a class type */ {
	rv := objc.Send[PasskeyCredentialExtensionInput](p_.ID, objc.Sel("extensionInput"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeycredentialrequest/extensioninput
func (p_ PasskeyCredentialRequest) SetExtensionInput(value PasskeyCredentialExtensionInput /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setExtensionInput:"), value)
}


// A list of cryptographic signature algorithms that the relying party supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeycredentialrequest/supportedalgorithms-74mad
func (p_ PasskeyCredentialRequest) SupportedAlgorithms() COSEAlgorithmIdentifier /* not a class type */ {
	rv := objc.Send[COSEAlgorithmIdentifier](p_.ID, objc.Sel("supportedAlgorithms"))
	return rv
}


// A list of cryptographic signature algorithms that the relying party supports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeycredentialrequest/supportedalgorithms-74mad
func (p_ PasskeyCredentialRequest) SetSupportedAlgorithms(value COSEAlgorithmIdentifier /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSupportedAlgorithms:"), value)
}


// The relying party’s user verification preference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeycredentialrequest/userverificationpreference
func (p_ PasskeyCredentialRequest) UserVerificationPreference() AuthorizationPublicKeyCredentialUserVerificationPreference /* not a class type */ {
	rv := objc.Send[AuthorizationPublicKeyCredentialUserVerificationPreference](p_.ID, objc.Sel("userVerificationPreference"))
	return rv
}


// The relying party’s user verification preference.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/aspasskeycredentialrequest/userverificationpreference
func (p_ PasskeyCredentialRequest) SetUserVerificationPreference(value AuthorizationPublicKeyCredentialUserVerificationPreference /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setUserVerificationPreference:"), value)
}



