// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSURLCredential */


/* debug [class_header]: Header for NSURLCredential */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for URLCredential */
// An interface definition for the [URLCredential] class.
type IURLCredential interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for URLCredential */
	// properties:
	Certificates() IArray
	HasPassword() bool
	Identity() objectivec.IObject
	Password() IString
	Persistence() URLCredentialPersistence
	User() IString
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for URLCredential */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for URLCredential */
// Alloc allocates a new instance without initialization.
func (uc _URLCredentialClass) Alloc() URLCredential {
	rv := objc.Send[URLCredential](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for URLCredential */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for URLCredential */

// Creates a URL credential instance for server trust authentication with a given accepted trust.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/init(forTrust:)
func NewURLCredentialForTrust(trust objectivec.IObject) URLCredential {
	rv := objc.Send[URLCredential](objc.ID(getURLCredentialClass().class), objc.Sel("credentialForTrust:"), trust)
	return rv
}/* debug [class_init_methods/constructor]: NewURLCredentialForTrust */


// Creates a URL credential instance for resolving a client certificate authentication challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/init(identity:certificates:persistence:)
func NewURLCredentialWithIdentityCertificatesPersistence(identity objectivec.IObject, certArray IArray, persistence URLCredentialPersistence) URLCredential {
	instance := getURLCredentialClass().Alloc()
	rv := objc.Send[URLCredential](instance.ID, objc.Sel("initWithIdentity:certificates:persistence:"), identity, certArray, persistence)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLCredentialWithIdentityCertificatesPersistence */


// Creates a URL credential instance for server trust authentication, initialized with a accepted trust.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/init(trust:)
func NewURLCredentialWithTrust(trust objectivec.IObject) URLCredential {
	instance := getURLCredentialClass().Alloc()
	rv := objc.Send[URLCredential](instance.ID, objc.Sel("initWithTrust:"), trust)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLCredentialWithTrust */


// Creates a URL credential instance initialized with a given user name and password, using a given persistence setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/init(user:password:persistence:)
func NewURLCredentialWithUserPasswordPersistence(user IString, password IString, persistence URLCredentialPersistence) URLCredential {
	instance := getURLCredentialClass().Alloc()
	rv := objc.Send[URLCredential](instance.ID, objc.Sel("initWithUser:password:persistence:"), user, password, persistence)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewURLCredentialWithUserPasswordPersistence */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for URLCredential */

// Creates a URL credential instance for resolving a client certificate authentication challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLCredential/credentialWithIdentity:certificates:persistence:
func (uc _URLCredentialClass) CredentialWithIdentityCertificatesPersistence(identity objectivec.IObject, certArray IArray, persistence URLCredentialPersistence) IURLCredential {
	rv := objc.Send[URLCredential](objc.ID(uc.class), objc.Sel("credentialWithIdentity:certificates:persistence:"), identity, certArray, persistence)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CredentialWithIdentityCertificatesPersistence) */


// Creates a URL credential instance for internet password authentication with a given user name and password, using a given persistence setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURLCredential/credentialWithUser:password:persistence:
func (uc _URLCredentialClass) CredentialWithUserPasswordPersistence(user IString, password IString, persistence URLCredentialPersistence) IURLCredential {
	rv := objc.Send[URLCredential](objc.ID(uc.class), objc.Sel("credentialWithUser:password:persistence:"), user, password, persistence)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CredentialWithUserPasswordPersistence) */


// Creates a URL credential instance for server trust authentication with a given accepted trust.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/init(forTrust:)
func (uc _URLCredentialClass) CredentialForTrust(trust objectivec.IObject) IURLCredential {
	rv := objc.Send[URLCredential](objc.ID(uc.class), objc.Sel("credentialForTrust:"), trust)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CredentialForTrust) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for URLCredential */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for URLCredential */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for URLCredential */

// The intermediate certificates of the credential, if it is a client certificate credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/certificates
func (u_ URLCredential) Certificates() IArray {
	rv := objc.Send[Array](u_.ID, objc.Sel("certificates"))
	return rv
}/* debug [instance_properties/getter]: certificates */


// A Boolean value that indicates whether the credential has a password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/hasPassword
func (u_ URLCredential) HasPassword() bool {
	rv := objc.Send[bool](u_.ID, objc.Sel("hasPassword"))
	return rv
}/* debug [instance_properties/getter]: hasPassword */


// The identity of this credential if it is a client certificate credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/identity
func (u_ URLCredential) Identity() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](u_.ID, objc.Sel("identity"))
	return rv
}/* debug [instance_properties/getter]: identity */


// The credential’s password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/password
func (u_ URLCredential) Password() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("password"))
	return rv
}/* debug [instance_properties/getter]: password */


// The credential’s persistence setting.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/persistence-swift.property
func (u_ URLCredential) Persistence() URLCredentialPersistence {
	rv := objc.Send[URLCredentialPersistence](u_.ID, objc.Sel("persistence"))
	return rv
}/* debug [instance_properties/getter]: persistence */


// The credential’s user name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLCredential/user
func (u_ URLCredential) User() IString {
	rv := objc.Send[String](u_.ID, objc.Sel("user"))
	return rv
}/* debug [instance_properties/getter]: user */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSURLCredential */


