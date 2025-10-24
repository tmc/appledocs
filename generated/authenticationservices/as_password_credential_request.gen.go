// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASPasswordCredentialRequest */


/* debug [class_header]: Header for ASPasswordCredentialRequest */
// The class instance for the [PasswordCredentialRequest] class.
var (
	PasswordCredentialRequestClass     _PasswordCredentialRequestClass
	PasswordCredentialRequestClassOnce sync.Once
)

func getPasswordCredentialRequestClass() _PasswordCredentialRequestClass {
	PasswordCredentialRequestClassOnce.Do(func() {
		PasswordCredentialRequestClass = _PasswordCredentialRequestClass{objc.GetClass("ASPasswordCredentialRequest")}
	})
	return PasswordCredentialRequestClass
}

type _PasswordCredentialRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PasswordCredentialRequest */
// An interface definition for the [PasswordCredentialRequest] class.
type IPasswordCredentialRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PasswordCredentialRequest */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PasswordCredentialRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PasswordCredentialRequest */
// Alloc allocates a new instance without initialization.
func (pc _PasswordCredentialRequestClass) Alloc() PasswordCredentialRequest {
	rv := objc.Send[PasswordCredentialRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PasswordCredentialRequestClass) New() PasswordCredentialRequest {
	rv := objc.Send[PasswordCredentialRequest](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasswordCredentialRequest) Init() PasswordCredentialRequest {
	rv := objc.Send[PasswordCredentialRequest](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasswordCredentialRequest) Autorelease() PasswordCredentialRequest {
	rv := objc.Send[PasswordCredentialRequest](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasswordCredentialRequest creates a new PasswordCredentialRequest instance.
func NewPasswordCredentialRequest() PasswordCredentialRequest {
	return getPasswordCredentialRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PasswordCredentialRequest */
// A class that represents a request to supply a password credential.


// A class that represents a request to supply a password credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialRequest
type PasswordCredentialRequest struct {
	objectivec.Object
}

// PasswordCredentialRequestFrom constructs a [PasswordCredentialRequest] from an unsafe.Pointer.
//
// A class that represents a request to supply a password credential.
func PasswordCredentialRequestFrom(ptr unsafe.Pointer) PasswordCredentialRequest {
	return PasswordCredentialRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PasswordCredentialRequest */

// Initializes a password credential request object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialRequest/init(credentialIdentity:)
func NewPasswordCredentialRequestWithCredentialIdentity(credentialIdentity IASPasswordCredentialIdentity) PasswordCredentialRequest {
	instance := getPasswordCredentialRequestClass().Alloc()
	rv := objc.Send[PasswordCredentialRequest](instance.ID, objc.Sel("initWithCredentialIdentity:"), credentialIdentity)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPasswordCredentialRequestWithCredentialIdentity */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PasswordCredentialRequest */

// Creates and initializes a password credential request object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasswordCredentialRequest/requestWithCredentialIdentity:
func (pc _PasswordCredentialRequestClass) RequestWithCredentialIdentity(credentialIdentity IASPasswordCredentialIdentity) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("requestWithCredentialIdentity:"), credentialIdentity)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=RequestWithCredentialIdentity) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PasswordCredentialRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PasswordCredentialRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PasswordCredentialRequest */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASPasswordCredentialRequest */


