// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRPowerSourceClusterWiredFaultChangeEvent */


/* debug [class_header]: Header for MTRPowerSourceClusterWiredFaultChangeEvent */
// The class instance for the [MTRPowerSourceClusterWiredFaultChangeEvent] class.
var (
	MTRPowerSourceClusterWiredFaultChangeEventClass     _MTRPowerSourceClusterWiredFaultChangeEventClass
	MTRPowerSourceClusterWiredFaultChangeEventClassOnce sync.Once
)

func getMTRPowerSourceClusterWiredFaultChangeEventClass() _MTRPowerSourceClusterWiredFaultChangeEventClass {
	MTRPowerSourceClusterWiredFaultChangeEventClassOnce.Do(func() {
		MTRPowerSourceClusterWiredFaultChangeEventClass = _MTRPowerSourceClusterWiredFaultChangeEventClass{objc.GetClass("MTRPowerSourceClusterWiredFaultChangeEvent")}
	})
	return MTRPowerSourceClusterWiredFaultChangeEventClass
}

type _MTRPowerSourceClusterWiredFaultChangeEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRPowerSourceClusterWiredFaultChangeEvent */
// An interface definition for the [MTRPowerSourceClusterWiredFaultChangeEvent] class.
type IMTRPowerSourceClusterWiredFaultChangeEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRPowerSourceClusterWiredFaultChangeEvent */
	// properties:
	Current() objc.IObject /* cross-framework: NSArray */
	SetCurrent(value objc.IObject /* cross-framework: NSArray */)
	Previous() objc.IObject /* cross-framework: NSArray */
	SetPrevious(value objc.IObject /* cross-framework: NSArray */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRPowerSourceClusterWiredFaultChangeEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRPowerSourceClusterWiredFaultChangeEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRPowerSourceClusterWiredFaultChangeEventClass) Alloc() MTRPowerSourceClusterWiredFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterWiredFaultChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRPowerSourceClusterWiredFaultChangeEventClass) New() MTRPowerSourceClusterWiredFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterWiredFaultChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRPowerSourceClusterWiredFaultChangeEvent) Init() MTRPowerSourceClusterWiredFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterWiredFaultChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRPowerSourceClusterWiredFaultChangeEvent) Autorelease() MTRPowerSourceClusterWiredFaultChangeEvent {
	rv := objc.Send[MTRPowerSourceClusterWiredFaultChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRPowerSourceClusterWiredFaultChangeEvent creates a new MTRPowerSourceClusterWiredFaultChangeEvent instance.
func NewMTRPowerSourceClusterWiredFaultChangeEvent() MTRPowerSourceClusterWiredFaultChangeEvent {
	return getMTRPowerSourceClusterWiredFaultChangeEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRPowerSourceClusterWiredFaultChangeEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterWiredFaultChangeEvent
type MTRPowerSourceClusterWiredFaultChangeEvent struct {
	objectivec.Object
}

// MTRPowerSourceClusterWiredFaultChangeEventFrom constructs a [MTRPowerSourceClusterWiredFaultChangeEvent] from an unsafe.Pointer.
func MTRPowerSourceClusterWiredFaultChangeEventFrom(ptr unsafe.Pointer) MTRPowerSourceClusterWiredFaultChangeEvent {
	return MTRPowerSourceClusterWiredFaultChangeEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRPowerSourceClusterWiredFaultChangeEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRPowerSourceClusterWiredFaultChangeEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRPowerSourceClusterWiredFaultChangeEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRPowerSourceClusterWiredFaultChangeEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRPowerSourceClusterWiredFaultChangeEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterWiredFaultChangeEvent/current
func (m_ MTRPowerSourceClusterWiredFaultChangeEvent) Current() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("current"))
	return rv
}/* debug [instance_properties/getter]: current */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterWiredFaultChangeEvent/current
func (m_ MTRPowerSourceClusterWiredFaultChangeEvent) SetCurrent(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrent:"), value)
}/* debug [instance_properties/setter]: current */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterWiredFaultChangeEvent/previous
func (m_ MTRPowerSourceClusterWiredFaultChangeEvent) Previous() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](m_.ID, objc.Sel("previous"))
	return rv
}/* debug [instance_properties/getter]: previous */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRPowerSourceClusterWiredFaultChangeEvent/previous
func (m_ MTRPowerSourceClusterWiredFaultChangeEvent) SetPrevious(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPrevious:"), value)
}/* debug [instance_properties/setter]: previous */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRPowerSourceClusterWiredFaultChangeEvent */



