// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CMAmbientPressureData */


/* debug [class_header]: Header for CMAmbientPressureData */
// The class instance for the [AmbientPressureData] class.
var (
	AmbientPressureDataClass     _AmbientPressureDataClass
	AmbientPressureDataClassOnce sync.Once
)

func getAmbientPressureDataClass() _AmbientPressureDataClass {
	AmbientPressureDataClassOnce.Do(func() {
		AmbientPressureDataClass = _AmbientPressureDataClass{objc.GetClass("CMAmbientPressureData")}
	})
	return AmbientPressureDataClass
}

type _AmbientPressureDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AmbientPressureData */
// An interface definition for the [AmbientPressureData] class.
type IAmbientPressureData interface {
	ILogItem
	
/* debug [class_interface_properties]: Properties for AmbientPressureData */
	// properties:
	Pressure() unsafe.Pointer
	Temperature() unsafe.Pointer
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AmbientPressureData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AmbientPressureData */
// Alloc allocates a new instance without initialization.
func (ac _AmbientPressureDataClass) Alloc() AmbientPressureData {
	rv := objc.Send[AmbientPressureData](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AmbientPressureDataClass) New() AmbientPressureData {
	rv := objc.Send[AmbientPressureData](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AmbientPressureData) Init() AmbientPressureData {
	rv := objc.Send[AmbientPressureData](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AmbientPressureData) Autorelease() AmbientPressureData {
	rv := objc.Send[AmbientPressureData](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAmbientPressureData creates a new AmbientPressureData instance.
func NewAmbientPressureData() AmbientPressureData {
	return getAmbientPressureDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AmbientPressureData */
// A measurement of the ambient pressure and temperature.


// A measurement of the ambient pressure and temperature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAmbientPressureData
type AmbientPressureData struct {
	LogItem
}

// AmbientPressureDataFrom constructs a [AmbientPressureData] from an unsafe.Pointer.
//
// A measurement of the ambient pressure and temperature.
func AmbientPressureDataFrom(ptr unsafe.Pointer) AmbientPressureData {
	return AmbientPressureData{
		LogItem: LogItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AmbientPressureData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AmbientPressureData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AmbientPressureData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AmbientPressureData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AmbientPressureData */

// The ambient pressure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAmbientPressureData/pressure
func (a_ AmbientPressureData) Pressure() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("pressure"))
	return rv
}/* debug [instance_properties/getter]: pressure */


// The temperature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAmbientPressureData/temperature
func (a_ AmbientPressureData) Temperature() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("temperature"))
	return rv
}/* debug [instance_properties/getter]: temperature */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMAmbientPressureData */



