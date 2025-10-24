// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRSwitchClusterShortReleaseEvent */


/* debug [class_header]: Header for MTRSwitchClusterShortReleaseEvent */
// The class instance for the [MTRSwitchClusterShortReleaseEvent] class.
var (
	MTRSwitchClusterShortReleaseEventClass     _MTRSwitchClusterShortReleaseEventClass
	MTRSwitchClusterShortReleaseEventClassOnce sync.Once
)

func getMTRSwitchClusterShortReleaseEventClass() _MTRSwitchClusterShortReleaseEventClass {
	MTRSwitchClusterShortReleaseEventClassOnce.Do(func() {
		MTRSwitchClusterShortReleaseEventClass = _MTRSwitchClusterShortReleaseEventClass{objc.GetClass("MTRSwitchClusterShortReleaseEvent")}
	})
	return MTRSwitchClusterShortReleaseEventClass
}

type _MTRSwitchClusterShortReleaseEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRSwitchClusterShortReleaseEvent */
// An interface definition for the [MTRSwitchClusterShortReleaseEvent] class.
type IMTRSwitchClusterShortReleaseEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRSwitchClusterShortReleaseEvent */
	// properties:
	PreviousPosition() objc.IObject /* cross-framework: NSNumber */
	SetPreviousPosition(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRSwitchClusterShortReleaseEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRSwitchClusterShortReleaseEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRSwitchClusterShortReleaseEventClass) Alloc() MTRSwitchClusterShortReleaseEvent {
	rv := objc.Send[MTRSwitchClusterShortReleaseEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRSwitchClusterShortReleaseEventClass) New() MTRSwitchClusterShortReleaseEvent {
	rv := objc.Send[MTRSwitchClusterShortReleaseEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSwitchClusterShortReleaseEvent) Init() MTRSwitchClusterShortReleaseEvent {
	rv := objc.Send[MTRSwitchClusterShortReleaseEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSwitchClusterShortReleaseEvent) Autorelease() MTRSwitchClusterShortReleaseEvent {
	rv := objc.Send[MTRSwitchClusterShortReleaseEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSwitchClusterShortReleaseEvent creates a new MTRSwitchClusterShortReleaseEvent instance.
func NewMTRSwitchClusterShortReleaseEvent() MTRSwitchClusterShortReleaseEvent {
	return getMTRSwitchClusterShortReleaseEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRSwitchClusterShortReleaseEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterShortReleaseEvent
type MTRSwitchClusterShortReleaseEvent struct {
	objectivec.Object
}

// MTRSwitchClusterShortReleaseEventFrom constructs a [MTRSwitchClusterShortReleaseEvent] from an unsafe.Pointer.
func MTRSwitchClusterShortReleaseEventFrom(ptr unsafe.Pointer) MTRSwitchClusterShortReleaseEvent {
	return MTRSwitchClusterShortReleaseEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRSwitchClusterShortReleaseEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRSwitchClusterShortReleaseEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRSwitchClusterShortReleaseEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRSwitchClusterShortReleaseEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRSwitchClusterShortReleaseEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterShortReleaseEvent/previousPosition
func (m_ MTRSwitchClusterShortReleaseEvent) PreviousPosition() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("previousPosition"))
	return rv
}/* debug [instance_properties/getter]: previousPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterShortReleaseEvent/previousPosition
func (m_ MTRSwitchClusterShortReleaseEvent) SetPreviousPosition(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreviousPosition:"), value)
}/* debug [instance_properties/setter]: previousPosition */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRSwitchClusterShortReleaseEvent */



