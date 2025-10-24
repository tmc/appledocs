// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPowerSourceClusterBatChargeFaultChangeEvent */


/* debug [class_header]: Header for MTRPowerSourceClusterBatChargeFaultChangeEvent */
// The class instance for the [MTRPowerSourceClusterBatChargeFaultChangeEvent] class.
var (
	MTRPowerSourceClusterBatChargeFaultChangeEventClass     _MTRPowerSourceClusterBatChargeFaultChangeEventClass
	MTRPowerSourceClusterBatChargeFaultChangeEventClassOnce sync.Once
)

func getMTRPowerSourceClusterBatChargeFaultChangeEventClass() _MTRPowerSourceClusterBatChargeFaultChangeEventClass {
	MTRPowerSourceClusterBatChargeFaultChangeEventClassOnce.Do(func() {
		MTRPowerSourceClusterBatChargeFaultChangeEventClass = _MTRPowerSourceClusterBatChargeFaultChangeEventClass{objc.GetClass("MTRPowerSourceClusterBatChargeFaultChangeEvent")}
	})
	return MTRPowerSourceClusterBatChargeFaultChangeEventClass
}

type _MTRPowerSourceClusterBatChargeFaultChangeEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPowerSourceClusterBatChargeFaultChangeEvent */
// An interface definition for the [MTRPowerSourceClusterBatChargeFaultChangeEvent] class.
type IMTRPowerSourceClusterBatChargeFaultChangeEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPowerSourceClusterBatChargeFaultChangeEvent */
	// properties:
	Current() objc.IObject /* cross-framework: NSArray */
	SetCurrent(value objc.IObject /* cross-framework: NSArray */)
	Previous() objc.IObject /* cross-framework: NSArray */
	SetPrevious(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPowerSourceClusterBatChargeFaultChangeEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPowerSourceClusterBatChargeFaultChangeEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPowerSourceClusterBatChargeFaultChangeEventClass) Alloc() MTRPowerSourceClusterBatChargeFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterBatChargeFaultChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPowerSourceClusterBatChargeFaultChangeEventClass) New() MTRPowerSourceClusterBatChargeFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterBatChargeFaultChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPowerSourceClusterBatChargeFaultChangeEvent) Init() MTRPowerSourceClusterBatChargeFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterBatChargeFaultChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPowerSourceClusterBatChargeFaultChangeEvent) Autorelease() MTRPowerSourceClusterBatChargeFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterBatChargeFaultChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPowerSourceClusterBatChargeFaultChangeEvent creates a new MTRPowerSourceClusterBatChargeFaultChangeEvent instance.
func NewMTRPowerSourceClusterBatChargeFaultChangeEvent() MTRPowerSourceClusterBatChargeFaultChangeEvent {
	return getMTRPowerSourceClusterBatChargeFaultChangeEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPowerSourceClusterBatChargeFaultChangeEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatChargeFaultChangeEvent
type MTRPowerSourceClusterBatChargeFaultChangeEvent struct {
	objectivec.Object
}

// MTRPowerSourceClusterBatChargeFaultChangeEventFrom constructs a [MTRPowerSourceClusterBatChargeFaultChangeEvent] from an unsafe.Pointer.
func MTRPowerSourceClusterBatChargeFaultChangeEventFrom(ptr unsafe.Pointer) MTRPowerSourceClusterBatChargeFaultChangeEvent {
	return MTRPowerSourceClusterBatChargeFaultChangeEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPowerSourceClusterBatChargeFaultChangeEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPowerSourceClusterBatChargeFaultChangeEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPowerSourceClusterBatChargeFaultChangeEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPowerSourceClusterBatChargeFaultChangeEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPowerSourceClusterBatChargeFaultChangeEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatChargeFaultChangeEvent/current
func (m_ MTRPowerSourceClusterBatChargeFaultChangeEvent) Current() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("current"))
	return rv
}/* debug [instance_properties/getter]: current */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatChargeFaultChangeEvent/current
func (m_ MTRPowerSourceClusterBatChargeFaultChangeEvent) SetCurrent(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrent:"), value)
}/* debug [instance_properties/setter]: current */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatChargeFaultChangeEvent/previous
func (m_ MTRPowerSourceClusterBatChargeFaultChangeEvent) Previous() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("previous"))
	return rv
}/* debug [instance_properties/getter]: previous */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatChargeFaultChangeEvent/previous
func (m_ MTRPowerSourceClusterBatChargeFaultChangeEvent) SetPrevious(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrevious:"), value)
}/* debug [instance_properties/setter]: previous */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPowerSourceClusterBatChargeFaultChangeEvent */



