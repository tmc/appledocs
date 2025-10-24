// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class ASAccountAuthenticationModificationExtensionContext */


/* debug [class_header]: Header for ASAccountAuthenticationModificationExtensionContext */
// The class instance for the [AccountAuthenticationModificationExtensionContext] class.
var (
	AccountAuthenticationModificationExtensionContextClass     _AccountAuthenticationModificationExtensionContextClass
	AccountAuthenticationModificationExtensionContextClassOnce sync.Once
)

func getAccountAuthenticationModificationExtensionContextClass() _AccountAuthenticationModificationExtensionContextClass {
	AccountAuthenticationModificationExtensionContextClassOnce.Do(func() {
		AccountAuthenticationModificationExtensionContextClass = _AccountAuthenticationModificationExtensionContextClass{objc.GetClass("ASAccountAuthenticationModificationExtensionContext")}
	})
	return AccountAuthenticationModificationExtensionContextClass
}

type _AccountAuthenticationModificationExtensionContextClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccountAuthenticationModificationExtensionContext */
// An interface definition for the [AccountAuthenticationModificationExtensionContext] class.
type IAccountAuthenticationModificationExtensionContext interface {
	foundation.IExtensionContext
	
/* debug [class_interface_properties]: Properties for AccountAuthenticationModificationExtensionContext */
	// properties:
	ASExtensionLocalizedFailureReasonErrorKey() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccountAuthenticationModificationExtensionContext */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccountAuthenticationModificationExtensionContext */
// Alloc allocates a new instance without initialization.
func (ac _AccountAuthenticationModificationExtensionContextClass) Alloc() AccountAuthenticationModificationExtensionContext {
	rv := objc.Send[AccountAuthenticationModificationExtensionContext](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccountAuthenticationModificationExtensionContextClass) New() AccountAuthenticationModificationExtensionContext {
	rv := objc.Send[AccountAuthenticationModificationExtensionContext](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccountAuthenticationModificationExtensionContext) Init() AccountAuthenticationModificationExtensionContext {
	rv := objc.Send[AccountAuthenticationModificationExtensionContext](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccountAuthenticationModificationExtensionContext) Autorelease() AccountAuthenticationModificationExtensionContext {
	rv := objc.Send[AccountAuthenticationModificationExtensionContext](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccountAuthenticationModificationExtensionContext creates a new AccountAuthenticationModificationExtensionContext instance.
func NewAccountAuthenticationModificationExtensionContext() AccountAuthenticationModificationExtensionContext {
	return getAccountAuthenticationModificationExtensionContextClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccountAuthenticationModificationExtensionContext */
// An object that you interact with to change an account’s password or to upgrade to Sign in with Apple.


// An object that you interact with to change an account’s password or to upgrade to Sign in with Apple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationExtensionContext
type AccountAuthenticationModificationExtensionContext struct {
	foundation.ExtensionContext
}

// AccountAuthenticationModificationExtensionContextFrom constructs a [AccountAuthenticationModificationExtensionContext] from an unsafe.Pointer.
//
// An object that you interact with to change an account’s password or to upgrade to Sign in with Apple.
func AccountAuthenticationModificationExtensionContextFrom(ptr unsafe.Pointer) AccountAuthenticationModificationExtensionContext {
	return AccountAuthenticationModificationExtensionContext{
		ExtensionContext: foundation.ExtensionContextFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccountAuthenticationModificationExtensionContext *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccountAuthenticationModificationExtensionContext */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccountAuthenticationModificationExtensionContext */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccountAuthenticationModificationExtensionContext */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccountAuthenticationModificationExtensionContext */

// A key that specifies a string value to show to the user when a request fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/authenticationservices/asextensionlocalizedfailurereasonerrorkey
func (a_ AccountAuthenticationModificationExtensionContext) ASExtensionLocalizedFailureReasonErrorKey() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("ASExtensionLocalizedFailureReasonErrorKey"))
	return rv
}/* debug [instance_properties/getter]: ASExtensionLocalizedFailureReasonErrorKey */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAccountAuthenticationModificationExtensionContext */


