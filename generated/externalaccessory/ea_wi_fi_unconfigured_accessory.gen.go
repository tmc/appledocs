// Code generated from Apple documentation for ExternalAccessory. DO NOT EDIT.

package externalaccessory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class EAWiFiUnconfiguredAccessory */


/* debug [class_header]: Header for EAWiFiUnconfiguredAccessory */
// The class instance for the [EAWiFiUnconfiguredAccessory] class.
var (
	EAWiFiUnconfiguredAccessoryClass     _EAWiFiUnconfiguredAccessoryClass
	EAWiFiUnconfiguredAccessoryClassOnce sync.Once
)

func getEAWiFiUnconfiguredAccessoryClass() _EAWiFiUnconfiguredAccessoryClass {
	EAWiFiUnconfiguredAccessoryClassOnce.Do(func() {
		EAWiFiUnconfiguredAccessoryClass = _EAWiFiUnconfiguredAccessoryClass{objc.GetClass("EAWiFiUnconfiguredAccessory")}
	})
	return EAWiFiUnconfiguredAccessoryClass
}

type _EAWiFiUnconfiguredAccessoryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for EAWiFiUnconfiguredAccessory */
// An interface definition for the [EAWiFiUnconfiguredAccessory] class.
type IEAWiFiUnconfiguredAccessory interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for EAWiFiUnconfiguredAccessory */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for EAWiFiUnconfiguredAccessory */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for EAWiFiUnconfiguredAccessory */
// Alloc allocates a new instance without initialization.
func (ec _EAWiFiUnconfiguredAccessoryClass) Alloc() EAWiFiUnconfiguredAccessory {
	rv := objc.Send[EAWiFiUnconfiguredAccessory](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ec _EAWiFiUnconfiguredAccessoryClass) New() EAWiFiUnconfiguredAccessory {
	rv := objc.Send[EAWiFiUnconfiguredAccessory](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ EAWiFiUnconfiguredAccessory) Init() EAWiFiUnconfiguredAccessory {
	rv := objc.Send[EAWiFiUnconfiguredAccessory](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ EAWiFiUnconfiguredAccessory) Autorelease() EAWiFiUnconfiguredAccessory {
	rv := objc.Send[EAWiFiUnconfiguredAccessory](e_.ID, objc.Sel("autorelease"))
	return rv
}

// NewEAWiFiUnconfiguredAccessory creates a new EAWiFiUnconfiguredAccessory instance.
func NewEAWiFiUnconfiguredAccessory() EAWiFiUnconfiguredAccessory {
	return getEAWiFiUnconfiguredAccessoryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for EAWiFiUnconfiguredAccessory */
// An object that provides information about an unconfigured MFi Wireless Accessory Configuration accessory.


// An object that provides information about an unconfigured MFi Wireless Accessory Configuration accessory.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ExternalAccessory/EAWiFiUnconfiguredAccessory
type EAWiFiUnconfiguredAccessory struct {
	objectivec.Object
}

// EAWiFiUnconfiguredAccessoryFrom constructs a [EAWiFiUnconfiguredAccessory] from an unsafe.Pointer.
//
// An object that provides information about an unconfigured MFi Wireless Accessory Configuration accessory.
func EAWiFiUnconfiguredAccessoryFrom(ptr unsafe.Pointer) EAWiFiUnconfiguredAccessory {
	return EAWiFiUnconfiguredAccessory{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for EAWiFiUnconfiguredAccessory *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for EAWiFiUnconfiguredAccessory */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for EAWiFiUnconfiguredAccessory */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for EAWiFiUnconfiguredAccessory */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for EAWiFiUnconfiguredAccessory */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class EAWiFiUnconfiguredAccessory */


