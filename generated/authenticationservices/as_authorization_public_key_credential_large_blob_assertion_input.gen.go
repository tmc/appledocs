// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPublicKeyCredentialLargeBlobAssertionInput */


/* debug [class_header]: Header for ASAuthorizationPublicKeyCredentialLargeBlobAssertionInput */
// The class instance for the [AuthorizationPublicKeyCredentialLargeBlobAssertionInput] class.
var (
	AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass     _AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass
	AuthorizationPublicKeyCredentialLargeBlobAssertionInputClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialLargeBlobAssertionInputClass() _AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass {
	AuthorizationPublicKeyCredentialLargeBlobAssertionInputClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass = _AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass{objc.GetClass("ASAuthorizationPublicKeyCredentialLargeBlobAssertionInput")}
	})
	return AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass
}

type _AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPublicKeyCredentialLargeBlobAssertionInput */
// An interface definition for the [AuthorizationPublicKeyCredentialLargeBlobAssertionInput] class.
type IAuthorizationPublicKeyCredentialLargeBlobAssertionInput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationPublicKeyCredentialLargeBlobAssertionInput */
	// properties:
	DataToWrite() objc.IObject /* cross-framework: NSData */
	SetDataToWrite(value objc.IObject /* cross-framework: NSData */)
	Operation() AuthorizationPublicKeyCredentialLargeBlobAssertionOperation
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPublicKeyCredentialLargeBlobAssertionInput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPublicKeyCredentialLargeBlobAssertionInput */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass) Alloc() AuthorizationPublicKeyCredentialLargeBlobAssertionInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationPublicKeyCredentialLargeBlobAssertionInputClass) New() AuthorizationPublicKeyCredentialLargeBlobAssertionInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionInput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialLargeBlobAssertionInput) Init() AuthorizationPublicKeyCredentialLargeBlobAssertionInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionInput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialLargeBlobAssertionInput) Autorelease() AuthorizationPublicKeyCredentialLargeBlobAssertionInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionInput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialLargeBlobAssertionInput creates a new AuthorizationPublicKeyCredentialLargeBlobAssertionInput instance.
func NewAuthorizationPublicKeyCredentialLargeBlobAssertionInput() AuthorizationPublicKeyCredentialLargeBlobAssertionInput {
	return getAuthorizationPublicKeyCredentialLargeBlobAssertionInputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPublicKeyCredentialLargeBlobAssertionInput */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionInput-c.class
type AuthorizationPublicKeyCredentialLargeBlobAssertionInput struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialLargeBlobAssertionInputFrom constructs a [AuthorizationPublicKeyCredentialLargeBlobAssertionInput] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialLargeBlobAssertionInputFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialLargeBlobAssertionInput {
	return AuthorizationPublicKeyCredentialLargeBlobAssertionInput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPublicKeyCredentialLargeBlobAssertionInput */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionInput-c.class/initWithOperation:
func NewAuthorizationPublicKeyCredentialLargeBlobAssertionInputWithOperation(operation AuthorizationPublicKeyCredentialLargeBlobAssertionOperation) AuthorizationPublicKeyCredentialLargeBlobAssertionInput {
	instance := getAuthorizationPublicKeyCredentialLargeBlobAssertionInputClass().Alloc()
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionInput](instance.ID, objc.Sel("initWithOperation:"), operation)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAuthorizationPublicKeyCredentialLargeBlobAssertionInputWithOperation */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPublicKeyCredentialLargeBlobAssertionInput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPublicKeyCredentialLargeBlobAssertionInput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPublicKeyCredentialLargeBlobAssertionInput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPublicKeyCredentialLargeBlobAssertionInput */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionInput-c.class/dataToWrite
func (a_ AuthorizationPublicKeyCredentialLargeBlobAssertionInput) DataToWrite() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("dataToWrite"))
	return rv
}/* debug [instance_properties/getter]: dataToWrite */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionInput-c.class/dataToWrite
func (a_ AuthorizationPublicKeyCredentialLargeBlobAssertionInput) SetDataToWrite(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDataToWrite:"), value)
}/* debug [instance_properties/setter]: dataToWrite */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobAssertionInput-c.class/operation
func (a_ AuthorizationPublicKeyCredentialLargeBlobAssertionInput) Operation() AuthorizationPublicKeyCredentialLargeBlobAssertionOperation {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionOperation](a_.ID, objc.Sel("operation"))
	return rv
}/* debug [instance_properties/getter]: operation */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPublicKeyCredentialLargeBlobAssertionInput */


