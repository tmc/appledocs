// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CMRecordedAccelerometerData */


/* debug [class_header]: Header for CMRecordedAccelerometerData */
// The class instance for the [RecordedAccelerometerData] class.
var (
	RecordedAccelerometerDataClass     _RecordedAccelerometerDataClass
	RecordedAccelerometerDataClassOnce sync.Once
)

func getRecordedAccelerometerDataClass() _RecordedAccelerometerDataClass {
	RecordedAccelerometerDataClassOnce.Do(func() {
		RecordedAccelerometerDataClass = _RecordedAccelerometerDataClass{objc.GetClass("CMRecordedAccelerometerData")}
	})
	return RecordedAccelerometerDataClass
}

type _RecordedAccelerometerDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for RecordedAccelerometerData */
// An interface definition for the [RecordedAccelerometerData] class.
type IRecordedAccelerometerData interface {
	IAccelerometerData
	
/* debug [class_interface_properties]: Properties for RecordedAccelerometerData */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for RecordedAccelerometerData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for RecordedAccelerometerData */
// Alloc allocates a new instance without initialization.
func (rc _RecordedAccelerometerDataClass) Alloc() RecordedAccelerometerData {
	rv := objc.Send[RecordedAccelerometerData](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RecordedAccelerometerDataClass) New() RecordedAccelerometerData {
	rv := objc.Send[RecordedAccelerometerData](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RecordedAccelerometerData) Init() RecordedAccelerometerData {
	rv := objc.Send[RecordedAccelerometerData](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RecordedAccelerometerData) Autorelease() RecordedAccelerometerData {
	rv := objc.Send[RecordedAccelerometerData](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRecordedAccelerometerData creates a new RecordedAccelerometerData instance.
func NewRecordedAccelerometerData() RecordedAccelerometerData {
	return getRecordedAccelerometerDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for RecordedAccelerometerData */
// A single piece of accelerometer data that was recorded by the device.
//
// You do not create instances of this class directly. Instead, you use a object to retrieve already recorded data from the system.


// A single piece of accelerometer data that was recorded by the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMotion/CMRecordedAccelerometerData
type RecordedAccelerometerData struct {
	AccelerometerData
}

// RecordedAccelerometerDataFrom constructs a [RecordedAccelerometerData] from an unsafe.Pointer.
//
// A single piece of accelerometer data that was recorded by the device.
func RecordedAccelerometerDataFrom(ptr unsafe.Pointer) RecordedAccelerometerData {
	return RecordedAccelerometerData{
		AccelerometerData: AccelerometerDataFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for RecordedAccelerometerData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for RecordedAccelerometerData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for RecordedAccelerometerData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for RecordedAccelerometerData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for RecordedAccelerometerData */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CMRecordedAccelerometerData */


