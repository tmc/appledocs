// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CMGyroData */


/* debug [class_header]: Header for CMGyroData */
// The class instance for the [GyroData] class.
var (
	GyroDataClass     _GyroDataClass
	GyroDataClassOnce sync.Once
)

func getGyroDataClass() _GyroDataClass {
	GyroDataClassOnce.Do(func() {
		GyroDataClass = _GyroDataClass{objc.GetClass("CMGyroData")}
	})
	return GyroDataClass
}

type _GyroDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GyroData */
// An interface definition for the [GyroData] class.
type IGyroData interface {
	ILogItem
	
/* debug [class_interface_properties]: Properties for GyroData */
	// properties:
	RotationRate() objc.IObject /* cross-framework: CMRotationRate */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GyroData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GyroData */
// Alloc allocates a new instance without initialization.
func (gc _GyroDataClass) Alloc() GyroData {
	rv := objc.Send[GyroData](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GyroDataClass) New() GyroData {
	rv := objc.Send[GyroData](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GyroData) Init() GyroData {
	rv := objc.Send[GyroData](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GyroData) Autorelease() GyroData {
	rv := objc.Send[GyroData](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGyroData creates a new GyroData instance.
func NewGyroData() GyroData {
	return getGyroDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GyroData */
// A single measurement of the device’s rotation rate.
//
// An application receives or samples objects at regular intervals after calling the method or the method of the class.


// A single measurement of the device’s rotation rate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMGyroData
type GyroData struct {
	LogItem
}

// GyroDataFrom constructs a [GyroData] from an unsafe.Pointer.
//
// A single measurement of the device’s rotation rate.
func GyroDataFrom(ptr unsafe.Pointer) GyroData {
	return GyroData{
		LogItem: LogItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GyroData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GyroData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GyroData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GyroData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GyroData */

// The rotation rate as measured by the device’s gyroscope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMGyroData/rotationRate
func (g_ GyroData) RotationRate() objc.IObject /* cross-framework: CMRotationRate */ {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("rotationRate"))
	return rv
}/* debug [instance_properties/getter]: rotationRate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMGyroData */



