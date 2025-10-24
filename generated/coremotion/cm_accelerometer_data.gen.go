// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class CMAccelerometerData */


/* debug [class_header]: Header for CMAccelerometerData */
// The class instance for the [AccelerometerData] class.
var (
	AccelerometerDataClass     _AccelerometerDataClass
	AccelerometerDataClassOnce sync.Once
)

func getAccelerometerDataClass() _AccelerometerDataClass {
	AccelerometerDataClassOnce.Do(func() {
		AccelerometerDataClass = _AccelerometerDataClass{objc.GetClass("CMAccelerometerData")}
	})
	return AccelerometerDataClass
}

type _AccelerometerDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AccelerometerData */
// An interface definition for the [AccelerometerData] class.
type IAccelerometerData interface {
	ILogItem
	
/* debug [class_interface_properties]: Properties for AccelerometerData */
	// properties:
	Acceleration() objc.IObject /* cross-framework: CMAcceleration */
	Timestamp() float64
	SetTimestamp(value float64)
	AccelerometerData() ICMAccelerometerData
	SetAccelerometerData(value ICMAccelerometerData)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AccelerometerData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AccelerometerData */
// Alloc allocates a new instance without initialization.
func (ac _AccelerometerDataClass) Alloc() AccelerometerData {
	rv := objc.Send[AccelerometerData](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AccelerometerDataClass) New() AccelerometerData {
	rv := objc.Send[AccelerometerData](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AccelerometerData) Init() AccelerometerData {
	rv := objc.Send[AccelerometerData](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AccelerometerData) Autorelease() AccelerometerData {
	rv := objc.Send[AccelerometerData](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAccelerometerData creates a new AccelerometerData instance.
func NewAccelerometerData() AccelerometerData {
	return getAccelerometerDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AccelerometerData */
// A data sample from the device’s three accelerometers.
//
// An application accesses objects through the block handler specified as the last parameter of the method and through the property, both declared by the class. The superclass of , , defines a property that records when the acceleration measurement was taken.


// A data sample from the device’s three accelerometers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAccelerometerData
type AccelerometerData struct {
	LogItem
}

// AccelerometerDataFrom constructs a [AccelerometerData] from an unsafe.Pointer.
//
// A data sample from the device’s three accelerometers.
func AccelerometerDataFrom(ptr unsafe.Pointer) AccelerometerData {
	return AccelerometerData{
		LogItem: LogItemFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AccelerometerData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AccelerometerData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AccelerometerData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AccelerometerData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AccelerometerData */

// The acceleration measured by the accelerometer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMAccelerometerData/acceleration
func (a_ AccelerometerData) Acceleration() objc.IObject /* cross-framework: CMAcceleration */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("acceleration"))
	return rv
}/* debug [instance_properties/getter]: acceleration */


// The time when the logged item is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmlogitem/timestamp
func (a_ AccelerometerData) Timestamp() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("timestamp"))
	return rv
}/* debug [instance_properties/getter]: timestamp */


// The time when the logged item is valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmlogitem/timestamp
func (a_ AccelerometerData) SetTimestamp(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTimestamp:"), value)
}/* debug [instance_properties/setter]: timestamp */


// The latest sample of accelerometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/accelerometerdata
func (a_ AccelerometerData) AccelerometerData() ICMAccelerometerData {
	rv := objc.Send[AccelerometerData](a_.ID, objc.Sel("accelerometerData"))
	return rv
}/* debug [instance_properties/getter]: accelerometerData */


// The latest sample of accelerometer data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremotion/cmmotionmanager/accelerometerdata
func (a_ AccelerometerData) SetAccelerometerData(value ICMAccelerometerData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAccelerometerData:"), value)
}/* debug [instance_properties/setter]: accelerometerData */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMAccelerometerData */



