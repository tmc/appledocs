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
	// properties:
	Certificates() unsafe.Pointer
	SetCertificates(value unsafe.Pointer)
	HasPassword() bool /* primitive/slice/pointer. */
	SetHasPassword(value bool /* primitive/slice/pointer. */)
	Identity() unsafe.Pointer
	SetIdentity(value unsafe.Pointer)
	Password() string /* primitive/slice/pointer. */
	SetPassword(value string /* primitive/slice/pointer. */)
	Persistence() unsafe.Pointer
	SetPersistence(value unsafe.Pointer)
	User() string /* primitive/slice/pointer. */
	SetUser(value string /* primitive/slice/pointer. */)
	// methods:
}

// n authentication credential consisting of information specific to the type of credential and the type of persistent storage to use, if any.
//
// The URL Loading System supports password-based user credentials, certificate-based user credentials, and certificate-based server credentials. When you create a credential, you can specify it for a single request, persist it temporarily (until your app quits), or persist it permanently. Permanent persistence can be local persistence in the keychain, or synchronized persistence across the user’s devices, based on their Apple ID.


// n authentication credential consisting of information specific to the type of credential and the type of persistent storage to use, if any.
//
// [Full Topic]
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



// The intermediate certificates of the credential, if it is a client certificate credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcredential/certificates
func (u_ URLCredential) Certificates() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("certificates"))
	return rv
}


// The intermediate certificates of the credential, if it is a client certificate credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcredential/certificates
func (u_ URLCredential) SetCertificates(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setCertificates:"), value)
}


// A Boolean value that indicates whether the credential has a password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcredential/haspassword
func (u_ URLCredential) HasPassword() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](u_.ID, objc.Sel("hasPassword"))
	return rv
}


// A Boolean value that indicates whether the credential has a password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcredential/haspassword
func (u_ URLCredential) SetHasPassword(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setHasPassword:"), value)
}


// The identity of this credential if it is a client certificate credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcredential/identity
func (u_ URLCredential) Identity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("identity"))
	return rv
}


// The identity of this credential if it is a client certificate credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcredential/identity
func (u_ URLCredential) SetIdentity(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setIdentity:"), value)
}


// The credential’s password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcredential/password
func (u_ URLCredential) Password() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](u_.ID, objc.Sel("password"))
	return rv
}


// The credential’s password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcredential/password
func (u_ URLCredential) SetPassword(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPassword:"), objc.String(value))
}


// The credential’s persistence setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcredential/persistence-swift.property
func (u_ URLCredential) Persistence() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](u_.ID, objc.Sel("persistence"))
	return rv
}


// The credential’s persistence setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcredential/persistence-swift.property
func (u_ URLCredential) SetPersistence(value unsafe.Pointer) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setPersistence:"), value)
}


// The credential’s user name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcredential/user
func (u_ URLCredential) User() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](u_.ID, objc.Sel("user"))
	return rv
}


// The credential’s user name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/urlcredential/user
func (u_ URLCredential) SetUser(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](u_.ID, objc.Sel("setUser:"), objc.String(value))
}



