// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPublicKeyCredentialLargeBlobAssertionOutput */


/* debug [class_header]: Header for ASAuthorizationPublicKeyCredentialLargeBlobAssertionOutput */
// The class instance for the [AuthorizationPublicKeyCredentialLargeBlobAssertionOutput] class.
var (
	AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass     _AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass
	AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass() _AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass {
	AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass = _AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass{objc.GetClass("ASAuthorizationPublicKeyCredentialLargeBlobAssertionOutput")}
	})
	return AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass
}

type _AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPublicKeyCredentialLargeBlobAssertionOutput */
// An interface definition for the [AuthorizationPublicKeyCredentialLargeBlobAssertionOutput] class.
type IAuthorizationPublicKeyCredentialLargeBlobAssertionOutput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationPublicKeyCredentialLargeBlobAssertionOutput */
	// properties:
	DidWrite() bool
	ReadData() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPublicKeyCredentialLargeBlobAssertionOutput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPublicKeyCredentialLargeBlobAssertionOutput */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass) Alloc() AuthorizationPublicKeyCredentialLargeBlobAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass) New() AuthorizationPublicKeyCredentialLargeBlobAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialLargeBlobAssertionOutput) Init() AuthorizationPublicKeyCredentialLargeBlobAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialLargeBlobAssertionOutput) Autorelease() AuthorizationPublicKeyCredentialLargeBlobAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialLargeBlobAssertionOutput creates a new AuthorizationPublicKeyCredentialLargeBlobAssertionOutput instance.
func NewAuthorizationPublicKeyCredentialLargeBlobAssertionOutput() AuthorizationPublicKeyCredentialLargeBlobAssertionOutput {
	return getAuthorizationPublicKeyCredentialLargeBlobAssertionOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPublicKeyCredentialLargeBlobAssertionOutput */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionOutput-c.class
type AuthorizationPublicKeyCredentialLargeBlobAssertionOutput struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialLargeBlobAssertionOutputFrom constructs a [AuthorizationPublicKeyCredentialLargeBlobAssertionOutput] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialLargeBlobAssertionOutputFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialLargeBlobAssertionOutput {
	return AuthorizationPublicKeyCredentialLargeBlobAssertionOutput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPublicKeyCredentialLargeBlobAssertionOutput *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPublicKeyCredentialLargeBlobAssertionOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPublicKeyCredentialLargeBlobAssertionOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPublicKeyCredentialLargeBlobAssertionOutput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPublicKeyCredentialLargeBlobAssertionOutput */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionOutput-c.class/didWrite
func (a_ AuthorizationPublicKeyCredentialLargeBlobAssertionOutput) DidWrite() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("didWrite"))
	return rv
}/* debug [instance_properties/getter]: didWrite */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionOutput-c.class/readData
func (a_ AuthorizationPublicKeyCredentialLargeBlobAssertionOutput) ReadData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("readData"))
	return rv
}/* debug [instance_properties/getter]: readData */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPublicKeyCredentialLargeBlobAssertionOutput */



