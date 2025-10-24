// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASPasswordCredential */


/* debug [class_header]: Header for ASPasswordCredential */
// The class instance for the [PasswordCredential] class.
var (
	PasswordCredentialClass     _PasswordCredentialClass
	PasswordCredentialClassOnce sync.Once
)

func getPasswordCredentialClass() _PasswordCredentialClass {
	PasswordCredentialClassOnce.Do(func() {
		PasswordCredentialClass = _PasswordCredentialClass{objc.GetClass("ASPasswordCredential")}
	})
	return PasswordCredentialClass
}

type _PasswordCredentialClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PasswordCredential */
// An interface definition for the [PasswordCredential] class.
type IPasswordCredential interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PasswordCredential */
	// properties:
	Password() objc.IObject /* cross-framework: NSString */
	User() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PasswordCredential */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PasswordCredential */
// Alloc allocates a new instance without initialization.
func (pc _PasswordCredentialClass) Alloc() PasswordCredential {
	rv := objc.Send[PasswordCredential](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PasswordCredentialClass) New() PasswordCredential {
	rv := objc.Send[PasswordCredential](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasswordCredential) Init() PasswordCredential {
	rv := objc.Send[PasswordCredential](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasswordCredential) Autorelease() PasswordCredential {
	rv := objc.Send[PasswordCredential](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasswordCredential creates a new PasswordCredential instance.
func NewPasswordCredential() PasswordCredential {
	return getPasswordCredentialClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PasswordCredential */
// A password credential.


// A password credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredential
type PasswordCredential struct {
	objectivec.Object
}

// PasswordCredentialFrom constructs a [PasswordCredential] from an unsafe.Pointer.
//
// A password credential.
func PasswordCredentialFrom(ptr unsafe.Pointer) PasswordCredential {
	return PasswordCredential{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PasswordCredential */

// Initializes a password credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredential/init(user:password:)
func NewPasswordCredentialWithUserPassword(user objc.IObject /* cross-framework: NSString */, password objc.IObject /* cross-framework: NSString */) PasswordCredential {
	instance := getPasswordCredentialClass().Alloc()
	rv := objc.Send[PasswordCredential](instance.ID, objc.Sel("initWithUser:password:"), user, password)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPasswordCredentialWithUserPassword */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PasswordCredential */

// Creates a password credential instance with a given user name and password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredential/credentialWithUser:password:
func (pc _PasswordCredentialClass) CredentialWithUserPassword(user objc.IObject /* cross-framework: NSString */, password objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("credentialWithUser:password:"), user, password)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CredentialWithUserPassword) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PasswordCredential */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PasswordCredential */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PasswordCredential */

// The password for a password credential object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredential/password
func (p_ PasswordCredential) Password() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("password"))
	return rv
}/* debug [instance_properties/getter]: password */


// The user for a password credential object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredential/user
func (p_ PasswordCredential) User() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("user"))
	return rv
}/* debug [instance_properties/getter]: user */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASPasswordCredential */


