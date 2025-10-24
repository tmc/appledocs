// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRSwitchClusterLongReleaseEvent */


/* debug [class_header]: Header for MTRSwitchClusterLongReleaseEvent */
// The class instance for the [MTRSwitchClusterLongReleaseEvent] class.
var (
	MTRSwitchClusterLongReleaseEventClass     _MTRSwitchClusterLongReleaseEventClass
	MTRSwitchClusterLongReleaseEventClassOnce sync.Once
)

func getMTRSwitchClusterLongReleaseEventClass() _MTRSwitchClusterLongReleaseEventClass {
	MTRSwitchClusterLongReleaseEventClassOnce.Do(func() {
		MTRSwitchClusterLongReleaseEventClass = _MTRSwitchClusterLongReleaseEventClass{objc.GetClass("MTRSwitchClusterLongReleaseEvent")}
	})
	return MTRSwitchClusterLongReleaseEventClass
}

type _MTRSwitchClusterLongReleaseEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRSwitchClusterLongReleaseEvent */
// An interface definition for the [MTRSwitchClusterLongReleaseEvent] class.
type IMTRSwitchClusterLongReleaseEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRSwitchClusterLongReleaseEvent */
	// properties:
	PreviousPosition() objc.IObject /* cross-framework: NSNumber */
	SetPreviousPosition(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRSwitchClusterLongReleaseEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRSwitchClusterLongReleaseEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRSwitchClusterLongReleaseEventClass) Alloc() MTRSwitchClusterLongReleaseEvent {
	rv := objc.Send[MTRSwitchClusterLongReleaseEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRSwitchClusterLongReleaseEventClass) New() MTRSwitchClusterLongReleaseEvent {
	rv := objc.Send[MTRSwitchClusterLongReleaseEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSwitchClusterLongReleaseEvent) Init() MTRSwitchClusterLongReleaseEvent {
	rv := objc.Send[MTRSwitchClusterLongReleaseEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSwitchClusterLongReleaseEvent) Autorelease() MTRSwitchClusterLongReleaseEvent {
	rv := objc.Send[MTRSwitchClusterLongReleaseEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSwitchClusterLongReleaseEvent creates a new MTRSwitchClusterLongReleaseEvent instance.
func NewMTRSwitchClusterLongReleaseEvent() MTRSwitchClusterLongReleaseEvent {
	return getMTRSwitchClusterLongReleaseEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRSwitchClusterLongReleaseEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterLongReleaseEvent
type MTRSwitchClusterLongReleaseEvent struct {
	objectivec.Object
}

// MTRSwitchClusterLongReleaseEventFrom constructs a [MTRSwitchClusterLongReleaseEvent] from an unsafe.Pointer.
func MTRSwitchClusterLongReleaseEventFrom(ptr unsafe.Pointer) MTRSwitchClusterLongReleaseEvent {
	return MTRSwitchClusterLongReleaseEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRSwitchClusterLongReleaseEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRSwitchClusterLongReleaseEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRSwitchClusterLongReleaseEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRSwitchClusterLongReleaseEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRSwitchClusterLongReleaseEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterLongReleaseEvent/previousPosition
func (m_ MTRSwitchClusterLongReleaseEvent) PreviousPosition() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("previousPosition"))
	return rv
}/* debug [instance_properties/getter]: previousPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterLongReleaseEvent/previousPosition
func (m_ MTRSwitchClusterLongReleaseEvent) SetPreviousPosition(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreviousPosition:"), value)
}/* debug [instance_properties/setter]: previousPosition */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRSwitchClusterLongReleaseEvent */



