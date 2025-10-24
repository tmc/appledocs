// Code generated from Apple documentation for AuthenticationServices. DO NOT EDIT.

package authenticationservices

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class ASAccountAuthenticationModificationController */


/* debug [class_header]: Header for ASAccountAuthenticationModificationController */
// The class instance for the [AccountAuthenticationModificationController] class.
var (
	AccountAuthenticationModificationControllerClass     _AccountAuthenticationModificationControllerClass
	AccountAuthenticationModificationControllerClassOnce sync.Once
)

func getAccountAuthenticationModificationControllerClass() _AccountAuthenticationModificationControllerClass {
	AccountAuthenticationModificationControllerClassOnce.Do(func() {
		AccountAuthenticationModificationControllerClass = _AccountAuthenticationModificationControllerClass{objc.GetClass("ASAccountAuthenticationModificationController")}
	})
	return AccountAuthenticationModificationControllerClass
}

type _AccountAuthenticationModificationControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccountAuthenticationModificationController */
// An interface definition for the [AccountAuthenticationModificationController] class.
type IAccountAuthenticationModificationController interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AccountAuthenticationModificationController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccountAuthenticationModificationController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccountAuthenticationModificationController */
// Alloc allocates a new instance without initialization.
func (ac _AccountAuthenticationModificationControllerClass) Alloc() AccountAuthenticationModificationController {
	rv := objc.Send[AccountAuthenticationModificationController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccountAuthenticationModificationControllerClass) New() AccountAuthenticationModificationController {
	rv := objc.Send[AccountAuthenticationModificationController](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccountAuthenticationModificationController) Init() AccountAuthenticationModificationController {
	rv := objc.Send[AccountAuthenticationModificationController](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccountAuthenticationModificationController) Autorelease() AccountAuthenticationModificationController {
	rv := objc.Send[AccountAuthenticationModificationController](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccountAuthenticationModificationController creates a new AccountAuthenticationModificationController instance.
func NewAccountAuthenticationModificationController() AccountAuthenticationModificationController {
	return getAccountAuthenticationModificationControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccountAuthenticationModificationController */
// An object that performs a request to modify an account’s authentication properties.


// An object that performs a request to modify an account’s authentication properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AuthenticationServices/ASAccountAuthenticationModificationController
type AccountAuthenticationModificationController struct {
	objectivec.Object
}

// AccountAuthenticationModificationControllerFrom constructs a [AccountAuthenticationModificationController] from an unsafe.Pointer.
//
// An object that performs a request to modify an account’s authentication properties.
func AccountAuthenticationModificationControllerFrom(ptr unsafe.Pointer) AccountAuthenticationModificationController {
	return AccountAuthenticationModificationController{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccountAuthenticationModificationController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccountAuthenticationModificationController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccountAuthenticationModificationController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccountAuthenticationModificationController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccountAuthenticationModificationController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class ASAccountAuthenticationModificationController */


