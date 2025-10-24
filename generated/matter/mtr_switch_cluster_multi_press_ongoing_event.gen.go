// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRSwitchClusterMultiPressOngoingEvent */


/* debug [class_header]: Header for MTRSwitchClusterMultiPressOngoingEvent */
// The class instance for the [MTRSwitchClusterMultiPressOngoingEvent] class.
var (
	MTRSwitchClusterMultiPressOngoingEventClass     _MTRSwitchClusterMultiPressOngoingEventClass
	MTRSwitchClusterMultiPressOngoingEventClassOnce sync.Once
)

func getMTRSwitchClusterMultiPressOngoingEventClass() _MTRSwitchClusterMultiPressOngoingEventClass {
	MTRSwitchClusterMultiPressOngoingEventClassOnce.Do(func() {
		MTRSwitchClusterMultiPressOngoingEventClass = _MTRSwitchClusterMultiPressOngoingEventClass{objc.GetClass("MTRSwitchClusterMultiPressOngoingEvent")}
	})
	return MTRSwitchClusterMultiPressOngoingEventClass
}

type _MTRSwitchClusterMultiPressOngoingEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRSwitchClusterMultiPressOngoingEvent */
// An interface definition for the [MTRSwitchClusterMultiPressOngoingEvent] class.
type IMTRSwitchClusterMultiPressOngoingEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRSwitchClusterMultiPressOngoingEvent */
	// properties:
	CurrentNumberOfPressesCounted() objc.IObject /* cross-framework: NSNumber */
	SetCurrentNumberOfPressesCounted(value objc.IObject /* cross-framework: NSNumber */)
	NewPosition() objc.IObject /* cross-framework: NSNumber */
	SetNewPosition(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRSwitchClusterMultiPressOngoingEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRSwitchClusterMultiPressOngoingEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRSwitchClusterMultiPressOngoingEventClass) Alloc() MTRSwitchClusterMultiPressOngoingEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressOngoingEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRSwitchClusterMultiPressOngoingEventClass) New() MTRSwitchClusterMultiPressOngoingEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressOngoingEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSwitchClusterMultiPressOngoingEvent) Init() MTRSwitchClusterMultiPressOngoingEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressOngoingEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSwitchClusterMultiPressOngoingEvent) Autorelease() MTRSwitchClusterMultiPressOngoingEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressOngoingEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSwitchClusterMultiPressOngoingEvent creates a new MTRSwitchClusterMultiPressOngoingEvent instance.
func NewMTRSwitchClusterMultiPressOngoingEvent() MTRSwitchClusterMultiPressOngoingEvent {
	return getMTRSwitchClusterMultiPressOngoingEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRSwitchClusterMultiPressOngoingEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterMultiPressOngoingEvent
type MTRSwitchClusterMultiPressOngoingEvent struct {
	objectivec.Object
}

// MTRSwitchClusterMultiPressOngoingEventFrom constructs a [MTRSwitchClusterMultiPressOngoingEvent] from an unsafe.Pointer.
func MTRSwitchClusterMultiPressOngoingEventFrom(ptr unsafe.Pointer) MTRSwitchClusterMultiPressOngoingEvent {
	return MTRSwitchClusterMultiPressOngoingEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRSwitchClusterMultiPressOngoingEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRSwitchClusterMultiPressOngoingEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRSwitchClusterMultiPressOngoingEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRSwitchClusterMultiPressOngoingEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRSwitchClusterMultiPressOngoingEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterMultiPressOngoingEvent/currentNumberOfPressesCounted
func (m_ MTRSwitchClusterMultiPressOngoingEvent) CurrentNumberOfPressesCounted() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("currentNumberOfPressesCounted"))
	return rv
}/* debug [instance_properties/getter]: currentNumberOfPressesCounted */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterMultiPressOngoingEvent/currentNumberOfPressesCounted
func (m_ MTRSwitchClusterMultiPressOngoingEvent) SetCurrentNumberOfPressesCounted(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrentNumberOfPressesCounted:"), value)
}/* debug [instance_properties/setter]: currentNumberOfPressesCounted */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterMultiPressOngoingEvent/newPosition
func (m_ MTRSwitchClusterMultiPressOngoingEvent) NewPosition() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newPosition"))
	return rv
}/* debug [instance_properties/getter]: newPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterMultiPressOngoingEvent/newPosition
func (m_ MTRSwitchClusterMultiPressOngoingEvent) SetNewPosition(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewPosition:"), value)
}/* debug [instance_properties/setter]: newPosition */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRSwitchClusterMultiPressOngoingEvent */



