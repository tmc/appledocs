// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class ASAccountAuthenticationModificationViewController */


/* debug [class_header]: Header for ASAccountAuthenticationModificationViewController */
// The class instance for the [AccountAuthenticationModificationViewController] class.
var (
	AccountAuthenticationModificationViewControllerClass     _AccountAuthenticationModificationViewControllerClass
	AccountAuthenticationModificationViewControllerClassOnce sync.Once
)

func getAccountAuthenticationModificationViewControllerClass() _AccountAuthenticationModificationViewControllerClass {
	AccountAuthenticationModificationViewControllerClassOnce.Do(func() {
		AccountAuthenticationModificationViewControllerClass = _AccountAuthenticationModificationViewControllerClass{objc.GetClass("ASAccountAuthenticationModificationViewController")}
	})
	return AccountAuthenticationModificationViewControllerClass
}

type _AccountAuthenticationModificationViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccountAuthenticationModificationViewController */
// An interface definition for the [AccountAuthenticationModificationViewController] class.
type IAccountAuthenticationModificationViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for AccountAuthenticationModificationViewController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccountAuthenticationModificationViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccountAuthenticationModificationViewController */
// Alloc allocates a new instance without initialization.
func (ac _AccountAuthenticationModificationViewControllerClass) Alloc() AccountAuthenticationModificationViewController {
	rv := objc.Send[AccountAuthenticationModificationViewController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccountAuthenticationModificationViewControllerClass) New() AccountAuthenticationModificationViewController {
	rv := objc.Send[AccountAuthenticationModificationViewController](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccountAuthenticationModificationViewController) Init() AccountAuthenticationModificationViewController {
	rv := objc.Send[AccountAuthenticationModificationViewController](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccountAuthenticationModificationViewController) Autorelease() AccountAuthenticationModificationViewController {
	rv := objc.Send[AccountAuthenticationModificationViewController](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccountAuthenticationModificationViewController creates a new AccountAuthenticationModificationViewController instance.
func NewAccountAuthenticationModificationViewController() AccountAuthenticationModificationViewController {
	return getAccountAuthenticationModificationViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccountAuthenticationModificationViewController */
// A view controller that can upgrade user passwords to strong passwords, or convert accounts to use Sign in with Apple.
//
// Adding an account modification extension lets your app seamlessly upgrade user passwords to strong passwords, or convert from using passwords to using Sign in with Apple. The entire process can be automatic, requiring no user interaction, or you can include interactions, such as two-factor authentication confirmation.


// A view controller that can upgrade user passwords to strong passwords, or convert accounts to use Sign in with Apple.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationViewController
type AccountAuthenticationModificationViewController struct {
	ViewController
}

// AccountAuthenticationModificationViewControllerFrom constructs a [AccountAuthenticationModificationViewController] from an unsafe.Pointer.
//
// A view controller that can upgrade user passwords to strong passwords, or convert accounts to use Sign in with Apple.
func AccountAuthenticationModificationViewControllerFrom(ptr unsafe.Pointer) AccountAuthenticationModificationViewController {
	return AccountAuthenticationModificationViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccountAuthenticationModificationViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccountAuthenticationModificationViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccountAuthenticationModificationViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccountAuthenticationModificationViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccountAuthenticationModificationViewController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAccountAuthenticationModificationViewController */


