// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [AuthorizationSecurityKeyPublicKeyCredentialDescriptor] class.
type IAuthorizationSecurityKeyPublicKeyCredentialDescriptor interface {
	objectivec.IObject
}

// An object that holds public key credential transport information.
//
// This class ties together a credential and its corresponding transport types (USB, NFC, Bluetooth, or all of them).
//
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

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationSecurityKeyPublicKeyCredentialDescriptorClass) Alloc() AuthorizationSecurityKeyPublicKeyCredentialDescriptor {
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialDescriptor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates the object with the credential ID and the array of transports.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialDescriptor/init(credentialID:transports:)
func NewAuthorizationSecurityKeyPublicKeyCredentialDescriptorWithCredentialIDTransports(credentialID unsafe.Pointer, allowedTransports unsafe.Pointer) AuthorizationSecurityKeyPublicKeyCredentialDescriptor {
	instance := getAuthorizationSecurityKeyPublicKeyCredentialDescriptorClass().Alloc()
	rv := objc.Send[AuthorizationSecurityKeyPublicKeyCredentialDescriptor](instance.ID, objc.Sel("initWithCredentialID:transports:"), credentialID, allowedTransports)
	rv.Autorelease()
	return rv
}


// The array of transport types.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialDescriptor/transports
func (a_ AuthorizationSecurityKeyPublicKeyCredentialDescriptor) Transports() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("transports"))
	return rv
}


// SetTransports sets the value of the transports property.
// The array of transport types.

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationSecurityKeyPublicKeyCredentialDescriptor/transports
func (a_ AuthorizationSecurityKeyPublicKeyCredentialDescriptor) SetTransports(value []string) {
	// Convert Go slice to NSArray
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
}


