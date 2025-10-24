// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAccountAuthenticationModificationRequest */


/* debug [class_header]: Header for ASAccountAuthenticationModificationRequest */
// The class instance for the [AccountAuthenticationModificationRequest] class.
var (
	AccountAuthenticationModificationRequestClass     _AccountAuthenticationModificationRequestClass
	AccountAuthenticationModificationRequestClassOnce sync.Once
)

func getAccountAuthenticationModificationRequestClass() _AccountAuthenticationModificationRequestClass {
	AccountAuthenticationModificationRequestClassOnce.Do(func() {
		AccountAuthenticationModificationRequestClass = _AccountAuthenticationModificationRequestClass{objc.GetClass("ASAccountAuthenticationModificationRequest")}
	})
	return AccountAuthenticationModificationRequestClass
}

type _AccountAuthenticationModificationRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccountAuthenticationModificationRequest */
// An interface definition for the [AccountAuthenticationModificationRequest] class.
type IAccountAuthenticationModificationRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AccountAuthenticationModificationRequest */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccountAuthenticationModificationRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccountAuthenticationModificationRequest */
// Alloc allocates a new instance without initialization.
func (ac _AccountAuthenticationModificationRequestClass) Alloc() AccountAuthenticationModificationRequest {
	rv := objc.Send[AccountAuthenticationModificationRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccountAuthenticationModificationRequestClass) New() AccountAuthenticationModificationRequest {
	rv := objc.Send[AccountAuthenticationModificationRequest](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccountAuthenticationModificationRequest) Init() AccountAuthenticationModificationRequest {
	rv := objc.Send[AccountAuthenticationModificationRequest](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccountAuthenticationModificationRequest) Autorelease() AccountAuthenticationModificationRequest {
	rv := objc.Send[AccountAuthenticationModificationRequest](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccountAuthenticationModificationRequest creates a new AccountAuthenticationModificationRequest instance.
func NewAccountAuthenticationModificationRequest() AccountAuthenticationModificationRequest {
	return getAccountAuthenticationModificationRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccountAuthenticationModificationRequest */
// A request to modify an account’s authentication properties.
//
// To initiate an account authentication modification request from your app, use either or .


// A request to modify an account’s authentication properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationRequest
type AccountAuthenticationModificationRequest struct {
	objectivec.Object
}

// AccountAuthenticationModificationRequestFrom constructs a [AccountAuthenticationModificationRequest] from an unsafe.Pointer.
//
// A request to modify an account’s authentication properties.
func AccountAuthenticationModificationRequestFrom(ptr unsafe.Pointer) AccountAuthenticationModificationRequest {
	return AccountAuthenticationModificationRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccountAuthenticationModificationRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccountAuthenticationModificationRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccountAuthenticationModificationRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccountAuthenticationModificationRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccountAuthenticationModificationRequest */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAccountAuthenticationModificationRequest */



