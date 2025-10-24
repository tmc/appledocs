// Code generated from Apple documentation for CoreBluetooth. DO NOT EDIT.

package corebluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CBManager */


/* debug [class_header]: Header for CBManager */
// The class instance for the [CBManager] class.
var (
	CBManagerClass     _CBManagerClass
	CBManagerClassOnce sync.Once
)

func getCBManagerClass() _CBManagerClass {
	CBManagerClassOnce.Do(func() {
		CBManagerClass = _CBManagerClass{objc.GetClass("CBManager")}
	})
	return CBManagerClass
}

type _CBManagerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CBManager */
// An interface definition for the [CBManager] class.
type ICBManager interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CBManager */
	// properties:
	Authorization() CBManagerAuthorization
	State() CBManagerState
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CBManager */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CBManager */
// Alloc allocates a new instance without initialization.
func (cc _CBManagerClass) Alloc() CBManager {
	rv := objc.Send[CBManager](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CBManagerClass) New() CBManager {
	rv := objc.Send[CBManager](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CBManager) Init() CBManager {
	rv := objc.Send[CBManager](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CBManager) Autorelease() CBManager {
	rv := objc.Send[CBManager](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCBManager creates a new CBManager instance.
func NewCBManager() CBManager {
	return getCBManagerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CBManager */
// The abstract base class that manages central and peripheral objects.


// The abstract base class that manages central and peripheral objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManager
type CBManager struct {
	objectivec.Object
}

// CBManagerFrom constructs a [CBManager] from an unsafe.Pointer.
//
// The abstract base class that manages central and peripheral objects.
func CBManagerFrom(ptr unsafe.Pointer) CBManager {
	return CBManager{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CBManager *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CBManager */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CBManager */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CBManager */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CBManager */

// The current authorization status for using Bluetooth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManager/authorization-swift.property
func (c_ CBManager) Authorization() CBManagerAuthorization {
	rv := objc.Send[CBManagerAuthorization](c_.ID, objc.Sel("authorization"))
	return rv
}/* debug [instance_properties/getter]: authorization */


// The current state of the manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreBluetooth/CBManager/state
func (c_ CBManager) State() CBManagerState {
	rv := objc.Send[CBManagerState](c_.ID, objc.Sel("state"))
	return rv
}/* debug [instance_properties/getter]: state */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CBManager */



