// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class ASAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest */


/* debug [class_header]: Header for ASAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest */
// The class instance for the [AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest] class.
var (
	AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestClass     _AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestClass
	AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestClassOnce sync.Once
)

func getAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestClass() _AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestClass {
	AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestClassOnce.Do(func() {
		AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestClass = _AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestClass{objc.GetClass("ASAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest")}
	})
	return AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestClass
}

type _AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest */
// An interface definition for the [AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest] class.
type IAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest interface {
	IAccountAuthenticationModificationRequest
	
/* debug [class_interface_properties]: Properties for AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest */
// Alloc allocates a new instance without initialization.
func (ac _AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestClass) Alloc() AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest {
	rv := objc.Send[AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestClass) New() AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest {
	rv := objc.Send[AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest) Init() AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest {
	rv := objc.Send[AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest) Autorelease() AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest {
	rv := objc.Send[AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest creates a new AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest instance.
func NewAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest() AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest {
	return getAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest */
// A request to upgrade from using a password to using Sign in with Apple.
//
// Your app uses this class to initiate an upgrade to Sign in with Apple. After creating the request, your app initiates the upgrade process by instantiating an object and calling on it. The system invokes your authentication modification extension to complete the upgrade.


// A request to upgrade from using a password to using Sign in with Apple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest
type AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest struct {
	AccountAuthenticationModificationRequest
}

// AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestFrom constructs a [AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest] from an unsafe.Pointer.
//
// A request to upgrade from using a password to using Sign in with Apple.
func AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestFrom(ptr unsafe.Pointer) AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest {
	return AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest{
		AccountAuthenticationModificationRequest: AccountAuthenticationModificationRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest */

// Creates a request to upgrade from using passwords to using Sign in with Apple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest/init(user:serviceIdentifier:userInfo:)
func NewAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestWithUserServiceIdentifierUserInfo(user objc.IObject /* cross-framework: NSString */, serviceIdentifier IASCredentialServiceIdentifier, userInfo objc.IObject /* cross-framework: NSDictionary */) AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest {
	instance := getAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestClass().Alloc()
	rv := objc.Send[AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest](instance.ID, objc.Sel("initWithUser:serviceIdentifier:userInfo:"), user, serviceIdentifier, userInfo)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequestWithUserServiceIdentifierUserInfo */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAccountAuthenticationModificationReplacePasswordWithSignInWithAppleRequest */


