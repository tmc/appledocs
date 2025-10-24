// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPublicKeyCredentialPRFAssertionOutput */


/* debug [class_header]: Header for ASAuthorizationPublicKeyCredentialPRFAssertionOutput */
// The class instance for the [AuthorizationPublicKeyCredentialPRFAssertionOutput] class.
var (
	AuthorizationPublicKeyCredentialPRFAssertionOutputClass     _AuthorizationPublicKeyCredentialPRFAssertionOutputClass
	AuthorizationPublicKeyCredentialPRFAssertionOutputClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialPRFAssertionOutputClass() _AuthorizationPublicKeyCredentialPRFAssertionOutputClass {
	AuthorizationPublicKeyCredentialPRFAssertionOutputClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialPRFAssertionOutputClass = _AuthorizationPublicKeyCredentialPRFAssertionOutputClass{objc.GetClass("ASAuthorizationPublicKeyCredentialPRFAssertionOutput")}
	})
	return AuthorizationPublicKeyCredentialPRFAssertionOutputClass
}

type _AuthorizationPublicKeyCredentialPRFAssertionOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPublicKeyCredentialPRFAssertionOutput */
// An interface definition for the [AuthorizationPublicKeyCredentialPRFAssertionOutput] class.
type IAuthorizationPublicKeyCredentialPRFAssertionOutput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationPublicKeyCredentialPRFAssertionOutput */
	// properties:
	First() objc.IObject /* cross-framework: NSData */
	Second() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPublicKeyCredentialPRFAssertionOutput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPublicKeyCredentialPRFAssertionOutput */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialPRFAssertionOutputClass) Alloc() AuthorizationPublicKeyCredentialPRFAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationPublicKeyCredentialPRFAssertionOutputClass) New() AuthorizationPublicKeyCredentialPRFAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialPRFAssertionOutput) Init() AuthorizationPublicKeyCredentialPRFAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialPRFAssertionOutput) Autorelease() AuthorizationPublicKeyCredentialPRFAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialPRFAssertionOutput creates a new AuthorizationPublicKeyCredentialPRFAssertionOutput instance.
func NewAuthorizationPublicKeyCredentialPRFAssertionOutput() AuthorizationPublicKeyCredentialPRFAssertionOutput {
	return getAuthorizationPublicKeyCredentialPRFAssertionOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPublicKeyCredentialPRFAssertionOutput */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFAssertionOutput-c.class
type AuthorizationPublicKeyCredentialPRFAssertionOutput struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialPRFAssertionOutputFrom constructs a [AuthorizationPublicKeyCredentialPRFAssertionOutput] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialPRFAssertionOutputFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialPRFAssertionOutput {
	return AuthorizationPublicKeyCredentialPRFAssertionOutput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPublicKeyCredentialPRFAssertionOutput *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPublicKeyCredentialPRFAssertionOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPublicKeyCredentialPRFAssertionOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPublicKeyCredentialPRFAssertionOutput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPublicKeyCredentialPRFAssertionOutput */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFAssertionOutput-c.class/first
func (a_ AuthorizationPublicKeyCredentialPRFAssertionOutput) First() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("first"))
	return rv
}/* debug [instance_properties/getter]: first */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFAssertionOutput-c.class/second
func (a_ AuthorizationPublicKeyCredentialPRFAssertionOutput) Second() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("second"))
	return rv
}/* debug [instance_properties/getter]: second */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPublicKeyCredentialPRFAssertionOutput */



