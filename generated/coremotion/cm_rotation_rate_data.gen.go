// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CMRotationRateData */


/* debug [class_header]: Header for CMRotationRateData */
// The class instance for the [RotationRateData] class.
var (
	RotationRateDataClass     _RotationRateDataClass
	RotationRateDataClassOnce sync.Once
)

func getRotationRateDataClass() _RotationRateDataClass {
	RotationRateDataClassOnce.Do(func() {
		RotationRateDataClass = _RotationRateDataClass{objc.GetClass("CMRotationRateData")}
	})
	return RotationRateDataClass
}

type _RotationRateDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RotationRateData */
// An interface definition for the [RotationRateData] class.
type IRotationRateData interface {
	ILogItem
	
/* debug [class_interface_properties]: Properties for RotationRateData */
	// properties:
	RotationRate() objc.IObject /* cross-framework: CMRotationRate */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RotationRateData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RotationRateData */
// Alloc allocates a new instance without initialization.
func (rc _RotationRateDataClass) Alloc() RotationRateData {
	rv := objc.Send[RotationRateData](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RotationRateDataClass) New() RotationRateData {
	rv := objc.Send[RotationRateData](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RotationRateData) Init() RotationRateData {
	rv := objc.Send[RotationRateData](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RotationRateData) Autorelease() RotationRateData {
	rv := objc.Send[RotationRateData](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRotationRateData creates a new RotationRateData instance.
func NewRotationRateData() RotationRateData {
	return getRotationRateDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RotationRateData */
// A data object that contains a single rotation-rate measurement.


// A data object that contains a single rotation-rate measurement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRotationRateData
type RotationRateData struct {
	LogItem
}

// RotationRateDataFrom constructs a [RotationRateData] from an unsafe.Pointer.
//
// A data object that contains a single rotation-rate measurement.
func RotationRateDataFrom(ptr unsafe.Pointer) RotationRateData {
	return RotationRateData{
		LogItem: LogItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RotationRateData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RotationRateData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RotationRateData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RotationRateData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RotationRateData */

// The rotation rate measured by the gyroscope.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRotationRateData/rotationRate
func (r_ RotationRateData) RotationRate() objc.IObject /* cross-framework: CMRotationRate */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("rotationRate"))
	return rv
}/* debug [instance_properties/getter]: rotationRate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMRotationRateData */





