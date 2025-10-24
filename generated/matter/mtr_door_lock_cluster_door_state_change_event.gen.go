// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDoorLockClusterDoorStateChangeEvent */


/* debug [class_header]: Header for MTRDoorLockClusterDoorStateChangeEvent */
// The class instance for the [MTRDoorLockClusterDoorStateChangeEvent] class.
var (
	MTRDoorLockClusterDoorStateChangeEventClass     _MTRDoorLockClusterDoorStateChangeEventClass
	MTRDoorLockClusterDoorStateChangeEventClassOnce sync.Once
)

func getMTRDoorLockClusterDoorStateChangeEventClass() _MTRDoorLockClusterDoorStateChangeEventClass {
	MTRDoorLockClusterDoorStateChangeEventClassOnce.Do(func() {
		MTRDoorLockClusterDoorStateChangeEventClass = _MTRDoorLockClusterDoorStateChangeEventClass{objc.GetClass("MTRDoorLockClusterDoorStateChangeEvent")}
	})
	return MTRDoorLockClusterDoorStateChangeEventClass
}

type _MTRDoorLockClusterDoorStateChangeEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDoorLockClusterDoorStateChangeEvent */
// An interface definition for the [MTRDoorLockClusterDoorStateChangeEvent] class.
type IMTRDoorLockClusterDoorStateChangeEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDoorLockClusterDoorStateChangeEvent */
	// properties:
	DoorState() objc.IObject /* cross-framework: NSNumber */
	SetDoorState(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDoorLockClusterDoorStateChangeEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDoorLockClusterDoorStateChangeEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRDoorLockClusterDoorStateChangeEventClass) Alloc() MTRDoorLockClusterDoorStateChangeEvent {
	rv := objc.Send[MTRDoorLockClusterDoorStateChangeEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDoorLockClusterDoorStateChangeEventClass) New() MTRDoorLockClusterDoorStateChangeEvent {
	rv := objc.Send[MTRDoorLockClusterDoorStateChangeEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDoorLockClusterDoorStateChangeEvent) Init() MTRDoorLockClusterDoorStateChangeEvent {
	rv := objc.Send[MTRDoorLockClusterDoorStateChangeEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDoorLockClusterDoorStateChangeEvent) Autorelease() MTRDoorLockClusterDoorStateChangeEvent {
	rv := objc.Send[MTRDoorLockClusterDoorStateChangeEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDoorLockClusterDoorStateChangeEvent creates a new MTRDoorLockClusterDoorStateChangeEvent instance.
func NewMTRDoorLockClusterDoorStateChangeEvent() MTRDoorLockClusterDoorStateChangeEvent {
	return getMTRDoorLockClusterDoorStateChangeEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDoorLockClusterDoorStateChangeEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterDoorStateChangeEvent
type MTRDoorLockClusterDoorStateChangeEvent struct {
	objectivec.Object
}

// MTRDoorLockClusterDoorStateChangeEventFrom constructs a [MTRDoorLockClusterDoorStateChangeEvent] from an unsafe.Pointer.
func MTRDoorLockClusterDoorStateChangeEventFrom(ptr unsafe.Pointer) MTRDoorLockClusterDoorStateChangeEvent {
	return MTRDoorLockClusterDoorStateChangeEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDoorLockClusterDoorStateChangeEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDoorLockClusterDoorStateChangeEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDoorLockClusterDoorStateChangeEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDoorLockClusterDoorStateChangeEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDoorLockClusterDoorStateChangeEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterDoorStateChangeEvent/doorState
func (m_ MTRDoorLockClusterDoorStateChangeEvent) DoorState() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("doorState"))
	return rv
}/* debug [instance_properties/getter]: doorState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDoorLockClusterDoorStateChangeEvent/doorState
func (m_ MTRDoorLockClusterDoorStateChangeEvent) SetDoorState(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDoorState:"), value)
}/* debug [instance_properties/setter]: doorState */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDoorLockClusterDoorStateChangeEvent */



