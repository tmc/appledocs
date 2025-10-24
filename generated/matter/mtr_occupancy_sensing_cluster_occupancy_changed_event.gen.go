// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROccupancySensingClusterOccupancyChangedEvent */


/* debug [class_header]: Header for MTROccupancySensingClusterOccupancyChangedEvent */
// The class instance for the [MTROccupancySensingClusterOccupancyChangedEvent] class.
var (
	MTROccupancySensingClusterOccupancyChangedEventClass     _MTROccupancySensingClusterOccupancyChangedEventClass
	MTROccupancySensingClusterOccupancyChangedEventClassOnce sync.Once
)

func getMTROccupancySensingClusterOccupancyChangedEventClass() _MTROccupancySensingClusterOccupancyChangedEventClass {
	MTROccupancySensingClusterOccupancyChangedEventClassOnce.Do(func() {
		MTROccupancySensingClusterOccupancyChangedEventClass = _MTROccupancySensingClusterOccupancyChangedEventClass{objc.GetClass("MTROccupancySensingClusterOccupancyChangedEvent")}
	})
	return MTROccupancySensingClusterOccupancyChangedEventClass
}

type _MTROccupancySensingClusterOccupancyChangedEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROccupancySensingClusterOccupancyChangedEvent */
// An interface definition for the [MTROccupancySensingClusterOccupancyChangedEvent] class.
type IMTROccupancySensingClusterOccupancyChangedEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROccupancySensingClusterOccupancyChangedEvent */
	// properties:
	Occupancy() objc.IObject /* cross-framework: NSNumber */
	SetOccupancy(value objc.IObject /* cross-framework: NSNumber */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROccupancySensingClusterOccupancyChangedEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROccupancySensingClusterOccupancyChangedEvent */
// Alloc allocates a new instance without initialization.
func (mc _MTROccupancySensingClusterOccupancyChangedEventClass) Alloc() MTROccupancySensingClusterOccupancyChangedEvent {
	rv := objc.Send[MTROccupancySensingClusterOccupancyChangedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROccupancySensingClusterOccupancyChangedEventClass) New() MTROccupancySensingClusterOccupancyChangedEvent {
	rv := objc.Send[MTROccupancySensingClusterOccupancyChangedEvent](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROccupancySensingClusterOccupancyChangedEvent) Init() MTROccupancySensingClusterOccupancyChangedEvent {
	rv := objc.Send[MTROccupancySensingClusterOccupancyChangedEvent](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROccupancySensingClusterOccupancyChangedEvent) Autorelease() MTROccupancySensingClusterOccupancyChangedEvent {
	rv := objc.Send[MTROccupancySensingClusterOccupancyChangedEvent](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROccupancySensingClusterOccupancyChangedEvent creates a new MTROccupancySensingClusterOccupancyChangedEvent instance.
func NewMTROccupancySensingClusterOccupancyChangedEvent() MTROccupancySensingClusterOccupancyChangedEvent {
	return getMTROccupancySensingClusterOccupancyChangedEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROccupancySensingClusterOccupancyChangedEvent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROccupancySensingClusterOccupancyChangedEvent
type MTROccupancySensingClusterOccupancyChangedEvent struct {
	objectivec.Object
}

// MTROccupancySensingClusterOccupancyChangedEventFrom constructs a [MTROccupancySensingClusterOccupancyChangedEvent] from an unsafe.Pointer.
func MTROccupancySensingClusterOccupancyChangedEventFrom(ptr unsafe.Pointer) MTROccupancySensingClusterOccupancyChangedEvent {
	return MTROccupancySensingClusterOccupancyChangedEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROccupancySensingClusterOccupancyChangedEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROccupancySensingClusterOccupancyChangedEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROccupancySensingClusterOccupancyChangedEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROccupancySensingClusterOccupancyChangedEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROccupancySensingClusterOccupancyChangedEvent */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROccupancySensingClusterOccupancyChangedEvent/occupancy
func (m_ MTROccupancySensingClusterOccupancyChangedEvent) Occupancy() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](m_.ID, objc.Sel("occupancy"))
	return rv
}/* debug [instance_properties/getter]: occupancy */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROccupancySensingClusterOccupancyChangedEvent/occupancy
func (m_ MTROccupancySensingClusterOccupancyChangedEvent) SetOccupancy(value objc.IObject /* cross-framework: NSNumber */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOccupancy:"), value)
}/* debug [instance_properties/setter]: occupancy */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROccupancySensingClusterOccupancyChangedEvent */



