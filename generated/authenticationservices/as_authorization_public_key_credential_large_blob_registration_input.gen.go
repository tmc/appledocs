// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput */


/* debug [class_header]: Header for ASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput */
// The class instance for the [AuthorizationPublicKeyCredentialLargeBlobRegistrationInput] class.
var (
	AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass     _AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass
	AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClassOnce sync.Once
)

func getAuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass() _AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass {
	AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClassOnce.Do(func() {
		AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass = _AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass{objc.GetClass("ASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput")}
	})
	return AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass
}

type _AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPublicKeyCredentialLargeBlobRegistrationInput */
// An interface definition for the [AuthorizationPublicKeyCredentialLargeBlobRegistrationInput] class.
type IAuthorizationPublicKeyCredentialLargeBlobRegistrationInput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationPublicKeyCredentialLargeBlobRegistrationInput */
	// properties:
	SupportRequirement() AuthorizationPublicKeyCredentialLargeBlobSupportRequirement
	SetSupportRequirement(value AuthorizationPublicKeyCredentialLargeBlobSupportRequirement)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPublicKeyCredentialLargeBlobRegistrationInput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPublicKeyCredentialLargeBlobRegistrationInput */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass) Alloc() AuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass) New() AuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationInput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPublicKeyCredentialLargeBlobRegistrationInput) Init() AuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationInput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPublicKeyCredentialLargeBlobRegistrationInput) Autorelease() AuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationInput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPublicKeyCredentialLargeBlobRegistrationInput creates a new AuthorizationPublicKeyCredentialLargeBlobRegistrationInput instance.
func NewAuthorizationPublicKeyCredentialLargeBlobRegistrationInput() AuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	return getAuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPublicKeyCredentialLargeBlobRegistrationInput */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput-c.class
type AuthorizationPublicKeyCredentialLargeBlobRegistrationInput struct {
	objectivec.Object
}

// AuthorizationPublicKeyCredentialLargeBlobRegistrationInputFrom constructs a [AuthorizationPublicKeyCredentialLargeBlobRegistrationInput] from an unsafe.Pointer.
func AuthorizationPublicKeyCredentialLargeBlobRegistrationInputFrom(ptr unsafe.Pointer) AuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	return AuthorizationPublicKeyCredentialLargeBlobRegistrationInput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPublicKeyCredentialLargeBlobRegistrationInput */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput-c.class/initWithSupportRequirement:
func NewAuthorizationPublicKeyCredentialLargeBlobRegistrationInputWithSupportRequirement(requirement AuthorizationPublicKeyCredentialLargeBlobSupportRequirement) AuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	instance := getAuthorizationPublicKeyCredentialLargeBlobRegistrationInputClass().Alloc()
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationInput](instance.ID, objc.Sel("initWithSupportRequirement:"), requirement)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAuthorizationPublicKeyCredentialLargeBlobRegistrationInputWithSupportRequirement */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPublicKeyCredentialLargeBlobRegistrationInput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPublicKeyCredentialLargeBlobRegistrationInput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPublicKeyCredentialLargeBlobRegistrationInput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPublicKeyCredentialLargeBlobRegistrationInput */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput-c.class/supportRequirement
func (a_ AuthorizationPublicKeyCredentialLargeBlobRegistrationInput) SupportRequirement() AuthorizationPublicKeyCredentialLargeBlobSupportRequirement {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobSupportRequirement](a_.ID, objc.Sel("supportRequirement"))
	return rv
}/* debug [instance_properties/getter]: supportRequirement */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput-c.class/supportRequirement
func (a_ AuthorizationPublicKeyCredentialLargeBlobRegistrationInput) SetSupportRequirement(value AuthorizationPublicKeyCredentialLargeBlobSupportRequirement) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSupportRequirement:"), value)
}/* debug [instance_properties/setter]: supportRequirement */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput */


