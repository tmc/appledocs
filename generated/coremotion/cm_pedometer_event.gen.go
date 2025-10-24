// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMPedometerEvent */


/* debug [class_header]: Header for CMPedometerEvent */
// The class instance for the [PedometerEvent] class.
var (
	PedometerEventClass     _PedometerEventClass
	PedometerEventClassOnce sync.Once
)

func getPedometerEventClass() _PedometerEventClass {
	PedometerEventClassOnce.Do(func() {
		PedometerEventClass = _PedometerEventClass{objc.GetClass("CMPedometerEvent")}
	})
	return PedometerEventClass
}

type _PedometerEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PedometerEvent */
// An interface definition for the [PedometerEvent] class.
type IPedometerEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PedometerEvent */
	// properties:
	Date() objc.IObject /* cross-framework: NSDate */
	Type() PedometerEventType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PedometerEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PedometerEvent */
// Alloc allocates a new instance without initialization.
func (pc _PedometerEventClass) Alloc() PedometerEvent {
	rv := objc.Send[PedometerEvent](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PedometerEventClass) New() PedometerEvent {
	rv := objc.Send[PedometerEvent](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PedometerEvent) Init() PedometerEvent {
	rv := objc.Send[PedometerEvent](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PedometerEvent) Autorelease() PedometerEvent {
	rv := objc.Send[PedometerEvent](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPedometerEvent creates a new PedometerEvent instance.
func NewPedometerEvent() PedometerEvent {
	return getPedometerEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PedometerEvent */
// A change in the user’s pedestrian activity.


// A change in the user’s pedestrian activity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerEvent
type PedometerEvent struct {
	objectivec.Object
}

// PedometerEventFrom constructs a [PedometerEvent] from an unsafe.Pointer.
//
// A change in the user’s pedestrian activity.
func PedometerEventFrom(ptr unsafe.Pointer) PedometerEvent {
	return PedometerEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PedometerEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PedometerEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PedometerEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PedometerEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PedometerEvent */

// The date on which the pedometer event was recorded.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerEvent/date
func (p_ PedometerEvent) Date() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("date"))
	return rv
}/* debug [instance_properties/getter]: date */


// The type of change that occurred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMPedometerEvent/type
func (p_ PedometerEvent) Type() PedometerEventType {
	rv := objc.Send[PedometerEventType](p_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMPedometerEvent */



