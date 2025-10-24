// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRTimeSynchronizationClusterTimeZoneStatusEvent */


/* debug [class_header]: Header for MTRTimeSynchronizationClusterTimeZoneStatusEvent */
// The class instance for the [MTRTimeSynchronizationClusterTimeZoneStatusEvent] class.
var (
	MTRTimeSynchronizationClusterTimeZoneStatusEventClass     _MTRTimeSynchronizationClusterTimeZoneStatusEventClass
	MTRTimeSynchronizationClusterTimeZoneStatusEventClassOnce sync.Once
)

func getMTRTimeSynchronizationClusterTimeZoneStatusEventClass() _MTRTimeSynchronizationClusterTimeZoneStatusEventClass {
	MTRTimeSynchronizationClusterTimeZoneStatusEventClassOnce.Do(func() {
		MTRTimeSynchronizationClusterTimeZoneStatusEventClass = _MTRTimeSynchronizationClusterTimeZoneStatusEventClass{objc.GetClass("MTRTimeSynchronizationClusterTimeZoneStatusEvent")}
	})
	return MTRTimeSynchronizationClusterTimeZoneStatusEventClass
}

type _MTRTimeSynchronizationClusterTimeZoneStatusEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRTimeSynchronizationClusterTimeZoneStatusEvent */
// An interface definition for the [MTRTimeSynchronizationClusterTimeZoneStatusEvent] class.
type IMTRTimeSynchronizationClusterTimeZoneStatusEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRTimeSynchronizationClusterTimeZoneStatusEvent */
	// properties:
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	Offset() objc.IObject /* cross-framework: NSNumber */
	SetOffset(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRTimeSynchronizationClusterTimeZoneStatusEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRTimeSynchronizationClusterTimeZoneStatusEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTRTimeSynchronizationClusterTimeZoneStatusEventClass) Alloc() MTRTimeSynchronizationClusterTimeZoneStatusEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneStatusEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRTimeSynchronizationClusterTimeZoneStatusEventClass) New() MTRTimeSynchronizationClusterTimeZoneStatusEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneStatusEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTimeSynchronizationClusterTimeZoneStatusEvent) Init() MTRTimeSynchronizationClusterTimeZoneStatusEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneStatusEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTimeSynchronizationClusterTimeZoneStatusEvent) Autorelease() MTRTimeSynchronizationClusterTimeZoneStatusEvent {
	rv := objc.Send[MTRTimeSynchronizationClusterTimeZoneStatusEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTimeSynchronizationClusterTimeZoneStatusEvent creates a new MTRTimeSynchronizationClusterTimeZoneStatusEvent instance.
func NewMTRTimeSynchronizationClusterTimeZoneStatusEvent() MTRTimeSynchronizationClusterTimeZoneStatusEvent {
	return getMTRTimeSynchronizationClusterTimeZoneStatusEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRTimeSynchronizationClusterTimeZoneStatusEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneStatusEvent
type MTRTimeSynchronizationClusterTimeZoneStatusEvent struct {
	objectivec.Object
}

// MTRTimeSynchronizationClusterTimeZoneStatusEventFrom constructs a [MTRTimeSynchronizationClusterTimeZoneStatusEvent] from an unsafe.Pointer.
func MTRTimeSynchronizationClusterTimeZoneStatusEventFrom(ptr unsafe.Pointer) MTRTimeSynchronizationClusterTimeZoneStatusEvent {
	return MTRTimeSynchronizationClusterTimeZoneStatusEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRTimeSynchronizationClusterTimeZoneStatusEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRTimeSynchronizationClusterTimeZoneStatusEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRTimeSynchronizationClusterTimeZoneStatusEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRTimeSynchronizationClusterTimeZoneStatusEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRTimeSynchronizationClusterTimeZoneStatusEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneStatusEvent/name
func (m_ MTRTimeSynchronizationClusterTimeZoneStatusEvent) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTimeSynchronizationClusterTimeZoneStatusEvent/name
func (m_ MTRTimeSynchronizationClusterTimeZoneStatusEvent) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonestatusevent/offset
func (m_ MTRTimeSynchronizationClusterTimeZoneStatusEvent) Offset() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrtimesynchronizationclustertimezonestatusevent/offset
func (m_ MTRTimeSynchronizationClusterTimeZoneStatusEvent) SetOffset(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOffset:"), value)
}/* debug [instance_properties/setter]: offset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRTimeSynchronizationClusterTimeZoneStatusEvent */



