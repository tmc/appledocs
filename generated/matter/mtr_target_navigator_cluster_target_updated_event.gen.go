// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRTargetNavigatorClusterTargetUpdatedEvent */


/* debug [class_header]: Header for MTRTargetNavigatorClusterTargetUpdatedEvent */
// The class instance for the [MTRTargetNavigatorClusterTargetUpdatedEvent] class.
var (
	MTRTargetNavigatorClusterTargetUpdatedEventClass     _MTRTargetNavigatorClusterTargetUpdatedEventClass
	MTRTargetNavigatorClusterTargetUpdatedEventClassOnce sync.Once
)

func getMTRTargetNavigatorClusterTargetUpdatedEventClass() _MTRTargetNavigatorClusterTargetUpdatedEventClass {
	MTRTargetNavigatorClusterTargetUpdatedEventClassOnce.Do(func() {
		MTRTargetNavigatorClusterTargetUpdatedEventClass = _MTRTargetNavigatorClusterTargetUpdatedEventClass{objc.GetClass("MTRTargetNavigatorClusterTargetUpdatedEvent")}
	})
	return MTRTargetNavigatorClusterTargetUpdatedEventClass
}

type _MTRTargetNavigatorClusterTargetUpdatedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTargetNavigatorClusterTargetUpdatedEvent */
// An interface definition for the [MTRTargetNavigatorClusterTargetUpdatedEvent] class.
type IMTRTargetNavigatorClusterTargetUpdatedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRTargetNavigatorClusterTargetUpdatedEvent */
	// properties:
	CurrentTarget() objc.IObject /* cross-framework: NSNumber */
	SetCurrentTarget(value objc.IObject /* cross-framework: NSNumber */)
	Data() objc.IObject /* cross-framework: NSData */
	SetData(value objc.IObject /* cross-framework: NSData */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTargetNavigatorClusterTargetUpdatedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTargetNavigatorClusterTargetUpdatedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRTargetNavigatorClusterTargetUpdatedEventClass) Alloc() MTRTargetNavigatorClusterTargetUpdatedEvent {
	rv := objc.Send[MTRTargetNavigatorClusterTargetUpdatedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTargetNavigatorClusterTargetUpdatedEventClass) New() MTRTargetNavigatorClusterTargetUpdatedEvent {
	rv := objc.Send[MTRTargetNavigatorClusterTargetUpdatedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTargetNavigatorClusterTargetUpdatedEvent) Init() MTRTargetNavigatorClusterTargetUpdatedEvent {
	rv := objc.Send[MTRTargetNavigatorClusterTargetUpdatedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTargetNavigatorClusterTargetUpdatedEvent) Autorelease() MTRTargetNavigatorClusterTargetUpdatedEvent {
	rv := objc.Send[MTRTargetNavigatorClusterTargetUpdatedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTargetNavigatorClusterTargetUpdatedEvent creates a new MTRTargetNavigatorClusterTargetUpdatedEvent instance.
func NewMTRTargetNavigatorClusterTargetUpdatedEvent() MTRTargetNavigatorClusterTargetUpdatedEvent {
	return getMTRTargetNavigatorClusterTargetUpdatedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTargetNavigatorClusterTargetUpdatedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterTargetUpdatedEvent
type MTRTargetNavigatorClusterTargetUpdatedEvent struct {
	objectivec.Object
}

// MTRTargetNavigatorClusterTargetUpdatedEventFrom constructs a [MTRTargetNavigatorClusterTargetUpdatedEvent] from an unsafe.Pointer.
func MTRTargetNavigatorClusterTargetUpdatedEventFrom(ptr unsafe.Pointer) MTRTargetNavigatorClusterTargetUpdatedEvent {
	return MTRTargetNavigatorClusterTargetUpdatedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTargetNavigatorClusterTargetUpdatedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTargetNavigatorClusterTargetUpdatedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTargetNavigatorClusterTargetUpdatedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTargetNavigatorClusterTargetUpdatedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTargetNavigatorClusterTargetUpdatedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterTargetUpdatedEvent/currentTarget
func (m_ MTRTargetNavigatorClusterTargetUpdatedEvent) CurrentTarget() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("currentTarget"))
	return rv
}/* debug [instance_properties/getter]: currentTarget */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterTargetUpdatedEvent/currentTarget
func (m_ MTRTargetNavigatorClusterTargetUpdatedEvent) SetCurrentTarget(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setCurrentTarget:"), value)
}/* debug [instance_properties/setter]: currentTarget */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterTargetUpdatedEvent/data
func (m_ MTRTargetNavigatorClusterTargetUpdatedEvent) Data() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterTargetUpdatedEvent/data
func (m_ MTRTargetNavigatorClusterTargetUpdatedEvent) SetData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setData:"), value)
}/* debug [instance_properties/setter]: data */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTargetNavigatorClusterTargetUpdatedEvent */



