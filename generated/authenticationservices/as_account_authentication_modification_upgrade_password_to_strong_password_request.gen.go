// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class ASAccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest */


/* debug [class_header]: Header for ASAccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest */
// The class instance for the [AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest] class.
var (
	AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestClass     _AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestClass
	AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestClassOnce sync.Once
)

func getAccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestClass() _AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestClass {
	AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestClassOnce.Do(func() {
		AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestClass = _AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestClass{objc.GetClass("ASAccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest")}
	})
	return AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestClass
}

type _AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest */
// An interface definition for the [AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest] class.
type IAccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest interface {
	IAccountAuthenticationModificationRequest
	
/* debug [class_interface_properties]: Properties for AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest */
// Alloc allocates a new instance without initialization.
func (ac _AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestClass) Alloc() AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest {
	rv := objc.Send[AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestClass) New() AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest {
	rv := objc.Send[AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest) Init() AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest {
	rv := objc.Send[AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest) Autorelease() AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest {
	rv := objc.Send[AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest creates a new AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest instance.
func NewAccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest() AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest {
	return getAccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest */
// A request to automatically upgrade from a weak password to a strong password.
//
// Your app uses this class to initiate an upgrade from a weak password to a strong system-generated one. After creating the request, your app initiates the upgrade process by instantiating an object and calling on it. The system invokes your authentication modification extension to complete the upgrade. For details about how to enforce requirements on the password, such as minimum length or requiring both letters and numbers, see .


// A request to automatically upgrade from a weak password to a strong password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest
type AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest struct {
	AccountAuthenticationModificationRequest
}

// AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestFrom constructs a [AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest] from an unsafe.Pointer.
//
// A request to automatically upgrade from a weak password to a strong password.
func AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestFrom(ptr unsafe.Pointer) AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest {
	return AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest{
		AccountAuthenticationModificationRequest: AccountAuthenticationModificationRequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest */

// Creates a request to upgrade from using a weak password to using a strong system-generated password.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest/init(user:serviceIdentifier:userInfo:)
func NewAccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestWithUserServiceIdentifierUserInfo(user objc.IObject /* cross-framework: NSString */, serviceIdentifier IASCredentialServiceIdentifier, userInfo objc.IObject /* cross-framework: NSDictionary */) AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest {
	instance := getAccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestClass().Alloc()
	rv := objc.Send[AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest](instance.ID, objc.Sel("initWithUser:serviceIdentifier:userInfo:"), user, serviceIdentifier, userInfo)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAccountAuthenticationModificationUpgradePasswordToStrongPasswordRequestWithUserServiceIdentifierUserInfo */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAccountAuthenticationModificationUpgradePasswordToStrongPasswordRequest */


