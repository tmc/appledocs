// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CTCarrier */


/* debug [class_header]: Header for CTCarrier */
// The class instance for the [Carrier] class.
var (
	CarrierClass     _CarrierClass
	CarrierClassOnce sync.Once
)

func getCarrierClass() _CarrierClass {
	CarrierClassOnce.Do(func() {
		CarrierClass = _CarrierClass{objc.GetClass("CTCarrier")}
	})
	return CarrierClass
}

type _CarrierClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Carrier */
// An interface definition for the [Carrier] class.
type ICarrier interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Carrier */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Carrier */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Carrier */
// Alloc allocates a new instance without initialization.
func (cc _CarrierClass) Alloc() Carrier {
	rv := objc.Send[Carrier](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CarrierClass) New() Carrier {
	rv := objc.Send[Carrier](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ Carrier) Init() Carrier {
	rv := objc.Send[Carrier](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ Carrier) Autorelease() Carrier {
	rv := objc.Send[Carrier](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCarrier creates a new Carrier instance.
func NewCarrier() Carrier {
	return getCarrierClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Carrier */
// Information about the user’s cellular service provider, such as its unique identifier and whether it allows VoIP calls on its network.


// Information about the user’s cellular service provider, such as its unique identifier and whether it allows VoIP calls on its network.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTCarrier
type Carrier struct {
	objectivec.Object
}

// CarrierFrom constructs a [Carrier] from an unsafe.Pointer.
//
// Information about the user’s cellular service provider, such as its unique identifier and whether it allows VoIP calls on its network.
func CarrierFrom(ptr unsafe.Pointer) Carrier {
	return Carrier{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Carrier *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Carrier */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Carrier */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Carrier */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Carrier */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CTCarrier */


