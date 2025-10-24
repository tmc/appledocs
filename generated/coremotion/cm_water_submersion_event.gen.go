// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMWaterSubmersionEvent */


/* debug [class_header]: Header for CMWaterSubmersionEvent */
// The class instance for the [WaterSubmersionEvent] class.
var (
	WaterSubmersionEventClass     _WaterSubmersionEventClass
	WaterSubmersionEventClassOnce sync.Once
)

func getWaterSubmersionEventClass() _WaterSubmersionEventClass {
	WaterSubmersionEventClassOnce.Do(func() {
		WaterSubmersionEventClass = _WaterSubmersionEventClass{objc.GetClass("CMWaterSubmersionEvent")}
	})
	return WaterSubmersionEventClass
}

type _WaterSubmersionEventClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WaterSubmersionEvent */
// An interface definition for the [WaterSubmersionEvent] class.
type IWaterSubmersionEvent interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WaterSubmersionEvent */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WaterSubmersionEvent */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WaterSubmersionEvent */
// Alloc allocates a new instance without initialization.
func (wc _WaterSubmersionEventClass) Alloc() WaterSubmersionEvent {
	rv := objc.Send[WaterSubmersionEvent](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WaterSubmersionEventClass) New() WaterSubmersionEvent {
	rv := objc.Send[WaterSubmersionEvent](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WaterSubmersionEvent) Init() WaterSubmersionEvent {
	rv := objc.Send[WaterSubmersionEvent](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WaterSubmersionEvent) Autorelease() WaterSubmersionEvent {
	rv := objc.Send[WaterSubmersionEvent](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWaterSubmersionEvent creates a new WaterSubmersionEvent instance.
func NewWaterSubmersionEvent() WaterSubmersionEvent {
	return getWaterSubmersionEventClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WaterSubmersionEvent */
// An event indicating that the device’s submersion state has changed.


// An event indicating that the device’s submersion state has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionEvent
type WaterSubmersionEvent struct {
	objectivec.Object
}

// WaterSubmersionEventFrom constructs a [WaterSubmersionEvent] from an unsafe.Pointer.
//
// An event indicating that the device’s submersion state has changed.
func WaterSubmersionEventFrom(ptr unsafe.Pointer) WaterSubmersionEvent {
	return WaterSubmersionEvent{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WaterSubmersionEvent *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WaterSubmersionEvent */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WaterSubmersionEvent */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WaterSubmersionEvent */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WaterSubmersionEvent */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMWaterSubmersionEvent */


