// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASOneTimeCodeCredential */


/* debug [class_header]: Header for ASOneTimeCodeCredential */
// The class instance for the [OneTimeCodeCredential] class.
var (
	OneTimeCodeCredentialClass     _OneTimeCodeCredentialClass
	OneTimeCodeCredentialClassOnce sync.Once
)

func getOneTimeCodeCredentialClass() _OneTimeCodeCredentialClass {
	OneTimeCodeCredentialClassOnce.Do(func() {
		OneTimeCodeCredentialClass = _OneTimeCodeCredentialClass{objc.GetClass("ASOneTimeCodeCredential")}
	})
	return OneTimeCodeCredentialClass
}

type _OneTimeCodeCredentialClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for OneTimeCodeCredential */
// An interface definition for the [OneTimeCodeCredential] class.
type IOneTimeCodeCredential interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for OneTimeCodeCredential */
	// properties:
	Code() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for OneTimeCodeCredential */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for OneTimeCodeCredential */
// Alloc allocates a new instance without initialization.
func (oc _OneTimeCodeCredentialClass) Alloc() OneTimeCodeCredential {
	rv := objc.Send[OneTimeCodeCredential](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OneTimeCodeCredentialClass) New() OneTimeCodeCredential {
	rv := objc.Send[OneTimeCodeCredential](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OneTimeCodeCredential) Init() OneTimeCodeCredential {
	rv := objc.Send[OneTimeCodeCredential](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OneTimeCodeCredential) Autorelease() OneTimeCodeCredential {
	rv := objc.Send[OneTimeCodeCredential](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOneTimeCodeCredential creates a new OneTimeCodeCredential instance.
func NewOneTimeCodeCredential() OneTimeCodeCredential {
	return getOneTimeCodeCredentialClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for OneTimeCodeCredential */
// A one-time passcode (OTP) credential.


// A one-time passcode (OTP) credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASOneTimeCodeCredential
type OneTimeCodeCredential struct {
	objectivec.Object
}

// OneTimeCodeCredentialFrom constructs a [OneTimeCodeCredential] from an unsafe.Pointer.
//
// A one-time passcode (OTP) credential.
func OneTimeCodeCredentialFrom(ptr unsafe.Pointer) OneTimeCodeCredential {
	return OneTimeCodeCredential{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for OneTimeCodeCredential */

// Creates a one-time passcode (OTP) credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASOneTimeCodeCredential/init(code:)
func NewOneTimeCodeCredentialWithCode(code objc.IObject /* cross-framework: NSString */) OneTimeCodeCredential {
	instance := getOneTimeCodeCredentialClass().Alloc()
	rv := objc.Send[OneTimeCodeCredential](instance.ID, objc.Sel("initWithCode:"), code)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewOneTimeCodeCredentialWithCode */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for OneTimeCodeCredential */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASOneTimeCodeCredential/credentialWithCode:
func (oc _OneTimeCodeCredentialClass) CredentialWithCode(code objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("credentialWithCode:"), code)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CredentialWithCode) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for OneTimeCodeCredential */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for OneTimeCodeCredential */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for OneTimeCodeCredential */

// The one-time passcode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASOneTimeCodeCredential/code
func (o_ OneTimeCodeCredential) Code() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](o_.ID, objc.Sel("code"))
	return rv
}/* debug [instance_properties/getter]: code */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASOneTimeCodeCredential */


