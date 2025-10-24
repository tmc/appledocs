// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASPasskeyAssertionCredentialExtensionInput */


/* debug [class_header]: Header for ASPasskeyAssertionCredentialExtensionInput */
// The class instance for the [PasskeyAssertionCredentialExtensionInput] class.
var (
	PasskeyAssertionCredentialExtensionInputClass     _PasskeyAssertionCredentialExtensionInputClass
	PasskeyAssertionCredentialExtensionInputClassOnce sync.Once
)

func getPasskeyAssertionCredentialExtensionInputClass() _PasskeyAssertionCredentialExtensionInputClass {
	PasskeyAssertionCredentialExtensionInputClassOnce.Do(func() {
		PasskeyAssertionCredentialExtensionInputClass = _PasskeyAssertionCredentialExtensionInputClass{objc.GetClass("ASPasskeyAssertionCredentialExtensionInput")}
	})
	return PasskeyAssertionCredentialExtensionInputClass
}

type _PasskeyAssertionCredentialExtensionInputClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PasskeyAssertionCredentialExtensionInput */
// An interface definition for the [PasskeyAssertionCredentialExtensionInput] class.
type IPasskeyAssertionCredentialExtensionInput interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PasskeyAssertionCredentialExtensionInput */
	// properties:
	LargeBlob() IASAuthorizationPublicKeyCredentialLargeBlobAssertionInput
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PasskeyAssertionCredentialExtensionInput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PasskeyAssertionCredentialExtensionInput */
// Alloc allocates a new instance without initialization.
func (pc _PasskeyAssertionCredentialExtensionInputClass) Alloc() PasskeyAssertionCredentialExtensionInput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionInput](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PasskeyAssertionCredentialExtensionInputClass) New() PasskeyAssertionCredentialExtensionInput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionInput](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PasskeyAssertionCredentialExtensionInput) Init() PasskeyAssertionCredentialExtensionInput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionInput](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PasskeyAssertionCredentialExtensionInput) Autorelease() PasskeyAssertionCredentialExtensionInput {
	rv := objc.Send[PasskeyAssertionCredentialExtensionInput](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPasskeyAssertionCredentialExtensionInput creates a new PasskeyAssertionCredentialExtensionInput instance.
func NewPasskeyAssertionCredentialExtensionInput() PasskeyAssertionCredentialExtensionInput {
	return getPasskeyAssertionCredentialExtensionInputClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PasskeyAssertionCredentialExtensionInput */
// This class encapsulates input for various WebAuthn extensions during passkey assertion.


// This class encapsulates input for various WebAuthn extensions during passkey assertion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredentialExtensionInput-c.class
type PasskeyAssertionCredentialExtensionInput struct {
	objectivec.Object
}

// PasskeyAssertionCredentialExtensionInputFrom constructs a [PasskeyAssertionCredentialExtensionInput] from an unsafe.Pointer.
//
// This class encapsulates input for various WebAuthn extensions during passkey assertion.
func PasskeyAssertionCredentialExtensionInputFrom(ptr unsafe.Pointer) PasskeyAssertionCredentialExtensionInput {
	return PasskeyAssertionCredentialExtensionInput{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PasskeyAssertionCredentialExtensionInput *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PasskeyAssertionCredentialExtensionInput */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PasskeyAssertionCredentialExtensionInput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PasskeyAssertionCredentialExtensionInput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PasskeyAssertionCredentialExtensionInput */

// Input for the extension in passkey assertion requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASPasskeyAssertionCredentialExtensionInput-c.class/largeBlob
func (p_ PasskeyAssertionCredentialExtensionInput) LargeBlob() IASAuthorizationPublicKeyCredentialLargeBlobAssertionInput {
	rv := objc.Send[AuthorizationPublicKeyCredentialLargeBlobAssertionInput](p_.ID, objc.Sel("largeBlob"))
	return rv
}/* debug [instance_properties/getter]: largeBlob */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASPasskeyAssertionCredentialExtensionInput */



