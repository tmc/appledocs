// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationPlatformPublicKeyCredentialDescriptor */


/* debug [class_header]: Header for ASAuthorizationPlatformPublicKeyCredentialDescriptor */
// The class instance for the [AuthorizationPlatformPublicKeyCredentialDescriptor] class.
var (
	AuthorizationPlatformPublicKeyCredentialDescriptorClass     _AuthorizationPlatformPublicKeyCredentialDescriptorClass
	AuthorizationPlatformPublicKeyCredentialDescriptorClassOnce sync.Once
)

func getAuthorizationPlatformPublicKeyCredentialDescriptorClass() _AuthorizationPlatformPublicKeyCredentialDescriptorClass {
	AuthorizationPlatformPublicKeyCredentialDescriptorClassOnce.Do(func() {
		AuthorizationPlatformPublicKeyCredentialDescriptorClass = _AuthorizationPlatformPublicKeyCredentialDescriptorClass{objc.GetClass("ASAuthorizationPlatformPublicKeyCredentialDescriptor")}
	})
	return AuthorizationPlatformPublicKeyCredentialDescriptorClass
}

type _AuthorizationPlatformPublicKeyCredentialDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationPlatformPublicKeyCredentialDescriptor */
// An interface definition for the [AuthorizationPlatformPublicKeyCredentialDescriptor] class.
type IAuthorizationPlatformPublicKeyCredentialDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationPlatformPublicKeyCredentialDescriptor */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationPlatformPublicKeyCredentialDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationPlatformPublicKeyCredentialDescriptor */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationPlatformPublicKeyCredentialDescriptorClass) Alloc() AuthorizationPlatformPublicKeyCredentialDescriptor {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationPlatformPublicKeyCredentialDescriptorClass) New() AuthorizationPlatformPublicKeyCredentialDescriptor {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationPlatformPublicKeyCredentialDescriptor) Init() AuthorizationPlatformPublicKeyCredentialDescriptor {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationPlatformPublicKeyCredentialDescriptor) Autorelease() AuthorizationPlatformPublicKeyCredentialDescriptor {
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationPlatformPublicKeyCredentialDescriptor creates a new AuthorizationPlatformPublicKeyCredentialDescriptor instance.
func NewAuthorizationPlatformPublicKeyCredentialDescriptor() AuthorizationPlatformPublicKeyCredentialDescriptor {
	return getAuthorizationPlatformPublicKeyCredentialDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationPlatformPublicKeyCredentialDescriptor */
// An object that holds the credential.
//
// This class holds the platform credential identifier.


// An object that holds the credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialDescriptor
type AuthorizationPlatformPublicKeyCredentialDescriptor struct {
	objectivec.Object
}

// AuthorizationPlatformPublicKeyCredentialDescriptorFrom constructs a [AuthorizationPlatformPublicKeyCredentialDescriptor] from an unsafe.Pointer.
//
// An object that holds the credential.
func AuthorizationPlatformPublicKeyCredentialDescriptorFrom(ptr unsafe.Pointer) AuthorizationPlatformPublicKeyCredentialDescriptor {
	return AuthorizationPlatformPublicKeyCredentialDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationPlatformPublicKeyCredentialDescriptor */

// Creates the descriptor with a credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationPlatformPublicKeyCredentialDescriptor/init(credentialID:)
func NewAuthorizationPlatformPublicKeyCredentialDescriptorWithCredentialID(credentialID objc.IObject /* cross-framework: NSData */) AuthorizationPlatformPublicKeyCredentialDescriptor {
	instance := getAuthorizationPlatformPublicKeyCredentialDescriptorClass().Alloc()
	rv := objc.Send[AuthorizationPlatformPublicKeyCredentialDescriptor](instance.ID, objc.Sel("initWithCredentialID:"), credentialID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAuthorizationPlatformPublicKeyCredentialDescriptorWithCredentialID */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationPlatformPublicKeyCredentialDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationPlatformPublicKeyCredentialDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationPlatformPublicKeyCredentialDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationPlatformPublicKeyCredentialDescriptor */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationPlatformPublicKeyCredentialDescriptor */


