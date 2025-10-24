// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CMWaterTemperature */


/* debug [class_header]: Header for CMWaterTemperature */
// The class instance for the [WaterTemperature] class.
var (
	WaterTemperatureClass     _WaterTemperatureClass
	WaterTemperatureClassOnce sync.Once
)

func getWaterTemperatureClass() _WaterTemperatureClass {
	WaterTemperatureClassOnce.Do(func() {
		WaterTemperatureClass = _WaterTemperatureClass{objc.GetClass("CMWaterTemperature")}
	})
	return WaterTemperatureClass
}

type _WaterTemperatureClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for WaterTemperature */
// An interface definition for the [WaterTemperature] class.
type IWaterTemperature interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for WaterTemperature */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for WaterTemperature */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for WaterTemperature */
// Alloc allocates a new instance without initialization.
func (wc _WaterTemperatureClass) Alloc() WaterTemperature {
	rv := objc.Send[WaterTemperature](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (wc _WaterTemperatureClass) New() WaterTemperature {
	rv := objc.Send[WaterTemperature](objc.ID(wc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (w_ WaterTemperature) Init() WaterTemperature {
	rv := objc.Send[WaterTemperature](w_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w_ WaterTemperature) Autorelease() WaterTemperature {
	rv := objc.Send[WaterTemperature](w_.ID, objc.Sel("autorelease"))
	return rv
}

// NewWaterTemperature creates a new WaterTemperature instance.
func NewWaterTemperature() WaterTemperature {
	return getWaterTemperatureClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for WaterTemperature */
// An update that contains data about the water temperature.


// An update that contains data about the water temperature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMWaterTemperature
type WaterTemperature struct {
	objectivec.Object
}

// WaterTemperatureFrom constructs a [WaterTemperature] from an unsafe.Pointer.
//
// An update that contains data about the water temperature.
func WaterTemperatureFrom(ptr unsafe.Pointer) WaterTemperature {
	return WaterTemperature{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for WaterTemperature *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for WaterTemperature */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for WaterTemperature */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for WaterTemperature */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for WaterTemperature */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMWaterTemperature */


