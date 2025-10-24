// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRSwitchClusterSwitchLatchedEvent */


/* debug [class_header]: Header for MTRSwitchClusterSwitchLatchedEvent */
// The class instance for the [MTRSwitchClusterSwitchLatchedEvent] class.
var (
	MTRSwitchClusterSwitchLatchedEventClass     _MTRSwitchClusterSwitchLatchedEventClass
	MTRSwitchClusterSwitchLatchedEventClassOnce sync.Once
)

func getMTRSwitchClusterSwitchLatchedEventClass() _MTRSwitchClusterSwitchLatchedEventClass {
	MTRSwitchClusterSwitchLatchedEventClassOnce.Do(func() {
		MTRSwitchClusterSwitchLatchedEventClass = _MTRSwitchClusterSwitchLatchedEventClass{objc.GetClass("MTRSwitchClusterSwitchLatchedEvent")}
	})
	return MTRSwitchClusterSwitchLatchedEventClass
}

type _MTRSwitchClusterSwitchLatchedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRSwitchClusterSwitchLatchedEvent */
// An interface definition for the [MTRSwitchClusterSwitchLatchedEvent] class.
type IMTRSwitchClusterSwitchLatchedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRSwitchClusterSwitchLatchedEvent */
	// properties:
	NewPosition() objc.IObject /* cross-framework: NSNumber */
	SetNewPosition(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRSwitchClusterSwitchLatchedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRSwitchClusterSwitchLatchedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRSwitchClusterSwitchLatchedEventClass) Alloc() MTRSwitchClusterSwitchLatchedEvent {
	rv := objc.Send[MTRSwitchClusterSwitchLatchedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRSwitchClusterSwitchLatchedEventClass) New() MTRSwitchClusterSwitchLatchedEvent {
	rv := objc.Send[MTRSwitchClusterSwitchLatchedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSwitchClusterSwitchLatchedEvent) Init() MTRSwitchClusterSwitchLatchedEvent {
	rv := objc.Send[MTRSwitchClusterSwitchLatchedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSwitchClusterSwitchLatchedEvent) Autorelease() MTRSwitchClusterSwitchLatchedEvent {
	rv := objc.Send[MTRSwitchClusterSwitchLatchedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSwitchClusterSwitchLatchedEvent creates a new MTRSwitchClusterSwitchLatchedEvent instance.
func NewMTRSwitchClusterSwitchLatchedEvent() MTRSwitchClusterSwitchLatchedEvent {
	return getMTRSwitchClusterSwitchLatchedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRSwitchClusterSwitchLatchedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterSwitchLatchedEvent
type MTRSwitchClusterSwitchLatchedEvent struct {
	objectivec.Object
}

// MTRSwitchClusterSwitchLatchedEventFrom constructs a [MTRSwitchClusterSwitchLatchedEvent] from an unsafe.Pointer.
func MTRSwitchClusterSwitchLatchedEventFrom(ptr unsafe.Pointer) MTRSwitchClusterSwitchLatchedEvent {
	return MTRSwitchClusterSwitchLatchedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRSwitchClusterSwitchLatchedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRSwitchClusterSwitchLatchedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRSwitchClusterSwitchLatchedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRSwitchClusterSwitchLatchedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRSwitchClusterSwitchLatchedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterSwitchLatchedEvent/newPosition
func (m_ MTRSwitchClusterSwitchLatchedEvent) NewPosition() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newPosition"))
	return rv
}/* debug [instance_properties/getter]: newPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterSwitchLatchedEvent/newPosition
func (m_ MTRSwitchClusterSwitchLatchedEvent) SetNewPosition(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewPosition:"), value)
}/* debug [instance_properties/setter]: newPosition */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRSwitchClusterSwitchLatchedEvent */



