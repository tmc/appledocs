// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMWaterSubmersionMeasurement */


/* debug [class_header]: Header for CMWaterSubmersionMeasurement */
// The class instance for the [WaterSubmersionMeasurement] class.
var (
	WaterSubmersionMeasurementClass     _WaterSubmersionMeasurementClass
	WaterSubmersionMeasurementClassOnce sync.Once
)

func getWaterSubmersionMeasurementClass() _WaterSubmersionMeasurementClass {
	WaterSubmersionMeasurementClassOnce.Do(func() {
		WaterSubmersionMeasurementClass = _WaterSubmersionMeasurementClass{objc.GetClass("CMWaterSubmersionMeasurement")}
	})
	return WaterSubmersionMeasurementClass
}

type _WaterSubmersionMeasurementClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WaterSubmersionMeasurement */
// An interface definition for the [WaterSubmersionMeasurement] class.
type IWaterSubmersionMeasurement interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WaterSubmersionMeasurement */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WaterSubmersionMeasurement */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WaterSubmersionMeasurement */
// Alloc allocates a new instance without initialization.
func (wc _WaterSubmersionMeasurementClass) Alloc() WaterSubmersionMeasurement {
	rv := objc.Send[WaterSubmersionMeasurement](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WaterSubmersionMeasurementClass) New() WaterSubmersionMeasurement {
	rv := objc.Send[WaterSubmersionMeasurement](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WaterSubmersionMeasurement) Init() WaterSubmersionMeasurement {
	rv := objc.Send[WaterSubmersionMeasurement](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WaterSubmersionMeasurement) Autorelease() WaterSubmersionMeasurement {
	rv := objc.Send[WaterSubmersionMeasurement](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWaterSubmersionMeasurement creates a new WaterSubmersionMeasurement instance.
func NewWaterSubmersionMeasurement() WaterSubmersionMeasurement {
	return getWaterSubmersionMeasurementClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WaterSubmersionMeasurement */
// An update that contains data about the pressure and depth.


// An update that contains data about the pressure and depth.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterSubmersionMeasurement
type WaterSubmersionMeasurement struct {
	objectivec.Object
}

// WaterSubmersionMeasurementFrom constructs a [WaterSubmersionMeasurement] from an unsafe.Pointer.
//
// An update that contains data about the pressure and depth.
func WaterSubmersionMeasurementFrom(ptr unsafe.Pointer) WaterSubmersionMeasurement {
	return WaterSubmersionMeasurement{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WaterSubmersionMeasurement *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WaterSubmersionMeasurement */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WaterSubmersionMeasurement */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WaterSubmersionMeasurement */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WaterSubmersionMeasurement */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMWaterSubmersionMeasurement */


