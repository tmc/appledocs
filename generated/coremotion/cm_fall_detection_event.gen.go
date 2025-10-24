// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMFallDetectionEvent */


/* debug [class_header]: Header for CMFallDetectionEvent */
// The class instance for the [FallDetectionEvent] class.
var (
	FallDetectionEventClass     _FallDetectionEventClass
	FallDetectionEventClassOnce sync.Once
)

func getFallDetectionEventClass() _FallDetectionEventClass {
	FallDetectionEventClassOnce.Do(func() {
		FallDetectionEventClass = _FallDetectionEventClass{objc.GetClass("CMFallDetectionEvent")}
	})
	return FallDetectionEventClass
}

type _FallDetectionEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FallDetectionEvent */
// An interface definition for the [FallDetectionEvent] class.
type IFallDetectionEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for FallDetectionEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FallDetectionEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FallDetectionEvent */
// Alloc allocates a new instance without initialization.
func (fc _FallDetectionEventClass) Alloc() FallDetectionEvent {
	rv := objc.Send[FallDetectionEvent](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FallDetectionEventClass) New() FallDetectionEvent {
	rv := objc.Send[FallDetectionEvent](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FallDetectionEvent) Init() FallDetectionEvent {
	rv := objc.Send[FallDetectionEvent](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FallDetectionEvent) Autorelease() FallDetectionEvent {
	rv := objc.Send[FallDetectionEvent](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFallDetectionEvent creates a new FallDetectionEvent instance.
func NewFallDetectionEvent() FallDetectionEvent {
	return getFallDetectionEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FallDetectionEvent */
// An object that contains data about a fall detection event.


// An object that contains data about a fall detection event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMFallDetectionEvent
type FallDetectionEvent struct {
	objectivec.Object
}

// FallDetectionEventFrom constructs a [FallDetectionEvent] from an unsafe.Pointer.
//
// An object that contains data about a fall detection event.
func FallDetectionEventFrom(ptr unsafe.Pointer) FallDetectionEvent {
	return FallDetectionEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FallDetectionEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FallDetectionEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FallDetectionEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FallDetectionEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FallDetectionEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMFallDetectionEvent */


