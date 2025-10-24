// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASPasskeyRegistrationCredentialExtensionInput */


/* debug [class_header]: Header for ASPasskeyRegistrationCredentialExtensionInput */
// The class instance for the [PasskeyRegistrationCredentialExtensionInput] class.
var (
	PasskeyRegistrationCredentialExtensionInputClass     _PasskeyRegistrationCredentialExtensionInputClass
	PasskeyRegistrationCredentialExtensionInputClassOnce sync.Once
)

func getPasskeyRegistrationCredentialExtensionInputClass() _PasskeyRegistrationCredentialExtensionInputClass {
	PasskeyRegistrationCredentialExtensionInputClassOnce.Do(func() {
		PasskeyRegistrationCredentialExtensionInputClass = _PasskeyRegistrationCredentialExtensionInputClass{objc.GetClass("ASPasskeyRegistrationCredentialExtensionInput")}
	})
	return PasskeyRegistrationCredentialExtensionInputClass
}

type _PasskeyRegistrationCredentialExtensionInputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PasskeyRegistrationCredentialExtensionInput */
// An interface definition for the [PasskeyRegistrationCredentialExtensionInput] class.
type IPasskeyRegistrationCredentialExtensionInput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PasskeyRegistrationCredentialExtensionInput */
	// properties:
	LargeBlob() IASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PasskeyRegistrationCredentialExtensionInput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PasskeyRegistrationCredentialExtensionInput */
// Alloc allocates a new instance without initialization.
func (pc _PasskeyRegistrationCredentialExtensionInputClass) Alloc() PasskeyRegistrationCredentialExtensionInput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionInput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PasskeyRegistrationCredentialExtensionInputClass) New() PasskeyRegistrationCredentialExtensionInput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionInput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasskeyRegistrationCredentialExtensionInput) Init() PasskeyRegistrationCredentialExtensionInput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionInput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasskeyRegistrationCredentialExtensionInput) Autorelease() PasskeyRegistrationCredentialExtensionInput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionInput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasskeyRegistrationCredentialExtensionInput creates a new PasskeyRegistrationCredentialExtensionInput instance.
func NewPasskeyRegistrationCredentialExtensionInput() PasskeyRegistrationCredentialExtensionInput {
	return getPasskeyRegistrationCredentialExtensionInputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PasskeyRegistrationCredentialExtensionInput */
// This class encapsulates input for various WebAuthn extensions during passkey registration.


// This class encapsulates input for various WebAuthn extensions during passkey registration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredentialExtensionInput-c.class
type PasskeyRegistrationCredentialExtensionInput struct {
	objectivec.Object
}

// PasskeyRegistrationCredentialExtensionInputFrom constructs a [PasskeyRegistrationCredentialExtensionInput] from an unsafe.Pointer.
//
// This class encapsulates input for various WebAuthn extensions during passkey registration.
func PasskeyRegistrationCredentialExtensionInputFrom(ptr unsafe.Pointer) PasskeyRegistrationCredentialExtensionInput {
	return PasskeyRegistrationCredentialExtensionInput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PasskeyRegistrationCredentialExtensionInput *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PasskeyRegistrationCredentialExtensionInput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PasskeyRegistrationCredentialExtensionInput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PasskeyRegistrationCredentialExtensionInput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PasskeyRegistrationCredentialExtensionInput */

// Input for the extension in passkey registration requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredentialExtensionInput-c.class/largeBlob
func (p_ PasskeyRegistrationCredentialExtensionInput) LargeBlob() IASAuthorizationPublicKeyCredentialLargeBlobRegistrationInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationInput](p_.ID, objc.Sel("largeBlob"))
	return rv
}/* debug [instance_properties/getter]: largeBlob */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASPasskeyRegistrationCredentialExtensionInput */



