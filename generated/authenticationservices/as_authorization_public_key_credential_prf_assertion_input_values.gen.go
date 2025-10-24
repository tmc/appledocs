// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPublicKeyCredentialPRFAssertionInputValues */


/* debug [class_header]: Header for ASAuthorizationPublicKeyCredentialPRFAssertionInputValues */
// The class instance for the [AuthorizationPublicKeyCredentialPRFAssertionInputValues] class.
var (
	AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass     _AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass
	AuthorizationPublicKeyCredentialPRFAssertionInputValuesClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialPRFAssertionInputValuesClass() _AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass {
	AuthorizationPublicKeyCredentialPRFAssertionInputValuesClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass = _AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass{objc.GetClass("ASAuthorizationPublicKeyCredentialPRFAssertionInputValues")}
	})
	return AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass
}

type _AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPublicKeyCredentialPRFAssertionInputValues */
// An interface definition for the [AuthorizationPublicKeyCredentialPRFAssertionInputValues] class.
type IAuthorizationPublicKeyCredentialPRFAssertionInputValues interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationPublicKeyCredentialPRFAssertionInputValues */
	// properties:
	SaltInput1() objc.IObject /* cross-framework: NSData */
	SaltInput2() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPublicKeyCredentialPRFAssertionInputValues */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPublicKeyCredentialPRFAssertionInputValues */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass) Alloc() AuthorizationPublicKeyCredentialPRFAssertionInputValues {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionInputValues](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationPublicKeyCredentialPRFAssertionInputValuesClass) New() AuthorizationPublicKeyCredentialPRFAssertionInputValues {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionInputValues](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialPRFAssertionInputValues) Init() AuthorizationPublicKeyCredentialPRFAssertionInputValues {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionInputValues](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialPRFAssertionInputValues) Autorelease() AuthorizationPublicKeyCredentialPRFAssertionInputValues {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionInputValues](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialPRFAssertionInputValues creates a new AuthorizationPublicKeyCredentialPRFAssertionInputValues instance.
func NewAuthorizationPublicKeyCredentialPRFAssertionInputValues() AuthorizationPublicKeyCredentialPRFAssertionInputValues {
	return getAuthorizationPublicKeyCredentialPRFAssertionInputValuesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPublicKeyCredentialPRFAssertionInputValues */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFAssertionInputValues
type AuthorizationPublicKeyCredentialPRFAssertionInputValues struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialPRFAssertionInputValuesFrom constructs a [AuthorizationPublicKeyCredentialPRFAssertionInputValues] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialPRFAssertionInputValuesFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialPRFAssertionInputValues {
	return AuthorizationPublicKeyCredentialPRFAssertionInputValues{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPublicKeyCredentialPRFAssertionInputValues */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFAssertionInputValues/initWithSaltInput1:saltInput2:
func NewAuthorizationPublicKeyCredentialPRFAssertionInputValuesWithSaltInput1SaltInput2(saltInput1 objc.IObject /* cross-framework: NSData */, saltInput2 objc.IObject /* cross-framework: NSData */) AuthorizationPublicKeyCredentialPRFAssertionInputValues {
	instance := getAuthorizationPublicKeyCredentialPRFAssertionInputValuesClass().Alloc()
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionInputValues](instance.ID, objc.Sel("initWithSaltInput1:saltInput2:"), saltInput1, saltInput2)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAuthorizationPublicKeyCredentialPRFAssertionInputValuesWithSaltInput1SaltInput2 */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPublicKeyCredentialPRFAssertionInputValues */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPublicKeyCredentialPRFAssertionInputValues */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPublicKeyCredentialPRFAssertionInputValues */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPublicKeyCredentialPRFAssertionInputValues */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFAssertionInputValues/saltInput1
func (a_ AuthorizationPublicKeyCredentialPRFAssertionInputValues) SaltInput1() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("saltInput1"))
	return rv
}/* debug [instance_properties/getter]: saltInput1 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFAssertionInputValues/saltInput2
func (a_ AuthorizationPublicKeyCredentialPRFAssertionInputValues) SaltInput2() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("saltInput2"))
	return rv
}/* debug [instance_properties/getter]: saltInput2 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPublicKeyCredentialPRFAssertionInputValues */


