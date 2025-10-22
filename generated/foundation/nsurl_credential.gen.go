// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [URLCredential] class.
var (
	URLCredentialClass     _URLCredentialClass
	URLCredentialClassOnce sync.Once
)

func getURLCredentialClass() _URLCredentialClass {
	URLCredentialClassOnce.Do(func() {
		URLCredentialClass = _URLCredentialClass{objc.GetClass("NSURLCredential")}
	})
	return URLCredentialClass
}

type _URLCredentialClass struct {
	class objc.Class
}

// An interface definition for the [URLCredential] class.
type IURLCredential interface {
	objectivec.IObject
	Certificates() objc.ID
	HasPassword() bool
	Identity() unsafe.Pointer
	Password() string
	Persistence() URLCredentialPersistence
	User() string
}

// n authentication credential consisting of information specific to the type of credential and the type of persistent storage to use, if any.
//
// The URL Loading System supports password-based user credentials, certificate-based user credentials, and certificate-based server credentials. When you create a credential, you can specify it for a single request, persist it temporarily (until your app quits), or persist it permanently. Permanent persistence can be local persistence in the keychain, or synchronized persistence across the user’s devices, based on their Apple ID.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential
type URLCredential struct {
	objectivec.Object
}

// URLCredentialFrom constructs a [URLCredential] from an unsafe.Pointer.
//
// n authentication credential consisting of information specific to the type of credential and the type of persistent storage to use, if any.
func URLCredentialFrom(ptr unsafe.Pointer) URLCredential {
	return URLCredential{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (uc _URLCredentialClass) Alloc() URLCredential {
	rv := objc.Send[URLCredential](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (uc _URLCredentialClass) New() URLCredential {
	rv := objc.Send[URLCredential](objc.ID(uc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (u_ URLCredential) Init() URLCredential {
	rv := objc.Send[URLCredential](u_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u_ URLCredential) Autorelease() URLCredential {
	rv := objc.Send[URLCredential](u_.ID, objc.Sel("autorelease"))
	return rv
}

// NewURLCredential creates a new URLCredential instance.
func NewURLCredential() URLCredential {
	return getURLCredentialClass().New()
}




// Creates a URL credential instance for server trust authentication with a given accepted trust.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/init(forTrust:)
func NewURLCredentialForTrust(trust unsafe.Pointer) URLCredential {
	rv := objc.Send[URLCredential](objc.ID(getURLCredentialClass().class), objc.Sel("credentialForTrust:"), trust)
	return rv
}



// Creates a URL credential instance for resolving a client certificate authentication challenge.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/init(identity:certificates:persistence:)
func NewURLCredentialWithIdentityCertificatesPersistence(identity unsafe.Pointer, certArray objectivec.IObject, persistence IURLCredentialPersistence) URLCredential {
	instance := getURLCredentialClass().Alloc()
	rv := objc.Send[URLCredential](instance.ID, objc.Sel("initWithIdentity:certificates:persistence:"), identity, certArray, persistence)
	rv.Autorelease()
	return rv
}



// Creates a URL credential instance for server trust authentication, initialized with a accepted trust.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/init(trust:)
func NewURLCredentialWithTrust(trust unsafe.Pointer) URLCredential {
	instance := getURLCredentialClass().Alloc()
	rv := objc.Send[URLCredential](instance.ID, objc.Sel("initWithTrust:"), trust)
	rv.Autorelease()
	return rv
}



// Creates a URL credential instance initialized with a given user name and password, using a given persistence setting.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/init(user:password:persistence:)
func NewURLCredentialWithUserPasswordPersistence(user string, password string, persistence IURLCredentialPersistence) URLCredential {
	instance := getURLCredentialClass().Alloc()
	rv := objc.Send[URLCredential](instance.ID, objc.Sel("initWithUser:password:persistence:"), objc.String(user), objc.String(password), persistence)
	rv.Autorelease()
	return rv
}


// Creates a URL credential instance for resolving a client certificate authentication challenge.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLCredential/credentialWithIdentity:certificates:persistence:
func (uc _URLCredentialClass) CredentialWithIdentityCertificatesPersistence(identity unsafe.Pointer, certArray objectivec.IObject, persistence IURLCredentialPersistence) URLCredential {
	rv := objc.Send[URLCredential](objc.ID(uc.class), objc.Sel("credentialWithIdentity:certificates:persistence:"), identity, certArray, persistence)
	return rv
}

// Creates a URL credential instance for internet password authentication with a given user name and password, using a given persistence setting.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLCredential/credentialWithUser:password:persistence:
func (uc _URLCredentialClass) CredentialWithUserPasswordPersistence(user string, password string, persistence IURLCredentialPersistence) URLCredential {
	rv := objc.Send[URLCredential](objc.ID(uc.class), objc.Sel("credentialWithUser:password:persistence:"), objc.String(user), objc.String(password), persistence)
	return rv
}

// Creates a URL credential instance for server trust authentication with a given accepted trust.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/init(forTrust:)
func (uc _URLCredentialClass) CredentialForTrust(trust unsafe.Pointer) URLCredential {
	rv := objc.Send[URLCredential](objc.ID(uc.class), objc.Sel("credentialForTrust:"), trust)
	return rv
}

// The intermediate certificates of the credential, if it is a client certificate credential.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/certificates
func (u_ URLCredential) Certificates() objc.ID {
	rv := objc.Send[objc.ID](u_.ID, objc.Sel("certificates"))
	return rv
}

// A Boolean value that indicates whether the credential has a password.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/hasPassword
func (u_ URLCredential) HasPassword() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("hasPassword"))
	return rv
}

// The identity of this credential if it is a client certificate credential.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/identity
func (u_ URLCredential) Identity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("identity"))
	return rv
}

// The credential’s password.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/password
func (u_ URLCredential) Password() string {
	rv := objc.Send[string](u_.ID, objc.Sel("password"))
	return rv
}

// The credential’s persistence setting.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/persistence-swift.property
func (u_ URLCredential) Persistence() URLCredentialPersistence {
	rv := objc.Send[URLCredentialPersistence](u_.ID, objc.Sel("persistence"))
	return rv
}

// The credential’s user name.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/user
func (u_ URLCredential) User() string {
	rv := objc.Send[string](u_.ID, objc.Sel("user"))
	return rv
}


