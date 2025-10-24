// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRBridgedDeviceBasicInformationClusterActiveChangedEvent */


/* debug [class_header]: Header for MTRBridgedDeviceBasicInformationClusterActiveChangedEvent */
// The class instance for the [MTRBridgedDeviceBasicInformationClusterActiveChangedEvent] class.
var (
	MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass     _MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass
	MTRBridgedDeviceBasicInformationClusterActiveChangedEventClassOnce sync.Once
)

func getMTRBridgedDeviceBasicInformationClusterActiveChangedEventClass() _MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass {
	MTRBridgedDeviceBasicInformationClusterActiveChangedEventClassOnce.Do(func() {
		MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass = _MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass{objc.GetClass("MTRBridgedDeviceBasicInformationClusterActiveChangedEvent")}
	})
	return MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass
}

type _MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRBridgedDeviceBasicInformationClusterActiveChangedEvent */
// An interface definition for the [MTRBridgedDeviceBasicInformationClusterActiveChangedEvent] class.
type IMTRBridgedDeviceBasicInformationClusterActiveChangedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRBridgedDeviceBasicInformationClusterActiveChangedEvent */
	// properties:
	PromisedActiveDuration() objc.IObject /* cross-framework: NSNumber */
	SetPromisedActiveDuration(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRBridgedDeviceBasicInformationClusterActiveChangedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRBridgedDeviceBasicInformationClusterActiveChangedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass) Alloc() MTRBridgedDeviceBasicInformationClusterActiveChangedEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterActiveChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRBridgedDeviceBasicInformationClusterActiveChangedEventClass) New() MTRBridgedDeviceBasicInformationClusterActiveChangedEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterActiveChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBridgedDeviceBasicInformationClusterActiveChangedEvent) Init() MTRBridgedDeviceBasicInformationClusterActiveChangedEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterActiveChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBridgedDeviceBasicInformationClusterActiveChangedEvent) Autorelease() MTRBridgedDeviceBasicInformationClusterActiveChangedEvent {
	rv := objc.Send[MTRBridgedDeviceBasicInformationClusterActiveChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBridgedDeviceBasicInformationClusterActiveChangedEvent creates a new MTRBridgedDeviceBasicInformationClusterActiveChangedEvent instance.
func NewMTRBridgedDeviceBasicInformationClusterActiveChangedEvent() MTRBridgedDeviceBasicInformationClusterActiveChangedEvent {
	return getMTRBridgedDeviceBasicInformationClusterActiveChangedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRBridgedDeviceBasicInformationClusterActiveChangedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterActiveChangedEvent
type MTRBridgedDeviceBasicInformationClusterActiveChangedEvent struct {
	objectivec.Object
}

// MTRBridgedDeviceBasicInformationClusterActiveChangedEventFrom constructs a [MTRBridgedDeviceBasicInformationClusterActiveChangedEvent] from an unsafe.Pointer.
func MTRBridgedDeviceBasicInformationClusterActiveChangedEventFrom(ptr unsafe.Pointer) MTRBridgedDeviceBasicInformationClusterActiveChangedEvent {
	return MTRBridgedDeviceBasicInformationClusterActiveChangedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRBridgedDeviceBasicInformationClusterActiveChangedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRBridgedDeviceBasicInformationClusterActiveChangedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRBridgedDeviceBasicInformationClusterActiveChangedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRBridgedDeviceBasicInformationClusterActiveChangedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRBridgedDeviceBasicInformationClusterActiveChangedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterActiveChangedEvent/promisedActiveDuration
func (m_ MTRBridgedDeviceBasicInformationClusterActiveChangedEvent) PromisedActiveDuration() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("promisedActiveDuration"))
	return rv
}/* debug [instance_properties/getter]: promisedActiveDuration */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBridgedDeviceBasicInformationClusterActiveChangedEvent/promisedActiveDuration
func (m_ MTRBridgedDeviceBasicInformationClusterActiveChangedEvent) SetPromisedActiveDuration(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPromisedActiveDuration:"), value)
}/* debug [instance_properties/setter]: promisedActiveDuration */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRBridgedDeviceBasicInformationClusterActiveChangedEvent */



