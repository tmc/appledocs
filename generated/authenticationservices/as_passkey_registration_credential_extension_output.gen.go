// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASPasskeyRegistrationCredentialExtensionOutput */


/* debug [class_header]: Header for ASPasskeyRegistrationCredentialExtensionOutput */
// The class instance for the [PasskeyRegistrationCredentialExtensionOutput] class.
var (
	PasskeyRegistrationCredentialExtensionOutputClass     _PasskeyRegistrationCredentialExtensionOutputClass
	PasskeyRegistrationCredentialExtensionOutputClassOnce sync.Once
)

func getPasskeyRegistrationCredentialExtensionOutputClass() _PasskeyRegistrationCredentialExtensionOutputClass {
	PasskeyRegistrationCredentialExtensionOutputClassOnce.Do(func() {
		PasskeyRegistrationCredentialExtensionOutputClass = _PasskeyRegistrationCredentialExtensionOutputClass{objc.GetClass("ASPasskeyRegistrationCredentialExtensionOutput")}
	})
	return PasskeyRegistrationCredentialExtensionOutputClass
}

type _PasskeyRegistrationCredentialExtensionOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PasskeyRegistrationCredentialExtensionOutput */
// An interface definition for the [PasskeyRegistrationCredentialExtensionOutput] class.
type IPasskeyRegistrationCredentialExtensionOutput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PasskeyRegistrationCredentialExtensionOutput */
	// properties:
	LargeBlobRegistrationOutput() IASAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PasskeyRegistrationCredentialExtensionOutput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PasskeyRegistrationCredentialExtensionOutput */
// Alloc allocates a new instance without initialization.
func (pc _PasskeyRegistrationCredentialExtensionOutputClass) Alloc() PasskeyRegistrationCredentialExtensionOutput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionOutput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PasskeyRegistrationCredentialExtensionOutputClass) New() PasskeyRegistrationCredentialExtensionOutput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionOutput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasskeyRegistrationCredentialExtensionOutput) Init() PasskeyRegistrationCredentialExtensionOutput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionOutput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasskeyRegistrationCredentialExtensionOutput) Autorelease() PasskeyRegistrationCredentialExtensionOutput {
	rv := objc.Send[PasskeyRegistrationCredentialExtensionOutput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasskeyRegistrationCredentialExtensionOutput creates a new PasskeyRegistrationCredentialExtensionOutput instance.
func NewPasskeyRegistrationCredentialExtensionOutput() PasskeyRegistrationCredentialExtensionOutput {
	return getPasskeyRegistrationCredentialExtensionOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PasskeyRegistrationCredentialExtensionOutput */
// This class encapsulates output for various WebAuthn extensions used during passkey registration.


// This class encapsulates output for various WebAuthn extensions used during passkey registration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredentialExtensionOutput-c.class
type PasskeyRegistrationCredentialExtensionOutput struct {
	objectivec.Object
}

// PasskeyRegistrationCredentialExtensionOutputFrom constructs a [PasskeyRegistrationCredentialExtensionOutput] from an unsafe.Pointer.
//
// This class encapsulates output for various WebAuthn extensions used during passkey registration.
func PasskeyRegistrationCredentialExtensionOutputFrom(ptr unsafe.Pointer) PasskeyRegistrationCredentialExtensionOutput {
	return PasskeyRegistrationCredentialExtensionOutput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PasskeyRegistrationCredentialExtensionOutput */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredentialExtensionOutput-c.class/initWithLargeBlobOutput:
func NewPasskeyRegistrationCredentialExtensionOutputWithLargeBlobOutput(largeBlob IASAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput) PasskeyRegistrationCredentialExtensionOutput {
	instance := getPasskeyRegistrationCredentialExtensionOutputClass().Alloc()
	rv := objc.Send[PasskeyRegistrationCredentialExtensionOutput](instance.ID, objc.Sel("initWithLargeBlobOutput:"), largeBlob)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPasskeyRegistrationCredentialExtensionOutputWithLargeBlobOutput */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PasskeyRegistrationCredentialExtensionOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PasskeyRegistrationCredentialExtensionOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PasskeyRegistrationCredentialExtensionOutput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PasskeyRegistrationCredentialExtensionOutput */

// Output for operation during passkey registration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyRegistrationCredentialExtensionOutput-c.class/largeBlobRegistrationOutput
func (p_ PasskeyRegistrationCredentialExtensionOutput) LargeBlobRegistrationOutput() IASAuthorizationPublicKeyCredentialLargeBlobRegistrationOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobRegistrationOutput](p_.ID, objc.Sel("largeBlobRegistrationOutput"))
	return rv
}/* debug [instance_properties/getter]: largeBlobRegistrationOutput */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASPasskeyRegistrationCredentialExtensionOutput */


