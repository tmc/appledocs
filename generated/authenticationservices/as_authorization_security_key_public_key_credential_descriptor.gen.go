// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationSecurityKeyPublicKeyCredentialDescriptor */


/* debug [class_header]: Header for ASAuthorizationSecurityKeyPublicKeyCredentialDescriptor */
// The class instance for the [AuthorizationSecurityKeyPublicKeyCredentialDescriptor] class.
var (
	AuthorizationSecurityKeyPublicKeyCredentialDescriptorClass     _AuthorizationSecurityKeyPublicKeyCredentialDescriptorClass
	AuthorizationSecurityKeyPublicKeyCredentialDescriptorClassOnce sync.Once
)

func getAuthorizationSecurityKeyPublicKeyCredentialDescriptorClass() _AuthorizationSecurityKeyPublicKeyCredentialDescriptorClass {
	AuthorizationSecurityKeyPublicKeyCredentialDescriptorClassOnce.Do(func() {
		AuthorizationSecurityKeyPublicKeyCredentialDescriptorClass = _AuthorizationSecurityKeyPublicKeyCredentialDescriptorClass{objc.GetClass("ASAuthorizationSecurityKeyPublicKeyCredentialDescriptor")}
	})
	return AuthorizationSecurityKeyPublicKeyCredentialDescriptorClass
}

type _AuthorizationSecurityKeyPublicKeyCredentialDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationSecurityKeyPublicKeyCredentialDescriptor */
// An interface definition for the [AuthorizationSecurityKeyPublicKeyCredentialDescriptor] class.
type IAuthorizationSecurityKeyPublicKeyCredentialDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationSecurityKeyPublicKeyCredentialDescriptor */
	// properties:
	Transports() []string
	SetTransports(value []string)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationSecurityKeyPublicKeyCredentialDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationSecurityKeyPublicKeyCredentialDescriptor */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSecurityKeyPublicKeyCredentialDescriptorClass) Alloc() AuthorizationSecurityKeyPublicKeyCredentialDescriptor {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationSecurityKeyPublicKeyCredentialDescriptorClass) New() AuthorizationSecurityKeyPublicKeyCredentialDescriptor {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialDescriptor](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationSecurityKeyPublicKeyCredentialDescriptor) Init() AuthorizationSecurityKeyPublicKeyCredentialDescriptor {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialDescriptor](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationSecurityKeyPublicKeyCredentialDescriptor) Autorelease() AuthorizationSecurityKeyPublicKeyCredentialDescriptor {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialDescriptor](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationSecurityKeyPublicKeyCredentialDescriptor creates a new AuthorizationSecurityKeyPublicKeyCredentialDescriptor instance.
func NewAuthorizationSecurityKeyPublicKeyCredentialDescriptor() AuthorizationSecurityKeyPublicKeyCredentialDescriptor {
	return getAuthorizationSecurityKeyPublicKeyCredentialDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationSecurityKeyPublicKeyCredentialDescriptor */
// An object that holds public key credential transport information.
//
// This class ties together a credential and its corresponding transport types (USB, NFC, Bluetooth, or all of them).


// An object that holds public key credential transport information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialDescriptor
type AuthorizationSecurityKeyPublicKeyCredentialDescriptor struct {
	objectivec.Object
}

// AuthorizationSecurityKeyPublicKeyCredentialDescriptorFrom constructs a [AuthorizationSecurityKeyPublicKeyCredentialDescriptor] from an unsafe.Pointer.
//
// An object that holds public key credential transport information.
func AuthorizationSecurityKeyPublicKeyCredentialDescriptorFrom(ptr unsafe.Pointer) AuthorizationSecurityKeyPublicKeyCredentialDescriptor {
	return AuthorizationSecurityKeyPublicKeyCredentialDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationSecurityKeyPublicKeyCredentialDescriptor */

// Creates the object with the credential ID and the array of transports.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialDescriptor/init(credentialID:transports:)
func NewAuthorizationSecurityKeyPublicKeyCredentialDescriptorWithCredentialIDTransports(credentialID objc.IObject /* cross-framework: NSData */, allowedTransports []string) AuthorizationSecurityKeyPublicKeyCredentialDescriptor {
	instance := getAuthorizationSecurityKeyPublicKeyCredentialDescriptorClass().Alloc()
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialDescriptor](instance.ID, objc.Sel("initWithCredentialID:transports:"), credentialID, allowedTransports)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAuthorizationSecurityKeyPublicKeyCredentialDescriptorWithCredentialIDTransports */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationSecurityKeyPublicKeyCredentialDescriptor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationSecurityKeyPublicKeyCredentialDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationSecurityKeyPublicKeyCredentialDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationSecurityKeyPublicKeyCredentialDescriptor */

// The array of transport types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialDescriptor/transports
func (a_ AuthorizationSecurityKeyPublicKeyCredentialDescriptor) Transports() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("transports"))
	return rv
}/* debug [instance_properties/getter]: transports */


// The array of transport types.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialDescriptor/transports
func (a_ AuthorizationSecurityKeyPublicKeyCredentialDescriptor) SetTransports(value []string) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setTransports:"), nsArray)
}/* debug [instance_properties/setter]: transports */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationSecurityKeyPublicKeyCredentialDescriptor */


