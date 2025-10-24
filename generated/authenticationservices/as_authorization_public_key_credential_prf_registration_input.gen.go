// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPublicKeyCredentialPRFRegistrationInput */


/* debug [class_header]: Header for ASAuthorizationPublicKeyCredentialPRFRegistrationInput */
// The class instance for the [AuthorizationPublicKeyCredentialPRFRegistrationInput] class.
var (
	AuthorizationPublicKeyCredentialPRFRegistrationInputClass     _AuthorizationPublicKeyCredentialPRFRegistrationInputClass
	AuthorizationPublicKeyCredentialPRFRegistrationInputClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialPRFRegistrationInputClass() _AuthorizationPublicKeyCredentialPRFRegistrationInputClass {
	AuthorizationPublicKeyCredentialPRFRegistrationInputClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialPRFRegistrationInputClass = _AuthorizationPublicKeyCredentialPRFRegistrationInputClass{objc.GetClass("ASAuthorizationPublicKeyCredentialPRFRegistrationInput")}
	})
	return AuthorizationPublicKeyCredentialPRFRegistrationInputClass
}

type _AuthorizationPublicKeyCredentialPRFRegistrationInputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPublicKeyCredentialPRFRegistrationInput */
// An interface definition for the [AuthorizationPublicKeyCredentialPRFRegistrationInput] class.
type IAuthorizationPublicKeyCredentialPRFRegistrationInput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationPublicKeyCredentialPRFRegistrationInput */
	// properties:
	InputValues() IASAuthorizationPublicKeyCredentialPRFAssertionInputValues
	ShouldCheckForSupport() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPublicKeyCredentialPRFRegistrationInput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPublicKeyCredentialPRFRegistrationInput */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialPRFRegistrationInputClass) Alloc() AuthorizationPublicKeyCredentialPRFRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationPublicKeyCredentialPRFRegistrationInputClass) New() AuthorizationPublicKeyCredentialPRFRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationInput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialPRFRegistrationInput) Init() AuthorizationPublicKeyCredentialPRFRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationInput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialPRFRegistrationInput) Autorelease() AuthorizationPublicKeyCredentialPRFRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationInput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialPRFRegistrationInput creates a new AuthorizationPublicKeyCredentialPRFRegistrationInput instance.
func NewAuthorizationPublicKeyCredentialPRFRegistrationInput() AuthorizationPublicKeyCredentialPRFRegistrationInput {
	return getAuthorizationPublicKeyCredentialPRFRegistrationInputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPublicKeyCredentialPRFRegistrationInput */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFRegistrationInput-c.class
type AuthorizationPublicKeyCredentialPRFRegistrationInput struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialPRFRegistrationInputFrom constructs a [AuthorizationPublicKeyCredentialPRFRegistrationInput] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialPRFRegistrationInputFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialPRFRegistrationInput {
	return AuthorizationPublicKeyCredentialPRFRegistrationInput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPublicKeyCredentialPRFRegistrationInput */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFRegistrationInput-c.class/initWithInputValues:
func NewAuthorizationPublicKeyCredentialPRFRegistrationInputWithInputValues(inputValues IASAuthorizationPublicKeyCredentialPRFAssertionInputValues) AuthorizationPublicKeyCredentialPRFRegistrationInput {
	instance := getAuthorizationPublicKeyCredentialPRFRegistrationInputClass().Alloc()
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFRegistrationInput](instance.ID, objc.Sel("initWithInputValues:"), inputValues)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAuthorizationPublicKeyCredentialPRFRegistrationInputWithInputValues */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPublicKeyCredentialPRFRegistrationInput */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFRegistrationInput-c.class/checkForSupport
func (ac _AuthorizationPublicKeyCredentialPRFRegistrationInputClass) CheckForSupport() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("checkForSupport"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CheckForSupport) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPublicKeyCredentialPRFRegistrationInput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPublicKeyCredentialPRFRegistrationInput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPublicKeyCredentialPRFRegistrationInput */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFRegistrationInput-c.class/inputValues
func (a_ AuthorizationPublicKeyCredentialPRFRegistrationInput) InputValues() IASAuthorizationPublicKeyCredentialPRFAssertionInputValues {
	rv := objc.Send[AuthorizationPublicKeyCredentialPRFAssertionInputValues](a_.ID, objc.Sel("inputValues"))
	return rv
}/* debug [instance_properties/getter]: inputValues */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialPRFRegistrationInput-c.class/shouldCheckForSupport
func (a_ AuthorizationPublicKeyCredentialPRFRegistrationInput) ShouldCheckForSupport() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldCheckForSupport"))
	return rv
}/* debug [instance_properties/getter]: shouldCheckForSupport */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPublicKeyCredentialPRFRegistrationInput */


