// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPowerSourceClusterBatFaultChangeEvent */


/* debug [class_header]: Header for MTRPowerSourceClusterBatFaultChangeEvent */
// The class instance for the [MTRPowerSourceClusterBatFaultChangeEvent] class.
var (
	MTRPowerSourceClusterBatFaultChangeEventClass     _MTRPowerSourceClusterBatFaultChangeEventClass
	MTRPowerSourceClusterBatFaultChangeEventClassOnce sync.Once
)

func getMTRPowerSourceClusterBatFaultChangeEventClass() _MTRPowerSourceClusterBatFaultChangeEventClass {
	MTRPowerSourceClusterBatFaultChangeEventClassOnce.Do(func() {
		MTRPowerSourceClusterBatFaultChangeEventClass = _MTRPowerSourceClusterBatFaultChangeEventClass{objc.GetClass("MTRPowerSourceClusterBatFaultChangeEvent")}
	})
	return MTRPowerSourceClusterBatFaultChangeEventClass
}

type _MTRPowerSourceClusterBatFaultChangeEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPowerSourceClusterBatFaultChangeEvent */
// An interface definition for the [MTRPowerSourceClusterBatFaultChangeEvent] class.
type IMTRPowerSourceClusterBatFaultChangeEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPowerSourceClusterBatFaultChangeEvent */
	// properties:
	Current() objc.IObject /* cross-framework: NSArray */
	SetCurrent(value objc.IObject /* cross-framework: NSArray */)
	Previous() objc.IObject /* cross-framework: NSArray */
	SetPrevious(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPowerSourceClusterBatFaultChangeEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPowerSourceClusterBatFaultChangeEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPowerSourceClusterBatFaultChangeEventClass) Alloc() MTRPowerSourceClusterBatFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterBatFaultChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPowerSourceClusterBatFaultChangeEventClass) New() MTRPowerSourceClusterBatFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterBatFaultChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPowerSourceClusterBatFaultChangeEvent) Init() MTRPowerSourceClusterBatFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterBatFaultChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPowerSourceClusterBatFaultChangeEvent) Autorelease() MTRPowerSourceClusterBatFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterBatFaultChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPowerSourceClusterBatFaultChangeEvent creates a new MTRPowerSourceClusterBatFaultChangeEvent instance.
func NewMTRPowerSourceClusterBatFaultChangeEvent() MTRPowerSourceClusterBatFaultChangeEvent {
	return getMTRPowerSourceClusterBatFaultChangeEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPowerSourceClusterBatFaultChangeEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatFaultChangeEvent
type MTRPowerSourceClusterBatFaultChangeEvent struct {
	objectivec.Object
}

// MTRPowerSourceClusterBatFaultChangeEventFrom constructs a [MTRPowerSourceClusterBatFaultChangeEvent] from an unsafe.Pointer.
func MTRPowerSourceClusterBatFaultChangeEventFrom(ptr unsafe.Pointer) MTRPowerSourceClusterBatFaultChangeEvent {
	return MTRPowerSourceClusterBatFaultChangeEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPowerSourceClusterBatFaultChangeEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPowerSourceClusterBatFaultChangeEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPowerSourceClusterBatFaultChangeEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPowerSourceClusterBatFaultChangeEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPowerSourceClusterBatFaultChangeEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatFaultChangeEvent/current
func (m_ MTRPowerSourceClusterBatFaultChangeEvent) Current() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("current"))
	return rv
}/* debug [instance_properties/getter]: current */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatFaultChangeEvent/current
func (m_ MTRPowerSourceClusterBatFaultChangeEvent) SetCurrent(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrent:"), value)
}/* debug [instance_properties/setter]: current */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatFaultChangeEvent/previous
func (m_ MTRPowerSourceClusterBatFaultChangeEvent) Previous() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("previous"))
	return rv
}/* debug [instance_properties/getter]: previous */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterBatFaultChangeEvent/previous
func (m_ MTRPowerSourceClusterBatFaultChangeEvent) SetPrevious(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrevious:"), value)
}/* debug [instance_properties/setter]: previous */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPowerSourceClusterBatFaultChangeEvent */



