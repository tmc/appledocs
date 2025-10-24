// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRSwitchClusterLongPressEvent */


/* debug [class_header]: Header for MTRSwitchClusterLongPressEvent */
// The class instance for the [MTRSwitchClusterLongPressEvent] class.
var (
	MTRSwitchClusterLongPressEventClass     _MTRSwitchClusterLongPressEventClass
	MTRSwitchClusterLongPressEventClassOnce sync.Once
)

func getMTRSwitchClusterLongPressEventClass() _MTRSwitchClusterLongPressEventClass {
	MTRSwitchClusterLongPressEventClassOnce.Do(func() {
		MTRSwitchClusterLongPressEventClass = _MTRSwitchClusterLongPressEventClass{objc.GetClass("MTRSwitchClusterLongPressEvent")}
	})
	return MTRSwitchClusterLongPressEventClass
}

type _MTRSwitchClusterLongPressEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRSwitchClusterLongPressEvent */
// An interface definition for the [MTRSwitchClusterLongPressEvent] class.
type IMTRSwitchClusterLongPressEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRSwitchClusterLongPressEvent */
	// properties:
	NewPosition() objc.IObject /* cross-framework: NSNumber */
	SetNewPosition(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRSwitchClusterLongPressEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRSwitchClusterLongPressEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRSwitchClusterLongPressEventClass) Alloc() MTRSwitchClusterLongPressEvent {
	rv := objc.Send[MTRSwitchClusterLongPressEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRSwitchClusterLongPressEventClass) New() MTRSwitchClusterLongPressEvent {
	rv := objc.Send[MTRSwitchClusterLongPressEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSwitchClusterLongPressEvent) Init() MTRSwitchClusterLongPressEvent {
	rv := objc.Send[MTRSwitchClusterLongPressEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSwitchClusterLongPressEvent) Autorelease() MTRSwitchClusterLongPressEvent {
	rv := objc.Send[MTRSwitchClusterLongPressEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSwitchClusterLongPressEvent creates a new MTRSwitchClusterLongPressEvent instance.
func NewMTRSwitchClusterLongPressEvent() MTRSwitchClusterLongPressEvent {
	return getMTRSwitchClusterLongPressEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRSwitchClusterLongPressEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterLongPressEvent
type MTRSwitchClusterLongPressEvent struct {
	objectivec.Object
}

// MTRSwitchClusterLongPressEventFrom constructs a [MTRSwitchClusterLongPressEvent] from an unsafe.Pointer.
func MTRSwitchClusterLongPressEventFrom(ptr unsafe.Pointer) MTRSwitchClusterLongPressEvent {
	return MTRSwitchClusterLongPressEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRSwitchClusterLongPressEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRSwitchClusterLongPressEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRSwitchClusterLongPressEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRSwitchClusterLongPressEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRSwitchClusterLongPressEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterLongPressEvent/newPosition
func (m_ MTRSwitchClusterLongPressEvent) NewPosition() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newPosition"))
	return rv
}/* debug [instance_properties/getter]: newPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterLongPressEvent/newPosition
func (m_ MTRSwitchClusterLongPressEvent) SetNewPosition(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewPosition:"), value)
}/* debug [instance_properties/setter]: newPosition */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRSwitchClusterLongPressEvent */



