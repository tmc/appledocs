// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [AuthorizationProviderExtensionLoginManager] class.
var (
	AuthorizationProviderExtensionLoginManagerClass     _AuthorizationProviderExtensionLoginManagerClass
	AuthorizationProviderExtensionLoginManagerClassOnce sync.Once
)

func getAuthorizationProviderExtensionLoginManagerClass() _AuthorizationProviderExtensionLoginManagerClass {
	AuthorizationProviderExtensionLoginManagerClassOnce.Do(func() {
		AuthorizationProviderExtensionLoginManagerClass = _AuthorizationProviderExtensionLoginManagerClass{objc.GetClass("ASAuthorizationProviderExtensionLoginManager")}
	})
	return AuthorizationProviderExtensionLoginManagerClass
}

type _AuthorizationProviderExtensionLoginManagerClass struct {
	class objc.Class
}

// An interface definition for the [AuthorizationProviderExtensionLoginManager] class.
type IAuthorizationProviderExtensionLoginManager interface {
	objectivec.IObject
	BeginKeyRotationForKeyType(keyType unsafe.Pointer) unsafe.Pointer
	CompleteKeyRotationForKeyType(keyType unsafe.Pointer)
	SaveCertificateKeyType(certificate unsafe.Pointer, keyType unsafe.Pointer)
}

// An interface to maintain platform single sign-on (SSO) during authentication and registration.
//
// Use this class to perform registration and authentication tasks, and to repair registrations.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager
type AuthorizationProviderExtensionLoginManager struct {
	objectivec.Object
}

// AuthorizationProviderExtensionLoginManagerFrom constructs a [AuthorizationProviderExtensionLoginManager] from an unsafe.Pointer.
//
// An interface to maintain platform single sign-on (SSO) during authentication and registration.
func AuthorizationProviderExtensionLoginManagerFrom(ptr unsafe.Pointer) AuthorizationProviderExtensionLoginManager {
	return AuthorizationProviderExtensionLoginManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AuthorizationProviderExtensionLoginManagerClass) Alloc() AuthorizationProviderExtensionLoginManager {
	rv := objc.Send[AuthorizationProviderExtensionLoginManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AuthorizationProviderExtensionLoginManagerClass) New() AuthorizationProviderExtensionLoginManager {
	rv := objc.Send[AuthorizationProviderExtensionLoginManager](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationProviderExtensionLoginManager) Init() AuthorizationProviderExtensionLoginManager {
	rv := objc.Send[AuthorizationProviderExtensionLoginManager](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationProviderExtensionLoginManager) Autorelease() AuthorizationProviderExtensionLoginManager {
	rv := objc.Send[AuthorizationProviderExtensionLoginManager](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationProviderExtensionLoginManager creates a new AuthorizationProviderExtensionLoginManager instance.
func NewAuthorizationProviderExtensionLoginManager() AuthorizationProviderExtensionLoginManager {
	return getAuthorizationProviderExtensionLoginManagerClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/beginKeyRotation(_:)
func (a_ AuthorizationProviderExtensionLoginManager) BeginKeyRotationForKeyType(keyType unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("beginKeyRotationForKeyType:"), keyType)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/completeKeyRotation(_:)
func (a_ AuthorizationProviderExtensionLoginManager) CompleteKeyRotationForKeyType(keyType unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("completeKeyRotationForKeyType:"), keyType)
}

// Saves the provided certificate for the key type.
//
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationProviderExtensionLoginManager/saveCertificate(_:keyType:)
func (a_ AuthorizationProviderExtensionLoginManager) SaveCertificateKeyType(certificate unsafe.Pointer, keyType unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("saveCertificate:keyType:"), certificate, keyType)
}



