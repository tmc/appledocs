// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationWebBrowserPublicKeyCredentialManager */


/* debug [class_header]: Header for ASAuthorizationWebBrowserPublicKeyCredentialManager */
// The class instance for the [AuthorizationWebBrowserPublicKeyCredentialManager] class.
var (
	AuthorizationWebBrowserPublicKeyCredentialManagerClass     _AuthorizationWebBrowserPublicKeyCredentialManagerClass
	AuthorizationWebBrowserPublicKeyCredentialManagerClassOnce sync.Once
)

func getAuthorizationWebBrowserPublicKeyCredentialManagerClass() _AuthorizationWebBrowserPublicKeyCredentialManagerClass {
	AuthorizationWebBrowserPublicKeyCredentialManagerClassOnce.Do(func() {
		AuthorizationWebBrowserPublicKeyCredentialManagerClass = _AuthorizationWebBrowserPublicKeyCredentialManagerClass{objc.GetClass("ASAuthorizationWebBrowserPublicKeyCredentialManager")}
	})
	return AuthorizationWebBrowserPublicKeyCredentialManagerClass
}

type _AuthorizationWebBrowserPublicKeyCredentialManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationWebBrowserPublicKeyCredentialManager */
// An interface definition for the [AuthorizationWebBrowserPublicKeyCredentialManager] class.
type IAuthorizationWebBrowserPublicKeyCredentialManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationWebBrowserPublicKeyCredentialManager */
	// properties:
	AuthorizationStateForPlatformCredentials() AuthorizationWebBrowserPublicKeyCredentialManagerAuthorizationState
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationWebBrowserPublicKeyCredentialManager */
	// methods:
	PlatformCredentialsForRelyingPartyCompletionHandler(relyingParty objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer)
	RequestAuthorizationForPublicKeyCredentials(completionHandler unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationWebBrowserPublicKeyCredentialManager */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationWebBrowserPublicKeyCredentialManagerClass) Alloc() AuthorizationWebBrowserPublicKeyCredentialManager {
	rv := objc.Send[AuthorizationWebBrowserPublicKeyCredentialManager](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationWebBrowserPublicKeyCredentialManagerClass) New() AuthorizationWebBrowserPublicKeyCredentialManager {
	rv := objc.Send[AuthorizationWebBrowserPublicKeyCredentialManager](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationWebBrowserPublicKeyCredentialManager) Init() AuthorizationWebBrowserPublicKeyCredentialManager {
	rv := objc.Send[AuthorizationWebBrowserPublicKeyCredentialManager](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationWebBrowserPublicKeyCredentialManager) Autorelease() AuthorizationWebBrowserPublicKeyCredentialManager {
	rv := objc.Send[AuthorizationWebBrowserPublicKeyCredentialManager](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationWebBrowserPublicKeyCredentialManager creates a new AuthorizationWebBrowserPublicKeyCredentialManager instance.
func NewAuthorizationWebBrowserPublicKeyCredentialManager() AuthorizationWebBrowserPublicKeyCredentialManager {
	return getAuthorizationWebBrowserPublicKeyCredentialManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationWebBrowserPublicKeyCredentialManager */
// A class that you use to request access to a person’s passkeys in a web browser, and that reports on the access status.


// A class that you use to request access to a person’s passkeys in a web browser, and that reports on the access status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPublicKeyCredentialManager
type AuthorizationWebBrowserPublicKeyCredentialManager struct {
	objectivec.Object
}

// AuthorizationWebBrowserPublicKeyCredentialManagerFrom constructs a [AuthorizationWebBrowserPublicKeyCredentialManager] from an unsafe.Pointer.
//
// A class that you use to request access to a person’s passkeys in a web browser, and that reports on the access status.
func AuthorizationWebBrowserPublicKeyCredentialManagerFrom(ptr unsafe.Pointer) AuthorizationWebBrowserPublicKeyCredentialManager {
	return AuthorizationWebBrowserPublicKeyCredentialManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationWebBrowserPublicKeyCredentialManager */
/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationWebBrowserPublicKeyCredentialManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationWebBrowserPublicKeyCredentialManager */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPublicKeyCredentialManager/isDeviceConfiguredForPasskeys-70sni
func (ac _AuthorizationWebBrowserPublicKeyCredentialManagerClass) IsDeviceConfiguredForPasskeys() bool {
	rv := objc.Send[bool](objc.ID(ac.class), objc.Sel("isDeviceConfiguredForPasskeys"))
	return rv
}/* debug [class_properties_class/property]: isDeviceConfiguredForPasskeys */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationWebBrowserPublicKeyCredentialManager */

// Gets a list of passkeys available for authenticating with the given relying party.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPublicKeyCredentialManager/platformCredentialsForRelyingParty:completionHandler:
func (a_ AuthorizationWebBrowserPublicKeyCredentialManager) PlatformCredentialsForRelyingPartyCompletionHandler(relyingParty objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("platformCredentialsForRelyingParty:completionHandler:"), relyingParty, completionHandler)
}/* debug [instance_methods/method]: PlatformCredentialsForRelyingPartyCompletionHandler */


// Requests a person’s permission to use their passkeys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPublicKeyCredentialManager/requestAuthorizationForPublicKeyCredentials(_:)
func (a_ AuthorizationWebBrowserPublicKeyCredentialManager) RequestAuthorizationForPublicKeyCredentials(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("requestAuthorizationForPublicKeyCredentials:"), completionHandler)
}/* debug [instance_methods/method]: RequestAuthorizationForPublicKeyCredentials */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationWebBrowserPublicKeyCredentialManager */

// Returns a value that indicates whether the browser app has access to a person’s passkeys.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPublicKeyCredentialManager/authorizationStateForPlatformCredentials
func (a_ AuthorizationWebBrowserPublicKeyCredentialManager) AuthorizationStateForPlatformCredentials() AuthorizationWebBrowserPublicKeyCredentialManagerAuthorizationState {
	rv := objc.Send[AuthorizationWebBrowserPublicKeyCredentialManagerAuthorizationState](a_.ID, objc.Sel("authorizationStateForPlatformCredentials"))
	return rv
}/* debug [instance_properties/getter]: authorizationStateForPlatformCredentials */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPublicKeyCredentialManager/isDeviceConfiguredForPasskeys-70sni
func (a_ AuthorizationWebBrowserPublicKeyCredentialManager) IsDeviceConfiguredForPasskeys() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isDeviceConfiguredForPasskeys"))
	return rv
}/* debug [instance_properties/getter]: isDeviceConfiguredForPasskeys */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationWebBrowserPublicKeyCredentialManager */


