// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAuthorizationWebBrowserPlatformPublicKeyCredential */


/* debug [class_header]: Header for ASAuthorizationWebBrowserPlatformPublicKeyCredential */
// The class instance for the [AuthorizationWebBrowserPlatformPublicKeyCredential] class.
var (
	AuthorizationWebBrowserPlatformPublicKeyCredentialClass     _AuthorizationWebBrowserPlatformPublicKeyCredentialClass
	AuthorizationWebBrowserPlatformPublicKeyCredentialClassOnce sync.Once
)

func getAuthorizationWebBrowserPlatformPublicKeyCredentialClass() _AuthorizationWebBrowserPlatformPublicKeyCredentialClass {
	AuthorizationWebBrowserPlatformPublicKeyCredentialClassOnce.Do(func() {
		AuthorizationWebBrowserPlatformPublicKeyCredentialClass = _AuthorizationWebBrowserPlatformPublicKeyCredentialClass{objc.GetClass("ASAuthorizationWebBrowserPlatformPublicKeyCredential")}
	})
	return AuthorizationWebBrowserPlatformPublicKeyCredentialClass
}

type _AuthorizationWebBrowserPlatformPublicKeyCredentialClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AuthorizationWebBrowserPlatformPublicKeyCredential */
// An interface definition for the [AuthorizationWebBrowserPlatformPublicKeyCredential] class.
type IAuthorizationWebBrowserPlatformPublicKeyCredential interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AuthorizationWebBrowserPlatformPublicKeyCredential */
	// properties:
	CredentialID() objc.IObject /* cross-framework: NSData */
	CustomTitle() objc.IObject /* cross-framework: NSString */
	Name() objc.IObject /* cross-framework: NSString */
	ProviderName() objc.IObject /* cross-framework: NSString */
	RelyingParty() objc.IObject /* cross-framework: NSString */
	UserHandle() objc.IObject /* cross-framework: NSData */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AuthorizationWebBrowserPlatformPublicKeyCredential */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AuthorizationWebBrowserPlatformPublicKeyCredential */
// Alloc allocates a new instance without initialization.
func (ac _AuthorizationWebBrowserPlatformPublicKeyCredentialClass) Alloc() AuthorizationWebBrowserPlatformPublicKeyCredential {
	rv := objc.Send[AuthorizationWebBrowserPlatformPublicKeyCredential](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AuthorizationWebBrowserPlatformPublicKeyCredentialClass) New() AuthorizationWebBrowserPlatformPublicKeyCredential {
	rv := objc.Send[AuthorizationWebBrowserPlatformPublicKeyCredential](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AuthorizationWebBrowserPlatformPublicKeyCredential) Init() AuthorizationWebBrowserPlatformPublicKeyCredential {
	rv := objc.Send[AuthorizationWebBrowserPlatformPublicKeyCredential](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AuthorizationWebBrowserPlatformPublicKeyCredential) Autorelease() AuthorizationWebBrowserPlatformPublicKeyCredential {
	rv := objc.Send[AuthorizationWebBrowserPlatformPublicKeyCredential](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAuthorizationWebBrowserPlatformPublicKeyCredential creates a new AuthorizationWebBrowserPlatformPublicKeyCredential instance.
func NewAuthorizationWebBrowserPlatformPublicKeyCredential() AuthorizationWebBrowserPlatformPublicKeyCredential {
	return getAuthorizationWebBrowserPlatformPublicKeyCredentialClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AuthorizationWebBrowserPlatformPublicKeyCredential */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPlatformPublicKeyCredential-c.class
type AuthorizationWebBrowserPlatformPublicKeyCredential struct {
	objectivec.Object
}

// AuthorizationWebBrowserPlatformPublicKeyCredentialFrom constructs a [AuthorizationWebBrowserPlatformPublicKeyCredential] from an unsafe.Pointer.
func AuthorizationWebBrowserPlatformPublicKeyCredentialFrom(ptr unsafe.Pointer) AuthorizationWebBrowserPlatformPublicKeyCredential {
	return AuthorizationWebBrowserPlatformPublicKeyCredential{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AuthorizationWebBrowserPlatformPublicKeyCredential *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AuthorizationWebBrowserPlatformPublicKeyCredential */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AuthorizationWebBrowserPlatformPublicKeyCredential */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AuthorizationWebBrowserPlatformPublicKeyCredential */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AuthorizationWebBrowserPlatformPublicKeyCredential */

// The identifier the operating system uses for this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPlatformPublicKeyCredential-c.class/credentialID
func (a_ AuthorizationWebBrowserPlatformPublicKeyCredential) CredentialID() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("credentialID"))
	return rv
}/* debug [instance_properties/getter]: credentialID */


// A string the person can supply to describe this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPlatformPublicKeyCredential-c.class/customTitle
func (a_ AuthorizationWebBrowserPlatformPublicKeyCredential) CustomTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("customTitle"))
	return rv
}/* debug [instance_properties/getter]: customTitle */


// The user name for the account associated with this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPlatformPublicKeyCredential-c.class/name
func (a_ AuthorizationWebBrowserPlatformPublicKeyCredential) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name of the app that manages this credential, or “iCloud Keychain” if it’s the operating system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPlatformPublicKeyCredential-c.class/providerName
func (a_ AuthorizationWebBrowserPlatformPublicKeyCredential) ProviderName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("providerName"))
	return rv
}/* debug [instance_properties/getter]: providerName */


// The relying party that issues challenges for this credential.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPlatformPublicKeyCredential-c.class/relyingParty
func (a_ AuthorizationWebBrowserPlatformPublicKeyCredential) RelyingParty() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("relyingParty"))
	return rv
}/* debug [instance_properties/getter]: relyingParty */


// A unique identifier for the user account at the relying party.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAuthorizationWebBrowserPlatformPublicKeyCredential-c.class/userHandle
func (a_ AuthorizationWebBrowserPlatformPublicKeyCredential) UserHandle() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](a_.ID, objc.Sel("userHandle"))
	return rv
}/* debug [instance_properties/getter]: userHandle */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAuthorizationWebBrowserPlatformPublicKeyCredential */



