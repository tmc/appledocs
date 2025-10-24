// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRSwitchClusterMultiPressCompleteEvent */


/* debug [class_header]: Header for MTRSwitchClusterMultiPressCompleteEvent */
// The class instance for the [MTRSwitchClusterMultiPressCompleteEvent] class.
var (
	MTRSwitchClusterMultiPressCompleteEventClass     _MTRSwitchClusterMultiPressCompleteEventClass
	MTRSwitchClusterMultiPressCompleteEventClassOnce sync.Once
)

func getMTRSwitchClusterMultiPressCompleteEventClass() _MTRSwitchClusterMultiPressCompleteEventClass {
	MTRSwitchClusterMultiPressCompleteEventClassOnce.Do(func() {
		MTRSwitchClusterMultiPressCompleteEventClass = _MTRSwitchClusterMultiPressCompleteEventClass{objc.GetClass("MTRSwitchClusterMultiPressCompleteEvent")}
	})
	return MTRSwitchClusterMultiPressCompleteEventClass
}

type _MTRSwitchClusterMultiPressCompleteEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRSwitchClusterMultiPressCompleteEvent */
// An interface definition for the [MTRSwitchClusterMultiPressCompleteEvent] class.
type IMTRSwitchClusterMultiPressCompleteEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRSwitchClusterMultiPressCompleteEvent */
	// properties:
	NewPosition() objc.IObject /* cross-framework: NSNumber */
	SetNewPosition(value objc.IObject /* cross-framework: NSNumber */)
	PreviousPosition() objc.IObject /* cross-framework: NSNumber */
	SetPreviousPosition(value objc.IObject /* cross-framework: NSNumber */)
	TotalNumberOfPressesCounted() objc.IObject /* cross-framework: NSNumber */
	SetTotalNumberOfPressesCounted(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRSwitchClusterMultiPressCompleteEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRSwitchClusterMultiPressCompleteEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRSwitchClusterMultiPressCompleteEventClass) Alloc() MTRSwitchClusterMultiPressCompleteEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressCompleteEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRSwitchClusterMultiPressCompleteEventClass) New() MTRSwitchClusterMultiPressCompleteEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressCompleteEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRSwitchClusterMultiPressCompleteEvent) Init() MTRSwitchClusterMultiPressCompleteEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressCompleteEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRSwitchClusterMultiPressCompleteEvent) Autorelease() MTRSwitchClusterMultiPressCompleteEvent {
	rv := objc.Send[MTRSwitchClusterMultiPressCompleteEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRSwitchClusterMultiPressCompleteEvent creates a new MTRSwitchClusterMultiPressCompleteEvent instance.
func NewMTRSwitchClusterMultiPressCompleteEvent() MTRSwitchClusterMultiPressCompleteEvent {
	return getMTRSwitchClusterMultiPressCompleteEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRSwitchClusterMultiPressCompleteEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterMultiPressCompleteEvent
type MTRSwitchClusterMultiPressCompleteEvent struct {
	objectivec.Object
}

// MTRSwitchClusterMultiPressCompleteEventFrom constructs a [MTRSwitchClusterMultiPressCompleteEvent] from an unsafe.Pointer.
func MTRSwitchClusterMultiPressCompleteEventFrom(ptr unsafe.Pointer) MTRSwitchClusterMultiPressCompleteEvent {
	return MTRSwitchClusterMultiPressCompleteEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRSwitchClusterMultiPressCompleteEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRSwitchClusterMultiPressCompleteEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRSwitchClusterMultiPressCompleteEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRSwitchClusterMultiPressCompleteEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRSwitchClusterMultiPressCompleteEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterMultiPressCompleteEvent/newPosition
func (m_ MTRSwitchClusterMultiPressCompleteEvent) NewPosition() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("newPosition"))
	return rv
}/* debug [instance_properties/getter]: newPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterMultiPressCompleteEvent/newPosition
func (m_ MTRSwitchClusterMultiPressCompleteEvent) SetNewPosition(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setNewPosition:"), value)
}/* debug [instance_properties/setter]: newPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterMultiPressCompleteEvent/previousPosition
func (m_ MTRSwitchClusterMultiPressCompleteEvent) PreviousPosition() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("previousPosition"))
	return rv
}/* debug [instance_properties/getter]: previousPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterMultiPressCompleteEvent/previousPosition
func (m_ MTRSwitchClusterMultiPressCompleteEvent) SetPreviousPosition(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPreviousPosition:"), value)
}/* debug [instance_properties/setter]: previousPosition */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterMultiPressCompleteEvent/totalNumberOfPressesCounted
func (m_ MTRSwitchClusterMultiPressCompleteEvent) TotalNumberOfPressesCounted() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("totalNumberOfPressesCounted"))
	return rv
}/* debug [instance_properties/getter]: totalNumberOfPressesCounted */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRSwitchClusterMultiPressCompleteEvent/totalNumberOfPressesCounted
func (m_ MTRSwitchClusterMultiPressCompleteEvent) SetTotalNumberOfPressesCounted(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setTotalNumberOfPressesCounted:"), value)
}/* debug [instance_properties/setter]: totalNumberOfPressesCounted */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRSwitchClusterMultiPressCompleteEvent */



