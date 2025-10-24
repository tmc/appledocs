// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPublicKeyCredentialPRFRegistrationOutput */


/* debug [class_header]: Header for ASAuthorizationPublicKeyCredentialPRFRegistrationOutput */
// The class instance for the [AuthorizationPublicKeyCredentialPRFRegistrationOutput] class.
var (
	AuthorizationPublicKeyCredentialPRFRegistrationOutputClass     _AuthorizationPublicKeyCredentialPRFRegistrationOutputClass
	AuthorizationPublicKeyCredentialPRFRegistrationOutputClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialPRFRegistrationOutputClass() _AuthorizationPublicKeyCredentialPRFRegistrationOutputClass {
	AuthorizationPublicKeyCredentialPRFRegistrationOutputClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialPRFRegistrationOutputClass = _AuthorizationPublicKeyCredentialPRFRegistrationOutputClass{objc.GetClass("ASAuthorizationPublicKeyCredentialPRFRegistrationOutput")}
	})
	return AuthorizationPublicKeyCredentialPRFRegistrationOutputClass
}

type _AuthorizationPublicKeyCredentialPRFRegistrationOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPublicKeyCredentialPRFRegistrationOutput */
// An interface definition for the [AuthorizationPublicKeyCredentialPRFRegistrationOutput] class.
type IAuthorizationPublicKeyCredentialPRFRegistrationOutput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationPublicKeyCredentialPRFRegistrationOutput */
	// properties:
	First() objc.IObject /* cross-framework: NSData */
	IsSupported() bool
	Second() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPublicKeyCredentialPRFRegistrationOutput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPublicKeyCredentialPRFRegistrationOutput */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialPRFRegistrationOutputClass) Alloc() AuthorizationPublicKeyCredentialPRFRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationPublicKeyCredentialPRFRegistrationOutputClass) New() AuthorizationPublicKeyCredentialPRFRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialPRFRegistrationOutput) Init() AuthorizationPublicKeyCredentialPRFRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialPRFRegistrationOutput) Autorelease() AuthorizationPublicKeyCredentialPRFRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialPRFRegistrationOutput creates a new AuthorizationPublicKeyCredentialPRFRegistrationOutput instance.
func NewAuthorizationPublicKeyCredentialPRFRegistrationOutput() AuthorizationPublicKeyCredentialPRFRegistrationOutput {
	return getAuthorizationPublicKeyCredentialPRFRegistrationOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPublicKeyCredentialPRFRegistrationOutput */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFRegistrationOutput-c.class
type AuthorizationPublicKeyCredentialPRFRegistrationOutput struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialPRFRegistrationOutputFrom constructs a [AuthorizationPublicKeyCredentialPRFRegistrationOutput] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialPRFRegistrationOutputFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialPRFRegistrationOutput {
	return AuthorizationPublicKeyCredentialPRFRegistrationOutput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPublicKeyCredentialPRFRegistrationOutput *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPublicKeyCredentialPRFRegistrationOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPublicKeyCredentialPRFRegistrationOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPublicKeyCredentialPRFRegistrationOutput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPublicKeyCredentialPRFRegistrationOutput */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFRegistrationOutput-c.class/first
func (a_ AuthorizationPublicKeyCredentialPRFRegistrationOutput) First() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("first"))
	return rv
}/* debug [instance_properties/getter]: first */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFRegistrationOutput-c.class/isSupported
func (a_ AuthorizationPublicKeyCredentialPRFRegistrationOutput) IsSupported() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSupported"))
	return rv
}/* debug [instance_properties/getter]: isSupported */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFRegistrationOutput-c.class/second
func (a_ AuthorizationPublicKeyCredentialPRFRegistrationOutput) Second() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("second"))
	return rv
}/* debug [instance_properties/getter]: second */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPublicKeyCredentialPRFRegistrationOutput */



