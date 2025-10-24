// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASPasskeyAssertionCredentialExtensionOutput */


/* debug [class_header]: Header for ASPasskeyAssertionCredentialExtensionOutput */
// The class instance for the [PasskeyAssertionCredentialExtensionOutput] class.
var (
	PasskeyAssertionCredentialExtensionOutputClass     _PasskeyAssertionCredentialExtensionOutputClass
	PasskeyAssertionCredentialExtensionOutputClassOnce sync.Once
)

func getPasskeyAssertionCredentialExtensionOutputClass() _PasskeyAssertionCredentialExtensionOutputClass {
	PasskeyAssertionCredentialExtensionOutputClassOnce.Do(func() {
		PasskeyAssertionCredentialExtensionOutputClass = _PasskeyAssertionCredentialExtensionOutputClass{objc.GetClass("ASPasskeyAssertionCredentialExtensionOutput")}
	})
	return PasskeyAssertionCredentialExtensionOutputClass
}

type _PasskeyAssertionCredentialExtensionOutputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PasskeyAssertionCredentialExtensionOutput */
// An interface definition for the [PasskeyAssertionCredentialExtensionOutput] class.
type IPasskeyAssertionCredentialExtensionOutput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PasskeyAssertionCredentialExtensionOutput */
	// properties:
	LargeBlobAssertionOutput() IASAuthorizationPublicKeyCredentialLargeBlobAssertionOutput
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PasskeyAssertionCredentialExtensionOutput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PasskeyAssertionCredentialExtensionOutput */
// Alloc allocates a new instance without initialization.
func (pc _PasskeyAssertionCredentialExtensionOutputClass) Alloc() PasskeyAssertionCredentialExtensionOutput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionOutput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PasskeyAssertionCredentialExtensionOutputClass) New() PasskeyAssertionCredentialExtensionOutput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionOutput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasskeyAssertionCredentialExtensionOutput) Init() PasskeyAssertionCredentialExtensionOutput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionOutput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasskeyAssertionCredentialExtensionOutput) Autorelease() PasskeyAssertionCredentialExtensionOutput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionOutput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasskeyAssertionCredentialExtensionOutput creates a new PasskeyAssertionCredentialExtensionOutput instance.
func NewPasskeyAssertionCredentialExtensionOutput() PasskeyAssertionCredentialExtensionOutput {
	return getPasskeyAssertionCredentialExtensionOutputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PasskeyAssertionCredentialExtensionOutput */
// This class encapsulates output for various WebAuthn extensions used during passkey assertion.


// This class encapsulates output for various WebAuthn extensions used during passkey assertion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredentialExtensionOutput-c.class
type PasskeyAssertionCredentialExtensionOutput struct {
	objectivec.Object
}

// PasskeyAssertionCredentialExtensionOutputFrom constructs a [PasskeyAssertionCredentialExtensionOutput] from an unsafe.Pointer.
//
// This class encapsulates output for various WebAuthn extensions used during passkey assertion.
func PasskeyAssertionCredentialExtensionOutputFrom(ptr unsafe.Pointer) PasskeyAssertionCredentialExtensionOutput {
	return PasskeyAssertionCredentialExtensionOutput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PasskeyAssertionCredentialExtensionOutput */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredentialExtensionOutput-c.class/initWithLargeBlobOutput:
func NewPasskeyAssertionCredentialExtensionOutputWithLargeBlobOutput(largeBlob IASAuthorizationPublicKeyCredentialLargeBlobAssertionOutput) PasskeyAssertionCredentialExtensionOutput {
	instance := getPasskeyAssertionCredentialExtensionOutputClass().Alloc()
	rv := objc.Send[PasskeyAssertionCredentialExtensionOutput](instance.ID, objc.Sel("initWithLargeBlobOutput:"), largeBlob)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewPasskeyAssertionCredentialExtensionOutputWithLargeBlobOutput */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PasskeyAssertionCredentialExtensionOutput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PasskeyAssertionCredentialExtensionOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PasskeyAssertionCredentialExtensionOutput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PasskeyAssertionCredentialExtensionOutput */

// Output for operation during passkey assertion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredentialExtensionOutput-c.class/largeBlobAssertionOutput
func (p_ PasskeyAssertionCredentialExtensionOutput) LargeBlobAssertionOutput() IASAuthorizationPublicKeyCredentialLargeBlobAssertionOutput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionOutput](p_.ID, objc.Sel("largeBlobAssertionOutput"))
	return rv
}/* debug [instance_properties/getter]: largeBlobAssertionOutput */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASPasskeyAssertionCredentialExtensionOutput */


